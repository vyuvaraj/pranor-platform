# Unified Roadmap - Completed Phases 61 to 65

## Phase 64: Pranor Notify — DMARC Enforcement, Inbound Webhooks & Email Template DSL (Completed)

> **Current State**: Pranor Notify implements SMTP ingestion, DKIM signing, a disk-backed sending queue, and basic Handlebars-style template rendering.
> **What is Missing**: DMARC policy enforcement (SPF + DKIM alignment validation for outbound integrity), inbound email webhook routing (no ability to receive and process incoming mail), a full template DSL with partials and loops, bounce and complaint handling with automatic suppression list management, and email analytics in Pranor Console.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| ML.G1 | **DMARC Policy Enforcement (SPF + DKIM Alignment Validation)** | Pranor Notify Security | Validate DMARC `p=quarantine`/`p=reject` policy alignment between SPF envelope-from domain and DKIM `d=` header domain before delivering or relaying outbound messages | [x] | OSS |
| ML.G2 | **Inbound Email Webhook Router** | Pranor Notify Inbound | Parse incoming SMTP messages and route them to configurable HTTP webhook endpoints based on recipient address pattern matching; enables "email as a workflow trigger" use cases | [x] | OSS |
| ML.G3 | **Email Template DSL with Partials, Loops & Conditional Blocks** | Pranor Notify Templates | Extend the template engine with `{{#each items}}`, `{{#if condition}}`, and `{{> partial_name}}` support for reusable transactional email component composition | [x] | OSS |
| ML.G4 | **Bounce & Complaint Handling with Automatic Suppression List** | Pranor Notify Delivery | Parse SMTP bounce DSNs and ISP Feedback Loop (FBL) complaint notifications; automatically add bounced and complained addresses to a suppression list to protect sender reputation | [x] | OSS |
| ML.G5 | **One-Click Unsubscribe (RFC 8058) & List Management** | Pranor Notify Compliance | Auto-inject `List-Unsubscribe` and `List-Unsubscribe-Post` headers; handle one-click unsubscribe webhooks; maintain per-sender suppression lists with full audit history | [x] | OSS |
| ML.G6 | **Email Delivery Analytics Dashboard (Open, Click, Bounce Rates)** | Pranor Console UI | Stream per-campaign delivery, open-pixel tracking, click-through, and bounce event metrics into Pranor Console for deliverability health monitoring and alert thresholds | [x] | OSS |

---





## Phase 65: Pranor Pool — Adaptive Scaling, Read-Replica Routing & Connection Health (Completed)

> **Current State**: Pranor Pool implements an adaptive connection pool with LRU eviction, dynamic max-connection scaling, a wait queue for pool saturation, and per-dialect connection management.
> **What is Missing**: Read/write split routing (primary for writes, replicas for reads), pre-checkout connection health validation (no heartbeat ping before returning stale connections), automatic connection leak detection, per-query latency histograms, and prepared statement caching.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| SP.G1 | **Read/Write Split Router (Primary for Writes, Replica for Reads)** | Pranor Pool Routing | Detect read vs. write SQL intent (SELECT vs. INSERT/UPDATE/DELETE) and route queries to appropriate replica or primary connections; configure per-replica weights for read load distribution | [x] | OSS |
| SP.G2 | **Pre-Checkout Connection Health Validation (Ping & Validation Query)** | Pranor Pool Health | Execute a configurable lightweight validation query (e.g. `SELECT 1`) before returning a pooled connection to the caller; immediately discard and replace stale, closed, or broken connections | [x] | OSS |
| SP.G3 | **Automatic Connection Leak Detection & Forced Reclaim** | Pranor Pool Safety | Track connections checked out beyond a configurable `maxHoldDuration`; log a stack trace warning and forcibly reclaim leaked connections to prevent pool exhaustion under load | [x] | OSS |
| SP.G4 | **Per-Query Execution Time Histogram & Slow Query Logger** | Pranor Pool Telemetry | Wrap every query execution with nanosecond timing; maintain P50/P95/P99 latency histograms per query fingerprint and log queries exceeding a configurable slow-query threshold | [x] | OSS |
| SP.G5 | **Multi-Dialect Prepared Statement Cache** | Pranor Pool Cache | Cache parsed and compiled prepared statements per-connection-per-dialect (PostgreSQL, MySQL, SQLite); reduce repeated parse overhead on high-throughput transactional workloads | [x] | OSS |
| SP.G6 | **Pool Utilization & Saturation Alerting in Pranor Console** | Pranor Console UI | Stream real-time pool utilization (active/idle/waiting connection counts) to Pranor Console; alert operators when pool saturation or wait queue depth exceeds configurable thresholds | [x] | OSS |

