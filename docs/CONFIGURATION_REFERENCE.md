# Configuration Reference

A complete index of environment variables, default ports, network parameters, and files used across all services in the Pranor ecosystem.

## Global Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | *(Service specific)* | Primary port the HTTP server binds to. |
| `PRANOR_JWT_SECRET` | `""` | Secret key used to sign and verify user JWTs. |
| `PRANOR_JWKS_URL` | `""` | JWKS endpoint path to fetch RSA validation keys. |
| `PRANOR_MESH_ADDR` | `http://localhost:8089` | Target URL of the active Pranor Mesh registry. |
| `PRANOR_OTLP_ENDPOINT` | `""` | OpenTelemetry collector pipeline receiver endpoint. |
| `LOG_LEVEL` | `info` | Filter log logs: `debug`, `info`, `warn`, `error`. |

---

## Service Configuration Details

### 1. Pranor Gate (API Gateway)
* **Default Port**: `8080`
* **Configuration File**: `config.json`
* **Key settings**:
  - `max_concurrent_requests`: Rate limit gate concurrency.
  - `client_cert_path` / `root_ca_path`: TLS cert files for backend mTLS calls.

### 2. Pranor Vault (Object Storage)
* **Default Port**: `8081`
* **Environment variables**:
  - `DATA_DIR`: Path to bucket storage contents on local disk (default: `./data`).

### 3. Pranor Pulse (Message Broker)
* **Default Port**: `8082`
* **Environment variables**:
  - `PERSISTENCE_ENABLED`: Write messages to journal files before dispatch (default: `true`).

### 4. Pranor Mesh (Service Mesh)
* **Default Port**: `8089` (HTTP), `9999` (UDP discovery)
* **Key settings**:
  - `registry.go` evicts nodes after default 10s timeout intervals if heartbeats fail.
