# Unified Roadmap - Completed Phases 71 to 75

## Phase 71: ServQueue — Schema Registry, DLQ Replay UI & Consumer Group Dashboard (Completed)

> **Current State**: ServQueue provides STOMP, MQTT, Kafka-protocol compatibility, WASM stream processing, CRDT geo-replication, OPFS browser queuing, Dead Letter Queues, ServConsole admin UI, and enterprise encryption.
> **What is Missing**: A Schema Registry for enforcing Avro/JSON Schema/Protobuf message contracts at the broker (currently no schema validation), a DLQ browser with one-click replay in ServConsole, per-consumer-group lag dashboards with partition-level offset inspection, Kafka-style log compaction, and W3C trace context propagation through messages.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| SQ.G1 | **Embedded Schema Registry for Message Contract Enforcement** | ServQueue Schema | Implement an embedded Schema Registry validating producer messages against registered Avro/JSON Schema/Protobuf contracts before routing; reject incompatible messages with structured schema validation errors | [x] | OSS |
| SQ.G2 | **DLQ Browser & One-Click Replay in ServConsole** | ServConsole UI | Render a DLQ inspector in ServConsole showing failed message payloads, error reasons, retry counts, and original topic metadata; allow operators to replay individual or bulk DLQ messages back into the source topic | [x] | OSS |
| SQ.G3 | **Consumer Group Lag Dashboard with Partition-Level Offset Inspection** | ServConsole UI | Display per-consumer-group, per-partition message lag (offset delta between latest produced and latest committed); surface groups falling behind their SLA processing rate with configurable lag alerts | [x] | OSS |
| SQ.G4 | **Topic Log Compaction Policy Engine (Key-Based Compaction & Retention)** | ServQueue Storage | Implement Kafka-style log compaction for topics with `cleanup.policy=compact`; retain only the latest record per message key; configurable time-based and size-based retention policies alongside compaction | [x] | OSS |
| SQ.G5 | **Per-Message W3C Trace Context Propagation (OTel traceparent)** | ServQueue Tracing | Inject W3C `traceparent` headers into each message at produce time; extract and continue the distributed trace on the consumer side; visualize full producer-to-consumer message flow in ServTrace | [x] | OSS |
| SQ.G6 | **Multi-Tab Browser OPFS Queue Leader Election via navigator.locks** | ServQueue OPFS | Implement browser multi-tab leader election using the Web Locks API (`navigator.locks`) so only one browser tab acts as OPFS queue primary; other tabs register as followers consuming via BroadcastChannel | [x] | OSS |

---







---

## Phase 72: Unified Platform — Single-Binary `servd`, WireGuard Mesh & Chaos Injection Engine (Completed)

> **Current State**: All Servverse modules are individually excellent standalone daemons. Operating a complete stack requires running 7+ separate processes with manual inter-service networking configuration.
> **What is Missing**: A unified single-binary runtime (`servd`) embedding all modules with zero-copy shared memory channels, automatic WireGuard cluster mesh for zero-trust inter-node networking, a cross-cutting Chaos Injection Engine for automated resilience validation, a unified cluster admin CLI (`servctl`), rollup health API, and official Docker Compose + Helm chart distribution.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| PL.G1 | **Single-Binary `servd` Unified Runtime (Embedded Monolith Mode)** | Platform Runtime | Build `servd` embedding ServGate, ServStore, ServQueue, ServCache, ServAuth, ServCron, and ServMesh; use zero-copy shared-memory channels for inter-module communication; scale from laptop dev mode to distributed cluster deployment | [x] | OSS |
| PL.G2 | **Automatic WireGuard Cluster Mesh Between `servd` Nodes** | Platform Network | Implement automatic WireGuard mesh key exchange between `servd` cluster peers using a DHT-based peer discovery protocol; establish encrypted kernel-level tunnels between all nodes without manual certificate provisioning | [x] | **EE** |
| PL.G3 | **Unified Chaos Injection Engine (Network, CPU, Memory, Disk, Clock Skew)** | Platform Chaos | Expose a Chaos Engineering API (`/api/v1/chaos`) on `servd` supporting: network latency/packet-drop injection, CPU stress, memory pressure, disk I/O throttling, and clock skew simulation; controlled per-node and per-service via `servctl` | [x] | OSS |
| PL.G4 | **`servctl` Cluster-Wide Administration CLI** | Platform CLI | Build `servctl` — a unified cluster administration CLI: `servctl cluster status`, `servctl chaos inject --service=servstore --type=latency --duration=60s`, `servctl deploy canary`, `servctl trace query --service=api` | [x] | OSS |
| PL.G5 | **Unified Health, Readiness & Rollup Metrics API** | Platform Observability | Expose `/api/v1/health` rollup on `servd` aggregating health status and key metrics across all embedded modules; consumable by Kubernetes liveness/readiness probes and external monitoring systems (Datadog, Grafana) | [x] | OSS |
| PL.G6 | **Official Docker Compose & Production Helm Chart Distribution** | Platform Distribution | Publish an official multi-service `docker-compose.yml` (local dev) and a production-grade `helm chart` deploying the complete Servverse stack to Kubernetes with a single `helm install servverse` command | [x] | OSS |

