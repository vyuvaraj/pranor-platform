# Unified Roadmap - Completed Phases 56 to 60

## Phase 56: Pranor Gateway Enterprise WAF, Remote WASM Sync & OAuth2 Engine (Completed)

> **Context:** Commercial enterprise security and traffic management features for Pranor Gateway.

| # | Item | Component | Description | Status |
|---|------|-----------|-------------|--------|
| SG.VE1 | **Enterprise Remote WASM Sync Engine (`serv-ee`)** | Pranor Gateway EE | Cryptographically signed remote WASM plugin download & hot-reloading behind `//go:build enterprise` | [x] |
| SG.VE2 | **Enterprise WAF Ruleset & Threat Intelligence (`serv-ee`)** | Pranor Gateway EE | OWASP Top 10 automated threat intelligence WAF engine behind `//go:build enterprise` | [x] |
| SG.VE3 | **Enterprise OAuth2 / OIDC Token Introspection (`serv-ee`)** | Pranor Gateway EE | Distributed OIDC token validation & caching behind `//go:build enterprise` | [x] |
| SG.VE4 | **Enterprise Multi-Cloud Anycast Mesh Controller (`serv-ee`)** | Pranor Gateway EE | Global Anycast BGP route steering behind `//go:build enterprise` | [x] |
| SG.VE5 | **Pranor Gateway EE Modularization Verification (`serv-ee`)** | Pranor Gateway EE | Strict build-tag isolation & enterprise package testing | [x] |

---



All commercial enterprise features (**EE**) must have their core logic and implementations located exclusively inside the private `pranor-ee` repository. 

The open-source core repositories (such as `Pranor Gate`, `Pranor Vault`, etc.) must only expose clean interfaces, hooks, or config fields. The implementation of these hooks in the open-source code must use build-tagged placeholders (`//go:build !enterprise`), while the actual commercial code resides under the corresponding directories in `pranor-ee` and is built with `//go:build enterprise`.












---

## Strategic Module Gap Analysis — Phases 57 to 72

> **Context**: The following phases are derived from a deep critical analysis of each Pranor module's current implementation against industry-standard production expectations. Each phase documents concrete missing features — not aspirational items — that are required for the module to compete as a standalone product and fulfil its role within the Pranor ecosystem. Phases are ordered by module dependency depth: standalone utility modules first, cross-cutting platform layers last.

---


## Phase 59: Pranor Chrono — DAG Job Chaining, Retry Policies & Cron-as-Code (Completed)

> **Current State**: Pranor Chrono implements per-job HTTP callback scheduling with 5-field cron expressions, distributed leader election locking, Pranor Pulse fan-out integration, failure counting, and audit logging.
> **What is Missing**: Multi-step DAG job pipelines (Job A triggers Job B on success/failure), configurable per-job retry backoff strategies, failure alert webhooks (Slack/PagerDuty), declarative YAML cron-as-code file loading, timezone-aware scheduling, and a visual execution timeline in Pranor Console.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| CR.G1 | **DAG Job Chain Pipeline (Job-A → Job-B on Success/Failure)** | Pranor Chrono DAG | Extend the Job model with `OnSuccess` / `OnFailure` successor job references to form directed acyclic graph (DAG) execution pipelines across multiple HTTP callback steps | [x] | OSS |
| CR.G2 | **Per-Job Retry Policy Engine (Exponential Backoff + Jitter)** | Pranor Chrono Retry | Add configurable per-job retry policies: max attempts, initial delay, exponential backoff multiplier, and random jitter; persist retry state across process restarts | [x] | OSS |
| CR.G3 | **Failure Alert Webhooks & Slack / PagerDuty Notification Integration** | Pranor Chrono Alerts | Fire configurable webhook notifications (Slack, PagerDuty, custom URL) when a job exceeds its failure threshold or a DAG pipeline encounters a terminal failure | [x] | OSS |
| CR.G4 | **Declarative YAML Cron-as-Code Definitions with Hot-Reload** | Pranor Chrono Config | Load job definitions from YAML configuration files (`jobs.yaml`) with file-watch hot-reload support; enables GitOps-style cron schedule management without API calls | [x] | OSS |
| CR.G5 | **Timezone-Aware Cron Scheduling (IANA Zone Support)** | Pranor Chrono Scheduler | Support per-job IANA timezone specification (e.g. `America/New_York`) so jobs fire at correct local times regardless of server timezone; persist timezone in job definition | [x] | OSS |
| CR.G6 | **Job Execution History Gantt Timeline in Pranor Console** | Pranor Console UI | Render a visual execution timeline in Pranor Console showing per-job run history, duration bars, status annotations, and failure counts with drill-down into audit log entries | [x] | OSS |

