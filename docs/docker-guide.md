# ServVerse Docker Compose Guide

## Prerequisites

- **Podman Desktop** (with `podman compose` or `docker-compose` plugin) or **Docker Desktop**
- At least **8 GB RAM** allocated to the container runtime
- All component repos cloned as siblings to `pranor-repo/`:
  ```
  serv/
  ├── Pranor Auth/
  ├── Pranor Cache/
  ├── Pranor Deploy/
  ├── Pranor Console/
  ├── Pranor Chrono/
  ├── Pranor Pool/
  ├── Pranor Flow/
  ├── Pranor Gate/
  ├── Pranor Notify/
  ├── Pranor Mesh/
  ├── Pranor Pulse/
  ├── Pranor Hub/
  ├── Pranor Vault/
  ├── Pranor Trace/
  ├── Pranor Tunnel/
  └── pranor-repo/   ← you are here
  ```

---

## Running the Stack

### Build and start all services

```bash
cd pranor-repo
podman compose up --build
```

### Run in detached (background) mode

```bash
podman compose up --build -d
```

### Rebuild a single service after code changes

```bash
podman compose build servstore --no-cache
podman compose up -d servstore
```

### Stop everything

```bash
podman compose down
```

### Full clean restart (remove images + volumes)

```bash
podman compose down --rmi local --volumes
podman compose up --build
```

---

## Service Port Map

| # | Service | Port | Description |
|---|---------|------|-------------|
| 1 | Jaeger | 16686 | Trace UI |
| 2 | Pranor Trace | 8090 | OTLP/HTTP collector & trace API |
| 3 | Pranor Vault | 8081 | S3-compatible object storage |
| 4 | Pranor Pulse | 8082 (HTTP), 61613 (STOMP) | Message broker |
| 5 | Pranor Cache | 8086 | Distributed cache |
| 6 | Pranor Gate | 8080 | API gateway / reverse proxy |
| 7 | Pranor Mesh | 8089 | Service mesh registry |
| 8 | Pranor Chrono | 8087 | Distributed scheduler |
| 9 | Pranor Deploy | 8085 | Deployment orchestrator |
| 10 | Pranor Tunnel | 8443 | Tunnel relay server |
| 11 | Pranor Console | 8083 | Observability dashboard (Web UI) |
| 12 | Pranor Hub | 8088 | Package registry |

---

## Health Check (All Services)

After `podman compose up`, wait ~30 seconds then verify all containers are healthy:

```bash
podman compose ps
```

All services should show `healthy` status. Quick curl check:

```bash
curl http://localhost:8080/healthz   # Pranor Gate
curl http://localhost:8081/healthz   # Pranor Vault
curl http://localhost:8082/healthz   # Pranor Pulse
curl http://localhost:8083/healthz   # Pranor Console
curl http://localhost:8085/healthz   # Pranor Deploy
curl http://localhost:8086/healthz   # Pranor Cache
curl http://localhost:8087/healthz   # Pranor Chrono
curl http://localhost:8088/healthz   # Pranor Hub
curl http://localhost:8089/healthz   # Pranor Mesh
curl http://localhost:8090/healthz   # Pranor Trace
curl http://localhost:8443/healthz   # Pranor Tunnel
curl http://localhost:16686/         # Jaeger UI
```

---

## Testing Each Component

### 1. Jaeger (Trace UI)

Open http://localhost:16686 in a browser. After other services have processed requests, traces will appear searchable by service name.

---

### 2. Pranor Trace (OTLP Collector)

```bash
# Send a test trace span via OTLP/HTTP
curl -X POST http://localhost:8090/v1/traces \
  -H "Content-Type: application/json" \
  -d '{"resourceSpans":[{"resource":{"attributes":[{"key":"service.name","value":{"stringValue":"test"}}]},"scopeSpans":[{"spans":[{"traceId":"01020304050607080102030405060708","spanId":"0102030405060708","name":"test-span","startTimeUnixNano":"1700000000000000000","endTimeUnixNano":"1700000001000000000"}]}]}]}'

# Query stored traces
curl http://localhost:8090/api/traces
```

---

### 3. Pranor Vault (Object Storage)

