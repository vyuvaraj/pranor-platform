# Serv Unified Ecosystem Roadmap - Completed Items (Phases 16-20)

This document preserves the archived history of completed items for Phase 16, Phase 17, Phase 18, Phase 19, and Phase 20.

---

## Phase 16: Operational Hardening & Production Readiness (Completed Items)

All backlog tasks targeting the upgrade of remaining components from **Stable** to **Production-Ready** by hardening their security, persistence, and observability layers:

### 📦 Pranor Vault
- **KMS Envelope Encryption** — Implement envelope encryption via AWS KMS / Google Cloud KMS for stored S3 objects to secure sensitive file payloads. [July 9, 2026]
- **OTel Performance Instrumentation** — Add OpenTelemetry metric tracking for S3 upload/download latency, throughput, and error budgets. [July 9, 2026]

### 📥 Pranor Pulse
- **mTLS Client Verification** — Enforce client certificate authentication (mTLS) for publishers and subscribers on enterprise topics. [July 9, 2026]
- **Prometheus Queue Lag Metrics** — Export message consumer lag, queue depth, and processing latency directly to Prometheus endpoints. [July 9, 2026]

### 💻 Pranor Console
- **Persistent Session Storage** — Implement PostgreSQL/Sqlite-based persistent storage for user sessions to ensure session survivability. [July 9, 2026]
- **Frontend Playwright E2E Tests** — Set up Playwright automated browser tests to validate critical UI flows and charts. [July 9, 2026]

### ⚡ Pranor Cache
- **Redis/Memcached Protocol TLS** — Support native TLS encryption for all Redis and Memcached client connections. [July 9, 2026]
- **OTel Cache Metrics** — Export cache hit/miss ratio, memory fragmentation, and key eviction counts to central OTel collectors. [July 9, 2026]

### ⏰ Pranor Chrono
- **API RBAC Enforcement** — Enforce Role-Based Access Control on job builder and trigger APIs, requiring admin privilege to register new crons. [July 9, 2026]
- **Execution Syslog Integration** — Direct cron job stdout, exit statuses, and durations to syslog or central log drains. [July 9, 2026]

### 🛡️ Pranor Auth
- **Persistent Token Storage** — Store issued refresh tokens in a database to enable remote token revocation and active session audits. [July 9, 2026]
- **Auth Audit Logs** — Generate structured JSON audit trails for all authentication events, login failures, and MFA setups. [July 9, 2026]

### 🗄️ Pranor Pool
- **mTLS Connection Checks** — Support mutual TLS verification for all database client connections. [July 9, 2026]
- **Pool Connection Stats** — Export connection pool usage, active/idle count, and wait duration metrics. [July 9, 2026]

### 📧 Pranor Notify
- **DKIM Outbound Signatures** — Add support for automated DKIM signature headers and SPF verification checks. [July 9, 2026]
- **Disk Queue Persistence** — Save outgoing mail queues to disk to prevent email loss during server restarts or crashes. [July 9, 2026]

### 🔄 Pranor Flow
- **Saga State DB Storage** — Persist saga states, transaction steps, and rollback progress in a distributed database instead of memory. [July 9, 2026]
- **Flow Latency telemetry** — Add OTel spans tracking execution duration for each step and overall workflow success rates. [July 9, 2026]

---

## Phase 17: Zero-Trust Clustering & Edge Serverless Evolution (Completed Items)

All zero-trust security mesh, serverless compilation, and resilient caching enhancements:

### 🛡️ Zero-Trust Mesh & Gateway Resilience
- **Distributed Rate-Limiting Backend** — Extend Pranor Gate to support dynamic Redis/Valkey rate-limiting stores instead of in-memory maps. [July 10, 2026]
- **Inter-Service Mesh mTLS** — Enforce automatic mutual TLS client verification for all inter-service mesh routes inside Pranor Mesh. [July 10, 2026]
- **Secure Enclave Isolation** — Add process execution support within secure enclaves (e.g. AWS Nitro Enclaves, Intel SGX). [July 10, 2026]

### 📦 S3 Durability & Pool Auto-Recovery
- **Write-Ahead Logging (WAL)** — Add WAL and fsync safety limits to Pranor Vault S3 layers to prevent dirty writes during unexpected node shutdowns. [July 10, 2026]
- **Connection Pool Leak Recovery** — Add automatic timeout reaping for deadlocked connection leases in Pranor Pool pools. [July 10, 2026]
- **LRU Cache Key Eviction** — Implement thread-safe Least Recently Used (LRU) key evictions in Pranor Cache memory stores. [July 10, 2026]

### ⚡ Edge Serverless & Code Execution
- **WASM Edge Compilation** — Compile Pranor code modules directly to WebAssembly components for zero-cold-start hosting on Wasmtime. [July 10, 2026]
- **AI Observability Pipelines** — Enable automatic scaling triggers and query cache rule mutations via Pranor Console observability hooks. [July 10, 2026]

---

## Phase 18: OSS-to-EE Boundary Alignment & Refactoring (Completed Items)

All commercial dual-license segmentation and hook separation boundaries:

### 📦 Pranor Vault & Pranor Pulse
- **KMS Enterprise Separation** — Migrate AWS KMS, Google Cloud KMS, and HashiCorp Vault implementations to EE, leaving only simple local key encryption in OSS. [July 10, 2026]
- **mTLS Enforcement Hooks** — Restrict client certificate authentication and PKI mappings to the EE broker overlay. [July 10, 2026]