---

## Phase 73: VS Code Extension — Modern Ecosystem Alignment (Completed)

> **Current State**: `serv-vscode` (v3.3.0) provides LSP client support, syntax highlighting, basic snippets, and exploratory panels.
> **What is Missing**: Full syntax and snippet support for `async`/`concurrent` primitives, `servctl` CLI integration inside VS Code, `serv diff` breaking change detector integration, multi-target codegen context menus (Rust/Python), platform chaos injection controls, and direct ServConsole / WASM Playground deep-linking.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| VS.G1 | **`async` / `concurrent` Syntax & Snippet Alignment** | VSCode Syntax | Update TextMate grammar (`serv.tmLanguage.json`) and snippets to fully support `async fn`, `async` call expressions, and `concurrent {}` parallel blocks | [x] | OSS |
| VS.G2 | **`servctl` Cluster Administration Integration** | VSCode Commands | Add VS Code command palette and sidebar actions for `servctl`: list services/nodes, restart services, apply configuration, and view `servd` health rollups | [x] | OSS |
| VS.G3 | **`serv diff` Breaking Change Detector Command** | VSCode Tooling | Add `Serv: Check Breaking API Changes` command running `serv diff` against git base branch (`main`) with inline diagnostic markers for field removals and type changes | [x] | OSS |
| VS.G4 | **Multi-Target Client Code Generation Context Menu** | VSCode Codegen | Add right-click context menu options on `.srv` files: `Serv: Generate Rust Client` (`--lang rust`) and `Serv: Generate Python Client` (`--lang python`) | [x] | OSS |
| VS.G5 | **Platform Chaos Injection Sidebar Control Panel** | VSCode Chaos | Add a Chaos Control View in the ServVerse sidebar to quickly trigger/abort network delay, CPU stress, and disk throttle experiments across cluster nodes | [x] | OSS |
| VS.G6 | **WASM Playground & ServConsole Deep-Link Export** | VSCode Integrations | Add `Serv: Export to WASM Playground` (opens snippet in `playground.servverse.dev`) and `Serv: Open in ServConsole` deep-linking commands | [x] | OSS |
| VS.G7 | **`servd` Unified Runtime Webview & Component Management Panel** | VSCode Webview | Upgrade `serv-vscode` sidebar panel and webview suite to support `servd` embedded monolith mode: auto-detect `servd` status via `/api/v1/servd/components`, provide a unified multi-tab `servd` Console Webview, and execute component actions via `servctl` | [x] | OSS |

---

## Phase 74: Developer Adoption & High-Impact Differentiators (Completed)