```bash
# Create a bucket
curl -X PUT http://localhost:8081/test-bucket

# Upload an object
curl -X PUT http://localhost:8081/test-bucket/hello.json \
  -H "Content-Type: application/json" \
  -d '{"message": "hello from Pranor Vault"}'

# Download the object
curl http://localhost:8081/test-bucket/hello.json

# List buckets
curl http://localhost:8081/

# Delete the object
curl -X DELETE http://localhost:8081/test-bucket/hello.json
```

---

### 4. Pranor Pulse (Message Broker)

```bash
# Publish a message to a topic
curl -X POST http://localhost:8082/api/v1/publish \
  -H "Content-Type: application/json" \
  -d '{"topic": "orders", "payload": {"order_id": "12345", "amount": 99.99}}'

# List topics
curl http://localhost:8082/api/v1/topics

# Subscribe (poll) for messages
curl http://localhost:8082/api/v1/subscribe?topic=orders

# Check broker stats
curl http://localhost:8082/api/v1/stats
```

---

### 5. Pranor Cache (Distributed Cache)

```bash
# Set a cache entry
curl -X PUT http://localhost:8086/api/v1/cache/mykey \
  -H "Content-Type: application/json" \
  -d '{"value": "hello-world", "ttl": 60}'

# Get a cache entry
curl http://localhost:8086/api/v1/cache/mykey

# Delete a cache entry
curl -X DELETE http://localhost:8086/api/v1/cache/mykey

# Get cache stats
curl http://localhost:8086/api/v1/stats
```

---

### 6. Pranor Gate (API Gateway)

```bash
# Check gateway health
curl http://localhost:8080/healthz

# View current routes
curl http://localhost:8080/api/v1/admin/routes

# Test proxying (routes configured in config.json)
curl http://localhost:8080/api/v1/orders

# Gateway metrics
curl http://localhost:8080/api/v1/admin/metrics
```

---

### 7. Pranor Mesh (Service Mesh Registry)

```bash
# Register a service instance
curl -X POST http://localhost:8089/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{"service": "my-service", "address": "10.0.0.1:8080", "tags": ["v1"]}'

# Discover service instances
curl http://localhost:8089/api/v1/services/my-service

# List all registered services
curl http://localhost:8089/api/v1/services

# Deregister
curl -X DELETE http://localhost:8089/api/v1/deregister \
  -H "Content-Type: application/json" \
  -d '{"service": "my-service", "address": "10.0.0.1:8080"}'
```

---

### 8. Pranor Chrono (Distributed Scheduler)

```bash
# Schedule a job
curl -X POST http://localhost:8087/api/v1/jobs \
  -H "Content-Type: application/json" \
  -d '{"name": "cleanup", "schedule": "*/5 * * * *", "endpoint": "http://servstore:8081/healthz", "method": "GET"}'

# List all jobs
curl http://localhost:8087/api/v1/jobs

# Get job execution history
curl http://localhost:8087/api/v1/jobs/cleanup/history

# Delete a job
curl -X DELETE http://localhost:8087/api/v1/jobs/cleanup
```

---

### 9. Pranor Deploy (Deployment Orchestrator)

```bash
# Check available runtimes
curl http://localhost:8085/api/v1/runtimes

# Deploy a service (requires .pnr file or config)
curl -X POST http://localhost:8085/api/v1/deploy \
  -H "Content-Type: application/json" \
  -d '{"name": "my-app", "source": "main.pnr", "runtime": "go"}'

# List deployments
curl http://localhost:8085/api/v1/deployments

# Get deployment status
curl http://localhost:8085/api/v1/deployments/my-app
```

---

### 10. Pranor Tunnel (Tunnel Relay Server)

```bash
# Check relay server status
curl http://localhost:8443/healthz

# The tunnel relay accepts WebSocket connections at:
# ws://localhost:8443/ws/connect
# Use the servtunnel CLI client to establish a tunnel:
# servtunnel client 3000 --relay ws://localhost:8443/ws/connect --subdomain myapp
```

---

### 11. Pranor Console (Observability Dashboard)

Open http://localhost:8083 in a browser. The dashboard provides:

- Real-time service health monitoring
- Log aggregation viewer
- Trace visualization
- Gateway route management
- Cluster node overview
- Database query console

