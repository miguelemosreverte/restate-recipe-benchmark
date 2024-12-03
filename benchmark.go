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
    "math"
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
    benchmarkTimeout := time.After(60 * time.Minute) // Extended for fine-tuning
    maxTotalIterations := 10
    currentIteration := 0
    
    state.cooldownPeriod = time.Second * 5  
    state.maxInFlightRetries = 20
    state.minUsers = 5
    state.startingUsers = 20
    state.maxUsers = 100

    state.currentUsers = state.startingUsers
    state.isRunning = true
    state.startTime = time.Now()
    state.metrics = make([]Metrics, 0)

    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()

    diskInfo, err := disk.Usage(state.restateDataPath)
    if err != nil {
        log.Printf("Warning: Could not get initial disk usage: %v", err)
    } else {
        state.diskUsageStart = diskInfo.Used
    }

    log.Printf("Monitoring Restate data at: %s\n", state.restateDataPath)
    log.Printf("Initial disk usage: %d bytes\n", state.diskUsageStart)

    // Search variables
    bestStableUsers := state.minUsers
    bestStableTPS := 0.0
    lastStableUsers := 0
    firstUnstableUsers := 0
    const minSamplesNeeded = 3    
    const tpsVarianceThreshold = 0.1 

    // Run initial binary search phase
    searchPhase: for state.isRunning {
        select {
        case <-benchmarkTimeout:
            log.Printf("Benchmark timeout reached")
            goto fineTuning
        default:
        }

        currentIteration++
        if currentIteration > maxTotalIterations {
            log.Printf("Reached maximum iterations (%d)", maxTotalIterations)
            goto fineTuning
        }

        log.Printf("Starting test iteration with %d users", state.currentUsers)
        
        iterationStopCh := make(chan bool)
        var iterationWg sync.WaitGroup
        recentTPS := make([]float64, 0, minSamplesNeeded)

        // Reset counters
        state.successRequests.Store(0)
        state.failedRequests.Store(0)
        state.inFlightRetries.Store(0)
        state.startTime = time.Now()

        // Start workers
        for i := 0; i < state.currentUsers; i++ {
            iterationWg.Add(1)
            go worker(state, iterationStopCh, &iterationWg)
        }

        // Monitor performance
        isStable := false
        var currentAvgTPS float64
        maxTestDuration := 10 * time.Second
        testStart := time.Now()
        testTicker := time.NewTicker(time.Second)
        monitoringDone := make(chan bool)

        go func() {
            defer close(monitoringDone)
            for time.Since(testStart) < maxTestDuration {
                select {
                case <-testTicker.C:
                    state.recordMetrics()

                    state.mutex.RLock()
                    metrics := state.metrics[len(state.metrics)-1]
                    state.mutex.RUnlock()

                    currentTPS := metrics.TPS
                    log.Printf("TPS: %.2f, In-flight Retries: %d, Users: %d",
                        currentTPS, metrics.InFlightRetries, state.currentUsers)

                    if metrics.InFlightRetries > 0 {
                        isStable = false
                        return
                    }

                    if currentTPS > 0 {
                        recentTPS = append(recentTPS, currentTPS)
                        if len(recentTPS) >= minSamplesNeeded {
                            if len(recentTPS) > minSamplesNeeded {
                                recentTPS = recentTPS[1:]
                            }
                            
                            avg := 0.0
                            for _, tps := range recentTPS {
                                avg += tps
                            }
                            avg /= float64(len(recentTPS))
                            currentAvgTPS = avg
                            
                            isStable = true
                            for _, tps := range recentTPS {
                                if math.Abs(tps-avg)/avg > tpsVarianceThreshold {
                                    isStable = false
                                    break
                                }
                            }
                        }
                    }
                }
            }
        }()

        <-monitoringDone
        testTicker.Stop()

        // Clean up workers
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

        time.Sleep(state.cooldownPeriod)

        if isStable {
            log.Printf("Test was stable with TPS %.2f", currentAvgTPS)
            lastStableUsers = state.currentUsers
            if currentAvgTPS > bestStableTPS {
                bestStableTPS = currentAvgTPS
                bestStableUsers = state.currentUsers
                log.Printf("New best configuration found!")
            }
            
            nextUsers := state.currentUsers * 2
            log.Printf("Doubling users from %d to %d", state.currentUsers, nextUsers)
            state.currentUsers = nextUsers
        } else {
            log.Printf("Test was unstable")
            firstUnstableUsers = state.currentUsers
            if lastStableUsers > 0 {
                log.Printf("Found bounds: %d (stable) to %d (unstable)", 
                    lastStableUsers, firstUnstableUsers)
                break searchPhase
            }
            
            state.currentUsers = int(float64(state.currentUsers) * 0.75)
        }
    }

