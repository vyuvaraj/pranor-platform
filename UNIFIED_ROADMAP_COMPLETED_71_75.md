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
