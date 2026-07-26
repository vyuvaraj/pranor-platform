# Unified Roadmap - Completed Phases 66 to 70

## Phase 67: Serv-lang — Rust & Python Code-Gen Targets & WASM Browser Playground (Completed)

> **Current State**: Serv-lang compiles `.serv` schema definitions to Go and TypeScript with full LSP IntelliSense, an import system, async runtime, WASM compilation support, and a plugin architecture.
> **What is Missing**: A Rust code-generation target (for high-performance edge services), a Python code-generation target (for ML/AI service consumers using FastAPI/Pydantic), a zero-install WASM browser playground for frictionless evaluation, multi-file schema cross-import type resolution, and a breaking change detector (`serv diff`).

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| SL.G1 | **Rust Code-Generation Target (`serv generate --lang rust`)** | Serv-lang Codegen | Add a Rust backend code-generator producing idiomatic Rust `struct` definitions, `serde` derive macros, trait implementations, and `axum` server stubs from `.serv` schema files | [x] | OSS |
| SL.G2 | **Python Code-Generation Target (`serv generate --lang python`)** | Serv-lang Codegen | Add a Python code-generator producing Pydantic v2 model definitions and FastAPI router stubs from `.serv` schema files; targeting ML/AI teams building LLM service consumers | [x] | OSS |
| SL.G3 | **Zero-Install WASM Browser Playground (`playground.servverse.dev`)** | Serv-lang Runtime | Compile the Serv-lang compiler toolchain to WebAssembly and host an interactive browser IDE where developers write, compile, and preview `.serv` definitions without any local installation | [x] | OSS |
| SL.G4 | **Multi-File Schema Import System with Cross-File Type Resolution** | Serv-lang Compiler | Implement cross-file type imports (`import "auth.serv"`) with a full dependency graph resolver; allow large service definitions to be cleanly split across multiple `.serv` schema files | [x] | OSS |
| SL.G5 | **Breaking Change Detector (`serv diff old.serv new.serv`)** | Serv-lang Tooling | Add a `serv diff` command that compares two `.serv` schema versions and outputs a structured compatibility report: new fields (safe), renamed fields (warning), removed fields (breaking) | [x] | OSS |
| SL.G6 | **`async` Task & `concurrent {}` Language Primitives** | Serv-lang Parser | Add `async task` and `concurrent { }` block keywords to the Serv-lang grammar enabling declarative concurrent step execution within service handlers, compiled to goroutines in Go | [x] | OSS |

---





## Phase 69: ServStore — Native HNSW Vector Index & Semantic Object Search (Completed)

> **Current State**: ServStore provides full S3 API compatibility, DuckDB SQL analytics, streaming S3 Select, OPFS browser sync, CoW bucket branching, WebTorrent P2P seeding, CRDT geo-replication, WORM locking, and FIPS KMS encryption.
> **What is Missing**: An embedded HNSW vector index engine for semantic similarity search over stored object embeddings (currently requires external Qdrant/Pinecone), automatic embedding generation on text object upload, hybrid keyword + vector search, and per-bucket vector index namespace management.

| # | Item | Component | Description | Status | Tier |
|---|------|-----------|-------------|--------|:---:|
| SS.A1 | **Embedded HNSW Vector Index Engine (In-Process Similarity Search)** | ServStore Vector | Implement an HNSW (Hierarchical Navigable Small World) vector index engine embedded directly in `servstored`; store and query float32 embedding vectors with configurable M and efConstruction graph parameters | [x] | OSS |
| SS.A2 | **Automatic Embedding Generation for Text Objects on PUT** | ServStore Embedding | When an object with `Content-Type: text/*` is uploaded, call a configured embedding model endpoint (local Ollama or OpenAI) to auto-generate and store a float32 vector alongside object metadata | [x] | OSS |
| SS.A3 | **Hybrid Keyword + Vector Semantic Search API** | ServStore Search | Expose a `/api/v1/search` endpoint combining BM25 keyword relevance scoring with HNSW cosine similarity for hybrid semantic + lexical ranked search over bucket object content | [x] | OSS |
| SS.A4 | **Per-Bucket Vector Index Namespace Management** | ServStore Namespaces | Maintain isolated per-bucket HNSW vector index namespaces; support API-driven creation, rebuilding, and deletion of vector indexes independently from object storage lifecycle | [x] | OSS |
| SS.A5 | **ANN Query API with k, Score Threshold & Metadata Filters** | ServStore Query | Expose `GET /api/v1/vectors/{bucket}/search?k=10&min_score=0.80&tag=finance` for top-k nearest neighbor retrieval with minimum score thresholding and metadata predicate pre-filtering | [x] | OSS |
| SS.A6 | **Persistent mmap-Backed HNSW Graph with Incremental Node Insertion** | ServStore Persistence | Persist HNSW graph state to disk using memory-mapped files; support incremental online insertion of new vector nodes without requiring full graph rebuild on each new object PUT | [x] | OSS |

---





