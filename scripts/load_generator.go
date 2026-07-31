package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

const (
	gatewayURL  = "http://localhost:8080"
	storeURL    = "http://localhost:8081"
	queueURL    = "http://localhost:8082"
	consoleURL  = "http://localhost:8083"
	cacheURL    = "http://localhost:8086"
	cronURL     = "http://localhost:8087"
	registryURL = "http://localhost:8088"
	authURL     = "http://localhost:8098"
	dbURL       = "http://localhost:8097"
	mailURL     = "http://localhost:8094"
	flowURL     = "http://localhost:8096"
	traceURL    = "http://localhost:8090"
)

func main() {
	fmt.Println("🚀 Starting Pranor Ecosystem Workload Generator...")
	fmt.Println("Press Ctrl+C to stop.")

	// Pre-setup steps (e.g. define a workflow, publish stdlib, etc.)
	setupEcosystem()

	// Tick every 15 seconds — aggressive enough to show live traffic
	// without flooding the terminal output.
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	client := &http.Client{Timeout: 3 * time.Second}

	for {
		select {
		case <-ticker.C:
			// Launch each activity with a small random jitter so they
			// don't all print at the same moment.
			go func() { time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond); simulateGatewayActivity(client) }()
			go func() { time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond); simulateCacheActivity(client) }()
			go func() { time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond); simulateQueueActivity(client) }()
			go func() { time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond); simulateDBActivity(client) }()
			go func() { time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond); simulateFlowActivity(client) }()
			go func() { time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond); simulateVectorSearchActivity(client) }()
		}
	}
}

func setupEcosystem() {
	client := &http.Client{Timeout: 5 * time.Second}

	// 1. Define Saga Workflow in PranorFlow
	workflow := map[string]interface{}{
		"id": "ecom_order_checkout",
		"tasks": []map[string]interface{}{
			{
				"name":   "validate_auth",
				"action": authURL + "/api/auth/keys/validate",
			},
			{
				"name":      "reserve_inventory",
				"depends_on": []string{"validate_auth"},
				"action":    dbURL + "/api/db/query",
			},
			{
				"name":      "dispatch_queue_notification",
				"depends_on": []string{"reserve_inventory"},
				"action":    queueURL + "/api/publish",
			},
			{
				"name":      "send_confirmation_email",
				"depends_on": []string{"dispatch_queue_notification"},
				"action":    mailURL + "/api/mail/send",
			},
		},
	}
	body, _ := json.Marshal(workflow)
	resp, err := client.Post(flowURL+"/api/workflows/define", "application/json", bytes.NewReader(body))
	if err == nil {
		resp.Body.Close()
		fmt.Println("  [SETUP] Defined saga workflow 'ecom_order_checkout' successfully.")
	} else {
		fmt.Printf("  [SETUP] Warning: could not define workflow in PranorFlow: %v\n", err)
	}

	// 2. Put a mock bucket in PranorVault
	resp, err = client.Post(storeURL+"/src-bucket", "application/json", nil)
	if err == nil {
		resp.Body.Close()
		fmt.Println("  [SETUP] Created bucket 'src-bucket' in PranorVault successfully.")
	}
}

func simulateGatewayActivity(client *http.Client) {
	// Call gateway status
	resp, err := client.Get(gatewayURL + "/healthz")
	if err == nil {
		resp.Body.Close()
		fmt.Println("  [GATEWAY] Checked healthz status: OK")
	}
}

func simulateCacheActivity(client *http.Client) {
	// Put cache key
	payload := map[string]string{
		"key":   fmt.Sprintf("user-session-%d", rand.Intn(100)),
		"value": fmt.Sprintf("active-%d", time.Now().Unix()),
	}
	body, _ := json.Marshal(payload)
	resp, err := client.Post(cacheURL+"/api/cache", "application/json", bytes.NewReader(body))
	if err == nil {
		resp.Body.Close()
		fmt.Println("  [CACHE] Set session key: Success")
	}
}

func simulateQueueActivity(client *http.Client) {
	// Publish message to queue
	payload := map[string]interface{}{
		"topic": "orders",
		"payload": map[string]interface{}{
			"order_id": fmt.Sprintf("ord-%d", rand.Intn(100000)),
			"amount":   rand.Float64() * 100.0,
			"status":   "pending",
		},
	}
	body, _ := json.Marshal(payload)
	resp, err := client.Post(queueURL+"/api/publish", "application/json", bytes.NewReader(body))
	if err == nil {
		resp.Body.Close()
		fmt.Println("  [QUEUE] Published order event: Success")
	}
}

func simulateDBActivity(client *http.Client) {
	// Query DB connections
	payload := map[string]string{
		"query": "SELECT * FROM inventory WHERE status = 'available' LIMIT 1;",
	}
	body, _ := json.Marshal(payload)
	resp, err := client.Post(dbURL+"/api/db/query", "application/json", bytes.NewReader(body))
	if err == nil {
		resp.Body.Close()
		fmt.Println("  [DB] Executed database validation check: Success")
	}
}

func simulateFlowActivity(client *http.Client) {
	// Execute workflow instance
	payload := map[string]interface{}{
		"workflow_id": "ecom_order_checkout",
		"input": map[string]interface{}{
			"user_id":  "usr-9982",
			"order_id": fmt.Sprintf("ord-%d", rand.Intn(100000)),
		},
	}
	body, _ := json.Marshal(payload)
	resp, err := client.Post(flowURL+"/api/workflows/execute", "application/json", bytes.NewReader(body))
	if err == nil {
		var res struct {
			InstanceID string `json:"instance_id"`
			Status     string `json:"status"`
		}
		if json.NewDecoder(resp.Body).Decode(&res) == nil {
			fmt.Printf("  [FLOW] Started saga run '%s': State = %s\n", res.InstanceID, res.Status)
		}
		resp.Body.Close()
	}
}

func simulateVectorSearchActivity(client *http.Client) {
	// Query vectors or perform semantic search mock
	payload := map[string]interface{}{
		"vector": []float32{rand.Float32(), rand.Float32(), rand.Float32()},
		"top_k":  3,
	}
	body, _ := json.Marshal(payload)
	// PranorVault serves vector endpoints
	resp, err := client.Post(storeURL+"/api/search", "application/json", bytes.NewReader(body))
	if err == nil {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		fmt.Println("  [VECTOR_STORE] Performed HNSW vector query check: Success")
	}
}
