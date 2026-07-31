# Architecture Decision Records (ADRs)

This document contains design logs outlining key architectural decisions in the Pranor ecosystem.

## ADR 001: Golang Code Generation Compiler Target
**Status**: Accepted  
**Context**: Serv needs to compile high-level declarations (routes, workers, agents) into performant executable binaries.
**Decision**: The `pranor` compiler parses statements and directly transpiles AST nodes into standard Go files, utilizing `go build` to generate the final optimized executable.
**Consequences**: 
* Inherits Go's performance, concurrency models, and standard library.
* Allows seamless import of external Go modules directly inside `.pnr` files.
* Build times are bound to Go compiler execution speeds.

## ADR 002: Library-Level Service Mesh Integration (Pranor Mesh)
**Status**: Accepted  
**Context**: Traditional service meshes (like Istio/Envoy) use sidecar proxies that increase container overhead, RAM consumption, and networking hops.
**Decision**: Implement a library-level service mesh where endpoints query a registry node (`Pranor Mesh`) and communicate directly over standard mTLS HTTP connections.
**Consequences**:
* Minimal overhead (no sidecar processes, direct service-to-service connections).
* Low memory footprints on resource-constrained targets.
* Requires compiled-in helpers inside the runtime library (`Pranor Core`).

## ADR 003: Store-Backed Persistent State Adapters
**Status**: Accepted  
**Context**: Services like `Pranor Auth`, `Pranor Pool`, and `Pranor Flow` require persistent state journals but running dedicated databases increases service complexity.
**Decision**: Bind state adapters directly to the `Pranor Vault` S3 storage API to write state checkpoints as JSON blobs.
**Consequences**:
* Simple persistence pattern with zero database dependency overhead.
* Recoverable on restart via standard object downloads.
* Not suited for highly concurrent transactional workloads (limitations of object writes).
