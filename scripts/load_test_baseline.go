package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type Target struct {
	Name      string
	URL       string
	Method    string
	Payload   interface{}
	TargetRPS int
}

func main() {
	fmt.Println("⚡ Running Pranor Workspace Load Test Baselines...")

	targets := []Target{
		{
			Name:      "ServAuth - Key Validation",
			URL:       "http://localhost:8098/api/auth/keys/validate",
			Method:    "POST",
			Payload:   map[string]string{"key": "test-key"},
			TargetRPS: 500,
		},
		{
			Name:      "ServDB - Query Execution",
			URL:       "http://localhost:8097/api/db/query",
			Method:    "POST",
			Payload:   map[string]string{"query": "SELECT 1"},
			TargetRPS: 1000,
		},
		{
			Name:      "ServRegistry - Service Resolution",
			URL:       "http://localhost:8088/api/resolve?service=servdb",
			Method:    "GET",
			Payload:   nil,
			TargetRPS: 800,
		},
		{
			Name:      "ServFlow - Define Workflow",
			URL:       "http://localhost:8096/api/workflows/define",
			Method:    "POST",
			Payload: map[string]interface{}{
				"id": "load_test_flow",
				"tasks": []map[string]interface{}{
					{"name": "Step1", "action": "mock-success"},
				},
			},
			TargetRPS: 200,
		},
	}

	client := &http.Client{
		Timeout: 2 * time.Second,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
		},
	}

	for _, target := range targets {
		runLoadTest(client, target, 5*time.Second)
	}
}

func runLoadTest(client *http.Client, target Target, duration time.Duration) {
	fmt.Printf("\n--- Benchmarking %s ---\n", target.Name)
	fmt.Printf("Target SLA Throughput: %d RPS\n", target.TargetRPS)

	var successCount int64
	var errorCount int64
	var totalDurationMs int64

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	var wg sync.WaitGroup
	concurrency := 20
	ch := make(chan struct{}, concurrency)

	start := time.Now()

	for {
		select {
		case <-ctx.Done():
			goto Done
		default:
			ch <- struct{}{}
			wg.Add(1)
			go func() {
				defer func() {
					<-ch
					wg.Done()
				}()

				reqStart := time.Now()
				var req *http.Request
				var err error

				if target.Payload != nil {
					bodyBytes, _ := json.Marshal(target.Payload)
					req, err = http.NewRequest(target.Method, target.URL, bytes.NewReader(bodyBytes))
				} else {
					req, err = http.NewRequest(target.Method, target.URL, nil)
				}

				if err != nil {
					atomic.AddInt64(&errorCount, 1)
					return
				}
				req.Header.Set("Content-Type", "application/json")

				resp, err := client.Do(req)
				if err != nil {
					atomic.AddInt64(&errorCount, 1)
					return
				}
				defer resp.Body.Close()

				duration := time.Since(reqStart).Milliseconds()
				atomic.AddInt64(&totalDurationMs, duration)

				if resp.StatusCode < 400 {
					atomic.AddInt64(&successCount, 1)
				} else {
					atomic.AddInt64(&errorCount, 1)
				}
			}()
		}
	}

Done:
	wg.Wait()
	elapsed := time.Since(start).Seconds()

	totalRequests := successCount + errorCount
	if totalRequests == 0 {
		fmt.Println("No requests completed.")
		return
	}

	rps := float64(totalRequests) / elapsed
	avgLatency := float64(totalDurationMs) / float64(totalRequests)
	errorRate := (float64(errorCount) / float64(totalRequests)) * 100

	fmt.Printf("Completed %d requests in %.2fs\n", totalRequests, elapsed)
	fmt.Printf("Achieved Throughput: %.2f RPS\n", rps)
	fmt.Printf("Average Latency: %.2fms\n", avgLatency)
	fmt.Printf("Error Rate: %.2f%%\n", errorRate)

	// Since services might not be running in a standard pipeline run,
	// we document the targets and do a soft check.
	if rps >= float64(target.TargetRPS) && errorRate < 1.0 {
		fmt.Println("✅ SLA Target Met!")
	} else {
		fmt.Printf("⚠️ SLA Target not met under test concurrency (Expected >= %d RPS, got %.2f RPS)\n", target.TargetRPS, rps)
	}
}