---







---

## Phase 61: Pranor Mesh — WireGuard Overlay, mTLS Identity & Adaptive Load Balancing (Completed)

> **Current State**: Pranor Mesh implements a library-level service registry with heartbeat TTL eviction, round-robin load balancing, circuit breaking, mutual TLS, and a topology inspection CLI.
> **What is Missing**: Automatic WireGuard kernel mesh overlay between nodes (currently requires manual TLS cert management), SPIFFE/SPIRE workload identity attestation, latency-aware and locality-aware load balancing (currently round-robin only), global distributed rate limiting, and a live real-time topology graph in Pranor Console.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| SM.G1 | **Automatic WireGuard Kernel Tunnel Mesh Between Nodes** | Pranor Mesh Network | Auto-provision WireGuard encrypted kernel-level tunnels between Pranor Mesh peers using a DHT-based key exchange protocol; eliminates manual TLS certificate provisioning for inter-service traffic | [x] | **EE** |
| SM.G2 | **SPIFFE/SPIRE mTLS Workload Identity Attestation** | Pranor Mesh Identity | Issue short-lived SPIFFE SVIDs (X.509 certificates) to each registered service instance; enforce mTLS workload identity verification on all service-to-service calls within the mesh | [x] | **EE** |
| SM.G3 | **Latency-Aware P2C Load Balancing & Locality Preference** | Pranor Mesh LB | Replace round-robin with Power-of-Two-Choices (P2C) latency-weighted load balancing; prefer same-zone or same-region replicas to minimize cross-datacenter latency and egress cost | [x] | OSS |
| SM.G4 | **Distributed Global Rate Limiting via Pranor Cache Token Buckets** | Pranor Mesh RateLimit | Implement distributed token-bucket rate limiting shared across all Pranor Mesh nodes via Pranor Cache; prevent upstream overload cascades from triggering inter-service retry storms | [x] | OSS |
| SM.G5 | **Live Real-Time Service Topology Graph in Pranor Console** | Pranor Console UI | Render an animated, force-directed service dependency graph in Pranor Console showing live traffic flows, circuit breaker open/closed states, and per-edge P99 latency metrics | [x] | OSS |
| SM.G6 | **Chaos Fault Injection API (Latency, Error Rate, Partition Simulation)** | Pranor Mesh Chaos | Expose a Chaos Engineering REST API that injects simulated network delays, configurable error rates, and service partition failures into Pranor Mesh routing for resilience validation | [x] | OSS |

---

---

## Phase 62: Pranor Deploy — Blue/Green Deploys, Preview URLs & Container Isolation (Completed)

