package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"

	"gopkg.in/yaml.v3"
)

var version = "dev"

// Config represents the pranor.yaml structure
type Config struct {
	Shared   map[string]string        `yaml:"shared"`
	Services map[string]ServiceConfig `yaml:"services"`
}

type ServiceConfig struct {
	Port    int               `yaml:"port"`
	Env     map[string]string `yaml:"env"`
	Disable bool              `yaml:"disable"`
	Args    []string          `yaml:"args"`
}

type Service struct {
	Name    string
	Binary  string
	Port    int
	Env     map[string]string
	Args    []string
	process *exec.Cmd
}

var defaultServices = []Service{
	{Name: "PranorGate", Binary: "servgate", Port: 8080},
	{Name: "PranorVault", Binary: "servstore", Port: 8081},
	{Name: "PranorPulse", Binary: "servqueue", Port: 8082},
	{Name: "ServConsole", Binary: "servconsole", Port: 8083},
	{Name: "ServCache", Binary: "servcache", Port: 8084},
	{Name: "ServCron", Binary: "servcron", Port: 8085},
	{Name: "ServCloud", Binary: "servcloud", Port: 8086},
	{Name: "ServMesh", Binary: "servmesh", Port: 8087},
	{Name: "ServRegistry", Binary: "servregistry", Port: 8088},
	{Name: "ServTrace", Binary: "servtrace", Port: 8090},
	{Name: "ServMail", Binary: "servmail", Port: 8094},
	{Name: "ServFlow", Binary: "servflow", Port: 8096},
	{Name: "ServPool", Binary: "ServPool", Port: 8097},
	{Name: "ServAuth", Binary: "servauth", Port: 8098},
	{Name: "ServTunnel", Binary: "servtunnel", Port: 8443},
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "up":
		cmdUp()
	case "status":
		cmdStatus()
	case "init":
		cmdInit()
	case "version":
		fmt.Printf("pranor %s\n", version)
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("pranor — Unified launcher for the Serv ecosystem")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  up [--only name1,name2]   Start all (or selected) services")
	fmt.Println("  status                    Check health of all services")
	fmt.Println("  init                      Generate default pranor.yaml config")
	fmt.Println("  version                   Print version")
	fmt.Println()
	fmt.Println("Configuration:")
	fmt.Println("  Reads from: ./pranor.yaml → ~/.pranor/pranor.yaml → defaults")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  pranor init                          # Generate config file")
	fmt.Println("  pranor up                            # Start all services")
	fmt.Println("  pranor up --only servgate,servstore  # Start subset")
	fmt.Println("  pranor status                        # Health check")
}

// loadConfig finds and parses pranor.yaml
func loadConfig() Config {
	cfg := Config{
		Shared:   make(map[string]string),
		Services: make(map[string]ServiceConfig),
	}

	// Search order: ./pranor.yaml → ~/.pranor/pranor.yaml
	candidates := []string{
		"pranor.yaml",
		filepath.Join(homeDir(), ".pranor", "pranor.yaml"),
	}

	var configPath string
	for _, p := range candidates {
		if _, err := os.Stat(p); err == nil {
			configPath = p
			break
		}
	}

	if configPath == "" {
		// No config file — use defaults
		return cfg
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Warning: could not read %s: %v\n", configPath, err)
		return cfg
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		fmt.Printf("Warning: could not parse %s: %v\n", configPath, err)
		return cfg
	}

	fmt.Printf("Config loaded: %s\n", configPath)
	return cfg
}

// expandEnv replaces ${VAR} references with environment variable values
func expandEnv(val string) string {
	re := regexp.MustCompile(`\$\{([^}]+)\}`)
	return re.ReplaceAllStringFunc(val, func(match string) string {
		key := match[2 : len(match)-1]
		if v, ok := os.LookupEnv(key); ok {
			return v
		}
		return match // keep as-is if not found
	})
}

// buildServiceList merges defaults with config
func buildServiceList(cfg Config) []Service {
	services := make([]Service, 0, len(defaultServices))
	for _, svc := range defaultServices {
		key := strings.ToLower(svc.Binary)
		scfg, hasConfig := cfg.Services[key]

		if hasConfig && scfg.Disable {
			continue
		}

		// Apply config overrides
		if hasConfig && scfg.Port != 0 {
			svc.Port = scfg.Port
		}

		// Build env map: shared + per-service
		env := make(map[string]string)
		for k, v := range cfg.Shared {
			env[k] = expandEnv(v)
		}
		if hasConfig {
			for k, v := range scfg.Env {
				env[k] = expandEnv(v)
			}
		}
		svc.Env = env

		if hasConfig && len(scfg.Args) > 0 {
			svc.Args = scfg.Args
		}

		services = append(services, svc)
	}
	return services
}

