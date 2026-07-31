# Pranor Flow — Workflow Orchestrator

> **Status:** 🟡 Stable | **Port:** 8096 | **Repository:** [Pranor Flow](https://github.com/vyuvaraj/pranor/tree/main/packages/Pranor Flow)

## Overview

Pranor Flow is a workflow orchestrator with DAG-based execution, durable checkpointing, saga compensation and rollback, human approval gates, and event-triggered execution via Pranor Pulse for building complex multi-step business processes.

## Key Features

- DAG-based workflow execution engine
- Durable checkpointing for crash recovery
- Saga pattern with compensation/rollback steps
- Human approval gates with timeout escalation
- Event-triggered workflows via Pranor Pulse
- Parallel step execution within DAG levels
- Workflow versioning and migration
- Execution history and audit trail

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | HTTP listen port | `8096` |
| `PRANOR_VAULT_URL` | Pranor Vault URL for state persistence | (required) |
| `PRANOR_PULSE_URL` | Pranor Pulse URL for event triggers | (required) |
| `PRANOR_OTLP_ENDPOINT` | OTel collector URL | (disabled) |

## Endpoints

| Endpoint | Description |
|----------|-------------|
| `GET /healthz` | Liveness probe |
| `POST /api/v1/workflows` | Create a workflow definition |
| `GET /api/v1/workflows/{id}` | Get workflow definition |
| `POST /api/v1/workflows/{id}/approve` | Approve a pending gate |
| `GET /api/v1/executions` | List workflow executions |

## Pranor Integration

```srv
// Native workflow block syntax in .pnr files
workflow "order-fulfillment" {
    step "validate" { ... }
    step "charge" { compensate { ... } }
    step "ship" { ... }
}
```