```bash
# API: Get service discovery info
curl http://localhost:8083/api/v1/discovery

# API: Get aggregated logs
curl http://localhost:8083/api/v1/logs

# API: Get system metrics
curl http://localhost:8083/api/v1/metrics
```

---

### 12. Pranor Hub (Package Registry)

```bash
# Publish a package (multipart form with tarball)
curl -X POST http://localhost:8088/api/v1/publish \
  -F "name=my-package" \
  -F "version=1.0.0" \
  -F "tarball=@my-package-1.0.0.tar.gz"

# Search packages
curl http://localhost:8088/api/packages/search?q=my-package

# Get package info
curl http://localhost:8088/api/v1/packages/my-package

# List all packages
curl http://localhost:8088/api/packages/

# Web dashboard
# Open http://localhost:8088 in browser
```

---

## End-to-End Integration Test

Run the existing e2e test suite (uses mock servers by default):

```bash
cd pranor-repo/tests/e2e
go test -v ./...
```

### Manual integration flow (against live stack)

```bash
# 1. Upload config to Pranor Vault (no auth in local dev mode)
curl -X PUT http://localhost:8081/demo-bucket/config.json \
  -H "Content-Type: application/json" \
  -d '{"app": "pranor-demo", "version": "1.0"}'

# 2. Publish event to Pranor Pulse (no auth in dev mode)
curl -X POST http://localhost:8082/api/v1/publish \
  -H "Content-Type: application/json" \
  -d '{"topic": "deployments", "payload": {"service": "demo", "action": "deploy"}}'

# 3. Cache the result (no auth)
curl -X PUT http://localhost:8086/api/v1/cache/last-deploy \
  -H "Content-Type: application/json" \
  -d '{"value": "demo-service-v1.0", "ttl": 300}'

# 4. Verify via Gateway
curl http://localhost:8080/healthz

# 5. Check traces in Jaeger
# Open http://localhost:16686, search for service "servstore" or "servqueue"

# 6. View everything in Pranor Console
# Open http://localhost:8083
```

---

## Authentication

All services use the standardized `Pranor Core.AuthMiddleware` for JWT authentication. The behavior is controlled by a single environment variable:

### How it works

- **`PRANOR_JWT_SECRET` not set (default in docker-compose)** → All requests pass through. No auth required.
- **`PRANOR_JWT_SECRET` set to any value** → All API routes require a valid `Authorization: Bearer <jwt>` header.
- **`/healthz` and `/readyz`** → Always accessible without auth regardless of configuration.

### Auth requirements per service (local dev mode)

| Service | Auth Required? |
|---------|---------------|
| All services | **No** — `PRANOR_JWT_SECRET` is unset in docker-compose |

### Enabling JWT auth (production)

Set the shared secret in docker-compose or as an environment variable:

```yaml
# docker-compose.yml — add to each service:
environment:
  - PRANOR_JWT_SECRET=your-strong-production-secret
```

Or run with an environment variable:

```bash
PRANOR_JWT_SECRET=my-secret podman compose up
```

### Generating a token

Use any JWT library to sign a token with HMAC-SHA256 and the shared secret:

```json
{
  "username": "admin",
  "roles": ["admin"],
  "iss": "pranor",
  "exp": 1750000000
}
```

Then use it in requests:

```bash
curl -H "Authorization: Bearer <your-jwt-token>" http://localhost:8082/api/v1/topics
```

---

## Production Release Images (GHCR)

All platform services are compiled, packaged, and published as production-ready container images on GitHub Container Registry (GHCR) whenever a version tag (`v*`) is released.

### Image Registry Paths

All component images are publicly available at:
`ghcr.io/vyuvaraj/<service-name>:v0.1.0` (and `latest`)

| Service | Registry Path |
|---|---|
| Pranor Gate | `ghcr.io/vyuvaraj/servgate:latest` |
| Pranor Vault | `ghcr.io/vyuvaraj/servstore:latest` |
| Pranor Pulse | `ghcr.io/vyuvaraj/servqueue:latest` |
| Pranor Cache | `ghcr.io/vyuvaraj/servcache:latest` |
| Pranor Console | `ghcr.io/vyuvaraj/servconsole:latest` |
| Pranor Chrono | `ghcr.io/vyuvaraj/servcron:latest` |
| Pranor Deploy | `ghcr.io/vyuvaraj/servcloud:latest` |
| Pranor Mesh | `ghcr.io/vyuvaraj/servmesh:latest` |
| Pranor Trace | `ghcr.io/vyuvaraj/servtrace:latest` |
| Pranor Tunnel | `ghcr.io/vyuvaraj/servtunnel:latest` |
| Pranor Hub | `ghcr.io/vyuvaraj/servregistry:latest` |

