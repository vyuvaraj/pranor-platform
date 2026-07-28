# Unified Roadmap - Completed Phases 61 to 65

## Phase 64: ServMail — DMARC Enforcement, Inbound Webhooks & Email Template DSL (Completed)

> **Current State**: ServMail implements SMTP ingestion, DKIM signing, a disk-backed sending queue, and basic Handlebars-style template rendering.
> **What is Missing**: DMARC policy enforcement (SPF + DKIM alignment validation for outbound integrity), inbound email webhook routing (no ability to receive and process incoming mail), a full template DSL with partials and loops, bounce and complaint handling with automatic suppression list management, and email analytics in ServConsole.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| ML.G1 | **DMARC Policy Enforcement (SPF + DKIM Alignment Validation)** | ServMail Security | Validate DMARC `p=quarantine`/`p=reject` policy alignment between SPF envelope-from domain and DKIM `d=` header domain before delivering or relaying outbound messages | [x] | OSS |
| ML.G2 | **Inbound Email Webhook Router** | ServMail Inbound | Parse incoming SMTP messages and route them to configurable HTTP webhook endpoints based on recipient address pattern matching; enables "email as a workflow trigger" use cases | [x] | OSS |
| ML.G3 | **Email Template DSL with Partials, Loops & Conditional Blocks** | ServMail Templates | Extend the template engine with `{{#each items}}`, `{{#if condition}}`, and `{{> partial_name}}` support for reusable transactional email component composition | [x] | OSS |
| ML.G4 | **Bounce & Complaint Handling with Automatic Suppression List** | ServMail Delivery | Parse SMTP bounce DSNs and ISP Feedback Loop (FBL) complaint notifications; automatically add bounced and complained addresses to a suppression list to protect sender reputation | [x] | OSS |
| ML.G5 | **One-Click Unsubscribe (RFC 8058) & List Management** | ServMail Compliance | Auto-inject `List-Unsubscribe` and `List-Unsubscribe-Post` headers; handle one-click unsubscribe webhooks; maintain per-sender suppression lists with full audit history | [x] | OSS |
| ML.G6 | **Email Delivery Analytics Dashboard (Open, Click, Bounce Rates)** | ServConsole UI | Stream per-campaign delivery, open-pixel tracking, click-through, and bounce event metrics into ServConsole for deliverability health monitoring and alert thresholds | [x] | OSS |

---





## Phase 65: ServPool — Adaptive Scaling, Read-Replica Routing & Connection Health (Completed)

> **Current State**: ServPool implements an adaptive connection pool with LRU eviction, dynamic max-connection scaling, a wait queue for pool saturation, and per-dialect connection management.
> **What is Missing**: Read/write split routing (primary for writes, replicas for reads), pre-checkout connection health validation (no heartbeat ping before returning stale connections), automatic connection leak detection, per-query latency histograms, and prepared statement caching.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| SP.G1 | **Read/Write Split Router (Primary for Writes, Replica for Reads)** | ServPool Routing | Detect read vs. write SQL intent (SELECT vs. INSERT/UPDATE/DELETE) and route queries to appropriate replica or primary connections; configure per-replica weights for read load distribution | [x] | OSS |
| SP.G2 | **Pre-Checkout Connection Health Validation (Ping & Validation Query)** | ServPool Health | Execute a configurable lightweight validation query (e.g. `SELECT 1`) before returning a pooled connection to the caller; immediately discard and replace stale, closed, or broken connections | [x] | OSS |
| SP.G3 | **Automatic Connection Leak Detection & Forced Reclaim** | ServPool Safety | Track connections checked out beyond a configurable `maxHoldDuration`; log a stack trace warning and forcibly reclaim leaked connections to prevent pool exhaustion under load | [x] | OSS |
| SP.G4 | **Per-Query Execution Time Histogram & Slow Query Logger** | ServPool Telemetry | Wrap every query execution with nanosecond timing; maintain P50/P95/P99 latency histograms per query fingerprint and log queries exceeding a configurable slow-query threshold | [x] | OSS |
| SP.G5 | **Multi-Dialect Prepared Statement Cache** | ServPool Cache | Cache parsed and compiled prepared statements per-connection-per-dialect (PostgreSQL, MySQL, SQLite); reduce repeated parse overhead on high-throughput transactional workloads | [x] | OSS |
| SP.G6 | **Pool Utilization & Saturation Alerting in ServConsole** | ServConsole UI | Stream real-time pool utilization (active/idle/waiting connection counts) to ServConsole; alert operators when pool saturation or wait queue depth exceeds configurable thresholds | [x] | OSS |

