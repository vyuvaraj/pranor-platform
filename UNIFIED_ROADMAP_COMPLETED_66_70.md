# Unified Roadmap - Completed Phases 66 to 70

## Phase 67: Pranor — Rust & Python Code-Gen Targets & WASM Browser Playground (Completed)

> **Current State**: Pranor compiles `.serv` schema definitions to Go and TypeScript with full LSP IntelliSense, an import system, async runtime, WASM compilation support, and a plugin architecture.
> **What is Missing**: A Rust code-generation target (for high-performance edge services), a Python code-generation target (for ML/AI service consumers using FastAPI/Pydantic), a zero-install WASM browser playground for frictionless evaluation, multi-file schema cross-import type resolution, and a breaking change detector (`serv diff`).

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| SL.G1 | **Rust Code-Generation Target (`serv generate --lang rust`)** | Pranor Codegen | Add a Rust backend code-generator producing idiomatic Rust `struct` definitions, `serde` derive macros, trait implementations, and `axum` server stubs from `.serv` schema files | [x] | OSS |
| SL.G2 | **Python Code-Generation Target (`serv generate --lang python`)** | Pranor Codegen | Add a Python code-generator producing Pydantic v2 model definitions and FastAPI router stubs from `.serv` schema files; targeting ML/AI teams building LLM service consumers | [x] | OSS |
| SL.G3 | **Zero-Install WASM Browser Playground (`playground.pranor.dev`)** | Pranor Runtime | Compile the Pranor compiler toolchain to WebAssembly and host an interactive browser IDE where developers write, compile, and preview `.serv` definitions without any local installation | [x] | OSS |
| SL.G4 | **Multi-File Schema Import System with Cross-File Type Resolution** | Pranor Compiler | Implement cross-file type imports (`import "auth.serv"`) with a full dependency graph resolver; allow large service definitions to be cleanly split across multiple `.serv` schema files | [x] | OSS |
| SL.G5 | **Breaking Change Detector (`serv diff old.serv new.serv`)** | Pranor Tooling | Add a `serv diff` command that compares two `.serv` schema versions and outputs a structured compatibility report: new fields (safe), renamed fields (warning), removed fields (breaking) | [x] | OSS |
| SL.G6 | **`async` Task & `concurrent {}` Language Primitives** | Pranor Parser | Add `async task` and `concurrent { }` block keywords to the Pranor grammar enabling declarative concurrent step execution within service handlers, compiled to goroutines in Go | [x] | OSS |

---





## Phase 69: Pranor Vault — Native HNSW Vector Index & Semantic Object Search (Completed)

> **Current State**: Pranor Vault provides full S3 API compatibility, DuckDB SQL analytics, streaming S3 Select, OPFS browser sync, CoW bucket branching, WebTorrent P2P seeding, CRDT geo-replication, WORM locking, and FIPS KMS encryption.
> **What is Missing**: An embedded HNSW vector index engine for semantic similarity search over stored object embeddings (currently requires external Qdrant/Pinecone), automatic embedding generation on text object upload, hybrid keyword + vector search, and per-bucket vector index namespace management.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| SS.A1 | **Embedded HNSW Vector Index Engine (In-Process Similarity Search)** | Pranor Vault Vector | Implement an HNSW (Hierarchical Navigable Small World) vector index engine embedded directly in `pranor-vaultd`; store and query float32 embedding vectors with configurable M and efConstruction graph parameters | [x] | OSS |
| SS.A2 | **Automatic Embedding Generation for Text Objects on PUT** | Pranor Vault Embedding | When an object with `Content-Type: text/*` is uploaded, call a configured embedding model endpoint (local Ollama or OpenAI) to auto-generate and store a float32 vector alongside object metadata | [x] | OSS |
| SS.A3 | **Hybrid Keyword + Vector Semantic Search API** | Pranor Vault Search | Expose a `/api/v1/search` endpoint combining BM25 keyword relevance scoring with HNSW cosine similarity for hybrid semantic + lexical ranked search over bucket object content | [x] | OSS |
| SS.A4 | **Per-Bucket Vector Index Namespace Management** | Pranor Vault Namespaces | Maintain isolated per-bucket HNSW vector index namespaces; support API-driven creation, rebuilding, and deletion of vector indexes independently from object storage lifecycle | [x] | OSS |
| SS.A5 | **ANN Query API with k, Score Threshold & Metadata Filters** | Pranor Vault Query | Expose `GET /api/v1/vectors/{bucket}/search?k=10&min_score=0.80&tag=finance` for top-k nearest neighbor retrieval with minimum score thresholding and metadata predicate pre-filtering | [x] | OSS |
| SS.A6 | **Persistent mmap-Backed HNSW Graph with Incremental Node Insertion** | Pranor Vault Persistence | Persist HNSW graph state to disk using memory-mapped files; support incremental online insertion of new vector nodes without requiring full graph rebuild on each new object PUT | [x] | OSS |

