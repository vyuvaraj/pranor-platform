# Serv Unified Ecosystem Roadmap - Completed Items (Phases 11-15)

This document preserves the archived history of completed items for Phases 11 through 15.

---

## Phase 11: Next-Level Component Hardening & Ecosystem Depth (Completed Items)

### 🏗️ Structural Debt — Monolith Decomposition
- **Pranor Console real decomposition (SD.1)** — Extracted reverse proxies, Websockets, and AI metrics to packages inside `pkg/`.
- **Pranor Auth package extraction (SD.2)** — Split store adapters, MFA verify, and OAuth validator into `pkg/handlers/`, `pkg/store/`, `pkg/oauth/`, `pkg/mfa/` with proper interfaces.
- **Pranor Hub package split (SD.3)** — Extracted semver resolvers and package index structures into `pkg/registry/`, `pkg/resolution/`, `pkg/web/`.
- **Pranor Flow package extraction (SD.4)** — Split DAG engine, API handlers, saga execution, and checkpoint logic.
- **Pranor Notify package extraction (SD.5)** — Split delivery, templates, and storage into `pkg/` packages.

### 🔐 Security Gaps
- **JWT Key Rotation via JWKS (SEC.S1)** — Replace single shared `PRANOR_JWT_SECRET` with RS256 keypair + `/.well-known/jwks.json` endpoint. Enable rotation without service restarts.
- **Secret Redaction in Logs (SEC.S2)** — Robust regex redaction of quoted/unquoted credentials (`SanitizeLog()`).
- **Secret Versioning in KMS (SEC.S3)** — Fallback decrypt across active KMS key versions.
- **Audit Event Coverage Enforcement (SEC.S4)** — EmitAuditEvent calls enforced and lint-checked.

### 🧪 Testing & Quality Gaps
- **ServDocs test suite (TQ.1)** — Table-driven OpenAPI tests for parser, generator, and OpenAPI output validation.
- **Pranor Pool migration.go real implementation (TQ.2)** — Real migration executor, table tracking, and rollback.
- **Pranor Flow .state file gitignore (TQ.3)** — Cleaned committed state files and ignored `.state` files.
- **Property-based tests for critical paths (TQ.4)** — Added property-based fuzz test for token verification, S3 signature verification, and encryption/decryption roundtrips.
- **Load test baselines for all services (TQ.5)** — Added `load_test_baseline.go` script with SLA validations.

### 📦 Missing Infrastructure
- **ServDocs Dockerfile (INF.1)** — Multi-stage builder containerization.
- **ServDocs CI pipeline (INF.2)** — Actions build/test workflow added.
- **Pranor Core README (INF.3)** — Added comprehensive readme guide.
- **Pranor Deploy roadmap cleanup (INF.4)** — Duplicate Phase 3 headings cleaned up.
- **Unified Makefile/Taskfile (INF.5)** — Unified Makefile orchestrating all services builds/tests.
- **Dependency version pinning (INF.6)** — Aligned Pranor Core versions across workspace go.mod files.

### 🔗 Integration Depth
- **Pranor Console topology auto-discovery (INT.1)** — Parse OTel trace spans to auto-build service dependency graph.
- **Pranor → Pranor Auth native keyword (INT.2)** — Support `servauth://` connection string with native APIs.
- **Pranor → Pranor Pool proxy keyword (INT.3)** — `database "Pranor Pool://"` routes through Pranor Pool pooler.
- **Pranor → Pranor Notify notify keyword (INT.4)** — Support `notify "servmail://"` with `notify.send()` API.
- **Pranor Pulse stream processing DSL (INT.5)** — `stream "orders" |> filter(...) |> window(5m) |> count()`.
- **Pranor Chrono → Pranor Pulse job chaining (INT.6)** — Trigger next job by publishing to topic on completion.

### 🛠️ Developer Experience
- **`serv cache inspect` CLI (DX.S1)** — Show per-namespace key counts, hit/miss ratios, top hot keys.
- **`servqueue tail` CLI (DX.S2)** — Stream live topic messages with JSON pretty-print and regex filter.
- **`serv trace search` CLI (DX.S3)** — Search traces with JSON or ASCII waterfall outputs.
- **`serv tunnel inspect` CLI (DX.S4)** — Expose active tunnels, throughput, recent request logs.
- **`serv cron list` CLI (DX.S5)** — List job details, consecutive failure count, next 5 projected runs.
- **Pranor Notify local mock dev server (DX.S6)** — Consolidate SMTP and HTTP mail mocks.
- **`serv auth inspect` CLI (DX.S7)** — Show registered clients, active sessions, expired rate limits.
- **`serv docs preview` CLI (DX.S8)** — Spin up local server rendering ServDocs interface.

