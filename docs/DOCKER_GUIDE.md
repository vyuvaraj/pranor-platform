# Docker Compose Deployment Guide (Multi-service Production)

This guide documents deploying the complete 15-service ecosystem via Docker Compose, including configuration variables, networking boundaries, and healthcheck mappings.

## Docker Compose Manifest

Create a `docker-compose.yml` to boot core components alongside the newer identity, database, and workflow agents:

```yaml
version: '3.8'

services:
  # ── Core Infrastructure ─────────────────────────────────────────────────────
  serv-store:
    image: pranor/pranor-vault:latest
    ports:
      - "8081:8081"
    environment:
      - PORT=8081
      - DATA_DIR=/data
    volumes:
      - store-data:/data
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8081/healthz"]
      interval: 10s

  serv-queue:
    image: pranor/pranor-pulse:latest
    ports:
      - "8082:8082"
    environment:
      - PORT=8082
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8082/healthz"]
      interval: 10s

  # ── Identity & Persistence (Phase 9/10) ─────────────────────────────────────
  serv-auth:
    image: pranor/pranor-auth:latest
    ports:
      - "8098:8098"
    environment:
      - PORT=8098
      - PRANOR_JWT_SECRET=my-jwt-shared-secret
      - PRANOR_STORE_URL=http://serv-store:8081
    depends_on:
      - serv-store
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8098/healthz"]
      interval: 10s

  serv-db:
    image: pranor/Pranor Pool:latest
    ports:
      - "8097:8097"
    environment:
      - PORT=8097
      - DATABASE_URL=sqlite:///data/app.db
    volumes:
      - db-data:/data
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8097/healthz"]
      interval: 10s

  # ── Workflow Engine & Schedulers ────────────────────────────────────────────
  serv-flow:
    image: pranor/pranor-flow:latest
    ports:
      - "8096:8096"
    environment:
      - PORT=8096
      - PRANOR_STORE_URL=http://serv-store:8081
    depends_on:
      - serv-store
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8096/healthz"]
      interval: 10s

  serv-mail:
    image: pranor/pranor-notify:latest
    ports:
      - "8094:8094"
    environment:
      - PORT=8094
      - SMTP_HOST=smtp.mailtrap.io
      - SMTP_PORT=2525
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8094/healthz"]
      interval: 10s

  # ── API Gateway & Ingress ───────────────────────────────────────────────────
  serv-gate:
    image: pranor/pranor-gate:latest
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
      - PRANOR_JWT_SECRET=my-jwt-shared-secret
    depends_on:
      - serv-auth
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/healthz"]
      interval: 10s

volumes:
  store-data:
  db-data:
```

## Security Configurations (TLS / JWT Hardening)

1. **Shared Secret Rotation**: Ensure `$PRANOR_JWT_SECRET` is synchronized between `serv-auth`, `serv-gate`, and your compiled `pranor` backend applications.
2. **mTLS Client Authorization**: Enable server-to-server TLS verification by mounting client certificates inside containers and setting `PRANOR_MUTUAL_TLS=true` on your service mesh.
