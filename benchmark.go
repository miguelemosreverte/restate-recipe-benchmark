package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/disk"
    "github.com/shirou/gopsutil/v3/mem"
    "html/template"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "sync"
    "sync/atomic"
    "time"
)

type SystemMetrics struct {
    CPUUsage    float64 `json:"cpuUsage"`
    MemoryUsage float64 `json:"memoryUsage"`
    DiskUsage   float64 `json:"diskUsage"`
    DiskIOPS    float64 `json:"diskIops"`
}

type Metrics struct {
    Timestamp       int64         `json:"timestamp"`
    TPS            float64       `json:"tps"`
    CurrentUsers   int          `json:"currentUsers"`
    SuccessRequests int64        `json:"successRequests"`
    FailureRate    float64       `json:"failureRate"`
    SystemMetrics  SystemMetrics `json:"systemMetrics"`
    InFlightRetries int64        `json:"inFlightRetries"`
}

type BenchmarkState struct {
    mutex            sync.RWMutex
    metrics          []Metrics
    currentUsers     int
    successRequests  atomic.Int64
    failedRequests   atomic.Int64
    inFlightRetries  atomic.Int64
    diskUsageStart   uint64
    isRunning        bool
    cooldownPeriod   time.Duration
    startTime        time.Time
    restateDataPath  string
    minStableTime    time.Duration
    maxInFlightRetries int64
    scaleDownFactor  float64
    minUsers         int
    maxUsers         int    // New field for binary search
    startingUsers    int
    stabilityThreshold time.Duration
    lastStableTPS    float64 // New field to track last stable TPS
}

func (s *BenchmarkState) collectSystemMetrics() (SystemMetrics, error) {
    cpuPercent, err := cpu.Percent(0, false)
    if err != nil {
        cpuPercent = []float64{0.0}
    }

    memInfo, err := mem.VirtualMemory()
    if err != nil {
        memInfo = &mem.VirtualMemoryStat{}
    }

    var diskUsageGB float64
    diskUsage, err := disk.Usage(s.restateDataPath)
    if err != nil {
        log.Printf("Warning: Failed to get disk usage: %v", err)
        dirSize, err := getDirSize(s.restateDataPath)
        if err != nil {
            log.Printf("Warning: Failed to get directory size: %v", err)
        } else {
            diskUsageGB = float64(dirSize) / (1024 * 1024 * 1024)
        }
    } else {
        diskUsageGB = float64(diskUsage.Used) / (1024 * 1024 * 1024)
    }

    var totalIOPS float64
    diskStats, err := disk.IOCounters()
    if err == nil {
        for _, stat := range diskStats {
            totalIOPS += float64(stat.ReadCount + stat.WriteCount)
        }
    }

    return SystemMetrics{
        CPUUsage:    cpuPercent[0],
        MemoryUsage: memInfo.UsedPercent,
        DiskUsage:   diskUsageGB,
        DiskIOPS:    totalIOPS,
    }, nil
}

func getDirSize(path string) (int64, error) {
    var size int64
    err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            size += info.Size()
        }
        return err
    })
    return size, err
}

func (s *BenchmarkState) recordMetrics() {
    sysMetrics, err := s.collectSystemMetrics()
    if err != nil {
        log.Printf("Warning: Error collecting system metrics: %v", err)
        if len(s.metrics) > 0 {
            sysMetrics = s.metrics[len(s.metrics)-1].SystemMetrics
        }
    }

    totalReqs := s.successRequests.Load() + s.failedRequests.Load()
    var failureRate float64
    if totalReqs > 0 {
        failureRate = float64(s.failedRequests.Load()) / float64(totalReqs) * 100
    }

    duration := time.Since(s.startTime).Seconds()
    tps := float64(s.successRequests.Load()) / duration

    metrics := Metrics{
        Timestamp:      time.Now().Unix(),
        TPS:           tps,
        CurrentUsers:  s.currentUsers,
        SuccessRequests: s.successRequests.Load(),
        FailureRate:   failureRate,
        SystemMetrics: sysMetrics,
        InFlightRetries: s.inFlightRetries.Load(),
    }

    s.mutex.Lock()
    s.metrics = append(s.metrics, metrics)
    s.mutex.Unlock()
}

func sendRequest(client *http.Client) (bool, error) {
    payload := []byte(`"user1"`)
    resp, err := client.Post("http://localhost:8080/Greeter/greet",
        "application/json",
        bytes.NewBuffer(payload))

    if err != nil {
        return false, err
    }
    defer resp.Body.Close()

    return resp.StatusCode == http.StatusOK, nil
}