### ⚡ Performance & Reliability
- **Store chunked multipart upload (PR.S1)** — Enforce chunked streaming for uploads >10MB.
- **Queue batch dequeue optimizations (PR.S2)** — Support batch pulling of messages.
- **Cache connection multiplexing (PR.S3)** — Share connection TCP pools.
- **Mesh circuit breaker trip levels (PR.S4)** — Dynamic backoffs and connection dropouts.
- **Trace compression buffer (PR.S5)** — Compress historical spans using zlib before storage.

### 📝 Documentation & Hygiene
- **Ecosystem architectural book (DOC.1)** — Authored `ARCHITECTURE.md` monorepo blueprint.
- **CLI command reference list (DOC.2)** — Published clean `COMMANDS.md` detailing all CLI options.
- **API status checklist matrix (DOC.3)** — Added detailed maturity dashboard showing API compliance status.
- **Clean up orphan tests (DOC.4)** — Unified all temporary unit test code under proper test files.

### 11.1 Integration Completeness
- **Full 15-service discovery (UC.1)** — Add CLI flags + `ServDiscovery` fields for Pranor Mesh, Pranor Chrono, etc.
- **Unified health aggregation (UC.2)** — Health loop monitors all connected services.
- **Pranor Mesh panel (UC.3)** — Registry, circuit breaker states.
- **Pranor Chrono panel (UC.4)** — Scheduled jobs builder and history.
- **Pranor Cache panel (UC.5)** — Hit/miss ratios, key inspection.
- **Pranor Deploy panel consolidation (UC.6)** — Resource quotas proxy.
- **Pranor Hub panel (UC.7)** — Download stats, package warnings.
- **ServDocs embedding (UC.8)** — Render documentation in dashboard.

### 11.2 Cross-Service Intelligence
- **End-to-end request flow visualization (UC.9)** — Timeline visualization.
- **Ecosystem dependency matrix (UC.10)** — Runtime dependency map.
- **Unified configuration editor (UC.11)** — Central config panel.
- **Cross-service log correlation (UC.12)** — Trace_id logs correlation.
- **Ecosystem upgrade dashboard (UC.13)** — Version alignment checks.

### 11.3 Operational Intelligence
- **Capacity planning view (UC.14)** — Trend projections.
- **Change correlation engine (UC.15)** — Incident timeline overlays.
- **Service comparison mode (UC.16)** — Latency, CPU canary compare.
- **Ecosystem startup orchestrator (UC.17)** — Dependency boot orchestrator.
- **Unified API documentation portal (UC.18)** — Interactive spec portals.

### 11.4 AI-Powered Operations
- **AI root cause analysis (UC.19)** — Root cause hypotheses generator.
- **Natural language operations query (UC.20)** — Incident NL searches.
- **Predictive alerting (UC.21)** — Trend warnings.
- **Automated incident playbooks (UC.22)** — Pattern runbook trigger.

### 11.6 OSS/EE Feature Boundary Enforcement
- **Multi-tenant resource isolation (EE.10)**
- **Pranor Vault federation (EE.11)** (`federation_ee.go` / `federation_oss.go`)
- **Pranor Pulse federation (EE.12)** (`federation_ee.go` / `federation_oss.go`)
- **SLO tracking & error budgets (EE.13)** (`enterprise_ee.go` / `enterprise_oss.go`)
- **Cost estimation panel (EE.14)** (`enterprise_oss.go`)
- **Runbook automation (EE.15)** (`enterprise_oss.go`)
- **Custom dashboard builder (EE.16)** (`enterprise_oss.go`)
- **Infrastructure provisioning (EE.17)** (`enterprise_oss.go`)
- **Diagnostics terminal (EE.18)** (`enterprise_oss.go`)
- **Multi-environment management (EE.19)** (`enterprise_oss.go`)
- **Deployment rollback (EE.20)** (`enterprise_oss.go`)

---

## Phase 12: Dual-Licensing, Monetization, & Enterprise Separation (Completed Items)