fineTuning:
    if lastStableUsers > 0 && firstUnstableUsers > 0 {
        log.Printf("Starting fine-tuning phase between %d and %d users", 
            lastStableUsers, firstUnstableUsers)

        // Two fine-tuning attempts
        for attempt := 1; attempt <= 2; attempt++ {
            candidateUsers := (lastStableUsers + firstUnstableUsers) / 2
            state.currentUsers = candidateUsers
            log.Printf("Fine-tuning attempt %d/2 with %d users", attempt, candidateUsers)

            // Run test with same monitoring logic as above
            iterationStopCh := make(chan bool)
            var iterationWg sync.WaitGroup
            recentTPS := make([]float64, 0, minSamplesNeeded)
            
            state.successRequests.Store(0)
            state.failedRequests.Store(0)
            state.inFlightRetries.Store(0)
            state.startTime = time.Now()

            for i := 0; i < state.currentUsers; i++ {
                iterationWg.Add(1)
                go worker(state, iterationStopCh, &iterationWg)
            }

            isStable := false
            var currentAvgTPS float64
            testStart := time.Now()
            testTicker := time.NewTicker(time.Second)
            monitoringDone := make(chan bool)

            go func() {
                defer close(monitoringDone)
                for time.Since(testStart) < 30*time.Second { // Longer test for fine-tuning
                    select {
                    case <-testTicker.C:
                        state.recordMetrics()

                        state.mutex.RLock()
                        metrics := state.metrics[len(state.metrics)-1]
                        state.mutex.RUnlock()

                        currentTPS := metrics.TPS
                        log.Printf("Fine-tuning TPS: %.2f, Retries: %d, Users: %d",
                            currentTPS, metrics.InFlightRetries, state.currentUsers)

                        if metrics.InFlightRetries > 0 {
                            isStable = false
                            return
                        }

                        if currentTPS > 0 {
                            recentTPS = append(recentTPS, currentTPS)
                            if len(recentTPS) >= minSamplesNeeded {
                                if len(recentTPS) > minSamplesNeeded {
                                    recentTPS = recentTPS[1:]
                                }
                                
                                avg := 0.0
                                for _, tps := range recentTPS {
                                    avg += tps
                                }
                                avg /= float64(len(recentTPS))
                                currentAvgTPS = avg
                                
                                isStable = true
                                for _, tps := range recentTPS {
                                    if math.Abs(tps-avg)/avg > tpsVarianceThreshold {
                                        isStable = false
                                        break
                                    }
                                }
                            }
                        }
                    }
                }
            }()

            <-monitoringDone
            testTicker.Stop()
            close(iterationStopCh)
            cleanup := make(chan bool)
            go func() {
                iterationWg.Wait()
                cleanup <- true
            }()

            select {
            case <-cleanup:
            case <-time.After(10 * time.Second):
                log.Printf("Fine-tuning worker cleanup timed out")
            }

            time.Sleep(state.cooldownPeriod)

            if isStable {
                if currentAvgTPS > bestStableTPS {
                    bestStableTPS = currentAvgTPS
                    bestStableUsers = state.currentUsers
                    log.Printf("New best configuration in fine-tuning: %d users at %.2f TPS",
                        bestStableUsers, bestStableTPS)
                }
                lastStableUsers = candidateUsers
            } else {
                firstUnstableUsers = candidateUsers
            }
        }

        // Final 5-minute stability demonstration
        log.Printf("Running final 5-minute stability test with %d users", bestStableUsers)
        state.currentUsers = bestStableUsers
        
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

        finalTestStart := time.Now()
        finalTestTicker := time.NewTicker(time.Second)
        for time.Since(finalTestStart) < 5*time.Minute {
            select {
            case <-finalTestTicker.C:
                state.recordMetrics()
                state.mutex.RLock()
                metrics := state.metrics[len(state.metrics)-1]
                state.mutex.RUnlock()
                log.Printf("Final stability test - TPS: %.2f, Retries: %d",
                    metrics.TPS, metrics.InFlightRetries)
            }
        }
        
        finalTestTicker.Stop()
        close(iterationStopCh)
        iterationWg.Wait()
    }

    log.Printf("Benchmark complete. Final configuration: %d users, %.2f TPS", 
        bestStableUsers, bestStableTPS)
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
