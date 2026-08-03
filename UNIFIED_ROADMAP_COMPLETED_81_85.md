# Unified Roadmap - Completed Phases 81 to 85

## Phase 81: Rebrand — Serv/Pranor → Pranor (Completed)

> **Goal:** Rename the entire ecosystem from "Serv/Pranor" to "Pranor" with a consistent product naming convention. Full rebrand across Go modules, packages, env vars, CLI binaries, Docker images, and documentation.

| # | Item | Scope | Description | Status |
|---|------|-------|-------------|--------|
| 81.1 | **Go Module & Package Rename** | Monorepo | Module path and imports updated across all 22 packages | [x] |
| 81.2 | **Compiler & Runtime Rename** | Pranor Compiler | `serv` → `pranor`, `.pnr` extension support | [x] |
| 81.3 | **Environment Variables & Configuration** | All daemons | `PRANOR_*` prefix with backward-compatible fallbacks | [x] |
| 81.4 | **Docker & Deployment Artifacts** | Infrastructure | Single binary daemon `pranord`, compose service tags | [x] |
| 81.5 | **Service Mesh & Internal Protocols** | Pranor Mesh | Custom `pranor://` routing scheme & discovery keys | [x] |
| 81.6 | **Package Distribution & Installers** | Installers | Homebrew formula, Scoop bucket, GoReleaser specs | [x] |
| 81.7 | **Documentation & Branding** | Platform Docs | Monorepo READMEs, index.html, guides rebranded | [x] |
| 81.8 | **GitHub & Repository Infrastructure** | GitHub | Repos & Actions workflows rebranded | [x] |
| 81.9 | **Pipeline Dashboard & Tooling** | Dashboard | Dashboard title & status monitors updated | [x] |
| 81.10 | **Enterprise Edition (EE)** | pranor-ee | EE package imports & build tag compatibility | [x] |

---

## Phase 82: Documentation Consolidation (Completed)

> **Goal:** Consolidate all documentation into a single `docs/` directory in the pranor monorepo. Create a unified, navigable documentation structure with mdBook static site generator setup.

| # | Item | Description | Status |
|---|------|-------------|--------|
| DC.1 | **Create docs/ directory structure** | Create all subdirectories and index files | [x] |
| DC.2 | **Consolidate module docs** | Move content from each module's README.md into `docs/modules/{name}.md` | [x] |
| DC.3 | **Write getting-started.md** | Install → first app → deploy → observe in 5 minutes | [x] |
| DC.4 | **Write language reference** | Merge lang/docs/ + lang/README into `docs/language/` | [x] |
| DC.5 | **Write deployment guides** | Consolidate docker-compose, K8s, standalone docs | [x] |
| DC.6 | **Write architecture overview** | System diagram, module interaction, data flow | [x] |
| DC.7 | **Write security guide** | Auth model, mTLS, RBAC, zero-trust | [x] |
| DC.8 | **Write observability guide** | Tracing, metrics, alerting across modules | [x] |
| DC.9 | **Consolidate changelogs** | Merge 17 per-module CHANGELOGs into one | [x] |
| DC.10 | **Shorten per-module READMEs** | Keep only quick reference + link to full docs | [x] |
| DC.11 | **Move blog content** | Migrate relevant technical content from pranor-platform/blog to docs/ | [x] |
| DC.12 | **Add CLI reference** | Complete `pranor` CLI command reference | [x] |
| DC.13 | **Add enterprise docs** | EE features, licensing, comparison table | [x] |
| DC.14 | **Add docs build system** | mdbook configuration (`book.toml`, `SUMMARY.md`) for static site generation | [x] |
| DC.15 | **Deploy docs site** | GitHub Pages deployment workflow configuration | [x] |

---

## Phase 83: Pranor LSP & IDE Intelligence Evolution (Completed)

> **Goal**: Upgrade `pranor-lsp` into an enterprise-grade language server with multi-file refactoring, workspace indexing, and intelligent code actions.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| LSP.1 | **Multi-File Workspace Rename & Refactoring Engine** | Pranor LSP | Implement `textDocument/rename` emitting `WorkspaceEdit` diffs across all imported `.pnr` files in the workspace with safety validation | [x] | OSS |
| LSP.2 | **Multi-File Code Actions & Quick-Fixes** | Pranor LSP | Implement `textDocument/codeAction` supporting auto-import insertion (`use std/...`), struct field stubbing, and missing error handler generation | [x] | OSS |
| LSP.3 | **Workspace-Wide Symbol Indexing & Search** | Pranor LSP | Background workspace symbol indexer supporting `workspace/symbol` (`Ctrl+T`) for instant symbol lookups across large multi-file projects | [x] | OSS |
| LSP.4 | **Call & Type Hierarchy Provider** | Pranor LSP | Support `textDocument/prepareCallHierarchy` allowing visual incoming/outgoing call tree navigation across functions and routes | [x] | OSS |
| LSP.5 | **Chained Type Inference & Member Auto-Completion** | Pranor LSP | Context-aware member completion on chained method invocations (e.g. `db.query().first().`) using type-inferred AST evaluation | [x] | OSS |
| LSP.6 | **Incremental AST Sync & Document Highlighting** | Pranor LSP | Use incremental document sync (`textDocument/didChange`) and `textDocument/documentHighlight` for zero-latency symbol occurrence highlighting | [x] | OSS |

---

## Phase 84: Pranor VS Code Extension Ecosystem & Control Plane Expansion (Completed)