### ⚖️ License & Policy Transition
- **Ecosystem CLA (Contributor License Agreement) (LIC.1)** — Drafted CLA.md and integrated CI checker.
- **License Re-assignment (v2.0.0+) (LIC.2)** — Transitioned all LICENSE files to AGPLv3.
- **Commercial License Terms (LIC.3)** — Authored EULA.md in pranor-repo.

### 📦 Codebase & Module Split
- **Private Enterprise Monorepo Setup (SPL.1)** — Initialized pranor-ee repository and premium plugins module.
- **Build Tag Integration (SPL.2)** — Extracted premium canary promotion engine logic to build-tagged source files.
- **Premium WASM Middleware Compilation (SPL.3)** — Implemented premium OIDC verification and PII redaction middleware packages in pranor-ee.
- **AI Diagnostics & Incident Panel Migration (SPL.4)** — Migrated out of public repos to private pranor-ee overlay.
- **GraphQL Schema Federation (SPL.5)** — Migrated out of public repos to private pranor-ee overlay.
- **Cost-Aware LLM Routing & Guardrails (SPL.6)** — Migrated out of public repos to private pranor-ee overlay.
- **Cold Storage Cloud Tiering (SPL.7)** — Migrated out of public repos to private pranor-ee overlay.

### 🚀 Enterprise Build Pipeline
- **Commercial CLI Builder (EE.1)** — Configured private CI pipeline with dynamic EE code overlays.
- **Licensed Artifact Verification (EE.2)** — Cryptographic HMAC-SHA256 license check enforced on Enterprise panels startup.

### 🛠️ Developer Experience
- **Unified CLI subcommands (DX.17)**
- **Multi-service compose generation (DX.19)**
- **`serv bench` load testing (DX.20)**
- **Schema-first development (DX.21)**
- **Observability-as-code (DX.23)**
- **`serv playground` IDE (DX.24)**
- **Cross-service config propagation (DX.25)**
- **`pranor dev` terminal dashboard (DX.26)**

---

## Phase 13: Language & Runtime Evolution (Completed Items)

- **Sum types / tagged unions (LANG.1)**
- **Native GraphQL support (LANG.2)**
- **Compiler plugin system (LANG.3)**
- **`serv migrate` rollback (LANG.4)**
- **Distributed lock primitive (LANG.5)**
- **Hot module replacement for stdlib (LANG.6)**
- **Security scanning (LANG.7)**
- **Interface satisfaction checking (LANG.8)**
- **Native Service Mesh Declarations (LANG.9)**
- **First-Class Storage Buckets (LANG.10)**
- **Declarative Event Handlers (LANG.11)**
- **Native Distributed Locking (LANG.12)**
- **Integrated Gateway Routing (LANG.13)**
- **Declarative Scheduled Workloads (LANG.14)**
- **Event Sourcing & CQRS Framework (CORE.9)**
- **Pranor Vault CDN mode (CORE.10)**

---

## Phase 14: AI-Native Ecosystem Deepening (Completed Items)

### Pranor (Compiler + Runtime)
- **RAG pipeline keyword (AI.10)** — `rag "servstore://docs" { embed: "openai", chunk: 512 }` declares retrieval-augmented generation as infrastructure. Auto-index on write, inject context on `ai.chat()`.
- **Structured output (JSON mode) (AI.11)** — `ai.complete(prompt, schema: UserSchema)` forces LLM responses to conform to a Serv struct. Compiler validates schema at build time.
- **Streaming responses (AI.12)** — `ai.stream(prompt, fn(chunk) { conn.send(chunk) })` for server-sent event streaming. Currently `ai.chat()` blocks until complete.
- **Prompt template library (AI.13)** — `import "stdlib/prompts"` with variable injection, versioning, and A/B testing hooks.
- **AI eval framework (AI.14)** — `test "quality" { assert ai.eval(prompt, expected, threshold: 0.8) }` for LLM output quality testing in `pranor test`.

### Pranor Gate (AI Gateway)
- **Token budget enforcement per route (AI.15)** — Max tokens/cost per API key per day. Reject when exhausted. Dashboard burn rate.
- **Prompt versioning & A/B routing (AI.16)** — Route % of traffic to different system prompts. Track outcome quality per version.
- **Response quality scoring (AI.17)** — Auto-score LLM responses for hallucination risk via factual grounding check against RAG context.
- **Multi-model fallback chain (AI.18)** — `models: [gpt-4o-mini, gpt-4o, claude-3-5-sonnet]` — try next on failure/timeout.
- **Semantic rate limiting (AI.19)** — Rate limit by semantic similarity of requests, not just IP. Prevent same question rephrased 100 ways.

