# Operational Runbooks — Recovery & Incident Response

This document provides checklists and runbooks for standard operational alert events in the Pranor ecosystem.

## Runbook 001: Pranor Gate High Latency / Backpressure
**Alert Trigger**: Inbound latency on `:8080` exceeds 1.5s, or request queues fill.

### 1. Diagnosis
1. Inspect Pranor Console trace list to see which downstream route is bottlenecking:
   ```bash
   serv status --json
   ```
2. Check active connection pool state. If pool size matches `max_open_conns`, verify downstream database health.

### 2. Resolution
1. Temporarily increase gateway concurrency limit by updating `max_concurrent_requests` in `config.json` and reloading:
   ```json
   "max_concurrent_requests": 100
   ```
2. If downstreams are timed out, apply a temporary circuit breaker:
   ```bash
   # Add emergency routing override in config.json
   "target": "http://localhost:8080/error-fallback"
   ```

---

## Runbook 002: Pranor Auth Verification Failures (Key Rotation)
**Alert Trigger**: HTTP 401/403 errors spike globally. Token validation fails on `Pranor Core/middleware.go`.

### 1. Diagnosis
1. Query key cache endpoint to check if rotated keys are public:
   ```bash
   curl -i http://localhost:8098/oauth/keys
   ```
2. Verify if the JWKS URL `:8098/oauth/keys` returns valid RSA keys.

### 2. Resolution
1. Force JWKS cache invalidation inside `Pranor Core` client by updating the cache expiration timestamp.
2. If a key leak is suspected, issue an emergency rotation:
   ```bash
   # Trigger key generator script
   ./scripts/rotate-keys.sh
   ```