### 🛡️ Pranor Auth & Pranor Pool
- **Session & Audit Log Isolation** — Move account lockout history, security audit trail generators, and remote session revocation control logic to EE. [July 10, 2026]
- **Database Replication Topologies** — Restrict replica pool routing, read/write splitting, and dynamic failover state machines to EE. [July 10, 2026]

### 📧 Pranor Notify & 🔄 Pranor Flow
- **DKIM Signing Delegation** — Delegate outbound DKIM header signing and SPF alignment checks to EE. [July 10, 2026]
- **Distributed Saga Checkpoints** — Separate distributed database state persistence from local file-based (`.state`) saga checkpoints. [July 10, 2026]

---

## Phase 19: Component Maturity Alignment (Completed Items)

All final component hardening, documentation, and unit test alignment tasks:

### 📦 Pranor Vault & 📥 Pranor Pulse
- **Documentation Hardening** — Write comprehensive API reference guides and operational recovery playbooks for S3 and STOMP message brokers. [July 10, 2026]
- **S3 Test Coverage Expansion** — Build automated failure-injection tests validating object upload behavior under network partition states. [July 10, 2026]

### 💻 Pranor Console
- **High-Availability Session Stores** — Upgrade local session stores to support clustered database replication backends. [July 10, 2026]
- **Playwright Test Matrix** — Expand Playwright browser test coverage to cover user permissions and custom widgets. [July 10, 2026]

### 🛡️ Pranor Auth, 🗄️ Pranor Pool, & 📧 Pranor Notify
- **E2E Trace Validation** — Verify spans propagate seamlessly through the Auth IDP, DB connection proxy, and Mail delivery pipelines. [July 10, 2026]
- **Persistent Audits Schema** — Implement long-term structured schemas and storage retention options for audit logs and mail event queues. [July 10, 2026]

### 🔄 Pranor Flow
- **Saga Verification Tests** — Create unit tests demonstrating complex Saga transactional rollbacks under network time-out and downstream component failure conditions. [July 10, 2026]

---

## Phase 20: OSS-to-EE Refactoring & Enterprise Migrations (Completed Items)

Refactored advanced OSS capabilities into clean, build-tagged hooks, migrating their enterprise implementations to the private `pranor-ee` overlay:

### 🗄️ Pranor Pool
- **Query Routing Optimizer** — Move the dynamic read/write splitting parser and query-caching optimization layer into the enterprise overlay, leaving OSS with a simple direct pool dispatcher. [July 10, 2026]

### 🤖 AI-Native Features (Ungated — Move to EE)
- **EE.23: Auto-Scaling Engine** (Pranor Deploy) — Move `autoscale.go` behind build tag; OSS has no auto-scaling capability. [July 11, 2026]
- **EE.25: Anomaly Explanation** (Pranor Trace) — Move `handleExplainAnomaly` behind build tag; OSS returns 403 EE required. [July 11, 2026]
- **EE.27: Multi-Relay Federation** (Pranor Tunnel) — Move federation peer routing behind build tag; OSS is single-relay only. [July 11, 2026]
- **EE.30: Semantic Message Routing** (Pranor Pulse) — Move semantic classification routing behind build tag; OSS uses exact topic match only. [July 11, 2026]
- **MFA TOTP Engine Separation** (Pranor Auth) — Move the MFA TOTP generator, secret storage, and validator logic into the enterprise tag, leaving OSS with basic password authentication. [July 11, 2026]
- **Social OAuth Provider Hooks** (Pranor Auth) — Extract third-party social OAuth integrations (Google, GitHub, GitLab) out of OSS, replacing them with generic pluggable enterprise authentication hooks. [July 11, 2026]
- **EE.28: Credential Stuffing Detection** (Pranor Auth) — Move stuffing detection behind build tag; OSS has basic lockout only. [July 11, 2026]
- **EE.29: Predictive Alerts** (Pranor Console) — Move `HandlePredictiveAlerts` behind build tag; OSS returns 403 EE required. [July 11, 2026]
- **Saga Parallel Coordinator** (Pranor Flow) — Refactor concurrent execution paths, DAG forks, and complex compensating workflows in Saga engines to the enterprise overlay. [July 11, 2026]
- **Dynamic HTML Templates** (Pranor Notify) — Separate the Go template engine compilation, MJML layout transpilation, and variable-bound rendering into the enterprise tag. [July 11, 2026]
- **Clustered Leader Election** (Pranor Chrono) — Extract distributed leader election and consensus-backed job execution scheduling for Cron clusters into the enterprise overlay. [July 11, 2026]
- **EE.21: Semantic API Caching** (Pranor Gate) — Move embedding-based cache logic to `ai_ee.go`; OSS always returns cache-miss. [July 11, 2026]
- **EE.22: Cost-Aware LLM Routing** (Pranor Gate) — Move model selection/fallback logic to `ai_ee.go`; OSS routes to single default model only. [July 11, 2026]
- **EE.24: Cold Tier Archival** (Pranor Trace) — Move Pranor Vault archival logic behind build tag; OSS keeps traces in-memory ring buffer only. [July 11, 2026]
- **EE.26: SLO Breach Prediction** (Pranor Trace) — Move `handleSloBreachPredict` behind build tag; OSS returns 403 EE required. [July 11, 2026]