### Pranor Vault (AI Storage)
- **Conversational object query (AI.20)** — `GET /bucket?ask=<question>` generates embedding, searches, synthesizes answer (RAG on stored objects).
- **Auto-summarization on upload (AI.21)** — Generate and store summaries as metadata on document upload. Browse-by-summary without downloading.
- **Similarity deduplication (AI.22)** — On upload, check if semantically similar document exists (cosine > 0.95). Warn or auto-deduplicate.
- **Classification tags on ingest (AI.23)** — Auto-classify uploaded objects (invoices, contracts, logs, images) via lightweight model. Searchable tags.

### Pranor Pulse (AI Messaging)
- **Semantic message routing (AI.24)** — Route messages to subscribers based on content meaning: `subscribe "support" where ai.classify(msg) == "billing"`.
- **DLQ auto-summarization (AI.25)** — When messages pile up in DLQ, generate summary.
- **Message pattern anomaly detection (AI.26)** — Learn normal throughput patterns. Alert on significant volume/content shifts.

### Pranor Console (AI Operations)
- **Natural language log search (AI.27)** — "Show errors from Pranor Auth where users couldn't login" → structured log query + filters.
- **Incident root cause synthesis (AI.28)** — On alert: analyze deploys, config changes, correlated metrics, similar past incidents. One-paragraph hypothesis.
- **Auto-generated runbooks (AI.29)** — Observe how operators respond to recurring alerts. After 3 manual resolutions, suggest automated runbook.
- **Anomaly explanation (AI.30)** — When metric spikes, explain why.

### Pranor Auth (AI Security)
- **Adaptive risk scoring (AI.31)** — Score login attempts: new device + unusual time + different geo = high risk → step-up to MFA.
- **Credential stuffing detection (AI.32)** — Behavioral clustering to detect many IPs using same password list. Auto-block suspicious cohorts.

### Pranor Trace (AI Observability)
- **Auto-correlate slow spans (AI.33)** — Identify root cause span and explain: "95% latency in Pranor Pool query — missing index on order_date".
- **Predictive SLO breach (AI.34)** — Given current error rate trajectory, predict when SLO will be violated. "Error budget exhausted in 3 days".

### Pranor Chrono & Pranor Flow (AI Automation)
- **Smart scheduling (AI.35)** — Analyze job execution history (duration, resource usage, conflicts). Suggest optimal scheduling windows.
- **AI decision steps in workflows (AI.36)** — `step "classify" { ai.classify(input, ["approve", "review", "reject"]) }` — AI-powered branching.
- **Workflow generation from description (AI.37)** — NL prompt → DAG definition: "receives order → validates payment → notifies warehouse → sends email".

---

## Phase 15: Component Backlog & Future Enhancements (Completed Items)

### 📄 ServDocs
- **Middleware chain documentation** — Show which middleware applies to which routes with order. [July 9, 2026]
- **Code examples in docs** — Include `.pnr` usage examples alongside route documentation. [July 9, 2026]
- **Versioned docs** — Generate docs per git tag; host multiple versions side-by-side. [July 9, 2026]
- **`serv docs serve --watch`** — File watcher that regenerates docs on `.pnr` file changes. Integrated with a build server that injects a Server-Sent Events (SSE) live-reload script into the generated HTML. [July 9, 2026]
- **Type Schema Rendering** — Renders struct definitions as expandable schema tables directly within route documentation dynamically using a pre-serialized JSON schema map. [July 9, 2026]
- **Client-side Full-text Search** — Implements an advanced, high-performance client-side search indexing system with real-time matching, text highlighting via `<mark>` elements, dynamic result counts, and automated sidebar navigation filtering. [July 9, 2026]

### 🛡️ Pranor Gate
- **GitOps Config Sync** — Pulls configuration updates automatically via webhook, validating route rules and signature authentication (supporting GitHub signature and GitLab tokens). [July 9, 2026]
- **Auto TLS Let's Encrypt** — Provisions certs automatically for comma-separated list of domains with a self-signed fallback mechanism for local development. [July 9, 2026]