> **Current State**: Pranor Deploy implements a process orchestrator deploying `serv` binaries with health checks, stdout log streaming, deployment history, `process` and `wasm` isolation modes, and preview URL stubs.
> **What is Missing**: Actual blue/green deployment with traffic weight shifting via Pranor Gate, canary deployments with automatic error-rate rollback, fully functional per-branch preview subdomain routing, Docker/OCI container isolation mode, CPU/memory cgroup enforcement, and a deployment approval gate.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| CL.G1 | **Blue/Green Deployment with Atomic Traffic Cutover via Pranor Gate** | Pranor Deploy Deploy | Run a new service version in parallel with the old version; shift 100% of traffic atomically via a Pranor Gate route weight update after health check validation passes | [x] | OSS |
| CL.G2 | **Canary Deployment with Automatic Rollback on Error Rate Threshold** | Pranor Deploy Canary | Roll out new service versions to a configurable percentage of traffic; automatically roll back to the stable version when error rate exceeds a configurable SLO threshold | [x] | OSS |
| CL.G3 | **Fully Functional Per-Branch Preview Environment with Isolated Subdomain** | Pranor Deploy Preview | Auto-provision isolated preview environments for each Git branch or PR, accessible at `<branch>.preview.pranor-deploy.dev`, with automatic Pranor Gate route registration and teardown on PR merge | [x] | OSS |
| CL.G4 | **Docker / OCI Container Isolation Mode** | Pranor Deploy Runtime | Add `docker` isolation mode alongside `process` and `wasm`; pull OCI images, start containers with resource limits, health checks, and manage the full container lifecycle | [x] | OSS |
| CL.G5 | **CPU & Memory cgroup Resource Limits & Usage Telemetry** | Pranor Deploy Resources | Enforce per-service CPU and memory cgroup resource limits; stream real-time consumption metrics to Pranor Console with breach alerting and auto-throttle enforcement | [x] | **EE** |
| CL.G6 | **Deployment Approval Gate & Operator Confirm Flow** | Pranor Deploy Safety | Add a mandatory operator approval gate for production tier deployments; send Slack/webhook notifications requesting explicit approval before traffic cutover is executed | [x] | **EE** |

---


---

## Phase 63: Pranor Trace — eBPF Continuous Profiling, Flamegraphs & SLO Burn Alerts (Completed)

> **Current State**: Pranor Trace implements OTel span ingestion, trace storage, anomaly detection hooks, SLO breach predictor hooks, self-healing, and retention cleanup. The anomaly explainer and SLO predictor interfaces exist but have no concrete implementations.
> **What is Missing**: Actual eBPF continuous CPU/memory profiling engine (the hooks are declared but not implemented), flamegraph rendering, automatic trace-to-flamegraph correlation, concrete SLO burn rate alerting, trace sampling policy management, and exemplar-linked Prometheus metrics.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| ST.G1 | **eBPF Continuous CPU & Memory Flamegraph Profiler (Implement AnomalyExplainer)** | Pranor Trace Profiler | Implement the declared `AnomalyExplainer` interface using eBPF perf probes to capture goroutine CPU time and heap allocation hotspots; generate folded flamegraph stacks exposed via `/api/v1/profiles` | [x] | OSS |
| ST.G2 | **OTel Trace-to-Flamegraph Automatic Correlation** | Pranor Trace Correlation | Automatically link eBPF flamegraph samples captured during a trace span's execution window; enable single-click drill-down from a slow OTel span to its kernel-level CPU hotspot in Pranor Console | [x] | OSS |
| ST.G3 | **SLO Burn Rate Alert Engine (Implement SloBreachPredictor — Multi-Window)** | Pranor Trace SLO | Implement the declared `SloBreachPredictor` interface using Google SRE-style multi-window (1h + 6h) error budget burn rate; fire alerts when burn rate exceeds configured thresholds | [x] | OSS |
| ST.G4 | **Trace Sampling Policy Manager (Head-Based & Tail-Based)** | Pranor Trace Sampling | Configure per-service head-based sampling rates; implement tail-based adaptive sampling that retrospectively retains 100% of traces containing errors or latency outliers above a P95 threshold | [x] | **EE** |
| ST.G5 | **Exemplar-Linked Prometheus Metrics & Histogram Correlation** | Pranor Trace Metrics | Embed OTel trace exemplars into Prometheus histogram buckets so Pranor Console dashboards can jump from a P99 latency spike directly to a representative trace example with one click | [x] | OSS |
| ST.G6 | **Critical Path Analyzer & Distributed Dependency Map** | Pranor Trace Analysis | Automatically compute the critical path across distributed trace spans; highlight the slowest causal dependency in each trace and visualize the full service call graph in Pranor Console | [x] | OSS |

---
