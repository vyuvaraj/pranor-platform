# Changelog

All notable changes to the Pranor background service ecosystem will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [1.8.0] - 2026-07-17

### Added
* **CLI.1**: Implemented `serv changelog` command with version and service filters to display the ecosystem evolution history.
* **CLI.2**: Added `--attach` context verification capability to the interactive REPL.
* **LINT.1**: Implemented cross-service dead route static checks querying the Pranor Mesh registry.
* **OPS.16**: Added visual message flow timeline tracing and auto-postmortem incident report generation endpoints to `Pranor Console`.
* **PKG.12**: Verified and finalized installer package configurations (Inno Setup, macOS PKG, Snap, MSIX, Chocolatey, and Winget).

## [1.7.0] - 2026-07-02

### Added
* **SEC.S2**: Enhanced Structured Log Sanitizer. Upgraded `SanitizeLog` in `Pranor Core` to redact quoted/unquoted credentials matching JSON structures.
* **TQ.1**: Added unit test coverage for `ServDocs` verifying parser, HTML generator, and OpenAPI generator logic.
* **TQ.3**: Configured `*.state` gitignore rules and cleaned dangling state files in `Pranor Flow`.
* **INF.1**: Added multi-stage builder `Dockerfile` to `ServDocs`.
* **INF.2**: Configured GitHub Actions CI pipeline for `ServDocs`.

## [1.6.0] - 2026-07-01

### Added
* **OPS.14**: Enterprise Control Plane (Tenant Routing). Implemented dynamic tenant policies configuring allowed regions/clusters under `/api/tenants/policies` in `Pranor Gate`, automatically checking incoming `X-Tenant-ID` header parameters.
* **OPS.15**: Production Digital Twin (Sandbox Config Generator). Built AST scanner and regex fallback generator inside `Pranor` (`serv generate sandbox <file>`) to produce a local sanitized digital twin configuration.
* Added validation test suites: `TestTenantControlPlanePolicies` in `Pranor Gate` and `TestSandboxConfigGeneration` in `Pranor`.

## [1.5.0] - 2026-07-01

### Added
* **SEC.16**: Zero-Trust mTLS Network Policies. Implemented `NetworkPolicy` registry configurations and gRPC client CN certificate identity checking in `Pranor Mesh` to reject unauthorized mesh access requests.
* **DX.10**: Scaffolding CLI & Sandbox. Expanded `serv generate` CLI commands in the `Pranor` compiler to scaffold standard REST APIs (`api`), database schemes (`db`), and distributed workflow patterns (`workflow`).
* **OPS.12**: Automated Canary Deployment Engine. Added in-memory error stats tracking and background auto-promotion / auto-rollback routines in `Pranor Gate`. Rolling canary promotions automatically revert weights to stable target (0% canary weight) if error rates exceed 5%.
* Added validation test suites: `TestZeroTrustNetworkPolicies` in `Pranor Mesh`, `TestScaffoldingCLI` in `Pranor`, and `TestAutomatedCanaryRollback` in `Pranor Gate`.

## [1.4.0] - 2026-07-01

### Added
* **HA.1**: Dynamic Active-Active Cluster Replication. Implemented object replication in `Pranor Vault` with LWW conflict resolution, and mutating query statement replication across cluster peers in `ServDB`.
* **PS.4**: Internal gRPC Mesh Transport. Implemented a zero-compile gRPC JSON-codec transport in `Pranor Mesh` enabling inter-service communication to execute over high-performance HTTP/2 multiplexed gRPC connections.
* Added a unit test suite validating Active-Active object conflict resolution (`TestS3ActiveActiveConflictResolution` in `Pranor Vault`), database statement replication (`TestDatabaseQueryReplication` in `ServDB`), and gRPC mesh forwarding (`TestGRPCMeshTransport` in `Pranor Mesh`).

## [1.3.0] - 2026-07-01