### 🌐 Pranor Mesh
- **Rate Limiting per Service Pair** — Client-side token-bucket rate limiting keyed on `caller/callee` identity; returns HTTP 429 on exhaustion. [July 9, 2026]
- **Service Versioning & Header Routing** — Route requests to specific versions via `X-Service-Version` header; falls back to full pool if version not found. [July 9, 2026]
- **Health-Aware Load Balancing** — OTel-feedback weighted routing; clients push metrics via `POST /api/health-metrics`; `GET /api/topology` exposes aggregated health state per instance. [July 9, 2026]
- **gRPC Support** — Extend the resolver and circuit breaker to handle gRPC connections natively. [July 9, 2026]
- **Local Dev Service Mesh** — One-command `serv mesh up` that starts a local registry + resolver with zero config for fast developer iteration. [July 9, 2026]
- **Mesh Topology CLI** — `go run ./cmd/inspect/` shows live table of services, instances, latency, error rate, and health state. Supports `--watch` auto-refresh and `--service` filter. [July 9, 2026]

### 📦 Pranor Hub
- **Private Namespace Support** — Scoped package namespaces (`@org/package`) with access control lists. [July 9, 2026]
- **Mirror & Offline Cache** — Local proxy mode that caches the public registry to a Pranor Vault bucket; enables air-gapped builds. [July 9, 2026]
- **Provenance Attestation** — Record build provenance (commit SHA, CI run ID, builder identity) alongside the package; verify with `serv verify --attestation`. [July 9, 2026]

### 🔍 Pranor Trace
- **Trace Sampling Strategies** — Head-based and tail-based sampling with configurable rates per service. [July 9, 2026]
- **Trace Comparison** — Compare two traces side-by-side to identify regression causes. [July 9, 2026]
- **Retention Policies** — Configurable TTL per service. Auto-archive old traces to Pranor Vault. [July 9, 2026]
- **Distributed Context Baggage** — Propagate custom key-value pairs across service boundaries via trace context. [July 9, 2026]
- **Continuous Profiling Integration** — Link pprof CPU/memory profiles to trace spans; surface hot-path profiles in the Pranor Console waterfall view. [July 9, 2026]
- **Adaptive Sampling Rate** — Dynamically raise sampling rate when error rate spikes and lower it when traffic is healthy. [July 9, 2026]

### 🚇 Pranor Tunnel
- **Multi-relay federation** — Distribute tunnels across regions. [July 9, 2026]
- **Usage analytics and billing integration** — Integrated usage tracking. [July 9, 2026]
- **Enterprise features** — SSO, audit logging, IP allowlists. [July 9, 2026]
- **Team Collaboration** — Share tunnel access with team members via token-based invite links. [July 9, 2026]
- **Persistent Tunnels** — Keep tunnels alive across client restarts with session resumption. [July 9, 2026]
- **Custom Domain Mapping** — Map production domains to local tunnels for realistic testing. [July 9, 2026]
- **Request Recording & Replay** — Record all requests through tunnel, replay them later for debugging. [July 9, 2026]
- **Bandwidth Throttling** — Simulate slow networks (3G, satellite) for mobile testing. [July 9, 2026]
- **Request Diff Mode** — Show a colored diff between the proxied request and original, highlighting header mutations, body modifications or injected WASM transforms. [July 9, 2026]
- **Tunnel Config-as-Code** — Declare tunnel rules in `.serv/tunnel.yaml` (name, auth, subdomain, filters). [July 9, 2026]

### 🗄️ Pranor Pool
- **Connection Draining** — Gracefully drains database connection pools during rolling deploys/SIGTERM. Implemented a shutdown sequence that stops background cleaner routines and waits for all active in-flight database queries to finish before closing connection pools. [July 9, 2026]
- **Multi-region Query Routing** — Routes `SELECT` queries to specific geo-local replica pools based on the `X-Region` request header (supporting e.g. `--region-replicas` CLI option), defaulting back to the standard replica pool if the region pool is missing. [July 9, 2026]

### 📥 Pranor Pulse
- **Topic Schema Linting** — Static analyzer in compiler verifying message payloads against topic schemas at compile time. [July 9, 2026]
- **Dead Letter Queue Inspector (DX.S9)** — Added `serv queue dlq inspect <topic>` command with payload preview, retry count, and error cause, supporting `--replay` flag to requeue messages back to their source topics. Added GET `/api/v1/topics/{topic}/dlq` endpoint to fetch & decode DLQ entries from the WAL. [July 9, 2026]
