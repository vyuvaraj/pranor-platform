# Pranor Console — Unified Observability Dashboard

> **Status:** ✅ Production | **Port:** 8083 | **Repository:** [Pranor Console](https://github.com/vyuvaraj/pranor/tree/main/packages/Pranor Console)

## Overview

Single-binary glassmorphic web dashboard providing unified observability, administration, and operations for all Pranor components. Reverse-proxies to connected services, aggregates health, and provides real-time metrics via WebSocket push.

## Key Features

- Reverse proxy to 9 services (Gate, Store, Queue, Trace, Tunnel, Auth, DB, Mail, Flow)
- Real-time ecosystem health aggregation
- WebSocket push for live metrics/logs/traces
- OTel trace waterfall visualization
- SQL query workbench (SQLite, PostgreSQL, MySQL, Oracle)
- DB schema inspector and migration auditing
- Gateway route editor with WASM hot-swap UI
- Consistent hash ring visualization (Pranor Vault)
- RBAC-based access control (viewer/operator/admin)
- OIDC/OAuth2 SSO login
- Alerting engine with Slack/email/webhook notifications
- Incident timeline auto-generation
- Log aggregation with full-text search and live tail
- Custom dashboard builder (drag-and-drop widgets)
- SLO/SLI tracking with error budgets
- Deployment tracking with one-click rollback
- Multi-environment management (dev/staging/prod)
- Cost estimation panel
- Runbook automation (auto-execute on alert)
- Embedded terminal for service diagnostics
- AI agent traffic observatory (MCP tool calls, token costs)
- Infrastructure provisioning (create buckets, topics from UI)
- Dark/light theme toggle
- Audit log dashboard

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PRANOR_DISCOVERY` | JSON manifest or file path for service URLs | (uses CLI flags) |
| `PRANOR_JWT_SECRET` | Shared JWT secret for console auth | (required for SSO) |
| `PRANOR_OTLP_ENDPOINT` | OTel collector | (disabled) |

## CLI Flags

```bash
pranor-console \
  --port 8083 \
  --gate-url http://localhost:8080 \
  --store-url http://localhost:8081 \
  --queue-url http://localhost:8082 \
  --trace-url http://localhost:8090 \
  --tunnel-url http://localhost:8443 \
  --auth-url http://localhost:8098 \
  --db-url http://localhost:8097 \
  --mail-url http://localhost:8094 \
  --flow-url http://localhost:8096
```

## Endpoints

| Endpoint | Description |
|----------|-------------|
| `GET /healthz` | Liveness probe |
| `GET /api/status` | Aggregated service health |
| `GET /api/events` | WebSocket event stream |
| `GET /api/routes` | Pranor Gate route management |
| `GET /api/traces/replay` | Trace request replay |
| `GET /api/alerts` | Alert list + ack |
| `GET /api/logs` | Log aggregation query |
| `GET /api/topology` | Service dependency graph |
| `GET /` | Web UI dashboard |