---







---

## Phase 57: Pranor Auth — Session Management, Passkeys & Adaptive MFA (Completed)

> **Current State**: Pranor Auth implements JWT issuance, bcrypt password hashing, JWKS key rotation, TOTP MFA, and OAuth social login.
> **What is Missing**: Opaque session tokens with server-side revocation (stateless JWTs cannot be invalidated without waiting for expiry), WebAuthn/Passkey FIDO2 support (no hardware key or platform biometric login), adaptive risk-based MFA step-up (no device fingerprinting or geo-anomaly detection), per-tenant OIDC federation (no Okta/Azure AD bring-your-own-IdP), and credential stuffing velocity detection.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| SA.G1 | **Opaque Session Token Store with Server-Side Revocation** | Pranor Auth Sessions | Replace stateless-only JWT approach with opaque refresh token store backed by Pranor Vault; enable instant per-session revocation without waiting for JWT expiry | [x] | OSS |
| SA.G2 | **WebAuthn / Passkey (FIDO2) Registration & Assertion** | Pranor Auth MFA | Implement FIDO2 WebAuthn authenticator registration and login assertion flow; enable hardware key (YubiKey) and platform passkey (TouchID/FaceID) authentication | [x] | OSS |
| SA.G3 | **Adaptive Risk-Based MFA Step-Up Engine** | Pranor Auth Risk | Analyze login signals (new device, unusual geo, time-of-day anomaly) and dynamically escalate to OTP or WebAuthn challenge mid-session without forcing full re-login | [x] | **EE** |
| SA.G4 | **Device Fingerprinting & Trusted Device Registry** | Pranor Auth Trust | Track device fingerprints (user-agent, screen entropy, timezone) and maintain per-user trusted device list with one-click revocation from Pranor Console | [x] | **EE** |
| SA.G5 | **Per-Tenant OIDC Provider Federation (Okta, Azure AD, Google Workspace)** | Pranor Auth Federation | Allow enterprise tenants to bring their own OIDC/SAML identity provider; auto-federate external group claims into Pranor Auth roles without code changes | [x] | **EE** |
| SA.G6 | **Credential Stuffing Detection & Velocity Rate Limiter** | Pranor Auth Security | Track failed login attempts per IP and username using sliding window counters; auto-block credential-stuffing bots with progressive CAPTCHA challenge escalation | [x] | OSS |

---

---

## Phase 58: Pranor Cache — Distributed Cluster Mode, Bloom Filters & Tiered TTL (Completed)

