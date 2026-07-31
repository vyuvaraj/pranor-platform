# Pranor Gate — Programmable API Gateway

> **Status:** ✅ Production | **Port:** 8080 | **Repository:** [Pranor Gate](https://github.com/vyuvaraj/pranor/tree/main/packages/Pranor Gate)

## Overview

Pranor Gate is a WASM-programmable reverse proxy and API gateway with AI-native capabilities. Single binary, zero external dependencies.

## Key Features

- Path-based reverse proxying with route prefix stripping
- WASM middleware hot-swap (zero-downtime deploys)
- WASM response filters and transform pipelines
- gRPC-Web gateway transpilation
- WebSocket proxying
- Load balancing (round-robin, least-connections)
- TLS termination with native HTTPS
- Sliding-window rate limiting
- Circuit breakers with automatic retries
- MCP (Model Context Protocol) native gateway
- Cost-aware LLM routing (cheap model → premium escalation)
- Semantic API caching (embedding-based)
- AI prompt guard and PII redaction
- OpenAPI auto-discovery from registered routes
- Developer portal (interactive API playground)
- Multi-tenant API key management
- Canary/blue-green traffic splitting
- JSON Schema request validation
- GraphQL federation proxy
- IP allowlisting/blocklisting
- Mutual TLS (mTLS) for backends
- Request queuing and backpressure

## Configuration

```json
{
  "addr": ":8080",
  "auth_token": "gateway-secret-token",
  "routes": [
    { "prefix": "/api/users", "target": "http://localhost:3000" },
    { "prefix": "/api/orders", "target": "http://localhost:3001", "rate_limit_rpm": 100 }
  ]
}
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PRANOR_JWT_SECRET` | JWT verification secret | (required) |
| `PRANOR_OTLP_ENDPOINT` | OTel collector URL | (disabled) |
| `PRANOR_GATE_CONFIG` | Path to config.json | `./config.json` |

## Endpoints

| Endpoint | Description |
|----------|-------------|
| `GET /healthz` | Liveness probe |
| `GET /readyz` | Readiness probe |
| `GET /api/v1/routes` | List registered routes |
| `POST /api/v1/routes/register` | Dynamic route registration |
| `POST /api/v1/middleware/upload` | Upload WASM middleware |
| `GET /api/docs` | Interactive API playground |

## Pranor Integration

```srv
server "8080"
// Routes auto-register with Pranor Gate on startup via /api/v1/routes/register
```

## Production & Operational Guidelines

### WASM Safety & Resource Limits
Pranor Gate runs isolated WebAssembly middleware filters via the wazero JIT compiler engine. The following safety ceilings are strictly enforced:
* **Execution Timeout:** Each WASM middleware execution context is bounded by a strict `50ms` timeout. Runaway loops or hung filters are automatically halted with a `504 Gateway Timeout` response returned to the client.
* **Memory Isolation:** Each WASM instance is sandbox-capped at a `16MB` memory ceiling. Any out-of-bounds allocation immediately traps, shielding the host Go gateway process from memory exhaustion or segmentation faults.

### Standard Observability & Telemetry
Pranor Gate integrates seamlessly with standard DevOps monitoring stacks:
* **Prometheus Metrics:** Native Prometheus metrics are exposed at `GET /metrics` covering request counts, latencies (p50/p90/p99), active circuit breaker statuses, and rate-limiting blocks.
* **OpenTelemetry Compatibility:** Complete W3C context propagation is supported. Pass `PRANOR_OTLP_ENDPOINT` to export traces directly to standard collectors (e.g. Datadog, Jaeger, Grafana Tempo).
