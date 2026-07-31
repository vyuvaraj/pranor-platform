# Pranor Deploy — Multi-Target Deployment Orchestrator

> **Status:** ✅ Production | **Port:** 8086 | **Repository:** [Pranor Deploy](https://github.com/vyuvaraj/pranor/tree/main/packages/Pranor Deploy)

## Overview

Pranor Deploy is a multi-target deployment orchestrator supporting process management, Docker container runner, WASM isolation, automatic port allocation, gateway route synchronization, and health monitoring across all deployed services.

## Key Features

- Multi-target deployment (Docker, Fly.io, Railway, Render)
- Process manager for native binaries
- Docker container lifecycle management
- WASM sandbox isolation for untrusted workloads
- Automatic port allocation and conflict avoidance
- Gateway route sync with Pranor Gate
- Health monitoring with automatic restarts
- Deployment logs and rollback support

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVCLOUD_PORT` | HTTP listen port | `8086` |
| `PRANOR_JWT_SECRET` | JWT verification secret | (required) |
| `SERVGATE_URL` | Pranor Gate URL for route sync | (optional) |
| `PRANOR_OTLP_ENDPOINT` | OTel collector URL | (disabled) |

## Endpoints

| Endpoint | Description |
|----------|-------------|
| `GET /healthz` | Liveness probe |
| `POST /api/v1/deploy` | Deploy a service |
| `GET /api/v1/services` | List deployed services |
| `DELETE /api/v1/services/{id}` | Stop and remove a service |
| `GET /api/v1/logs/{id}` | Stream service logs |

## Pranor Integration

```srv
pranor deploy --target docker|fly|railway|render
```
