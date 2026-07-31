# Pranor Documentation

> The complete reference for the Serv ecosystem — language, components, operations, and architecture.

---

## Getting Started

| Doc | Description |
|-----|-------------|
| [Getting Started](getting-started.md) | Install Serv, write your first service, run it |
| [Language Guide](language-guide.md) | Full language tutorial — 30 sections covering all features |
| [Examples](examples.md) | Categorized code examples with explanations |

---

## Reference

| Doc | Description |
|-----|-------------|
| [Language Reference](language-reference.md) | Detailed syntax specification and type system |
| [Built-in Functions](builtins.md) | `log`, `db`, `cache`, `http`, `json`, `ai`, `store`, `broker` |
| [Standard Library](stdlib.md) | 48 importable `.pnr` modules (auth, jwt, retry, pagination, etc.) |
| [CLI Reference](cli.md) | All `serv` commands with flags and usage |

---

## Components

| Doc | Description |
|-----|-------------|
| [Component Catalog](components/README.md) | All 16 services with status, ports, and architecture |
| [Pranor Gate](components/Pranor Gate.md) | API Gateway — WASM middleware, AI routing, MCP support |
| [Pranor Vault](components/Pranor Vault.md) | Object Storage — S3-compatible, semantic search, time-travel |
| [Pranor Pulse](components/Pranor Pulse.md) | Message Broker — STOMP, WASM transforms, DLQ, tiered storage |
| [Pranor Console](components/Pranor Console.md) | Dashboard — unified observability, SQL workbench, alerting |
| [Pranor Mesh](components/Pranor Mesh.md) | Service Mesh — library-level, mTLS, circuit breaking |
| [Pranor Cache](components/Pranor Cache.md) | Cache — Redis/in-memory, namespacing, TTL |
| [Pranor Chrono](components/Pranor Chrono.md) | Scheduler — leader election, cron syntax, Pranor Vault persistence |
| [Pranor Deploy](components/Pranor Deploy.md) | Deployment — process orchestration, Docker, gateway sync |
| [Pranor Trace](components/Pranor Trace.md) | Tracing — OTLP ingestion, waterfall UI, anomaly detection |
| [Pranor Tunnel](components/Pranor Tunnel.md) | Tunneling — WebSocket relay, request inspection, subdomain routing |
| [Pranor Auth](components/Pranor Auth.md) | Identity — OAuth2/OIDC, MFA, RBAC, social login |
| [Pranor Pool](components/Pranor Pool.md) | Database Proxy — pooling, routing, query analytics |
| [Pranor Notify](components/Pranor Notify.md) | Notifications — SMTP, Slack, SMS, templates |
| [Pranor Flow](components/Pranor Flow.md) | Workflows — DAG execution, sagas, approval gates |
| [Pranor Hub](components/Pranor Hub.md) | Packages — semver resolution, signing, Pranor Vault backend |
| [ServDocs](components/ServDocs.md) | Documentation — auto-generated from `.pnr` source |
| [Pranor Core](components/Pranor Core.md) | Common Library — health probes, OTel, JWT middleware |
| [servlockctl](https://github.com/vyuvaraj/pranor/tree/main/packages/servlockctl) | Lock CLI — distributed lock acquisition, renewal, and deadlock inspection |
| [servsecretctl](https://github.com/vyuvaraj/pranor/tree/main/packages/servsecretctl) | Secrets CLI — key rotation, secret injection, and Shamir unseal operations |

---

## Operations

| Doc | Description |
|-----|-------------|
| [Deployment Guide](deployment.md) | Docker, TLS, multi-target deploy, production config |
| [Docker Compose Guide](docker-guide.md) | Run the full 16-service stack locally |
| [Architecture](architecture.md) | Runtime dependencies, service interactions, layers |
| [Roadmap](../UNIFIED_ROADMAP.md) | What's done, what's next, maturity matrix |

---

## Port Allocation

| Service | Port | Protocol |
|---------|------|----------|
| Pranor Gate | 8080 | HTTP/HTTPS |
| Pranor Vault | 8081 | HTTP (S3) |
| Pranor Pulse | 8082 / 61613 | HTTP + STOMP |
| Pranor Console | 8083 | HTTP |
| Pranor Cache | 8084 | HTTP |
| Pranor Chrono | 8085 | HTTP |
| Pranor Deploy | 8086 | HTTP |
| Pranor Mesh | 8087 | HTTP |
| Pranor Hub | 8088 | HTTP |
| ServDocs | 8089 | HTTP |
| Pranor Trace | 8090 | HTTP (OTLP) |
| Pranor Notify | 8094 | HTTP |
| Pranor Flow | 8096 | HTTP |
| Pranor Pool | 8097 | HTTP |
| Pranor Auth | 8098 | HTTP |
| Pranor Tunnel | 8443 | WebSocket |
