# Unified Roadmap - Completed Phases 76 to 80

## Phase 76: Serv-lang Compiler & Developer Experience Hardening (Completed)

> **Current State**: Serv-lang compiles `.serv` files into Go code, features an LSP server, WASM playground, multi-file imports, and test runner. However, error reporting lacks precise caret positions, codegen produces generic `interface{}` types, and database migrations lack state tracking.
> **What is Missing**: 10 targeted developer experience enhancements covering compiler error line/col carets, human-readable Go codegen with `.serv` line annotations, typed Go codegen, `.serv` source-mapped runtime errors, migration state tracking (`_serv_migrations`), multi-file LSP go-to-definition, non-crashing watch reloader, per-test DB transaction isolation, and WASM playground side-by-side Go code view.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| **Serv-lang DX & Compiler Hardening** | | | | | |
| SL.H1 | **Readable Generated Go Code & Source Line Annotations** | Serv-lang Codegen | Generate human-readable handler/cron names (`handleGET_api_tasks`) and inject `// line <file>:<line>` source map comments into generated Go code | [x] | OSS |
| SL.H2 | **Type-Aware Go Code Generator** | Serv-lang Codegen | Emit real Go types (`int`, `string`, structs) instead of generic `interface{}` to eliminate runtime type assertions and improve performance | [x] | OSS |
| SL.H3 | **Runtime Errors with `.serv` Source Context** | Serv-lang Runtime | Wrap runtime panic recovery and DB/HTTP library errors with `.serv` file, line, and route context instead of raw Go stacktraces | [x] | OSS |
| SL.H4 | **Migration State Tracking & Rollback Engine** | Serv-lang Storage | Track applied migrations in `_serv_migrations` table, prevent duplicate `ALTER TABLE` runs, and support `up {}` / `down {}` migration blocks | [x] | OSS |
| SL.H5 | **Workspace-Wide LSP Go-to-Definition & References** | Serv-lang Tooling | Extend LSP to jump across imported files, perform multi-file `Find-All-References`, and propagate symbol renames | [x] | OSS |
| SL.H6 | **Non-Crashing Hot Reload with Inline Error Diagnostics** | Serv-lang CLI | Keep previous good binary serving traffic when `serv run --watch` encounters compilation errors, pushing diagnostics inline to terminal and LSP | [x] | OSS |
| SL.H7 | **Isolated Per-Test Transaction & Rollback Engine** | Serv-lang Testing | Wrap each `test` block in an isolated database transaction that automatically rolls back, and scope `mock` statements to individual test blocks | [x] | OSS |
| SL.H8 | **WASM Playground Side-by-Side Go Output & Live Server** | Serv-lang Web | Upgrade web playground to show generated Go code side-by-side, render debounced syntax errors, and execute in-browser HTTP requests | [x] | OSS |
| SL.H9 | **`serv.toml` Package Dependency & Resolution Engine** | Serv-lang Package | Add `serv.toml` package declaration manifest and `serv add <pkg>` command with clear resolution order (local -> stdlib -> registry) | [x] | OSS |
| SL.H10 | **Precise AST Token Position & Source Caret Reporter** | Serv-lang Parser | Propagate lexer line/column positions to AST nodes and format compiler errors with source code lines and caret pointers (`^`) | [x] | OSS |

---
