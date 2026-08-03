# Changelog

All notable changes to the Pranor ecosystem and background microservices are documented below.

---

## [v1.0.0] - Initial Production Release

### Rebrand & Ecosystem Alignment (Phase 81)
- Full rebrand from Serv/Pranor to **Pranor** across all 22 monorepo packages, Go imports, CLI binaries (`pranor`), and file extensions (`.pnr`).
- Unified single-binary daemon `pranord` embedding all 17 microservices.
- Custom `pranor://` URL protocol scheme for zero-trust service mesh routing.

### Developer Experience & Onboarding Automation (Phase 79)
- **`pranor quickstart`**: Interactive CLI setup wizard for instant service scaffolding.
- **`pranor doctor`**: Automated infrastructure diagnostic command for Go, Docker, environment vars, and service discovery verification.
- **Pranor Web Playground**: Zero-install browser IDE hosted at `playground.pranor.dev`.
- **Unified Error Code Registry**: Comprehensive 140+ error code reference (`docs/error_codes.md`) with fix suggestions.

### Ecosystem Connectors & Integrations (Phase 78)
- **`terraform-provider-pranor`**: Terraform Provider for declarative management of buckets, topics, cron jobs, and routes.
- **`pranor/deploy-action@v1`**: Official GitHub Action for CI/CD compilation and blue/green deployments.
- **Prometheus Remote Write Receiver**: Direct ingestion of Prometheus scrapers into `Pranor Trace`.
- **OpenTelemetry Collector Exporter**: OTLP/HTTP exporter pipeline support for external OTel collectors.
- **Grafana Datasource Plugin**: Visualizing Pranor Trace spans and Pranor Pulse queue metrics natively in Grafana.

### Language Server & VS Code IDE Intelligence (Phase 83 & 84)
- **Multi-File Workspace Rename**: `textDocument/rename` emitting `WorkspaceEdit` diffs across workspace `.pnr` files.
- **Workspace Symbol Indexer**: High-performance fuzzy symbol lookup (`workspace/symbol`).
- **Call Hierarchy Provider**: Incoming and outgoing call tree inspection (`textDocument/prepareCallHierarchy`).
- **Control Plane Extension (`pranor-vscode`)**: Built-in VS Code webview panels for Interactive API Client (Gate), Live Event Tailer (Pulse), S3/Vector Search Explorer (Vault), Flamegraph Log Viewer (Trace), Secret Manager, and Multi-Cluster Dashboard (Deploy).

### End-to-End Conformance & Reliability (Phase 76 & 77)
- **S3 Conformance Suite**: Verified Mint S3 compliance for Pranor Vault.
- **STOMP/Kafka/MQTT Conformance**: Protocol decoders validated against standard conformance test suites in Pranor Pulse.
- **Compiler Hardening**: Precise source caret pointers (`^`), typed Go codegen, and source-mapped `.pnr` runtime error stack traces.

---

## [v0.9.0] - Microservice Architecture & Core Modules

### Pranor Gate (API Gateway & Ingress Router)
- Edge HTTP/gRPC ingress routing with dynamic mTLS certificate rotation.
- Token bucket rate limiting per IP / API key.
- Sandboxed WebAssembly (Wazero) middleware execution.

### Pranor Pulse (Async Event Broker & Message Queue)
- Multi-protocol message broker supporting Kafka wire format, STOMP WebSockets, and MQTT.
- Automatic Dead Letter Queue (DLQ) isolation and 1-click replay API.

### Pranor Vault (S3 Storage & Vector Search Engine)
- AWS S3 API compatibility with multipart uploads and presigned URLs.
- Native HNSW vector similarity search engine (Cosine, Euclidean, Dot-product).

### Pranor Flow, Auth, Chrono, Mesh & Trace
- Durable saga workflow orchestrator with automated compensations (`Pranor Flow`).
- Identity & access management with JWKS, TOTP MFA, and social OAuth (`Pranor Auth`).
- Scheduled cron task runner with leader election (`Pranor Chrono`).
- Zero-trust in-memory service discovery mesh (`Pranor Mesh`).
- OTLP distributed tracing and flamegraph collector (`Pranor Trace`).