func worker(state *BenchmarkState, stopCh chan bool, wg *sync.WaitGroup) {
    defer wg.Done()
    client := &http.Client{Timeout: time.Second * 60}

    for {
        select {
        case <-stopCh:
            return
        default:
            if state.inFlightRetries.Load() > state.maxInFlightRetries {
                time.Sleep(time.Millisecond * 100)
                continue
            }

            success, err := sendRequest(client)
            if success {
                state.successRequests.Add(1)
                if state.inFlightRetries.Load() > 0 {
                    state.inFlightRetries.Add(-1)
                }
            } else {
                state.failedRequests.Add(1)
                if err != nil {
                    state.inFlightRetries.Add(1)
                }
            }

            time.Sleep(time.Millisecond * 50)
        }
    }
}

func metricsHtmlHandler(state *BenchmarkState) http.HandlerFunc {
    tmpl := template.Must(template.ParseFiles("templates/metrics.html"))

    return func(w http.ResponseWriter, r *http.Request) {
        state.mutex.RLock()
        defer state.mutex.RUnlock()

        if len(state.metrics) == 0 {
            w.Write([]byte("No metrics available yet"))
            return
        }

        w.Header().Set("Content-Type", "text/html")
        err := tmpl.Execute(w, state.metrics[len(state.metrics)-1])
        if err != nil {
            log.Printf("Error executing template: %v", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }
    }
}

func runBenchmark(state *BenchmarkState) {
    benchmarkTimeout := time.After(60 * time.Minute)
    currentIteration := 0
    maxTotalIterations := 50
    targetTPS := 10000.0 // Target TPS

    state.cooldownPeriod = time.Second * 5
    state.startingUsers = 20
    state.maxUsers = 20000
    state.currentUsers = state.startingUsers
    state.isRunning = true
    state.startTime = time.Now()

    state.metrics = make([]Metrics, 0)

    log.Printf("Benchmark initialized with target TPS: %.2f, max users: %d, cooldown period: %v", targetTPS, state.maxUsers, state.cooldownPeriod)

    bestTPS := 0.0
    bestUsers := state.startingUsers
    stableUsers := state.startingUsers
    unstableUsers := state.maxUsers

    for state.isRunning {
        select {
        case <-benchmarkTimeout:
            log.Printf("Benchmark timeout reached. Exiting...")
            state.isRunning = false
            break
        default:
        }

        if currentIteration >= maxTotalIterations {
            log.Printf("Maximum iterations reached. Exiting...")
            break
        }

        currentIteration++

        log.Printf("Starting iteration %d with %d users", currentIteration, state.currentUsers)
        iterationStopCh := make(chan bool)
        var iterationWg sync.WaitGroup

        state.successRequests.Store(0)
        state.failedRequests.Store(0)
        state.inFlightRetries.Store(0)
        state.startTime = time.Now()

        for i := 0; i < state.currentUsers; i++ {
            iterationWg.Add(1)
            go worker(state, iterationStopCh, &iterationWg)
        }

        monitoringDone := make(chan bool)
        recentTPS := make([]float64, 0, 5)
        testDuration := time.Second * 15

        go func() {
            defer close(monitoringDone)
            testStart := time.Now()
            ticker := time.NewTicker(time.Second)

            for time.Since(testStart) < testDuration {
                <-ticker.C
                state.recordMetrics()

                state.mutex.RLock()
                metrics := state.metrics[len(state.metrics)-1]
                state.mutex.RUnlock()

                log.Printf("TPS: %.2f, Timeouts: %d, Users: %d", metrics.TPS, metrics.InFlightRetries, state.currentUsers)

                if metrics.TPS > 0 {
                    recentTPS = append(recentTPS, metrics.TPS)
                    if len(recentTPS) > 5 {
                        recentTPS = recentTPS[1:]
                    }
                }
            }

            ticker.Stop()
        }()

        <-monitoringDone

        close(iterationStopCh)
        cleanup := make(chan bool)
        go func() {
            iterationWg.Wait()
            cleanup <- true
        }()

        select {
        case <-cleanup:
        case <-time.After(10 * time.Second):
            log.Printf("Worker cleanup timed out")
        }

        avgTPS := 0.0
        for _, tps := range recentTPS {
            avgTPS += tps
        }
        if len(recentTPS) > 0 {
            avgTPS /= float64(len(recentTPS))
        }

        timeoutRate := float64(state.failedRequests.Load()) / float64(state.successRequests.Load()+state.failedRequests.Load())

        log.Printf("Iteration %d complete. Avg TPS: %.2f, Timeout Rate: %.2f", currentIteration, avgTPS, timeoutRate)

        if avgTPS > bestTPS {
            bestTPS = avgTPS
            bestUsers = state.currentUsers
            log.Printf("New best configuration: %d users, %.2f TPS", bestUsers, bestTPS)
        }

        if avgTPS >= targetTPS {
            log.Printf("Target TPS %.2f achieved with %d users!", targetTPS, state.currentUsers)
            break
        }

        if timeoutRate > 0.5 {
            log.Printf("High timeout rate detected (%.2f). Marking %d users as unstable.", timeoutRate, state.currentUsers)
            unstableUsers = state.currentUsers
        } else {
            log.Printf("Stable configuration detected at %d users.", state.currentUsers)
            stableUsers = state.currentUsers
        }

        // Binary search for convergence
        if unstableUsers > stableUsers+1 {
            state.currentUsers = (stableUsers + unstableUsers) / 2
            log.Printf("Adjusting users to midpoint: %d", state.currentUsers)
        } else {
            log.Printf("Converged to optimal configuration: %d users, %.2f TPS", bestUsers, bestTPS)
            break
        }

        time.Sleep(state.cooldownPeriod)
    }

    log.Printf("Benchmark completed. Best configuration: %d users, %.2f TPS", bestUsers, bestTPS)
    generateReport(state)
}




func generateReport(state *BenchmarkState) {
    if err := os.MkdirAll("reports", 0755); err != nil {
        log.Fatal(err)
    }

    tmpl := template.Must(template.ParseFiles("templates/index.html"))

    type TemplateData struct {
        RecommendedTPS   float64
        RecommendedUsers int
        MaxCPU          float64
        MaxMemory       float64
        MaxDiskIO       uint64
        Timestamps      []string
        TPSValues       []float64
        CPUValues       []float64
        MemoryValues    []float64
        DiskReadValues  []uint64
        DiskWriteValues []uint64
        RetryDebtValues []int64
    }

    state.mutex.RLock()
    // Find the highest stable TPS from the metrics
    var maxStableTPS float64
    for _, m := range state.metrics {
        if m.TPS > maxStableTPS && m.InFlightRetries == 0 {
            maxStableTPS = m.TPS
        }
    }

    data := TemplateData{
        RecommendedTPS:   maxStableTPS,  // Use the actual maximum stable TPS
        RecommendedUsers: state.currentUsers,
    }

    for _, m := range state.metrics {
        timeStr := time.Unix(m.Timestamp, 0).Format(time.RFC3339)
        data.Timestamps = append(data.Timestamps, timeStr)
        data.TPSValues = append(data.TPSValues, m.TPS)
        data.CPUValues = append(data.CPUValues, m.SystemMetrics.CPUUsage)
        data.MemoryValues = append(data.MemoryValues, m.SystemMetrics.MemoryUsage)
        data.DiskReadValues = append(data.DiskReadValues, uint64(m.SystemMetrics.DiskIOPS))
        data.DiskWriteValues = append(data.DiskWriteValues, uint64(m.SystemMetrics.DiskIOPS))
        data.RetryDebtValues = append(data.RetryDebtValues, m.InFlightRetries)

        if m.SystemMetrics.CPUUsage > data.MaxCPU {
            data.MaxCPU = m.SystemMetrics.CPUUsage
        }
        if m.SystemMetrics.MemoryUsage > float64(data.MaxMemory) {
            data.MaxMemory = m.SystemMetrics.MemoryUsage
        }
        if uint64(m.SystemMetrics.DiskIOPS) > data.MaxDiskIO {
            data.MaxDiskIO = uint64(m.SystemMetrics.DiskIOPS)
        }
    }
    state.mutex.RUnlock()

    reportFile := fmt.Sprintf("reports/benchmark-report-%s.html",
        time.Now().Format("2006-01-02-15-04-05"))

    f, err := os.Create(reportFile)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    if err := tmpl.Execute(f, data); err != nil {
        log.Fatal(err)
    }

    log.Printf("Report generated: %s\n", reportFile)
}

func main() {
    restateDataPath, err := filepath.Abs("restate-data")
    if err != nil {
        log.Fatal(err)
    }

    if _, err := os.Stat(restateDataPath); os.IsNotExist(err) {
        log.Fatal("Restate data directory not found: ", restateDataPath)
    }

    state := &BenchmarkState{
        restateDataPath: restateDataPath,
    }

    http.HandleFunc("/metrics.html", metricsHtmlHandler(state))
    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        state.mutex.RLock()
        if len(state.metrics) > 0 {
            json.NewEncoder(w).Encode(state.metrics[len(state.metrics)-1])
        }
        state.mutex.RUnlock()
    })

    go http.ListenAndServe(":3000", nil)

    fmt.Println("Starting benchmark...")
    runBenchmark(state)
}