> **Current State**: ServStore, ServGate, and ServQueue possess deep core features, but lack immediate 60-second time-to-value friction reducers (built-in benchmarking, migration import, SQL metadata queries, live terminal dashboards, plain HTTP/SSE & WebSocket pub/sub, embedded topic inspect UI, and scheduled messages).
> **What is Missing**: 16 specific high-impact differentiators across ServStore (`servstore bench`, `servstore import`, SQL metadata query API, zero-dependency webhooks, `servstore serve-static`), ServGate (upstream LLM dispatch for smart router, `servgate dashboard` TUI, OpenAPI request validation, response diff replay, per-route cost dashboard, `servgate generate-config` learn mode), and ServQueue (embedded web UI at `/ui/`, HTTP/SSE consumer API, cron-scheduled messages, `servqueue benchmark`, SQLite storage backend, plain WebSocket pub/sub).

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| **ServStore Differentiators** | | | | | |
| ST.D1 | **`servstore bench` Built-in Performance Benchmarking Command** | ServStore CLI | Add built-in CLI benchmark tool `servstore bench --ops 10000 --concurrency 16 --object-size 4KB` measuring PUT/GET/LIST ops/sec, P50, P95, and P99 latencies without external dependencies | [x] | OSS |
| ST.D2 | **`servstore import` One-Command Cloud Migration Mirror** | ServStore Migration | Add built-in S3 migration command `servstore import --from s3://my-aws-bucket` to mirror external AWS S3/GCS buckets into ServStore seamlessly | [x] | OSS |
| ST.D3 | **Object Metadata SQL Query API Endpoint (`GET /{bucket}?sql=...`)** | ServStore Query | Allow querying object metadata like a database directly via SQL expressions (`SELECT key, size, last_modified FROM objects WHERE size > 1048576 ORDER BY last_modified DESC`) | [x] | OSS |
| ST.D4 | **Zero-Dependency Webhook & Event Notifications** | ServStore Events | Simple REST API subscription for object creation/deletion webhooks (`{"bucket":"uploads","events":["s3:ObjectCreated:*"],"webhook":"https://myapp.com/hook"}`) without requiring SQS or Kafka | [x] | OSS |
| ST.D5 | **`servstore serve-static` Instant Static Site Hosting** | ServStore Web | Turn any bucket into a static website host (`servstore serve-static --bucket my-site --port 3000`) with proper MIME detection, index.html fallback, and 404 handling | [x] | OSS |
| **ServGate Differentiators** | | | | | |
| SG.D1 | **Upstream LLM Dispatch for Smart Router** | ServGate AI | Wire `smart_router.go` to execute real upstream HTTP completions to Ollama/OpenAI/Anthropic instead of returning stub mock responses | [x] | OSS |
| SG.D2 | **`servgate dashboard` Terminal Live Traffic Dashboard (TUI)** | ServGate CLI | Build a terminal-based live htop-style dashboard showing real-time RPS, active connections, circuit breaker state, cache hit rate, and P99 latency per route | [x] | OSS |
| SG.D3 | **OpenAPI-Aware Request Validation & Route Enforcement** | ServGate Validation | Validate incoming request bodies, query params, and headers against OpenAPI 3.0 spec (`openapi_spec: ./openapi.yaml`) before proxying | [x] | OSS |
| SG.D4 | **Request/Response Recording & Shadow Diff Replay** | ServGate Replay | Record live traffic, replay against a candidate backend version, and compute structural diffs highlighting changed API responses before production deployment | [x] | OSS |
| SG.D5 | **Per-Route AI Token & Cost Attribution Dashboard** | ServGate FinOps | Live cost tracking dashboard showing real-time spend per route, request count, and financial savings generated via semantic caching and smart routing | [x] | OSS |
| SG.D6 | **`servgate generate-config` Zero-Config Learn Mode** | ServGate Config | Run ServGate in learn mode to inspect unrouted traffic and auto-generate `config.json` with discovered routes, rate limits, and circuit breaker thresholds | [x] | OSS |
| **ServQueue Differentiators** | | | | | |
| SQ.D1 | **Standalone Embedded Web Management UI (`/ui/`)** | ServQueue UI | Lightweight browser UI embedded directly in `servqueued` at `/ui/` showing active topics, consumer lag, DLQ browser with one-click replay, and schema registry | [x] | OSS |
| SQ.D2 | **Plain HTTP/SSE Consumer Protocol (`/api/v1/subscribe/...`)** | ServQueue API | Long-polling Server-Sent Events (SSE) consumer endpoint allowing any HTTP client to subscribe to topic streams without STOMP/Kafka drivers | [x] | OSS |
| SQ.D3 | **Delayed & Cron-Scheduled Message Engine** | ServQueue Scheduling | First-class delayed and cron-scheduled message publishing API (`{"topic":"reports","schedule":"0 9 * * MON"}`) using internal TimeWheel engine | [x] | OSS |
| SQ.D4 | **`servqueue benchmark` Built-in Queue Throughput Test** | ServQueue CLI | Add built-in CLI benchmark tool `servqueue benchmark --messages 100000 --producers 8 --consumers 4` measuring msg/sec throughput and latency percentiles | [x] | OSS |
| SQ.D5 | **SQLite-Backed Persistent Storage Mode** | ServQueue Storage | Offer single-file SQLite storage backend for single-node deployments providing ACID durability and SQL queryability over message history | [x] | OSS |
| SQ.D6 | **Plain WebSocket Native Pub/Sub (`/ws/subscribe/...`)** | ServQueue Protocol | Native raw WebSocket pub/sub endpoint (`ws://localhost:9090/ws/subscribe/orders`) broadcasting JSON messages to frontend apps without STOMP framing | [x] | OSS |

