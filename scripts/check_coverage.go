package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var totalPercentRegex = regexp.MustCompile(`total:\s+\(statements\)\s+(\d+(?:\.\d+)?)%`)

func main() {
	// Baseline thresholds for existing repositories to prevent regression,
	// while new components are enforced at the strict 60% threshold.
	baselines := map[string]float64{
		"Pranor":    17.0,
		"PranorAuth":     0.0,
		"PranorCache":    40.0,
		"PranorDeploy":    15.0,
		"PranorConsole":  0.0,
		"ServDocs":     40.0,
		"PranorFlow":     24.0,
		"PranorGate":     20.0,
		"PranorNotify":     20.0,
		"PranorMesh":     10.0,
		"PranorPool":     40.0,
		"PranorPulse":    0.0,
		"PranorHub": 0.0,
		"PranorCore":   25.0,
		"PranorVault":    40.0,
		"PranorTrace":    19.0,
		"PranorTunnel":   2.0,
	}

	repos := []string{
		"Pranor",
		"PranorAuth",
		"PranorCache",
		"PranorDeploy",
		"PranorConsole",
		"ServDocs",
		"PranorFlow",
		"PranorGate",
		"PranorNotify",
		"PranorMesh",
		"PranorPool",
		"PranorPulse",
		"PranorHub",
		"PranorCore",
		"PranorVault",
		"PranorTrace",
		"PranorTunnel",
	}

	failed := false
	defaultThreshold := 60.0

	fmt.Println("==================================================")
	fmt.Println("Running Test Coverage Gates...")
	fmt.Println("==================================================")

	for _, repo := range repos {
		repoPath := filepath.Join("..", repo)
		if _, err := os.Stat(repoPath); os.IsNotExist(err) {
			repoPath = filepath.Join(".", repo)
			if _, err := os.Stat(repoPath); os.IsNotExist(err) {
				continue
			}
		}

		threshold := defaultThreshold
		if b, ok := baselines[repo]; ok {
			threshold = b
		}

		fmt.Printf("Checking coverage for: %s (required: %.1f%%)...\n", repo, threshold)
		covFile := filepath.Join(repoPath, "coverage.out")

		// Remove old coverage profile
		os.Remove(covFile)

		// Run tests with coverprofile
		cmd := exec.Command("go", "test", "-coverprofile=coverage.out", "./...")
		cmd.Dir = repoPath
		_ = cmd.Run()

		if _, err := os.Stat(covFile); os.IsNotExist(err) {
			if threshold > 0.0 {
				fmt.Printf("  ❌ FAIL: No tests/coverage generated for %s but required %.1f%%\n\n", repo, threshold)
				failed = true
			} else {
				fmt.Printf("  ✅ PASS: %s has no tests (threshold is 0%%)\n\n", repo)
			}
			continue
		}
		defer os.Remove(covFile)

		// Extract total statement coverage
		covCmd := exec.Command("go", "tool", "cover", "-func=coverage.out")
		covCmd.Dir = repoPath
		var out bytes.Buffer
		covCmd.Stdout = &out
		if err := covCmd.Run(); err != nil {
			fmt.Printf("  ❌ Failed to calculate coverage for %s: %v\n\n", repo, err)
			failed = true
			continue
		}

		// Parse the total percentage line
		lines := strings.Split(out.String(), "\n")
		var totalLine string
		for _, line := range lines {
			if strings.HasPrefix(line, "total:") {
				totalLine = line
				break
			}
		}

		if totalLine == "" {
			fmt.Printf("  ❌ Could not find total coverage percentage line for %s\n\n", repo)
			failed = true
			continue
		}

		matches := totalPercentRegex.FindStringSubmatch(totalLine)
		if len(matches) < 2 {
			fmt.Printf("  ❌ Could not parse coverage percentage: %q\n\n", totalLine)
			failed = true
			continue
		}

		percent, err := strconv.ParseFloat(matches[1], 64)
		if err != nil {
			fmt.Printf("  ❌ Failed to parse percentage value: %v\n\n", err)
			failed = true
			continue
		}

		if percent < threshold {
			fmt.Printf("  ❌ FAIL: %s coverage is %.1f%% (under the %.1f%% threshold)\n\n", repo, percent, threshold)
			failed = true
		} else {
			fmt.Printf("  ✅ PASS: %s coverage is %.1f%%\n\n", repo, percent)
		}
	}

	if failed {
		fmt.Println("❌ Test coverage check FAILED. Build blocked.")
		os.Exit(1)
	} else {
		fmt.Println("✅ All modules passed the minimum coverage check!")
		os.Exit(0)
	}
}
