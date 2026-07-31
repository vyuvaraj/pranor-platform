# Troubleshooting Guide

Common issues and solutions when building, running, and deploying Pranor modules.

## Issue 1: Compiler Mismatched Signatures
**Symptom**: `go test ./...` or `pranor build` fails with:
`undefined: SomeFunction` or `mismatched argument types`.

### Resolution
1. Verify the `go.work` file is configured at the workspace root to include all local subpackages.
2. Clean compiler build artifacts:
   ```bash
   go clean -cache -testcache
   ```
3. Regenerate packages list:
   ```bash
   serv packages --update
   ```

## Issue 2: Service Registry (Pranor Mesh) Offline / Heartbeat Dropped
**Symptom**: Backend services fail to register, printing:
`[ERROR] failed to connect to mesh registry on http://localhost:8089`.

### Resolution
1. Verify that `Pranor Mesh` is running:
   ```bash
   serv status
   ```
2. Check firewall or Docker network bindings. The default mesh discovery port requires UDP `:9999` to be open for multicast discovery.
3. If multicast fails in your hosting environment, bypass discovery and explicitly declare the registry target:
   ```bash
   export PRANOR_MESH_ADDR="http://127.0.0.1:8089"
   ```
