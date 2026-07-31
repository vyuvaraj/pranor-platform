# Unified Roadmap - Completed Phases 76 to 80

## Phase 76: Pranor Compiler & Developer Experience Hardening (Completed)

> **Current State**: Pranor compiles `.serv` files into Go code, features an LSP server, WASM playground, multi-file imports, and test runner. However, error reporting lacks precise caret positions, codegen produces generic `interface{}` types, and database migrations lack state tracking.
> **What is Missing**: 10 targeted developer experience enhancements covering compiler error line/col carets, human-readable Go codegen with `.serv` line annotations, typed Go codegen, `.serv` source-mapped runtime errors, migration state tracking (`_serv_migrations`), multi-file LSP go-to-definition, non-crashing watch reloader, per-test DB transaction isolation, and WASM playground side-by-side Go code view.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| **Pranor DX & Compiler Hardening** | | | | | |
| SL.H1 | **Readable Generated Go Code & Source Line Annotations** | Pranor Codegen | Generate human-readable handler/cron names (`handleGET_api_tasks`) and inject `// line <file>:<line>` source map comments into generated Go code | [x] | OSS |
| SL.H2 | **Type-Aware Go Code Generator** | Pranor Codegen | Emit real Go types (`int`, `string`, structs) instead of generic `interface{}` to eliminate runtime type assertions and improve performance | [x] | OSS |
| SL.H3 | **Runtime Errors with `.serv` Source Context** | Pranor Runtime | Wrap runtime panic recovery and DB/HTTP library errors with `.serv` file, line, and route context instead of raw Go stacktraces | [x] | OSS |
| SL.H4 | **Migration State Tracking & Rollback Engine** | Pranor Storage | Track applied migrations in `_serv_migrations` table, prevent duplicate `ALTER TABLE` runs, and support `up {}` / `down {}` migration blocks | [x] | OSS |
| SL.H5 | **Workspace-Wide LSP Go-to-Definition & References** | Pranor Tooling | Extend LSP to jump across imported files, perform multi-file `Find-All-References`, and propagate symbol renames | [x] | OSS |
| SL.H6 | **Non-Crashing Hot Reload with Inline Error Diagnostics** | Pranor CLI | Keep previous good binary serving traffic when `pranor run --watch` encounters compilation errors, pushing diagnostics inline to terminal and LSP | [x] | OSS |
| SL.H7 | **Isolated Per-Test Transaction & Rollback Engine** | Pranor Testing | Wrap each `test` block in an isolated database transaction that automatically rolls back, and scope `mock` statements to individual test blocks | [x] | OSS |
| SL.H8 | **WASM Playground Side-by-Side Go Output & Live Server** | Pranor Web | Upgrade web playground to show generated Go code side-by-side, render debounced syntax errors, and execute in-browser HTTP requests | [x] | OSS |
| SL.H9 | **`serv.toml` Package Dependency & Resolution Engine** | Pranor Package | Add `serv.toml` package declaration manifest and `serv add <pkg>` command with clear resolution order (local -> stdlib -> registry) | [x] | OSS |
| SL.H10 | **Precise AST Token Position & Source Caret Reporter** | Pranor Parser | Propagate lexer line/column positions to AST nodes and format compiler errors with source code lines and caret pointers (`^`) | [x] | OSS |

---

## Phase 77: End-to-End Integration Testing & Conformance Suites (Completed)

> **Goal**: Prove the system works as advertised with automated conformance tests that run in CI and produce publishable badges.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| IT.1 | **S3 Conformance Test Suite (Mint-compatible)** | Pranor Vault | Run the open-source Mint S3 conformance suite against Pranor Vault. Publish pass rate as CI badge. Target: 90%+ pass rate. | [x] | OSS |
| IT.2 | **Cross-Component Integration Test Harness** | pranor-repo | Automated test that starts all 15 services via `pranor up`, runs a full user journey (deploy `.pnr` app → hit API → see trace → check queue → verify cache), assert all pass. Run in CI on every push. | [x] | OSS |
| IT.3 | **STOMP/Kafka/MQTT Protocol Conformance Tests** | Pranor Pulse | Run standard protocol test suites: Apache ActiveMQ STOMP tests, Kafka protocol decoder tests, Eclipse Paho MQTT test suite. Publish conformance percentages. | [x] | OSS |
| IT.4 | **Load Test Baseline with Published Results** | Pranor Vault, Pranor Gate, Pranor Pulse | Run `servstore bench`, `servgate` performance tests, `servqueue benchmark` in CI. Publish baseline numbers in README badges. Detect regressions automatically. | [x] | OSS |
| IT.5 | **Upgrade Path Smoke Test (v0.x → v1.0 data migration)** | All | Automated test: create data with current version, upgrade binary to next version, verify data still accessible. Catches breaking changes in storage format. | [x] | OSS |
| IT.6 | **Chaos Engineering Regression Suite** | Pranor Console, Pranor Mesh | Inject network partitions, kill random services, verify the system recovers within SLA. Run weekly in CI. Publish mean-time-to-recovery. | [x] | OSS |

---

