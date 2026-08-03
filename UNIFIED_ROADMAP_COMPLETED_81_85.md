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