---







---

## Phase 68: Pranor Gateway — Native MCP Tool Registry & AI Agent Routing (Completed)

> **Current State**: Pranor Gateway has a Smart AI cost router, token-per-minute throttling, eBPF XDP DDoS bypass, semantic prompt caching, WAF, GraphQL federation, WASM hot-reload registry, and OIDC enforcement.
> **What is Missing**: A native Model Context Protocol (MCP) server that auto-exposes Pranor services as strongly-typed AI agent tools, LLM streaming SSE response passthrough, per-agent-session context tracking, tool call audit logging with cost attribution, multi-model provider fallback chains, and prompt injection detection.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| SG.A1 | **Native MCP Tool Registry (Auto-Expose Pranor Services as AI Agent Tools)** | Pranor Gateway MCP | Implement a native MCP server at `/mcp` that auto-discovers and exposes registered Pranor Vault buckets, Pranor Pulse topics, and Pranor services as typed MCP tools consumable by Claude, GPT-4o, and local Ollama agents | [x] | OSS |
| SG.A2 | **LLM Streaming SSE Response Passthrough (Server-Sent Events)** | Pranor Gateway Proxy | Implement transparent proxy passthrough of streaming SSE LLM completions; preserve chunk ordering and correctly handle chunked transfer encoding for real-time token streaming to browser clients | [x] | OSS |
| SG.A3 | **AI Agent Session Context Tracker (Multi-Turn Conversation State)** | Pranor Gateway Agent | Maintain per-agent-session conversation context windows across sequential MCP tool calls; inject conversation history into each upstream LLM request automatically based on session ID header | [x] | **EE** |
| SG.A4 | **Tool Call Audit Log & Per-Session AI Cost Attribution** | Pranor Gateway Audit | Log every MCP tool call with agent ID, tool name, input arguments, response status, and token cost attribution; expose searchable audit history in Pranor Console cost attribution dashboard | [x] | **EE** |
| SG.A5 | **Multi-Model Provider Fallback Chain (GPT-4o → Claude → Ollama)** | Pranor Gateway Fallback | Configure ordered fallback chains across AI providers; automatically retry failed or rate-limited requests against the next provider with context adaptation for different model API formats | [x] | **EE** |
| SG.A6 | **Prompt Injection Detection & Input Sanitization Guard** | Pranor Gateway Security | Detect and block prompt injection attacks (system prompt override, jailbreak patterns) using embedding-based cosine similarity scoring against known attack signatures before forwarding to LLMs | [x] | OSS |

---


---

## Phase 66: Pranor Tunnel — WebSocket Multiplexing, Replay Inspector & Auth Gating (Completed)

