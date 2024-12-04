// main.go
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type WrkMetrics struct {
	RequestsPerSecond float64
	LatencyAvg        float64
	LatencyP50        float64
	LatencyP99        float64
}

type SystemMetrics struct {
	CPUUsage     float64
	MemoryUsage  float64
	DiskUsage    float64
	DiskReadOps  float64
	DiskWriteOps float64
}

type BenchmarkReport struct {
	// Recommended values
	RecommendedTPS   float64
	RecommendedUsers int
	MaxCPU          float64
	MaxMemory       float64
	MaxDiskIO       float64

	// Time series data for plotting
	Timestamps      []string   `json:"timestamps"`
	TPSValues       []float64  `json:"tpsValues"`
	CPUValues       []float64  `json:"cpuValues"`
	MemoryValues    []float64  `json:"memoryValues"`
	DiskReadValues  []float64  `json:"diskReadValues"`
	DiskWriteValues []float64  `json:"diskWriteValues"`
	RetryDebtValues []float64  `json:"retryDebtValues"`
}

const reportTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Restate Benchmark Report</title>
    <script src="https://cdn.plot.ly/plotly-latest.min.js"></script>
    <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css" rel="stylesheet">
</head>
<body class="bg-gray-50">
    <div class="container mx-auto px-4 py-8">
        <h1 class="text-3xl font-bold mb-8">Restate Benchmark Results</h1>
        
        <div class="bg-yellow-100 border-l-4 border-yellow-500 p-4 mb-8">
            <p class="font-bold">Recommended Configuration</p>
            <p>Maximum Stable TPS: {{printf "%.2f" .RecommendedTPS}}</p>
            <p>Recommended Concurrent Users: {{.RecommendedUsers}}</p>
            <p>Hardware Requirements:</p>
            <ul class="list-disc ml-5">
                <li>CPU: {{printf "%.1f%%" .MaxCPU}}% utilization</li>
                <li>Memory: {{printf "%.1f" .MaxMemory}}GB</li>
                <li>Disk I/O: {{.MaxDiskIO}} operations/sec</li>
            </ul>
        </div>
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
            <div class="bg-white p-6 rounded-lg shadow">
                <h2 class="text-xl font-semibold mb-4">Performance Timeline</h2>
                <div id="tpsChart" class="h-80"></div>
            </div>
            
            <div class="bg-white p-6 rounded-lg shadow">
                <h2 class="text-xl font-semibold mb-4">System CPU & Memory</h2>
                <div id="systemChart" class="h-80"></div>
            </div>
            <div class="bg-white p-6 rounded-lg shadow">
                <h2 class="text-xl font-semibold mb-4">Disk Operations</h2>
                <div id="diskChart" class="h-80"></div>
            </div>
            <div class="bg-white p-6 rounded-lg shadow">
                <h2 class="text-xl font-semibold mb-4">Retry Debt</h2>
                <div id="retryChart" class="h-80"></div>
            </div>
        </div>
    </div>
    <script>
        // TPS Timeline with Retry Debt
        const tpsData = [{
            x: {{.Timestamps}},
            y: {{.TPSValues}},
            name: 'TPS',
            type: 'scatter'
        }];
        
        Plotly.newPlot('tpsChart', tpsData, {
            title: 'Transactions Per Second',
            yaxis: { title: 'TPS' }
        });
        // System Metrics
        const systemData = [{
            x: {{.Timestamps}},
            y: {{.CPUValues}},
            name: 'CPU %',
            type: 'scatter'
        }, {
            x: {{.Timestamps}},
            y: {{.MemoryValues}},
            name: 'Memory %',
            type: 'scatter'
        }];
        Plotly.newPlot('systemChart', systemData, {
            title: 'System Resource Usage',
            yaxis: { title: '%' }
        });
        // Disk Operations
        const diskData = [{
            x: {{.Timestamps}},
            y: {{.DiskReadValues}},
            name: 'Reads/s',
            type: 'scatter'
        }, {
            x: {{.Timestamps}},
            y: {{.DiskWriteValues}},
            name: 'Writes/s',
            type: 'scatter'
        }];
        Plotly.newPlot('diskChart', diskData, {
            title: 'Disk Operations',
            yaxis: { title: 'Operations/s' }
        });
        // Retry Debt
        const retryData = [{
            x: {{.Timestamps}},
            y: {{.RetryDebtValues}},
            name: 'Retry Debt',
            type: 'scatter'
        }];
        Plotly.newPlot('retryChart', retryData, {
            title: 'Retry Debt Over Time',
            yaxis: { title: 'Outstanding Retries' }
        });
    </script>