func cmdUp() {
	cfg := loadConfig()
	services := buildServiceList(cfg)

	// Parse --only flag
	for i, arg := range os.Args[2:] {
		if arg == "--only" && i+1 < len(os.Args[2:]) {
			names := strings.Split(os.Args[2:][i+1], ",")
			services = filterServices(services, names)
			break
		}
	}

	binDir := getBinDir()
	var wg sync.WaitGroup
	var running []*exec.Cmd

	fmt.Printf("Starting %d services...\n\n", len(services))

	for i := range services {
		svc := &services[i]
		binaryPath := filepath.Join(binDir, svc.Binary+getExt())

		if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
			fmt.Printf("  %-14s SKIP (binary not found)\n", svc.Name)
			continue
		}

		args := append([]string{"--port", fmt.Sprintf("%d", svc.Port)}, svc.Args...)
		cmd := exec.Command(binaryPath, args...)

		// Build environment: inherit parent + shared + per-service
		cmd.Env = os.Environ()
		for k, v := range svc.Env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
		}
		// Always inject PORT for services that read it
		cmd.Env = append(cmd.Env, fmt.Sprintf("PORT=%d", svc.Port))

		cmd.Stdout = nil
		cmd.Stderr = nil

		if err := cmd.Start(); err != nil {
			fmt.Printf("  %-14s FAILED (%v)\n", svc.Name, err)
			continue
		}

		svc.process = cmd
		running = append(running, cmd)
		fmt.Printf("  %-14s started on :%d (PID %d)\n", svc.Name, svc.Port, cmd.Process.Pid)

		wg.Add(1)
		go func(c *exec.Cmd) {
			defer wg.Done()
			c.Wait()
		}(cmd)
	}

	fmt.Printf("\n%d services running. Press Ctrl+C to stop all.\n", len(running))

	if len(svc_env_summary(cfg)) > 0 {
		fmt.Printf("Shared env: %s\n", svc_env_summary(cfg))
	}

	// Wait for Ctrl+C
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	fmt.Println("\nShutting down...")
	for _, cmd := range running {
		if cmd.Process != nil {
			cmd.Process.Signal(syscall.SIGTERM)
		}
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("All services stopped.")
	case <-time.After(5 * time.Second):
		fmt.Println("Force killing remaining processes...")
		for _, cmd := range running {
			if cmd.Process != nil {
				cmd.Process.Kill()
			}
		}
	}
}

func svc_env_summary(cfg Config) string {
	keys := []string{}
	for k := range cfg.Shared {
		keys = append(keys, k)
	}
	if len(keys) > 4 {
		return fmt.Sprintf("%s + %d more", strings.Join(keys[:4], ", "), len(keys)-4)
	}
	return strings.Join(keys, ", ")
}

func cmdStatus() {
	cfg := loadConfig()
	services := buildServiceList(cfg)

	fmt.Printf("%-14s %-10s %-6s %s\n", "SERVICE", "STATUS", "PORT", "LATENCY")
	fmt.Println(strings.Repeat("─", 52))

	client := &http.Client{Timeout: 500 * time.Millisecond}
	passing, failing := 0, 0

	for _, svc := range services {
		url := fmt.Sprintf("http://localhost:%d/healthz", svc.Port)
		start := time.Now()
		resp, err := client.Get(url)
		latency := time.Since(start)

		status := "\033[31m● offline\033[0m"
		latStr := "-"
		if err == nil {
			resp.Body.Close()
			if resp.StatusCode == 200 {
				status = "\033[32m● healthy\033[0m"
				latStr = fmt.Sprintf("%dms", latency.Milliseconds())
				passing++
			} else {
				status = "\033[33m● degraded\033[0m"
				latStr = fmt.Sprintf("%dms", latency.Milliseconds())
				failing++
			}
		} else {
			failing++
		}

		fmt.Printf("%-14s %s  %-6d %s\n", svc.Name, status, svc.Port, latStr)
	}

	fmt.Println(strings.Repeat("─", 52))
	fmt.Printf("%d healthy, %d offline (%d total)\n", passing, failing, passing+failing)
}