---

## Phase 75: Production Hardening & Scale Bottleneck Fixes (Completed)

> **Current State**: ServStore, ServGate, and ServQueue are production-ready for small-to-medium workloads (<100 req/sec, <50 services), but exhibit key architectural bottlenecks when handling high-concurrency, multi-GB streaming uploads, deep AI vector indexing, or high-throughput consumer rebalancing.
> **What is Missing**: 12 critical scale & concurrency hardening items across ServStore (fine-grained per-bucket/object read-write lock striping, zero-memory-copy streaming `PutObject` pipeline, Quorum write confirmation in erasure coding), ServGate (vector-indexed semantic cache, NLP entity recognition PII scrubber, advanced prompt injection classifier), and ServQueue (disk-persisted consumer offset log, lock-free consumer slice dispatching, durable subscription state, binary byte slice payload framing).

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| **ServStore Concurrency & Memory Hardening** | | | | | |
| ST.H1 | **Fine-Grained Per-Bucket/Key Striped Lock Manager** | ServStore Core | Replace global `sync.RWMutex` with a striped bucket/key lock manager to eliminate global write serialization and allow parallel `PutObject` calls across different keys | [x] | OSS |
| ST.H2 | **Zero-Memory-Copy Streaming `PutObject` Pipeline** | ServStore Storage | Replace `io.ReadAll` in `PutObject` with a streaming temp-file disk pipeline (`io.Copy`) to handle multi-GB object uploads with a constant ~64KB memory buffer | [x] | OSS |
| ST.H3 | **Erasure Coding Quorum Write Confirmation Engine** | ServStore Erasure | Enforce synchronous Quorum write confirmations ($M+K$ shards verified) across cluster nodes before returning `200 OK` on erasure-coded uploads to prevent partial state corruption | [x] | OSS |
| ST.H4 | **High-Scale Prefix-Scan Iterator for `ListObjects`** | ServStore Index | Optimize PebbleDB key iteration for `ListObjects` with prefix seek bounds and keyspace caching for high-density buckets containing >1M objects | [x] | OSS |
| **ServGate AI Gateway Hardening & Accuracy** | | | | | |
| SG.H1 | **Vector-Indexed HNSW Semantic Cache Engine** | ServGate AI | Upgrade linear TF-IDF prompt cache scan to an HNSW vector index with LRU memory eviction to maintain sub-millisecond cache lookups at 100K+ cached prompts | [x] | OSS |
| SG.H2 | **Advanced Multi-Turn Prompt Injection Classifier** | ServGate AI | Replace 4-regex pattern matcher with a multi-turn contextual prompt injection classifier detecting encoded, indirect, and adversarial jailbreak payloads | [x] | OSS |
| SG.H3 | **NLP Entity-Aware PII Redaction Engine** | ServGate AI | Extend regex PII scrubber with NLP named-entity recognition (NER) for contextual redaction of personal names, addresses, and sensitive organizations | [x] | OSS |
| SG.H4 | **Smart Model Complexity Classifier Engine** | ServGate AI | Replace naive word-count heuristic in `smart_router.go` with a multi-feature prompt complexity scoring engine evaluating code syntax, reasoning depth, and context length | [x] | OSS |
| **ServQueue Durable Messaging & Concurrency Scale** | | | | | |
| SQ.H1 | **Disk-Persisted Consumer Offset Log & Acknowledgment Engine** | ServQueue Engine | Replace in-memory channel delivery with a disk-backed append-only consumer log and offset tracker to guarantee zero message loss on process crashes | [x] | OSS |
| SQ.H2 | **Lock-Free Broadcast & Subscriber Slice Dispatcher** | ServQueue Engine | Replace single-mutex `map[string][]chan string` with lock-free atomic subscriber channels to scale concurrent subscriber dispatching to 10,000+ topics | [x] | OSS |
| SQ.H3 | **Durable Consumer Subscription Reconnect Engine** | ServQueue Session | Persist consumer group offsets to disk so disconnected STOMP/MQTT/Kafka clients resume consumption seamlessly from their exact last offset | [x] | OSS |
| SQ.H4 | **Native Binary Payload (`[]byte`) Frame Pipeline** | ServQueue Protocol | Upgrade internal string message payloads to native `[]byte` slice framing to support Protobuf, Avro, and raw binary streams without base64 overhead | [x] | OSS |