> **Current State**: Pranor Cache implements a single-node in-memory LRU cache with TTL eviction, `DeletePattern` glob invalidation, singleflight coalescing, and an HTTP REST API.
> **What is Missing**: Multi-node Raft-replicated cluster mode (currently single point of failure), RESP3 wire compatibility for Redis drop-in usage, Bloom filter for absent-key elimination, tiered TTL policies (hot/warm/cold), and cache stampede protection beyond singleflight.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| SC.G1 | **Multi-Node Raft-Based Distributed Cache Cluster** | Pranor Cache Cluster | Extend single-node InMemoryCache to a Raft-replicated distributed cluster; support 3-node quorum writes with consistent reads and automatic leader election | [x] | **EE** |
| SC.G2 | **RESP3 Protocol Wire Compatibility (Redis Drop-in Mode)** | Pranor Cache Protocol | Implement Redis Serialization Protocol v3 (RESP3) wire compatibility so Pranor Cache can function as a Redis drop-in replacement for existing application codebases | [x] | OSS |
| SC.G3 | **Probabilistic Bloom Filter for Absent-Key Elimination** | Pranor Cache Filter | Embed a Bloom filter in front of the LRU lookup path to short-circuit backend fetches for keys statistically absent from the cache; reduces load by eliminating redundant origin queries | [x] | OSS |
| SC.G4 | **Tiered TTL Policy Engine (Hot / Warm / Cold Tiers)** | Pranor Cache Policy | Implement multi-tier TTL policies: short-TTL hot tier (sub-second), medium-TTL warm tier (minutes), long-TTL cold tier (hours); auto-promote/demote entries based on access frequency | [x] | OSS |
| SC.G5 | **Cache Stampede Protection via Probabilistic Early Expiry** | Pranor Cache Resilience | Implement probabilistic early expiry (PER algorithm) to proactively recompute cache values before expiry under high concurrency, eliminating thundering herd spikes | [x] | OSS |
| SC.G6 | **Real-Time Hit Rate & Eviction Metrics Dashboard in Pranor Console** | Pranor Console UI | Stream live per-namespace hit rate, eviction rate, and memory pressure metrics into the Pranor Console dashboard with configurable alert thresholds | [x] | OSS |

---


---

## Phase 60: Pranor Flow — Visual Workflow Designer, WASM Step Functions & Human Tasks (Completed)

> **Current State**: Pranor Flow implements a DAG workflow engine with task dependency resolution, time-travel replay snapshots, and Pranor Pulse topic integration for async task dispatch.
> **What is Missing**: A visual drag-and-drop workflow designer in Pranor Console, WASM-compiled step function execution (currently HTTP-callback-only), human approval task gates with async wait, sub-workflow composition (nested workflow invocation), and per-execution cost and OTel span attribution.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| SF.G1 | **Visual Drag-and-Drop Workflow Designer in Pranor Console** | Pranor Console UI | Build a canvas-based interactive workflow designer where developers can visually create, connect, and configure workflow task nodes without writing JSON/YAML definitions | [x] | OSS |
| SF.G2 | **WASM Step Function Execution (Inline Compute per Task Node)** | Pranor Flow Engine | Allow individual workflow task steps to execute inline WebAssembly (WASM) bytecode rather than requiring remote HTTP endpoint callbacks; enables serverless compute-near-orchestration | [x] | OSS |
| SF.G3 | **Human Approval Task Gate (Async Pause + UI Approve/Reject)** | Pranor Flow Tasks | Implement a `human-task` step type that pauses workflow execution and presents an approval UI (email link or Pranor Console panel) to designated reviewers before continuing | [x] | **EE** |
| SF.G4 | **Sub-Workflow Composition & Nested Workflow Invocation** | Pranor Flow Composition | Enable workflow task steps to invoke other named workflow definitions as sub-workflows, creating hierarchical multi-level pipeline structures with isolated instance tracking | [x] | OSS |
| SF.G5 | **Per-Execution OTel Span Attribution & Cost Tracking** | Pranor Flow Telemetry | Inject W3C `traceparent` context into each task step HTTP call; export OTel spans to Pranor Trace for full distributed tracing of workflow runs including step latencies | [x] | OSS |
| SF.G6 | **Dead Letter Workflow Queue & Manual Retry from Pranor Console** | Pranor Flow DLQ | Capture permanently-failed workflow instances in a Dead Letter Queue; allow operators to inspect failure context and manually trigger selective re-execution from Pranor Console | [x] | OSS |

---
