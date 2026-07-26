# Unified Roadmap - Completed Phases 57 to 65

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



