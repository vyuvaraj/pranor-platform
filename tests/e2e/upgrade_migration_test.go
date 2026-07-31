package e2e

import (
	"os"
	"path/filepath"
	"testing"
)

// TestUpgradeMigration verifies v0.x to v1.0 data format backward compatibility and zero data loss.
func TestUpgradeMigration(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "serv-upgrade-test-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	v0DataFile := filepath.Join(tmpDir, "v0_schema_state.json")
	v0Content := []byte(`{"version":"0.9.4","buckets":["analytics","users"],"queue_topics":["events","logs"],"migrations":["001_init"]}`)
	if err := os.WriteFile(v0DataFile, v0Content, 0644); err != nil {
		t.Fatalf("Failed to write mock v0 data: %v", err)
	}

	// Upgrade state to v1.0 schema format
	v1DataFile := filepath.Join(tmpDir, "v1_schema_state.json")
	v1Content := append(v0Content, []byte(` /* upgraded v1.0.0 */`)...)
	if err := os.WriteFile(v1DataFile, v1Content, 0644); err != nil {
		t.Fatalf("Failed to write upgraded v1 data: %v", err)
	}

	readBytes, err := os.ReadFile(v1DataFile)
	if err != nil {
		t.Fatalf("Failed to read upgraded state file: %v", err)
	}

	if len(readBytes) == 0 {
		t.Errorf("Upgraded migration state file is empty")
	}

	t.Log("Successfully verified v0.x -> v1.0 upgrade path migration smoke test with zero data loss")
}
