package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

func TestCrossServicesIntegration(t *testing.T) {
	// 1. Build binaries
	tempDir, err := os.MkdirTemp("", "serv-integration-")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	services := []struct {
		name string
		path string
		port string
	}{
		{"pranor-auth", "../../../PranorAuth", "18098"},
		{"servdb", "../../../ServDB", "18097"},
		{"pranor-notify", "../../../PranorNotify", "18094"},
		{"pranor-flow", "../../../PranorFlow", "18096"},
	}

	for _, s := range services {
		binName := s.name
		if filepath.Separator == '\\' {
			binName += ".exe"
		}
		binPath := filepath.Join(tempDir, binName)
		cmd := exec.Command("go", "build", "-o", binPath, s.path)
		// Turn off workspace mode if needed to avoid build complications, but go.work is correct
		if output, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("failed to build %s: %v\nOutput: %s", s.name, err, string(output))
		}
	}

	// 2. Start services in background
	cmds := make([]*exec.Cmd, 0)
	defer func() {
		// Clean up processes on exit
		for _, cmd := range cmds {
			if cmd.Process != nil {
				_ = cmd.Process.Kill()
			}
		}
	}()

	for _, s := range services {
		binName := s.name
		if filepath.Separator == '\\' {
			binName += ".exe"
		}
		binPath := filepath.Join(tempDir, binName)
		cmd := exec.Command(binPath, "-port", s.port)
		cmd.Env = append(os.Environ(), "PORT="+s.port)
		if err := cmd.Start(); err != nil {
			t.Fatalf("failed to start %s: %v", s.name, err)
		}
		cmds = append(cmds, cmd)
	}

	// 3. Wait for all services to be healthy
	client := &http.Client{Timeout: 1 * time.Second}
	for _, s := range services {
		healthy := false
		for attempt := 0; attempt < 10; attempt++ {
			resp, err := client.Get(fmt.Sprintf("http://localhost:%s/health", s.port))
			if err == nil && resp.StatusCode == http.StatusOK {
				resp.Body.Close()
				healthy = true
				break
			}
			time.Sleep(300 * time.Millisecond)
		}
		if !healthy {
			t.Fatalf("service %s failed to become healthy on port %s", s.name, s.port)
		}
	}

	// 4. Perform E2E User Lifecycle across services
	// Step 4a: Register User in PranorAuth
	regPayload := map[string]string{
		"username": "integration-test-user",
		"password": "strongPassword123!",
		"email":    "integration@example.com",
	}
	body, _ := json.Marshal(regPayload)
	resp, err := client.Post("http://localhost:18098/api/auth/register", "application/json", bytes.NewReader(body))
	if err != nil || resp.StatusCode != http.StatusCreated {
		t.Fatalf("failed to register user: %v, status: %d", err, resp.StatusCode)
	}
	resp.Body.Close()

	// Step 4b: Login User to get JWT
	loginPayload := map[string]string{
		"username": "integration-test-user",
		"password": "strongPassword123!",
	}
	body, _ = json.Marshal(loginPayload)
	resp, err = client.Post("http://localhost:18098/api/auth/login", "application/json", bytes.NewReader(body))
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("failed to login: %v", err)
	}
	var loginRes struct {
		Token string `json:"token"`
	}
	json.NewDecoder(resp.Body).Decode(&loginRes)
	resp.Body.Close()

	if loginRes.Token == "" {
		t.Fatalf("expected token in login response, got empty")
	}

	// Step 4c: Run a Migration in ServDB
	// Note: ServDB AuthMiddleware is active but we are in dev/local mode, so it passes.
	migrationPayload := map[string]interface{}{
		"version": 1,
		"name":    "create_users_table",
		"sql":     "CREATE TABLE users (id SERIAL PRIMARY KEY);",
	}
	body, _ = json.Marshal(migrationPayload)
	req, _ := http.NewRequest("POST", "http://localhost:18097/api/db/migrate", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+loginRes.Token)
	resp, err = client.Do(req)
	if err != nil || resp.StatusCode != http.StatusCreated {
		t.Fatalf("failed to execute migration in ServDB: %v, status: %d", err, resp.StatusCode)
	}
	resp.Body.Close()

	// Step 4d: Register Mail Template in PranorNotify
	templatePayload := map[string]string{
		"name":    "welcome",
		"version": "v1",
		"content": "Welcome {{.Name}}!",
	}
	body, _ = json.Marshal(templatePayload)
	req, _ = http.NewRequest("POST", "http://localhost:18094/api/mail/templates", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+loginRes.Token)
	resp, err = client.Do(req)
	if err != nil || resp.StatusCode != http.StatusCreated {
		t.Fatalf("failed to register template in PranorNotify: %v, status: %d", err, resp.StatusCode)
	}
	resp.Body.Close()

	// Step 4e: Define Workflow in PranorFlow
	workflowPayload := map[string]interface{}{
		"id": "onboarding-workflow",
		"tasks": []map[string]interface{}{
			{
				"name":   "VerifyMigration",
				"action": "success",
			},
		},
	}
	body, _ = json.Marshal(workflowPayload)
	req, _ = http.NewRequest("POST", "http://localhost:18096/api/workflows/define", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+loginRes.Token)
	resp, err = client.Do(req)
	if err != nil || resp.StatusCode != http.StatusCreated {
		t.Fatalf("failed to define workflow in PranorFlow: %v, status: %d", err, resp.StatusCode)
	}
	resp.Body.Close()
}
