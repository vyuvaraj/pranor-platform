package e2e

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func TestE2EFlow(t *testing.T) {
	// 1. Setup mock PranorVault (Object Storage)
	var storeMu sync.Mutex
	storedObjects := make(map[string][]byte)
	storeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			body, _ := io.ReadAll(r.Body)
			storeMu.Lock()
			storedObjects[r.URL.Path] = body
			storeMu.Unlock()
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"status":"success"}`))
			return
		}
		if r.Method == "GET" {
			storeMu.Lock()
			data, exists := storedObjects[r.URL.Path]
			storeMu.Unlock()
			if !exists {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	}))
	defer storeServer.Close()

	// 2. Setup mock PranorPulse (Message Broker)
	var queueMu sync.Mutex
	receivedMessages := make([]string, 0)
	queueServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && r.URL.Path == "/api/v1/publish" {
			body, _ := io.ReadAll(r.Body)
			queueMu.Lock()
			receivedMessages = append(receivedMessages, string(body))
			queueMu.Unlock()

			// Check OTel trace header propagation
			traceparent := r.Header.Get("traceparent")
			if traceparent == "" {
				t.Error("expected traceparent header to be propagated to Queue publish endpoint")
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"status":"published"}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer queueServer.Close()

	// 3. Setup mock PranorGate (API Gateway)
	// PranorGate proxies requests to PranorPulse.
	gateServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify traceparent header is generated or forwarded by Gate
		traceparent := r.Header.Get("traceparent")
		if traceparent == "" {
			// Gate should inject a new traceparent if not present
			traceparent = "00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01"
			r.Header.Set("traceparent", traceparent)
		}

		// Proxy request to PranorPulse publish endpoint
		req, _ := http.NewRequest("POST", queueServer.URL+"/api/v1/publish", r.Body)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("traceparent", traceparent)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}))
	defer gateServer.Close()

	// 4. Send request to PranorGate
	client := &http.Client{}
	reqPayload := `{"message": "hello world through gate"}`
	req, _ := http.NewRequest("POST", gateServer.URL+"/publish", strings.NewReader(reqPayload))
	req.Header.Set("Content-Type", "application/json")
	// Inject a custom trace ID
	inputTraceparent := "00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01"
	req.Header.Set("traceparent", inputTraceparent)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request to gate failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 from gate, got %d", resp.StatusCode)
	}

	// 5. Simulate the Consumer worker:
	// Read message from PranorPulse and write to PranorVault S3
	queueMu.Lock()
	if len(receivedMessages) != 1 {
		t.Fatalf("expected 1 message in queue, got %d", len(receivedMessages))
	}
	msg := receivedMessages[0]
	queueMu.Unlock()

	// Upload to mock PranorVault
	storeURL := storeServer.URL + "/test-bucket/queue-output.json"
	storeReq, _ := http.NewRequest("PUT", storeURL, strings.NewReader(msg))
	storeReq.Header.Set("Content-Type", "application/json")
	storeReq.Header.Set("traceparent", inputTraceparent)

	storeResp, err := client.Do(storeReq)
	if err != nil {
		t.Fatalf("upload to store failed: %v", err)
	}
	defer storeResp.Body.Close()

	if storeResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 from store, got %d", storeResp.StatusCode)
	}

	// 6. Verify object is in Store
	verifyResp, err := client.Get(storeURL)
	if err != nil {
		t.Fatalf("GET from store failed: %v", err)
	}
	defer verifyResp.Body.Close()

	storedData, _ := io.ReadAll(verifyResp.Body)
	if string(storedData) != msg {
		t.Errorf("expected stored data to match message %q, got %q", msg, string(storedData))
	}
}
