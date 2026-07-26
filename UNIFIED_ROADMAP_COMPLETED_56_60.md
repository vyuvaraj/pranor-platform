# Unified Roadmap - Completed Phases 56 to 60

## Phase 56: ServGateway Enterprise WAF, Remote WASM Sync & OAuth2 Engine (Completed)

> **Context:** Commercial enterprise security and traffic management features for ServGateway.

| # | Item | Component | Description | Status |
|---|------|-----------|-------------|--------|
| SG.VE1 | **Enterprise Remote WASM Sync Engine (`serv-ee`)** | ServGateway EE | Cryptographically signed remote WASM plugin download & hot-reloading behind `//go:build enterprise` | [x] |
| SG.VE2 | **Enterprise WAF Ruleset & Threat Intelligence (`serv-ee`)** | ServGateway EE | OWASP Top 10 automated threat intelligence WAF engine behind `//go:build enterprise` | [x] |
| SG.VE3 | **Enterprise OAuth2 / OIDC Token Introspection (`serv-ee`)** | ServGateway EE | Distributed OIDC token validation & caching behind `//go:build enterprise` | [x] |
| SG.VE4 | **Enterprise Multi-Cloud Anycast Mesh Controller (`serv-ee`)** | ServGateway EE | Global Anycast BGP route steering behind `//go:build enterprise` | [x] |
| SG.VE5 | **ServGateway EE Modularization Verification (`serv-ee`)** | ServGateway EE | Strict build-tag isolation & enterprise package testing | [x] |

---



All commercial enterprise features (**EE**) must have their core logic and implementations located exclusively inside the private `servverse-ee` repository. 

The open-source core repositories (such as `ServGate`, `ServStore`, etc.) must only expose clean interfaces, hooks, or config fields. The implementation of these hooks in the open-source code must use build-tagged placeholders (`//go:build !enterprise`), while the actual commercial code resides under the corresponding directories in `servverse-ee` and is built with `//go:build enterprise`.












---

## Strategic Module Gap Analysis — Phases 57 to 72

> **Context**: The following phases are derived from a deep critical analysis of each Servverse module's current implementation against industry-standard production expectations. Each phase documents concrete missing features — not aspirational items — that are required for the module to compete as a standalone product and fulfil its role within the Servverse ecosystem. Phases are ordered by module dependency depth: standalone utility modules first, cross-cutting platform layers last.

---


## Phase 59: ServCron — DAG Job Chaining, Retry Policies & Cron-as-Code (Completed)

> **Current State**: ServCron implements per-job HTTP callback scheduling with 5-field cron expressions, distributed leader election locking, ServQueue fan-out integration, failure counting, and audit logging.
> **What is Missing**: Multi-step DAG job pipelines (Job A triggers Job B on success/failure), configurable per-job retry backoff strategies, failure alert webhooks (Slack/PagerDuty), declarative YAML cron-as-code file loading, timezone-aware scheduling, and a visual execution timeline in ServConsole.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| CR.G1 | **DAG Job Chain Pipeline (Job-A → Job-B on Success/Failure)** | ServCron DAG | Extend the Job model with `OnSuccess` / `OnFailure` successor job references to form directed acyclic graph (DAG) execution pipelines across multiple HTTP callback steps | [x] | OSS |
| CR.G2 | **Per-Job Retry Policy Engine (Exponential Backoff + Jitter)** | ServCron Retry | Add configurable per-job retry policies: max attempts, initial delay, exponential backoff multiplier, and random jitter; persist retry state across process restarts | [x] | OSS |
| CR.G3 | **Failure Alert Webhooks & Slack / PagerDuty Notification Integration** | ServCron Alerts | Fire configurable webhook notifications (Slack, PagerDuty, custom URL) when a job exceeds its failure threshold or a DAG pipeline encounters a terminal failure | [x] | OSS |
| CR.G4 | **Declarative YAML Cron-as-Code Definitions with Hot-Reload** | ServCron Config | Load job definitions from YAML configuration files (`jobs.yaml`) with file-watch hot-reload support; enables GitOps-style cron schedule management without API calls | [x] | OSS |
| CR.G5 | **Timezone-Aware Cron Scheduling (IANA Zone Support)** | ServCron Scheduler | Support per-job IANA timezone specification (e.g. `America/New_York`) so jobs fire at correct local times regardless of server timezone; persist timezone in job definition | [x] | OSS |
| CR.G6 | **Job Execution History Gantt Timeline in ServConsole** | ServConsole UI | Render a visual execution timeline in ServConsole showing per-job run history, duration bars, status annotations, and failure counts with drill-down into audit log entries | [x] | OSS |

---





