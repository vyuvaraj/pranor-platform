# Pranor Core — Common Service Library

> **Status:** ✅ Production | **Used by:** All Pranor services | **Repository:** [Pranor Core](https://github.com/vyuvaraj/pranor/tree/main/packages/Pranor Core)

## Overview

Pranor Core is the common Go library imported by all Pranor services. It provides standardized health probes (`/healthz`, `/readyz`), OpenTelemetry tracer initialization, JWT authentication middleware, structured JSON logging, and service token generation utilities.

## Key Features

- Standardized `/healthz` and `/readyz` probe handlers
- OpenTelemetry tracer initialization with OTLP export
- JWT authentication middleware (verify + extract claims)
- Structured JSON logging with request correlation
- Service-to-service token generation
- User token generation with configurable claims
- Graceful shutdown helpers
- Common error response formatting

## Key Exports

| Export | Description |
|--------|-------------|
| `HealthzHandler` | HTTP handler for liveness probes |
| `ReadyzHandler` | HTTP handler for readiness probes |
| `InitTrace()` | Initialize OpenTelemetry tracer |
| `AuthMiddleware()` | JWT verification middleware |
| `GenerateServiceToken()` | Create service-to-service JWT |
| `GenerateUserToken()` | Create user-scoped JWT |

## Usage

Pranor Core is not a standalone service — it is imported as a Go module by all Pranor services.

```go
import "github.com/vyuvaraj/Pranor Core/pkg/shared"

// Health probes
mux.HandleFunc("/healthz", shared.HealthzHandler)
mux.HandleFunc("/readyz", shared.ReadyzHandler)

// OTel init
shutdown := shared.InitTrace("my-service")
defer shutdown(ctx)

// Auth middleware
protected := shared.AuthMiddleware()(handler)
```