</body>
</html>
`

func collectSystemMetrics(path string) SystemMetrics {
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		cpuPercent = []float64{0.0}
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		memInfo = &mem.VirtualMemoryStat{}
	}

	var diskUsageGB float64
	diskUsage, err := disk.Usage(path)
	if err != nil {
		log.Printf("Warning: Failed to get disk usage: %v", err)
	} else {
		diskUsageGB = float64(diskUsage.Used) / (1024 * 1024 * 1024)
	}

	var readOps, writeOps float64
	diskStats, err := disk.IOCounters()
	if err == nil {
		for _, stat := range diskStats {
			readOps += float64(stat.ReadCount)
			writeOps += float64(stat.WriteCount)
		}
	}

	return SystemMetrics{
		CPUUsage:     cpuPercent[0],
		MemoryUsage:  memInfo.UsedPercent,
		DiskUsage:    diskUsageGB,
		DiskReadOps:  readOps,
		DiskWriteOps: writeOps,
	}
}

func parseWrkOutput(output []byte) (WrkMetrics, error) {
	var metrics WrkMetrics
	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Requests/sec:") {
			fmt.Sscanf(line, "Requests/sec: %f", &metrics.RequestsPerSecond)
		} else if strings.Contains(line, "Latency") {
			fields := strings.Fields(line)
			if len(fields) >= 4 {
				fmt.Sscanf(fields[1], "%fms", &metrics.LatencyAvg)
				fmt.Sscanf(fields[2], "50%%:%fms", &metrics.LatencyP50)
				fmt.Sscanf(fields[3], "99%%:%fms", &metrics.LatencyP99)
			}
		}
	}
	return metrics, scanner.Err()
}

func runWrkInterval(url string, duration time.Duration, threads, connections int) (WrkMetrics, error) {
	cmd := exec.Command("wrk",
		fmt.Sprintf("-d%ds", int(duration.Seconds())),
		fmt.Sprintf("-t%d", threads),
		fmt.Sprintf("-c%d", connections),
		"-s", "post.lua",
		url,
	)
	
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		return WrkMetrics{}, fmt.Errorf("wrk error: %v", err)
	}
	
	return parseWrkOutput(output.Bytes())
}

func generateReport(report BenchmarkReport) error {
	tmpl, err := template.New("report").Parse(reportTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	file, err := os.Create("benchmark_report.html")
	if err != nil {
		return fmt.Errorf("error creating report file: %v", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, report); err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	log.Println("Report generated: benchmark_report.html")
	return nil
}

func main() {
	const (
		totalDuration    = 5 * time.Minute
		intervalDuration = 15 * time.Second
		url             = "http://localhost:8080/Greeter/greet"
		threads         = 16
		connections     = 400
	)

	var (
		tpsValues       []float64
		cpuValues       []float64
		memoryValues    []float64
		diskReadValues  []float64
		diskWriteValues []float64
		retryDebtValues []float64
		timestamps      []string
		maxTPS         float64
	)

	restateDataPath, _ := filepath.Abs(".")
	intervalCount := int(totalDuration / intervalDuration)
	
	// Channel for system metrics collection
	metricsChan := make(chan SystemMetrics)
	done := make(chan bool)

	// Start system metrics collection
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				metricsChan <- collectSystemMetrics(restateDataPath)
			}
		}
	}()

	log.Printf("Starting benchmark: %d intervals of %v each\n", intervalCount, intervalDuration)
	
	// Run intervals
	for i := 0; i < intervalCount; i++ {
		log.Printf("Running interval %d/%d\n", i+1, intervalCount)
		
		// Run wrk for this interval
		wrkMetrics, err := runWrkInterval(url, intervalDuration, threads, connections)
		if err != nil {
			log.Printf("Error in interval %d: %v", i+1, err)
			continue
		}

		// Collect system metrics during the interval
		var intervalCPU, intervalMemory, intervalDiskRead, intervalDiskWrite float64
		var metricsCount int
		
		// Collect metrics for the duration of the interval
		timeout := time.After(intervalDuration)
	metricLoop:
		for {
			select {
			case <-timeout:
				break metricLoop
			case metrics := <-metricsChan:
				intervalCPU += metrics.CPUUsage
				intervalMemory += metrics.MemoryUsage
				intervalDiskRead += metrics.DiskReadOps
				intervalDiskWrite += metrics.DiskWriteOps
				metricsCount++
			}
		}

		// Average the metrics for this interval
		if metricsCount > 0 {
			intervalCPU /= float64(metricsCount)
			intervalMemory /= float64(metricsCount)
			intervalDiskRead /= float64(metricsCount)
			intervalDiskWrite /= float64(metricsCount)
		}

		// Append data points
		timestamps = append(timestamps, time.Now().Format(time.RFC3339))
		tpsValues = append(tpsValues, wrkMetrics.RequestsPerSecond)
		cpuValues = append(cpuValues, intervalCPU)
		memoryValues = append(memoryValues, intervalMemory)
		diskReadValues = append(diskReadValues, intervalDiskRead)
		diskWriteValues = append(diskWriteValues, intervalDiskWrite)
		retryDebtValues = append(retryDebtValues, 0) // Placeholder

		// Track maximum TPS
		if wrkMetrics.RequestsPerSecond > maxTPS {
			maxTPS = wrkMetrics.RequestsPerSecond
		}
	}

	// Stop metrics collection
	done <- true

	// Calculate maximums for recommendations
	maxCPU := 0.0
	maxMemory := 0.0
	maxDiskIO := 0.0
	for i := range cpuValues {
		if cpuValues[i] > maxCPU {
			maxCPU = cpuValues[i]
		}
		if memoryValues[i] > maxMemory {
			maxMemory = memoryValues[i]
		}
		diskIO := diskReadValues[i] + diskWriteValues[i]
		if diskIO > maxDiskIO {
			maxDiskIO = diskIO
		}
	}

	// Generate final report
	report := BenchmarkReport{
		RecommendedTPS:    maxTPS * 0.8, // 80% of max as recommended
		RecommendedUsers:  int(float64(connections) * 0.8), // 80% of max connections
		MaxCPU:           maxCPU,
		MaxMemory:        maxMemory,
		MaxDiskIO:        maxDiskIO,
		Timestamps:       timestamps,
		TPSValues:        tpsValues,
		CPUValues:        cpuValues,
		MemoryValues:     memoryValues,
		DiskReadValues:   diskReadValues,
		DiskWriteValues:  diskWriteValues,
		RetryDebtValues:  retryDebtValues,
	}

	if err := generateReport(report); err != nil {
		log.Fatal("Error generating report:", err)
	}

	log.Println("Benchmark completed successfully")
}