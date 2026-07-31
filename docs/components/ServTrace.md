# Pranor Trace — Distributed Tracing Backend

> **Status:** ✅ Production | **Port:** 8090 | **Repository:** [Pranor Trace](https://github.com/vyuvaraj/pranor/tree/main/packages/Pranor Trace)

## Overview

Pranor Trace is a lightweight OTLP/HTTP distributed tracing backend. It reconstructs hierarchical trace trees, renders waterfall UI visualizations, performs span anomaly detection, archives cold-tier data to Pranor Vault, and supports configurable sampling policies.

## Key Features

- OTLP/HTTP span ingestion (OpenTelemetry native)
- Hierarchical trace tree reconstruction
- Waterfall UI for trace visualization
- Span duration anomaly detection
- Cold-tier archival to Pranor Vault
- Configurable sampling policies (head/tail)
- Trace search by service, operation, and duration
- Span tag indexing for fast lookups

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PRANOR_TRACE_PORT` | HTTP listen port | `8090` |
| `PRANOR_VAULT_URL` | Pranor Vault URL for cold storage | (required) |
| `PRANOR_OTLP_ENDPOINT` | OTel collector for self-tracing | (disabled) |

## Endpoints

| Endpoint | Description |
|----------|-------------|
| `GET /healthz` | Liveness probe |
| `POST /v1/traces` | OTLP trace ingestion endpoint |
| `GET /api/v1/traces` | Search and list traces |
| `GET /api/v1/traces/{id}` | Get full trace detail with spans |

## Pranor Integration

```srv
otel "my-service"
// automatic tracing — all routes and outbound calls are instrumented
```