---







---

## Phase 61: ServMesh — WireGuard Overlay, mTLS Identity & Adaptive Load Balancing (Completed)

> **Current State**: ServMesh implements a library-level service registry with heartbeat TTL eviction, round-robin load balancing, circuit breaking, mutual TLS, and a topology inspection CLI.
> **What is Missing**: Automatic WireGuard kernel mesh overlay between nodes (currently requires manual TLS cert management), SPIFFE/SPIRE workload identity attestation, latency-aware and locality-aware load balancing (currently round-robin only), global distributed rate limiting, and a live real-time topology graph in ServConsole.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| SM.G1 | **Automatic WireGuard Kernel Tunnel Mesh Between Nodes** | ServMesh Network | Auto-provision WireGuard encrypted kernel-level tunnels between ServMesh peers using a DHT-based key exchange protocol; eliminates manual TLS certificate provisioning for inter-service traffic | [x] | **EE** |
| SM.G2 | **SPIFFE/SPIRE mTLS Workload Identity Attestation** | ServMesh Identity | Issue short-lived SPIFFE SVIDs (X.509 certificates) to each registered service instance; enforce mTLS workload identity verification on all service-to-service calls within the mesh | [x] | **EE** |
| SM.G3 | **Latency-Aware P2C Load Balancing & Locality Preference** | ServMesh LB | Replace round-robin with Power-of-Two-Choices (P2C) latency-weighted load balancing; prefer same-zone or same-region replicas to minimize cross-datacenter latency and egress cost | [x] | OSS |
| SM.G4 | **Distributed Global Rate Limiting via ServCache Token Buckets** | ServMesh RateLimit | Implement distributed token-bucket rate limiting shared across all ServMesh nodes via ServCache; prevent upstream overload cascades from triggering inter-service retry storms | [x] | OSS |
| SM.G5 | **Live Real-Time Service Topology Graph in ServConsole** | ServConsole UI | Render an animated, force-directed service dependency graph in ServConsole showing live traffic flows, circuit breaker open/closed states, and per-edge P99 latency metrics | [x] | OSS |
| SM.G6 | **Chaos Fault Injection API (Latency, Error Rate, Partition Simulation)** | ServMesh Chaos | Expose a Chaos Engineering REST API that injects simulated network delays, configurable error rates, and service partition failures into ServMesh routing for resilience validation | [x] | OSS |

---

---

## Phase 62: ServCloud — Blue/Green Deploys, Preview URLs & Container Isolation (Completed)

> **Current State**: ServCloud implements a process orchestrator deploying `serv` binaries with health checks, stdout log streaming, deployment history, `process` and `wasm` isolation modes, and preview URL stubs.
> **What is Missing**: Actual blue/green deployment with traffic weight shifting via ServGate, canary deployments with automatic error-rate rollback, fully functional per-branch preview subdomain routing, Docker/OCI container isolation mode, CPU/memory cgroup enforcement, and a deployment approval gate.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| CL.G1 | **Blue/Green Deployment with Atomic Traffic Cutover via ServGate** | ServCloud Deploy | Run a new service version in parallel with the old version; shift 100% of traffic atomically via a ServGate route weight update after health check validation passes | [x] | OSS |
| CL.G2 | **Canary Deployment with Automatic Rollback on Error Rate Threshold** | ServCloud Canary | Roll out new service versions to a configurable percentage of traffic; automatically roll back to the stable version when error rate exceeds a configurable SLO threshold | [x] | OSS |
| CL.G3 | **Fully Functional Per-Branch Preview Environment with Isolated Subdomain** | ServCloud Preview | Auto-provision isolated preview environments for each Git branch or PR, accessible at `<branch>.preview.servcloud.dev`, with automatic ServGate route registration and teardown on PR merge | [x] | OSS |
| CL.G4 | **Docker / OCI Container Isolation Mode** | ServCloud Runtime | Add `docker` isolation mode alongside `process` and `wasm`; pull OCI images, start containers with resource limits, health checks, and manage the full container lifecycle | [x] | OSS |
| CL.G5 | **CPU & Memory cgroup Resource Limits & Usage Telemetry** | ServCloud Resources | Enforce per-service CPU and memory cgroup resource limits; stream real-time consumption metrics to ServConsole with breach alerting and auto-throttle enforcement | [x] | **EE** |
| CL.G6 | **Deployment Approval Gate & Operator Confirm Flow** | ServCloud Safety | Add a mandatory operator approval gate for production tier deployments; send Slack/webhook notifications requesting explicit approval before traffic cutover is executed | [x] | **EE** |

---