func cmdInit() {
	configPath := "pranor.yaml"
	if _, err := os.Stat(configPath); err == nil {
		fmt.Printf("pranor.yaml already exists. Overwrite? (y/N): ")
		var answer string
		fmt.Scanln(&answer)
		if strings.ToLower(answer) != "y" {
			fmt.Println("Aborted.")
			return
		}
	}

	defaultConfig := `# Pranor Configuration
# Generated by: pranor init
# Docs: https://github.com/vyuvaraj/pranor/blob/main/docs/index.md

# Environment variables shared across ALL services
shared:
  PRANOR_JWT_SECRET: "pranor-local-dev-secret-CHANGE-IN-PRODUCTION"
  PRANOR_OTLP_ENDPOINT: "http://localhost:8090/v1/traces"

# Per-service configuration
services:
  servgate:
    port: 8080
    env:
      SERVGATE_AUTH_TOKEN: "gateway-secret-token"

  servstore:
    port: 8081
    env:
      SERVSTORE_DATA_DIR: "./data/store"
      SERVSTORE_ACCESS_KEY: "admin"
      SERVSTORE_SECRET_KEY: "password"

  servqueue:
    port: 8082

  servconsole:
    port: 8083
    env:
      PRANOR_JWT_SECRET: ""  # empty = no login required (dev mode). Set to shared secret for SSO.

  servcache:
    port: 8084
    env:
      SERVCACHE_REDIS_URL: ""  # empty = in-memory mode

  servcron:
    port: 8085
    env:
      SERVCRON_REDIS_URL: ""  # empty = standalone leader mode
      SERVSTORE_URL: "http://localhost:8081"

  servcloud:
    port: 8086
    env:
      SERVGATE_URL: "http://localhost:8080"

  servmesh:
    port: 8087

  servregistry:
    port: 8088
    env:
      SERVSTORE_URL: "http://localhost:8081"

  servtrace:
    port: 8090

  servmail:
    port: 8094
    env:
      SERVSTORE_URL: "http://localhost:8081"
      SERVMAIL_SMTP_HOST: ""  # empty = mock mode (no real emails sent)

  servflow:
    port: 8096
    env:
      SERVSTORE_URL: "http://localhost:8081"
      SERVQUEUE_URL: "http://localhost:8082"

  ServPool:
    port: 8097
    env:
      ServPool_PRIMARY_DSN: "sqlite://./data/ServPool.db"

  servauth:
    port: 8098
    env:
      SERVSTORE_URL: "http://localhost:8081"

  servtunnel:
    port: 8443
    env:
      SERVTUNNEL_DOMAIN: "localhost"

  # Disable a service by uncommenting:
  # servdocs:
  #   port: 8089
  #   disable: true
`

	if err := os.WriteFile(configPath, []byte(defaultConfig), 0644); err != nil {
		fmt.Printf("Error writing %s: %v\n", configPath, err)
		os.Exit(1)
	}

	// Create data directories
	os.MkdirAll("./data/store", 0755)
	os.MkdirAll("./data", 0755)

	fmt.Printf("Created %s with default local-dev configuration.\n", configPath)
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("  1. Edit pranor.yaml to set PRANOR_JWT_SECRET")
	fmt.Println("  2. Run: pranor up")
	fmt.Println()
	fmt.Println("Tip: Set GITHUB_TOKEN env var for pipeline dashboard API access.")
}

func filterServices(services []Service, names []string) []Service {
	nameSet := make(map[string]bool)
	for _, n := range names {
		nameSet[strings.ToLower(strings.TrimSpace(n))] = true
	}
	var result []Service
	for _, svc := range services {
		if nameSet[strings.ToLower(svc.Binary)] || nameSet[strings.ToLower(svc.Name)] {
			result = append(result, svc)
		}
	}
	return result
}

func getBinDir() string {
	exe, err := os.Executable()
	if err != nil {
		return "."
	}
	return filepath.Dir(exe)
}

func getExt() string {
	if os.PathSeparator == '\\' {
		return ".exe"
	}
	return ""
}

func homeDir() string {
	h, err := os.UserHomeDir()
	if err != nil {
		return "."
	}
	return h
}

// Suppress unused import
var _ = json.Marshal