> **Goal**: Transform `pranor-vscode` into a full-featured visual control plane for the entire Pranor ecosystem.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| VS.E1 | **Interactive API Client Panel (Pranor Gate)** | VSCode Webview | Auto-generate an in-editor HTTP/gRPC test client from route declarations with environment variables, header templates, and response formatting | [x] | OSS |
| VS.E2 | **Live Event Stream & DLQ Tailer Panel (Pranor Pulse)** | VSCode Webview | Tail topic events in real-time inside VS Code, filter by key/headers, and trigger one-click DLQ message replays | [x] | OSS |
| VS.E3 | **S3 & Vector Search Workspace Explorer (Pranor Vault)** | VSCode Explorer | Interactive tree view for browsing S3 buckets, drag-and-drop file uploads, and running natural language semantic vector queries | [x] | OSS |
| VS.E4 | **Live Distributed Flamegraph & Correlated Log Viewer (Pranor Trace)** | VSCode Webview | Render live CPU/latency flamegraphs for slow spans with instant side-by-side log correlation by `trace_id` | [x] | OSS |
| VS.E5 | **Visual Secret & Key Management Console (Pranor Secret)** | VSCode Control | Secure webview to unseal vault stores, manage environment secret maps, and inspect key rotation policies across clusters | [x] | **EE** |
| VS.E6 | **Multi-Cluster Infrastructure & Deployment Dashboard (Pranor Deploy)** | VSCode Dashboard | Manage multi-region cluster deployments, view real-time node resource utilization, and perform rollback/canary promotions directly in VS Code | [x] | OSS |

---

## Phase 85: Enterprise Commercial Tier Expansion & Monetization Moats (Completed)

> **Goal**: Define and implement 10-15 high-value Enterprise (EE) capabilities per module (Pranor Gate, Vault, Pulse, Flow, Trace, Console, Auth, Secret, Chrono, Mesh, Deploy) to establish a compelling commercial monetization model for enterprise engineering teams.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| EE.85.1 | **Kernel eBPF XDP DDoS Bypass** | Pranor Gate | 100Gbps network packet filtering at the Linux kernel level | [x] | EE |
| EE.85.2 | **Geo-IP Latency Anycast Steering** | Pranor Gate | Edge routing steering global client traffic to lowest latency datacenters | [x] | EE |
| EE.85.3 | **GraphQL Schema Stitching & Federation** | Pranor Gate | Unifies backend GraphQL endpoints into single gateway schema | [x] | EE |
| EE.85.4 | **Immutable Object Access Audit Trail** | Pranor Vault | Append-only immutable log recording object reads/writes with identity | [x] | EE |
| EE.85.5 | **Active-Active Multi-Region Replication** | Pranor Vault | Cross-cloud replication with LWW conflict resolution | [x] | EE |
| EE.85.6 | **Copy-on-Write (CoW) Bucket Branching** | Pranor Vault | Branch terabyte buckets instantly for sandbox development | [x] | EE |
| EE.85.7 | **Sovereign Client Envelope Encryption** | Pranor Vault | Zero-knowledge client envelope encryption | [x] | EE |
| EE.85.8 | **Multi-Region MirrorMaker Sync** | Pranor Pulse | Active-active cross-cloud event topic mirroring across AWS/GCP/Azure | [x] | EE |
| EE.85.9 | **Hardware Payload Encryption at Rest** | Pranor Pulse | KMS/HSM envelope encryption before writing to disk WAL | [x] | EE |
| EE.85.10 | **Schema Registry & Breaking Change Guard** | Pranor Pulse | Strict schema validation blocking breaking event payloads | [x] | EE |
| EE.85.11 | **Visual Workflow Builder & Live Replay** | Pranor Flow | Interactive drag-and-drop designer with step-by-step state diff playback | [x] | EE |
| EE.85.12 | **High-Availability Coordinator Cluster** | Pranor Flow | Multi-region workflow state replication across Raft clusters | [x] | EE |
| EE.85.13 | **Anomaly Auto-Remediation Runbooks** | Pranor Trace | Executes auto-remediation webhooks on SLA trace anomaly breaches | [x] | EE |
| EE.85.14 | **Long-Term Tail Sampling Storage Tier** | Pranor Trace | Tail sampling archiving 99.9% routine spans to cold S3 storage | [x] | EE |
| EE.85.15 | **SOC2 / HIPAA Immutable Audit Exporter** | Pranor Console | WORM audit logging with automated SOC2 compliance reporting | [x] | EE |
| EE.85.16 | **Single Sign-On (SAML 2.0 / OIDC) & RBAC** | Pranor Console | Enterprise SSO (Okta, Azure AD) with fine-grained RBAC | [x] | EE |
| EE.85.17 | **Passkey & WebAuthn Native Server** | Pranor Auth | FIDO2 / Passkey passwordless authentication server | [x] | EE |
| EE.85.18 | **Enterprise Directory Sync (AD / LDAP)** | Pranor Auth | Active Directory & LDAP identity sync with group mapping | [x] | EE |
| EE.85.19 | **Smart SLA-Aware Cron Window Optimizer** | Pranor Chrono | AI cron window optimizer shifting jobs to off-peak infrastructure | [x] | EE |
| EE.85.20 | **Multi-Cluster Fleet DR Failover** | Pranor Deploy | 1-click active-passive region failover and DNS update sync | [x] | EE |
| EE.85.21 | **Multi-Cloud Cross-VPC Mesh Peering** | Pranor Mesh | Encrypted mesh tunnels connecting AWS EKS, GCP GKE, and on-premise | [x] | EE |

