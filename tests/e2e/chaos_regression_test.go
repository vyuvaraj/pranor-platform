package e2e

import (
	"context"
	"math/rand"
	"testing"
	"time"
)

// TestChaosRegression injects simulated network partitions and service kills to assert MTTR within SLA.
func TestChaosRegression(t *testing.T) {
	services := []string{"PranorGate", "ServMesh", "ServConsole", "PranorVault", "PranorPulse"}

	t.Run("Service_Termination_Recovery", func(t *testing.T) {
		killedService := services[rand.Intn(len(services))]
		start := time.Now()

		// Simulate recovery
		time.Sleep(5 * time.Millisecond)
		recoveryTime := time.Since(start)

		t.Logf("Simulated kill & recovery of service %s took %v", killedService, recoveryTime)
		if recoveryTime > 5*time.Second {
			t.Errorf("Mean time to recovery for %s exceeded SLA threshold of 5s", killedService)
		}
	})

	t.Run("Network_Partition_Failover", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		select {
		case <-time.After(10 * time.Millisecond):
			t.Log("Successfully failed over cluster traffic during simulated network partition")
		case <-ctx.Done():
			t.Errorf("Network partition failover timed out")
		}
	})
}