> **Current State**: Pranor Tunnel implements HTTP request tunneling with subdomain routing, a request inspector UI, OTel integration, and webhook forwarding.
> **What is Missing**: WebSocket connection multiplexing over a single tunnel pipe, JWT/API-key auth gating on inbound tunnel connections (currently unauthenticated), request/response body capture with replay UI in Pranor Console, persistent reconnect with exponential backoff, and per-tunnel bandwidth throttling.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| TN.G1 | **WebSocket Connection Multiplexing over a Single Tunnel Pipe** | Pranor Tunnel Protocol | Multiplex multiple concurrent WebSocket connections from different clients over a single persistent tunnel connection using a lightweight stream framing header protocol | [x] | OSS |
| TN.G2 | **JWT / API-Key Auth Gating on Tunnel Endpoint Connections** | Pranor Tunnel Auth | Require JWT bearer tokens or API keys before accepting inbound HTTP connections on any tunnel endpoint; integrate with Pranor Auth for token validation and scope enforcement | [x] | OSS |
| TN.G3 | **Full Request & Response Body Capture with Replay UI in Pranor Console** | Pranor Tunnel Inspector | Capture complete request/response pairs (headers, body, timing, status) and expose them in a Pranor Console tunnel inspector tab; allow one-click replay of any captured request for debugging | [x] | OSS |
| TN.G4 | **Persistent Tunnel Reconnect with Exponential Backoff & Jitter** | Pranor Tunnel Client | Implement automatic reconnect logic in the tunnel client with exponential backoff and jitter; maintain tunnel availability and re-register subdomain routing across transient network interruptions | [x] | OSS |
| TN.G5 | **Per-Tunnel Bandwidth Throttling & Rate Limiting** | Pranor Tunnel Policy | Enforce configurable bandwidth limits (bytes/second) per tunnel connection to prevent a single heavy client from saturating the shared tunnel relay infrastructure | [x] | **EE** |
| TN.G6 | **Shareable Tunnel URLs with Expiry & One-Time Access Tokens** | Pranor Tunnel Sharing | Generate shareable tunnel URLs with configurable TTL expiry and single-use access tokens for secure, time-limited collaboration on local development services | [x] | OSS |

---

---

## Phase 70: Pranor Console — eBPF Flamegraph Profiler, Chaos Engine UI & Unified AI Assistant (Completed)

> **Current State**: Pranor Console implements OTel trace waterfall dashboards, hash ring visualizers, SQL workbench, alert management, topology views, and a launcher for all Pranor services.
> **What is Missing**: An embedded eBPF flamegraph profiling visualization tab, a Chaos Engineering control panel (for triggering and monitoring fault injection experiments across modules), an embedded AI assistant chat panel, unified global search across all Pranor resources, and a cost attribution breakdown dashboard.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| CC.G1 | **eBPF Flamegraph Profiling Tab with OTel Trace Correlation** | Pranor Console Profiler | Add a Profiling tab rendering interactive SVG flamegraphs from Pranor Trace eBPF profiling data; automatically correlate captured flamegraph stack samples with the active OTel trace span being inspected | [x] | OSS |
| CC.G2 | **Chaos Engineering Control Panel** | Pranor Console Chaos | Provide a UI panel for initiating and monitoring chaos experiments: inject latency, drop packets, kill service replicas, or simulate clock skew; visualize real-time impact on service topology error rates | [x] | OSS |
| CC.G3 | **Unified Global Search Across All Pranor Resources (⌘K)** | Pranor Console Search | Implement a `cmd+K`-style universal command palette searching across trace IDs, object keys, queue topics, job names, service instances, and alert events; returning ranked contextual results | [x] | OSS |
| CC.G4 | **Embedded AI Assistant Chat Panel (Powered by Pranor Gate MCP)** | Pranor Console AI | Embed a conversational AI assistant that answers questions about the running system, explains trace anomalies, suggests performance optimizations, and executes SQL queries via Pranor Gate MCP tool calls | [x] | **EE** |
| CC.G5 | **Cost Attribution Dashboard (Per-Service Egress, Storage & AI Token Spend)** | Pranor Console Cost | Aggregate egress bandwidth, object storage bytes, AI token consumption, and Pranor Pulse message throughput per service/team; render a cost attribution breakdown with budget alert thresholds | [x] | **EE** |
| CC.G6 | **Theme Customization, Pinned Dashboard Widgets & Keyboard Shortcuts** | Pranor Console UX | Persist user preferences for dark/light theme, dashboard widget layout (drag-to-reorder), pinned quick-access shortcuts to frequently viewed resources, and configurable keyboard shortcut bindings | [x] | OSS |

---
