package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestDockerComposeE2E(t *testing.T) {
	// Skip if DOCKER_COMPOSE_TEST is not set to prevent running it in environments without Docker
	if os.Getenv("DOCKER_COMPOSE_TEST") == "" {
		t.Skip("Skipping Docker Compose E2E test. Set DOCKER_COMPOSE_TEST=true to run.")
	}

	// 1. Build and start services using Docker Compose
	t.Log("Starting services with docker-compose...")
	cmd := exec.Command("docker-compose", "up", "-d", "--build")
	cmd.Dir = "../.." // pranor-repo directory
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to run docker-compose up: %v\nOutput: %s", err, string(output))
	}

	defer func() {
		t.Log("Stopping services with docker-compose...")
		downCmd := exec.Command("docker-compose", "down")
		downCmd.Dir = "../.."
		_ = downCmd.Run()
	}()

	// 2. Wait for services to be healthy
	services := []struct {
		name string
		port string
	}{
		{"pranor-auth", "8098"},
		{"servdb", "8097"},
		{"pranor-notify", "8094"},
		{"pranor-flow", "8096"},
	}

	client := &http.Client{Timeout: 1 * time.Second}
	for _, s := range services {
		healthy := false
		// Give it up to 45 seconds to spin up and compile in Docker
		for attempt := 0; attempt < 45; attempt++ {
			resp, err := client.Get(fmt.Sprintf("http://localhost:%s/healthz", s.port))
			if err == nil && resp.StatusCode == http.StatusOK {
				resp.Body.Close()
				healthy = true
				break
			}
			time.Sleep(1 * time.Second)
		}
		if !healthy {
			t.Fatalf("service %s failed to become healthy on port %s", s.name, s.port)
		}
	}

	t.Log("All services healthy. Executing E2E flow...")

	// 3. Perform contract assertions
	// Step A: Register User in PranorAuth
	regPayload := map[string]string{
		"username": "docker-test-user",
		"password": "strongPassword123!",
		"email":    "docker@example.com",
	}
	body, _ := json.Marshal(regPayload)
	resp, err := client.Post("http://localhost:8098/api/auth/register", "application/json", bytes.NewReader(body))
	if err != nil || resp.StatusCode != http.StatusCreated {
		t.Fatalf("failed to register user: %v, status: %d", err, resp.StatusCode)
	}
	resp.Body.Close()

	// Step B: Login User to get JWT
	loginPayload := map[string]string{
		"username": "docker-test-user",
		"password": "strongPassword123!",
	}
	body, _ = json.Marshal(loginPayload)
	resp, err = client.Post("http://localhost:8098/api/auth/login", "application/json", bytes.NewReader(body))
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

	// Step C: Run a Migration in ServDB
	migrationPayload := map[string]interface{}{
		"version": 1,
		"name":    "create_users_table",
		"sql":     "CREATE TABLE users (id SERIAL PRIMARY KEY);",
	}
	body, _ = json.Marshal(migrationPayload)
	req, _ := http.NewRequest("POST", "http://localhost:8097/api/db/migrate", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+loginRes.Token)
	resp, err = client.Do(req)
	if err != nil || resp.StatusCode != http.StatusCreated {
		t.Fatalf("failed to execute migration in ServDB: %v, status: %d", err, resp.StatusCode)
	}
	resp.Body.Close()

	// Step D: Register Mail Template in PranorNotify
	templatePayload := map[string]string{
		"name":    "welcome",
		"version": "v1",
		"content": "Welcome {{.Name}}!",
	}
	body, _ = json.Marshal(templatePayload)
	req, _ = http.NewRequest("POST", "http://localhost:8094/api/mail/templates", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+loginRes.Token)
	resp, err = client.Do(req)
	if err != nil || resp.StatusCode != http.StatusCreated {
		t.Fatalf("failed to register template in PranorNotify: %v, status: %d", err, resp.StatusCode)
	}
	resp.Body.Close()

	// Step E: Define Workflow in PranorFlow
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
	req, _ = http.NewRequest("POST", "http://localhost:8096/api/workflows/define", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+loginRes.Token)
	resp, err = client.Do(req)
	if err != nil || resp.StatusCode != http.StatusCreated {
		t.Fatalf("failed to define workflow in PranorFlow: %v, status: %d", err, resp.StatusCode)
	}
	resp.Body.Close()

	t.Log("E2E Docker Compose Contract Flow completed successfully!")
}