### Added
* **ARCH.5**: Shared package extraction and strict constructor dependency injection in `Pranor Notify`. Handlers are now structured as methods on `MailServer`.
* **DX.9**: Offline mock mode support. Adds a concurrent TCP-based mock SMTP server (listening on port 1025) in `Pranor Notify` and offline S3 mock API mode (activated via `--mock` or `SERVSTORE_MOCK=true`) in `Pranor Vault`.
* **OPS.5**: GitOps configuration sync webhook endpoint `/api/gitops/webhook` and `/api/v1/gitops/webhook` in `Pranor Gate` to pull changes and dynamically reload routes.
* **OPS.6**: Integrated ACME / Let's Encrypt autocert client in `Pranor Gate` supporting port 80 HTTP-01 challenge redirect and automated certificate renewals on port 443.
* **CORE.2**: Durable Sagas rollback engine in `Pranor Flow` which executes compensation actions (with support for HTTP endpoints), updates intermediate statuses to `"compensating"`, and durably checkpoints state so rollbacks resume on startup.
* **OPS.11**: Performance Regression CI Gates. Integrated PR benchmark gating workflow ([perf-gates.yml](file:///c:/Mine/try/serv/pranor-repo/.github/workflows/perf-gates.yml)) and SLA validators ([verify_perf_sla.py](file:///c:/Mine/try/serv/pranor-repo/scripts/verify_perf_sla.py)) verifying latency (<20ms) and error margins.
* **CORE.3**: Asynchronous Event-Driven Sagas. Implemented STOMP messaging-based saga compensation notifications over `Pranor Pulse` topics and a REST continuation callback API in `Pranor Flow`.
* **CORE.5**: First-Class Ecosystem Standard Library. Added `cache.pnr` and `db.pnr`, and updated `auth.pnr` and `queue.pnr` to export native bindings directly in `pranor`.
* **PS.3**: Dynamic Backpressure Routing. Added backpressure load balancer strategy in `Pranor Gate` routing load dynamically away from busy target nodes based on `X-Backpressure` headers.
* **SEC.15**: Dynamic IAM Policy Hot-Reloading. Enabled session revocation and dynamic token refresh signaling via `X-Token-Refresh` responses on stale policy versions in `Pranor Gate`.
* **ARCH.8**: Domain-Driven Decomposition. Added compile-time boundary coupling check in `Pranor` compiler/linter preventing direct cross-domain helper calls in route handlers.
* **OPS.10**: Zero-Configuration Mesh Service Discovery. Added automatic multicast UDP UDP-announce and query-broadcast auto-discovery loops inside `Pranor Mesh` registry.
* Added a unit test suite testing S3 mock gateways (`TestS3MockMode` in `Pranor Vault`), GitOps webhooks (`TestGitOpsConfigSyncWebhook` in `Pranor Gate`), event-driven saga compensations (`TestEventDrivenSagaCompensation` in `Pranor Flow`), dynamic backpressure routing (`TestDynamicBackpressureRouting` in `Pranor Gate`), policy reloading (`TestDynamicIAMPolicyHotReloading` in `Pranor Gate`), compiler boundary rules (`TestDomainDrivenDecompositionLinter` in `Pranor`), and mesh multicast service discovery (`TestMulticastServiceDiscovery` in `Pranor Mesh`).

### Fixed
* Fixed base64 URL decoding type mismatch bug in `base64UrlDecode` utility.
* Fixed vendor dependency resolution in `Pranor Flow` container builds.
* Fixed Printf format verb compilation warning in `Pranor` status command ([cmd_status.go](file:///c:/Mine/try/serv/Pranor/cmd_status.go)).
* Fixed lint warning on tagged switches for `selected` in `advanced_features_test.go` and `r.Method` in `main.go`.
* Fixed test port collision flakiness in `TestRateLimiting` and `TestDirectMemoryPassingAndResponseFilters` inside `Pranor Gate/main_test.go` by dynamic port reassignment.

---

## [1.2.0] - 2026-06-30

### Added
* **SEC.8**: KMS Secrets Envelope Key Rotation worker and SHA-256 API key hashing in `Pranor Auth`.
* **CORE.1**: HNSW Vector Search Graph implementation and comparative performance benchmark in `Pranor Vault`.
* **PS.1**: Dynamic Connection Pool Tuning (adaptive limit scaling and stale connection invalidation janitor) in `ServDB` and `Pranor Cache`.
* **DX.8**: Regular expression matching support with substring fallback in `Pranor Console` live log tailing.
* **PS.2**: WASM Memory Optimization (Wazero directory compilation caching and stateless guest module instance recycling) in `Pranor Gate` and `Pranor Pulse`.
* **SEC.7**: Automated zero-downtime mTLS certificate rotation utilizing dynamic TLS callbacks in `Pranor Mesh`.
* **OPS.7**: Ecosystem Performance Suite (Go native micro-benchmarks for Pranor Auth, ServDB, Pranor Mesh, Pranor Gate, and Pranor Pulse).
* **Phase 7 Audit — API Contract Enforcement**: Strict database dialect validation and query placeholder syntax checking in `ServDB`.
* **Phase 7 Audit — Secrets & Token Security**: JWT token expiry assertions, cryptographic hashing of API keys, and automated key rotation schemas.
* **Phase 7 Audit — Multi-Tenancy Enforcement**: Strict tenant isolation across HTTP and STOMP routes with dedicated database and queue pools.
* **Phase 7 Audit — API Versioning**: `/api/v1` API route structure compliance and backward-compatibility validations.

### Fixed
* Resolved flakiness in WebSocket HTTP handshake upgrade tests.
* Fixed file handle leaks and TempDir cleanup failures in `TestS3BatchDelete` operations.
* Corrected log security sanitization to prevent sensitive authorization tokens and session cookies from leaking into standard logs.

---

## [1.1.0] - 2026-06-30

### Added
* **TEST.8**: Go 1.18+ HTTP endpoint fuzzing engine in `Pranor Core`.
* **TEST.9**: Chaos Recovery and dependency dropout validation E2E integration test.
* **SEC.14**: Tenant switch API `/api/tenant/switch` with JWT session scope rotation in `Pranor Console`.
* **API.5**: Response header decoration middleware `DeprecationMiddleware` in `Pranor Core`.
* **API.6 / DOC.7**: Backward-compatible checking script `check_backward_compat.go`.
* **OPS.9**: Integrated `serv status` command returning metrics and health summaries.
* **DOC.6**: Release-tagger GitHub Action automating semantic release version tagging on branch merges.

### Fixed
* Decomposed console monolith handlers in `Pranor Console` into sub-packages cleanly.
* Hardened log output strings from exposing tokens, secrets, and keys.

---

## [1.0.0] - 2026-06-28

### Added
* Initial production release of the Pranor background service ecosystem.
* Shared validation middlewares, JWKS dynamic key verification, and tenant header assertion checks.
