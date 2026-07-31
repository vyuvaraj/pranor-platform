# Pranor Architecture

## Layers

```
┌─────────────────────────────────────────────────────────────────────┐
│                       DEVELOPER TOOLS                               │
│  Pranor Compiler │ VS Code LSP │ ServDocs │ Pranor Hub         │
├─────────────────────────────────────────────────────────────────────┤
│                       PLATFORM LAYER                                │
│  Pranor Gate │ Pranor Mesh │ Pranor Deploy │ Pranor Tunnel │ Pranor Console          │
├─────────────────────────────────────────────────────────────────────┤
│                     INFRASTRUCTURE LAYER                            │
│  Pranor Vault │ Pranor Pulse │ Pranor Cache │ Pranor Pool │ Pranor Auth               │
│  Pranor Notify  │ Pranor Chrono  │ Pranor Flow                                   │
├─────────────────────────────────────────────────────────────────────┤
│                     FOUNDATION                                      │
│  Pranor Core (common library — health, OTel, JWT, logging)           │
│  Pranor Trace (distributed tracing backend)                            │
└─────────────────────────────────────────────────────────────────────┘
```

## Runtime Dependency Flow

```mermaid
graph TD
    Client[External Client] -->|HTTPS| Pranor Gate
    Pranor Gate -->|route + proxy| Pranor Mesh
    Pranor Mesh -->|resolve + LB| Services[Service Instances]
    
    Services -->|persist| Pranor Vault
    Services -->|enqueue| Pranor Pulse
    Services -->|cache| Pranor Cache
    Services -->|query| Pranor Pool
    Services -->|authenticate| Pranor Auth
    Services -->|notify| Pranor Notify
    Services -->|schedule| Pranor Chrono
    Services -->|workflow| Pranor Flow
    
    Services -.->|traces| Pranor Trace
    Pranor Trace -->|cold tier| Pranor Vault
    Pranor Chrono -->|triggers| Services
    Pranor Pulse -->|delivers| Services
    Pranor Notify -->|DLQ retry| Pranor Pulse
    Pranor Flow -->|events| Pranor Pulse
    Pranor Flow -->|checkpoints| Pranor Vault
    Pranor Auth -->|users| Pranor Vault
    
    Pranor Console -->|aggregates| Pranor Gate
    Pranor Console -->|aggregates| Pranor Vault
    Pranor Console -->|aggregates| Pranor Pulse
    Pranor Console -->|aggregates| Pranor Trace
    Pranor Console -->|aggregates| Pranor Auth
    Pranor Console -->|aggregates| Pranor Pool
    Pranor Console -->|aggregates| Pranor Notify
    Pranor Console -->|aggregates| Pranor Flow
    Pranor Console -->|aggregates| Pranor Tunnel
    
    Pranor Deploy -->|deploys| Services
    Pranor Deploy -->|registers routes| Pranor Gate
    Pranor Hub -->|stores packages| Pranor Vault
```

## Service Discovery

All services locate each other via the `PRANOR_DISCOVERY` environment variable — a JSON manifest (or file path) mapping service names to URLs:

```json
{
  "gate": "http://localhost:8080",
  "store": "http://localhost:8081",
  "queue": "http://localhost:8082",
  "console_port": 8083,
  "cache": "http://localhost:8084",
  "cron": "http://localhost:8085",
  "cloud": "http://localhost:8086",
  "mesh": "http://localhost:8087",
  "registry": "http://localhost:8088",
  "docs": "http://localhost:8089",
  "trace": "http://localhost:8090",
  "mail": "http://localhost:8094",
  "flow": "http://localhost:8096",
  "db": "http://localhost:8097",
  "auth": "http://localhost:8098",
  "tunnel": "http://localhost:8443",
  "otlp_endpoint": "http://localhost:8090/v1/traces",
  "jwt_secret": "shared-secret"
}
```

## Shared Conventions

All services follow these patterns (enforced by Pranor Core):

| Convention | Implementation |
|------------|----------------|
| Health probe | `GET /healthz` → 200 OK |
| Readiness probe | `GET /readyz` → 200 OK |
| Error format | `{"error": "msg", "code": "ERR_CODE", "trace_id": "..."}` |
| Auth | Bearer JWT verified via `PRANOR_JWT_SECRET` |
| Tracing | OTel spans exported to `PRANOR_OTLP_ENDPOINT` |
| Logging | Structured JSON to stdout |
| Shutdown | Graceful on SIGTERM (drain + 5s timeout) |
| API versioning | `/api/v1/` prefix on all management endpoints |

## Communication Patterns

| Pattern | Used By |
|---------|---------|
| HTTP REST (sync) | All services for API calls |
| STOMP TCP (async) | Pranor Pulse for pub/sub messaging |
| WebSocket (push) | Pranor Console for real-time dashboards, Pranor Tunnel for tunneling |
| `serv://` resolver | Pranor Mesh for inter-service calls |
| S3 protocol | Pranor Vault for object storage |
| OTLP/HTTP | Pranor Trace for span ingestion |
