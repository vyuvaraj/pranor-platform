package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	// Matches "/api/..." but not "/api/v1/..."
	nonV1APIRegex = regexp.MustCompile(`"/api/(?i)(?:[^v]|v[^1])[^"]*"`)
	// Matches error JSON structures or keys
	errorStructRegex = regexp.MustCompile(`(?i)(?:error|message)`)
	// Matches deprecation comment
	deprecatedCommentRegex = regexp.MustCompile(`(?i)@deprecated|deprecated`)
	// Matches deprecation headers
	deprecationHeaderRegex = regexp.MustCompile(`(?i)X-Deprecated|Deprecated|Sunset`)
)

type APIIssue struct {
	File        string
	Line        int
	Code        string
	Description string
}

func main() {
	repos := []string{
		"ServAuth",
		"ServCache",
		"ServCloud",
		"ServConsole",
		"ServDocs",
		"ServFlow",
		"PranorGate",
		"ServMail",
		"ServMesh",
		"ServPool",
		"PranorPulse",
		"ServRegistry",
		"ServShared",
		"PranorVault",
		"ServTrace",
		"ServTunnel",
	}

	var issues []APIIssue

	fmt.Println("==================================================")
	fmt.Println("Running API Consistency Linter...")
	fmt.Println("==================================================")

	for _, repo := range repos {
		repoPath := filepath.Join("..", repo)
		if _, err := os.Stat(repoPath); os.IsNotExist(err) {
			repoPath = filepath.Join(".", repo)
			if _, err := os.Stat(repoPath); os.IsNotExist(err) {
				continue
			}
		}

		_ = filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() || !strings.HasSuffix(info.Name(), ".go") {
				return nil
			}
			// Skip test files, vendors
			if strings.HasSuffix(info.Name(), "_test.go") || strings.Contains(path, "vendor") {
				return nil
			}

			contentBytes, err := os.ReadFile(path)
			if err != nil {
				return nil
			}
			content := string(contentBytes)
			lines := strings.Split(content, "\n")

			hasDeprecComment := deprecatedCommentRegex.MatchString(content)
			hasDeprecHeader := deprecationHeaderRegex.MatchString(content)

			if hasDeprecComment && !hasDeprecHeader {
				issues = append(issues, APIIssue{
					File:        path,
					Line:        1,
					Description: "File contains 'deprecated' comments but no deprecation headers (X-Deprecated or Deprecated)",
				})
			}

			for i, line := range lines {
				trimmed := strings.TrimSpace(line)

				// 1. Check for /api/v1 prefix violation
				if matches := nonV1APIRegex.FindAllString(trimmed, -1); len(matches) > 0 {
					// We'll warn on this rather than fail, or check specific ones
					issues = append(issues, APIIssue{
						File:        path,
						Line:        i + 1,
						Code:        trimmed,
						Description: fmt.Sprintf("API route prefix mismatch: found %s (expected /api/v1/...)", strings.Join(matches, ", ")),
					})
				}
			}

			return nil
		})
	}

	fmt.Printf("Scan complete. Found %d API consistency warnings/issues.\n\n", len(issues))

	// Print reports
	for _, issue := range issues {
		fmt.Printf("⚠️  %s:%d\n   Issue: %s\n", issue.File, issue.Line, issue.Description)
		if issue.Code != "" {
			fmt.Printf("   Code:  %s\n", issue.Code)
		}
		fmt.Println()
	}

	// Always succeed or warning only to ensure backward compatibility and gradual migration of older APIs,
	// but fail if critical deprecation headers are missing on explicitly deprecated files.
	criticalFailure := false
	for _, issue := range issues {
		if strings.Contains(issue.Description, "but no deprecation headers") {
			criticalFailure = true
		}
	}

	if criticalFailure {
		fmt.Println("❌ API consistency checks FAILED due to missing deprecation headers.")
		os.Exit(1)
	}

	fmt.Println("✅ API consistency linter executed successfully!")
	os.Exit(0)
}