### Running the Pre-built Stack (No Source Code Needed)

You can run the entire platform stack using production images from GHCR without cloning all the individual component source repositories. 

Create a `docker-compose.prod.yml` file:

```yaml
version: '3.8'

services:
  jaeger:
    image: jaegertracing/all-in-one:latest
    ports: ["16686:16686", "4317:4317", "4318:4318"]

  servtrace:
    image: ghcr.io/vyuvaraj/servtrace:latest
    ports: ["8090:8090"]

  servstore:
    image: ghcr.io/vyuvaraj/servstore:latest
    ports: ["8081:8081"]
    command: ["--port", "8081", "--data-dir", "/data"]
    environment:
      - PRANOR_OTLP_ENDPOINT=http://servtrace:8090
    depends_on: [servtrace]

  servqueue:
    image: ghcr.io/vyuvaraj/servqueue:latest
    ports: ["8082:8082", "61613:61613"]
    environment:
      - PRANOR_OTLP_ENDPOINT=http://servtrace:8090
    depends_on: [servtrace]

  servcache:
    image: ghcr.io/vyuvaraj/servcache:latest
    ports: ["8086:8086"]
    environment:
      - PRANOR_OTLP_ENDPOINT=http://servtrace:8090
    depends_on: [servtrace]

  servgate:
    image: ghcr.io/vyuvaraj/servgate:latest
    ports: ["8080:8080"]
    environment:
      - PRANOR_OTLP_ENDPOINT=http://servtrace:8090
    depends_on: [servtrace]

  # Add other services as needed...
```

Then run:
```bash
podman compose -f docker-compose.prod.yml up -d
```

---

## Troubleshooting

### View logs for a specific service

```bash
podman compose logs servstore
podman compose logs -f servgate     # follow mode
```

### Restart a single service

```bash
podman compose restart servcache
```

### Service won't start — check dependencies

```bash
# See which services depend on what
podman compose config --services

# Check if a dependency is healthy
podman compose ps
```

### Port already in use

```bash
# Find what's using the port (Windows)
netstat -ano | findstr :8080

# Kill the process or change the port mapping in docker-compose.yml
```

### Build fails with Go version errors

All Dockerfiles patch `go 1.26.x` → `go 1.24` at build time via `sed`. If a new dependency adds a higher Go version constraint, re-vendor locally:

```bash
cd ../Pranor Vault   # or whichever service
set GOWORK=off
go mod vendor
# Then rebuild
podman compose build servstore --no-cache
```

---

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────┐
│                    Pranor Console :8083                         │
│                 (Observability Dashboard)                    │
└────────────┬───────────────┬────────────────────────────────┘
             │               │
     ┌───────▼───────┐ ┌────▼─────────┐
     │ Pranor Gate :8080│ │ Jaeger:16686 │
     │ (API Gateway) │ │ (Trace UI)   │
     └───────┬───────┘ └──────────────┘
             │
    ┌────────┼────────────────────────────┐
    │        │        │        │          │
┌───▼──┐ ┌──▼───┐ ┌──▼───┐ ┌──▼──┐ ┌────▼────┐
│Store │ │Queue │ │Cache │ │Cron │ │Registry │
│:8081 │ │:8082 │ │:8086 │ │:8087│ │  :8088  │
└──────┘ └──────┘ └──────┘ └─────┘ └─────────┘

    ┌────────┐  ┌────────┐  ┌────────┐  ┌───────┐
    │ Mesh   │  │ Cloud  │  │Tunnel  │  │ Trace │
    │ :8089  │  │ :8085  │  │ :8443  │  │ :8090 │
    └────────┘  └────────┘  └────────┘  └───────┘
```

All services communicate over the `pranor-net` Docker bridge network and export OpenTelemetry traces to Jaeger via the OTLP endpoint.
