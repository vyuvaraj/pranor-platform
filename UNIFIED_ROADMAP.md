# Serv Unified Ecosystem Roadmap & Architect Analysis


> Single source of truth for the **Serv** ecosystem: Serv-lang, ServGate, ServStore, ServQueue, ServConsole, ServCache, ServMesh, ServCron, ServCloud, ServTrace, ServTunnel, ServAuth, ServPool, ServMail, ServFlow, and the Servverse vision.  

> Last updated: July 9, 2026


---


## Ecosystem Completion Status


All items in Phases 1 through 14 have been fully implemented, verified, and pushed.


- For completed details of Phases 1 to 5: Refer to the git history and repository CHANGELOG.

- For completed details of Phases 6 to 10: See [UNIFIED_ROADMAP_COMPLETED_6_10.md](file:///c:/Mine/try/serv/servverse-repo/UNIFIED_ROADMAP_COMPLETED_6_10.md).

- For completed details of Phases 11 to 15: See [UNIFIED_ROADMAP_COMPLETED_11_15.md](file:///c:/Mine/try/serv/servverse-repo/UNIFIED_ROADMAP_COMPLETED_11_15.md).

- For completed details of Phase 16-19: See [UNIFIED_ROADMAP_COMPLETED_16_20.md](file:///c:/Mine/try/serv/servverse-repo/UNIFIED_ROADMAP_COMPLETED_16_20.md).

- For completed details of Phase 31-35: See [UNIFIED_ROADMAP_COMPLETED_31_35.md](UNIFIED_ROADMAP_COMPLETED_31_35.md).

- For completed details of Phase 36-38: See [UNIFIED_ROADMAP_COMPLETED_36_40.md](UNIFIED_ROADMAP_COMPLETED_36_40.md).


### Completion Tracker


| Initiative Area | Total Items | Completed | Pending | Progress | Status Bar |

|-----------------|-------------|-----------|---------|----------|------------|

| **Phase 9: Scale & Enterprise Hardening** | 13 | 13 | 0 | **100%** | ████████████████████ |

| **Phase 10: Productization & Cloud PaaS** | 32 | 32 | 0 | **100%** | ████████████████████ |

| **Phase 11: Unified Dashboard & Console** | 33 | 33 | 0 | **100%** | ████████████████████ |

| **Phase 12: Dual-Licensing & EE Split** | 19 | 19 | 0 | **100%** | ████████████████████ |

| **Phase 13: Language & Runtime Evolution**| 18 | 18 | 0 | **100%** | ████████████████████ |

| **Phase 14: AI-Native Ecosystem** | 28 | 28 | 0 | **100%** | ████████████████████ |

| **Phase 16: Operational Hardening & Production Readiness** | 18 | 18 | 0 | **100%** | ████████████████████ |

| **Phase 17: Zero-Trust Clustering & Edge Serverless** | 8 | 8 | 0 | **100%** | ████████████████████ |

| **Phase 18: OSS-to-EE Boundary Alignment & Refactoring** | 6 | 6 | 0 | **100%** | ████████████████████ |

| **Phase 19: Component Maturity Alignment** | 7 | 7 | 0 | **100%** | ████████████████████ |

| **Phase 20: OSS-to-EE Refactoring & Enterprise Migrations** | 6 | 6 | 0 | **100%** | ████████████████████ |

| **Phase 21: Enterprise Ecosystem Scale & Next-Gen** | 6 | 6 | 0 | **100%** | ████████████████████ |

| **Phase 22: Quality, Credibility & Code Health** | 20 | 20 | 0 | **100%** | ████████████████████ |

| **Phase 23: Developer Adoption & Growth** | 14 | 11 | 3 | **79%** | ███████████████░░░░ |

| **Phase 24: Standalone Component Independence** | 20 | 16 | 4 | **80%** | ████████████████░░░░ |

| **Phase 25: Component Depth & Production Hardening** | 60 | 60 | 0 | **100%** | ████████████████████ |

| **Phase 26: Competitive Differentiation** | 107 | 107 | 0 | **100%** | ████████████████████ |

| **Phase 27: v1.0 Release Readiness** | 14 | 13 | 1 | **93%** | ██████████████████░░ |

| **Phase 28: Distribution & Installer Packaging** | 11 | 11 | 0 | **100%** |

| **Phase 29: LSP IntelliSense & Developer Tooling** | 16 | 16 | 0 | **100%** |

| **Phase 30: ServLock & ServSecret Hardening** | 10 | 10 | 0 | **100%** | ████████████████████ |

| **Phase 31: ServLock & ServSecret Ecosystem Integration** | 5 | 5 | 0 | **100%** | ████████████████████ |

| **Phase 32: ServLock & ServSecret Standalone & Hardening** | 8 | 8 | 0 | **100%** | ████████████████████ |

| **Phase 33: ServLock & ServSecret Advanced Capabilities** | 10 | 10 | 0 | **100%** | ████████████████████ |

| **Phase 34: ServLock & ServSecret Enterprise & UI** | 6 | 6 | 0 | **100%** | ████████████████████ |

| **Phase 35: Serv-lang Language Ergonomics** | 40 | 40 | 0 | **100%** | ████████████████████ |
| **Phase 39: ServQueue Embedded & OPFS Browser Queue** | 6 | 6 | 0 | **100%** | ████████████████████ |
| **Phase 40: ServQueue Browser WASM Hardening & Multi-Tab Resilience** | 6 | 6 | 0 | **100%** | ████████████████████ |
| **Phase 41: ServQueue Next-Gen Enterprise Stream Engine** | 8 | 8 | 0 | **100%** | ████████████████████ |
| **Phase 42: ServQueue Beyond-Enterprise Security & Sovereign Stream Engine** | 8 | 2 | 6 | **25%** | █████░░░░░░░░░░░░░░░ |

| **TOTAL ECOSYSTEM WORK** | **538** | **532** | **6** | **99%** | ███████████████████░ |


---


## Phase 15: Component Backlog & Future Enhancements (Completed)


All backlog and component enhancement items for Phase 15 have been fully completed, verified, and pushed.


- For completed details of Phase 15: See [UNIFIED_ROADMAP_COMPLETED_11_15.md](file:///c:/Mine/try/serv/servverse-repo/UNIFIED_ROADMAP_COMPLETED_11_15.md).


---


## Appendix A: Cross-Service Runtime Dependency Diagram


```mermaid

graph TD

%% CLI & Compiler Connector

ServCLI["Serv-lang CLI"] -->|compiles & deploys| ServCloud

ServCLI -->|runs tests| ServLocal


%% Gateway Edge Layer

ServGate["ServGate (API Gateway)"] -->|proxies client requests| ServMesh

ServGate -->|reports traffic metrics| ServCloud


%% Service Mesh Layer

ServMesh["ServMesh (Service Mesh)"] -->|routes traffic| ServMeshInstances["Service Instances"]

ServMeshInstances -->|registers to| ServRegistry

ServMeshInstances -->|publishes traces| ServTrace

ServMeshInstances -->|accesses cache| ServCache

ServMeshInstances -->|enqueues tasks| ServQueue

ServMeshInstances -->|schedules jobs| ServCron

ServMeshInstances -->|sends mail| ServMail


%% Control Plane & Orchestrator

ServCloud["ServCloud (Orchestrator)"] -->|orchestrates processes| ServMeshInstances

ServCloud -->|manages routing rules| ServMesh

ServCloud -->|reads metrics & autoscales| ServGate


%% Auxiliary Core Services

ServCron -->|triggers target HTTP hooks| ServMeshInstances

ServQueue -->|dispatches messages to| ServMeshInstances

ServTrace -->|stores & indexes telemetry| ServStore

ServRegistry -->|indexes service packages| ServStore

ServStore["ServStore (S3 Storage)"]

```


---


## Appendix B: Component Maturity Matrix


> Updated July 14, 2026 � based on actual code metrics (test counts, pkg structure, standalone flags, EE gating, main.go line counts).


| Component | Tests | Code Structure | Security | Observability | Standalone | EE Gated | Overall |

|-----------|-------|----------------|----------|---------------|-----------|----------|---------|

| **Serv-lang** | ?? 112 funcs | ?? compiler/, runtime/, lsp/, stdlib/ | ?? Null safety, type checking | ?? OTel codegen | ? N/A | ? N/A | **Production** |

| **ServStore** | ?? 93 funcs | ?? cmd/ + 11 packages | ?? SigV4 + TLS + OIDC + LDAP | ?? OTel + slog | ?? A+ Zero-config | ? Federation + cold tier | **Production** |

| **ServGate** | ?? 50 funcs | ?? 3 packages (proxy, wasm, otel) | ?? JWT + mTLS + ACME + policy | ?? OTel + access logs | ?? A- config.json | ? AI cache + LLM routing | **Production** |

| **ServConsole** | ?? 56 funcs | ?? 12 packages | ?? OIDC + RBAC + JWT | ?? OTel | ? Aggregator | ? SLO, cost, runbooks, exec | **Production** |

| **ServCache** | ?? 46 funcs | ?? 3 packages | ?? Token auth | ?? OTel | ?? A Standalone flag | ? Namespace isolation | **Production** |

| **ServFlow** | ?? 43 funcs | ?? 3 packages (engine, handlers, storage) | ?? JWT | ?? OTel | ?? A Standalone flag | ? Saga hooks | **Production** |

| **ServPool** | ?? 40 funcs | ?? 4 packages | ?? JWT | ?? OTel | ?? B+ docs only | ? N/A | **Production** |

| **ServMesh** | ?? 37 funcs | ?? 7 packages | ?? mTLS + JWT + auto-rotate | ?? OTel | ?? B+ needs services | ? N/A | **Production** |

| **ServTunnel** | ?? 37 funcs | ?? 6 packages | ?? TLS + token + rate limit | ?? OTel | ?? A- generic tunnel | ? Federation | **Production** |

| **ServQueue** | ?? 36 funcs | ?? 5 packages | ?? TLS + STOMP auth | ?? OTel + spans | ?? A Zero-config | ? Federation + semantic route | **Production** |

| **ServMail** | ?? 34 funcs | ?? 5 packages | ?? JWT | ?? OTel | ?? A Standalone flag | ? N/A | **Production** |

| **ServDocs** | ?? 34 funcs | ?? 3 packages (generator, openapi, parser) | ? N/A | ? N/A | ?? B+ .srv-specific | ? N/A | **Stable** |

| **ServCloud** | ?? 31 funcs | ?? 3 packages | ?? JWT | ?? OTel | ?? B Serv-specific | ? Autoscale | **Stable** |

| **ServShared** | ?? 30 funcs | ?? 4 packages (datafabric, middleware, outbox, policy) | ?? JWT + mTLS + tenant | ?? OTel init | ? Library | ? Tenant isolation | **Production** |

| **ServTrace** | ?? 17 funcs | ?? 2 packages | ?? Basic auth | ?? Self-traces | ?? A- OTLP collector | ? Cold tier, NL, anomaly | **Stable** |

| **ServAuth** | ?? 16 funcs | ?? 6 packages | ?? bcrypt + AES + MFA + OIDC | ?? OTel | ?? B No standalone flag | ? Stuffing detection | **Stable** |

| **ServCron** | ?? 13 funcs | ?? 3 packages | ?? JWT + Redis lease | ?? OTel | ?? A Standalone flag | ? N/A | **Stable** |

| **ServRegistry** | ?? 12 funcs | ?? 4 packages (registry, resolution, signing, web) | ?? JWT + crypto signing | ?? OTel | ?? B+ Standalone flag | ? N/A | **Stable** |

| **ServLock** | ?? 2 funcs | ?? 2 packages (handlers, storage) | ?? JWT | ?? Basic | ?? B+ Embedded | ? N/A | **Beta** |
| **ServSecret** | ?? 2 funcs | ?? 2 packages (handlers, storage) | ?? AES-GCM + JWT | ?? Basic | ?? B+ Standalone | ? N/A | **Beta** |

### Summary


| Rating | Count | Components |

|--------|-------|-----------|

| **Production** | 13 | Serv-lang, ServStore, ServGate, ServConsole, ServCache, ServFlow, ServPool, ServMesh, ServTunnel, ServQueue, ServMail, ServShared |

| **Stable** | 6 | ServDocs, ServCloud, ServTrace, ServAuth, ServCron, ServRegistry |

| **Beta** | 2 | ServLock, ServSecret |

**Legend:** ?? Strong | ?? Adequate | ?? Needs work | ? Not applicable | ? EE features gated


---


## Phase 16: Operational Hardening & Production Readiness (Completed)


All backlog tasks for Phase 16 have been fully completed, verified, and archived.

- For completed details of Phase 16: See [UNIFIED_ROADMAP_COMPLETED_16_20.md](file:///c:/Mine/try/serv/servverse-repo/UNIFIED_ROADMAP_COMPLETED_16_20.md).


---


## Phase 17: Zero-Trust Clustering & Edge Serverless Evolution (Completed)


All backlog tasks for Phase 17 have been fully completed, verified, and archived.

- For completed details of Phase 17: See [UNIFIED_ROADMAP_COMPLETED_16_20.md](file:///c:/Mine/try/serv/servverse-repo/UNIFIED_ROADMAP_COMPLETED_16_20.md).


---


## Phase 18: OSS-to-EE Boundary Alignment & Refactoring (Completed)


All backlog tasks for Phase 18 have been fully completed, verified, and archived.

- For completed details of Phase 18: See [UNIFIED_ROADMAP_COMPLETED_16_20.md](file:///c:/Mine/try/serv/servverse-repo/UNIFIED_ROADMAP_COMPLETED_16_20.md).


---


## Phase 19: Component Maturity Alignment (Completed)


All backlog tasks for Phase 19 have been fully completed, verified, and archived.

- For completed details of Phase 19: See [UNIFIED_ROADMAP_COMPLETED_16_20.md](file:///c:/Mine/try/serv/servverse-repo/UNIFIED_ROADMAP_COMPLETED_16_20.md).


---


## Phase 20: OSS-to-EE Refactoring & Enterprise Migrations (Completed)


All backlog tasks for Phase 20 have been fully completed, verified, and archived.

- For completed details of Phase 20: See [UNIFIED_ROADMAP_COMPLETED_16_20.md](file:///c:/Mine/try/serv/servverse-repo/UNIFIED_ROADMAP_COMPLETED_16_20.md).


## Phase 21: Enterprise Ecosystem Scale & Next-Gen Capabilities (Completed)


All backlog tasks for Phase 21 have been fully completed, verified, and archived.

- For completed details of Phase 21: See [UNIFIED_ROADMAP_COMPLETED_21_25.md](file:///c:/Mine/try/serv/servverse-repo/UNIFIED_ROADMAP_COMPLETED_21_25.md).


## Phase 22: Quality, Credibility & Code Health (Completed)


All backlog tasks for Phase 22 have been fully completed, verified, and archived.

- For completed details of Phase 22: See [UNIFIED_ROADMAP_COMPLETED_21_25.md](file:///c:/Mine/try/serv/servverse-repo/UNIFIED_ROADMAP_COMPLETED_21_25.md).


## Phase 23: Developer Adoption & Growth (Pending)

> **Context:** The platform is feature-complete but has zero external users. This phase focuses on removing friction, building community, and proving production-readiness.

- For completed details of Phase 23: See [UNIFIED_ROADMAP_COMPLETED_21_25.md](file:///F:/Don/servverse/servverse/UNIFIED_ROADMAP_COMPLETED_21_25.md).

### Pending / Deferred Items

| # | Item | Component | Description | Status |
|---|------|-----------|-------------|--------|
| AG.4 | **10-minute demo video** | servverse-repo | Screen recording: install → write service → deploy → observe in console. Hosted on YouTube + embedded in GitHub Pages | [Deferred] |
| AG.5 | **Discord/community server** | - | Developer community for questions, showcases, and contributors | [Deferred] |
| AG.12 | **Customer pilot program** | - | Find 2-3 teams to run in staging. Gather real feedback on DX, performance, gaps | [Deferred] |

## Phase 24: Standalone Component Independence (Completed)


All backlog tasks for Phase 24 have been fully completed, verified, and archived.

- For completed details of Phase 24: See [UNIFIED_ROADMAP_COMPLETED_21_25.md](file:///c:/Mine/try/serv/servverse-repo/UNIFIED_ROADMAP_COMPLETED_21_25.md).


---


## Phase 24.1: Standalone Hardening to A+ (Completed)

All backlog tasks for Phase 24.1 have been fully completed, verified, and archived.
- For completed details of Phase 24.1: See [UNIFIED_ROADMAP_COMPLETED_21_25.md](file:///F:/Don/servverse/servverse/UNIFIED_ROADMAP_COMPLETED_21_25.md).

---


## Phase 25: Component Depth & Production Hardening (Completed)


All backlog tasks for Phase 25 (D.1 - D.60) have been fully completed, verified, and archived.

- For completed details of Phase 25: See [UNIFIED_ROADMAP_COMPLETED_21_25.md](file:///c:/Mine/try/serv/servverse-repo/UNIFIED_ROADMAP_COMPLETED_21_25.md).





---


## Phase 26: Competitive Differentiation (Completed)

All backlog tasks for Phase 26 have been fully completed, verified, and archived.
- For completed details of Phase 26: See [UNIFIED_ROADMAP_COMPLETED_26_30.md](file:///F:/Don/servverse/servverse/UNIFIED_ROADMAP_COMPLETED_26_30.md).

## Phase 27: v1.0 Release Readiness (Pending)

> **Goal:** Close the consistency gaps identified in the API maturity audit. These are mechanical fixes (not design changes) required to confidently tag v1.0.0.

- For completed details of Phase 27: See [UNIFIED_ROADMAP_COMPLETED_26_30.md](file:///F:/Don/servverse/servverse/UNIFIED_ROADMAP_COMPLETED_26_30.md).

### Pending / Deferred Items

| # | Item | Description | Status |
|---|------|-------------|--------|
| V1.9 | **API freeze period** | 4 weeks with zero breaking changes after all V1.1-V1.6 are done. Monitor for issues | [Deferred] |

## Phase 28: Distribution & Installer Packaging (Completed)


All backlog tasks for Phase 28 have been fully completed, verified, and archived.

- For completed details of Phase 28: See [UNIFIED_ROADMAP_COMPLETED_26_30.md](file:///F:/Don/servverse/servverse/UNIFIED_ROADMAP_COMPLETED_26_30.md).


## Phase 29: LSP IntelliSense & Developer Tooling (Completed)


All backlog tasks for Phase 29 have been fully completed, verified, and archived.

- For completed details of Phase 29: See [UNIFIED_ROADMAP_COMPLETED_26_30.md](file:///F:/Don/servverse/servverse/UNIFIED_ROADMAP_COMPLETED_26_30.md).

## Phase 30: ServLock & ServSecret Hardening (Completed)

All backlog tasks for Phase 30 have been fully completed, verified, and archived.

- For completed details of Phase 30: See [UNIFIED_ROADMAP_COMPLETED_26_30.md](file:///F:/Don/servverse/servverse/UNIFIED_ROADMAP_COMPLETED_26_30.md).

## Phase 31: ServLock & ServSecret Ecosystem Integration (Completed)

All backlog tasks for Phase 31 have been fully completed, verified, and archived.

- For completed details of Phase 31: See [UNIFIED_ROADMAP_COMPLETED_31_35.md](file:///F:/Don/servverse/servverse/UNIFIED_ROADMAP_COMPLETED_31_35.md).

## Phase 32: ServLock & ServSecret Standalone & Hardening (Completed)

All backlog tasks for Phase 32 have been fully completed, verified, and archived.

- For completed details of Phase 32: See [UNIFIED_ROADMAP_COMPLETED_31_35.md](file:///F:/Don/servverse/servverse/UNIFIED_ROADMAP_COMPLETED_31_35.md).

## Phase 33: ServLock & ServSecret Advanced Capabilities (Completed)

All backlog tasks for Phase 33 have been fully completed, verified, and archived.

- For completed details of Phase 33: See [UNIFIED_ROADMAP_COMPLETED_31_35.md](file:///F:/Don/servverse/servverse/UNIFIED_ROADMAP_COMPLETED_31_35.md).

## Phase 34: ServLock & ServSecret Enterprise & UI (Completed)

All backlog tasks for Phase 34 have been fully completed, verified, and archived.

- For completed details of Phase 31-35: See [UNIFIED_ROADMAP_COMPLETED_31_35.md](UNIFIED_ROADMAP_COMPLETED_31_35.md).

- For completed details of Phase 36-38: See [UNIFIED_ROADMAP_COMPLETED_36_40.md](UNIFIED_ROADMAP_COMPLETED_36_40.md).

---

## Phase 36: Component Maturity & Ecosystem Security (Completed)

All backlog tasks for Phase 36 have been fully completed, verified, and archived.
- For completed details of Phase 36: See [UNIFIED_ROADMAP_COMPLETED_36_40.md](UNIFIED_ROADMAP_COMPLETED_36_40.md).

## Phase 37: Serv-lang Niche Positioning & DX Evolution (Completed)

All backlog tasks for Phase 37 have been fully completed, verified, and archived.
- For completed details of Phase 37: See [UNIFIED_ROADMAP_COMPLETED_36_40.md](UNIFIED_ROADMAP_COMPLETED_36_40.md).

## Phase 38: WASM Plugin Ecosystem & Community Repository (Completed)

All backlog tasks for Phase 38 have been fully completed, verified, and archived.
- For completed details of Phase 38: See [UNIFIED_ROADMAP_COMPLETED_36_40.md](UNIFIED_ROADMAP_COMPLETED_36_40.md).

---

## Phase 39: ServQueue Embedded & OPFS Browser Event Streaming (Completed)

All backlog tasks for Phase 39 have been fully completed, verified, and archived.
- For completed details of Phase 39: See [UNIFIED_ROADMAP_COMPLETED_36_40.md](UNIFIED_ROADMAP_COMPLETED_36_40.md).

---

## Phase 40: ServQueue Browser WASM Hardening & Multi-Tab Resilience (Completed)

All backlog tasks for Phase 40 have been fully completed, verified, and archived.
- For completed details of Phase 40: See [UNIFIED_ROADMAP_COMPLETED_36_40.md](UNIFIED_ROADMAP_COMPLETED_36_40.md).

---

## Phase 41: ServQueue Next-Gen Enterprise Stream Engine (Completed)

All backlog tasks for Phase 41 have been fully completed, verified, and archived.
- For completed details of Phase 41: See [UNIFIED_ROADMAP_COMPLETED_41_45.md](UNIFIED_ROADMAP_COMPLETED_41_45.md).

---

## Phase 42: ServQueue Beyond-Enterprise Security & Sovereign Stream Engine (New Initiative — Proposed Q2 2027)

> **Context:** Elevate ServQueue to financial and defense-grade sovereign streaming software featuring FIPS 140-3 boundary integration, blind-broker E2EE, post-quantum cryptography (PQC), Merkle audit ledgers, inline WASM AI guardrails, and eBPF kernel bypass.

### Detailed Items

| # | Item | Component | Description | Status |
|---|------|-----------|-------------|--------|
| SQ.S1 | **FIPS 140-3 & HSM Key Unsealing** | ServQueue Security | Native PKCS#11 HSM integration (AWS CloudHSM, YubiHSM2) for FIPS 140-3 key operations | [ ] |
| SQ.S2 | **Blind Broker End-to-End Encryption (E2EE)** | ServQueue Core | Producer-side payload encryption; broker indexes offsets without possessing decryption keys | [x] |
| SQ.S3 | **Post-Quantum Hybrid Cryptography (PQC)** | ServQueue Security | NIST Kyber768 key exchange & Dilithium signatures to protect against quantum decryption | [ ] |
| SQ.S4 | **Tamper-Evident Merkle Audit Ledger** | ServQueue Security | Cryptographic Merkle tree hash chain over admin commands & WAL segment commits for audit immutability | [x] |
| SQ.S5 | **Inline WASM AI Guardrails & Interceptor** | ServQueue WASM | Embedded ONNX/WASM AI filters to intercept prompt injections, PII leaks & data exfiltration in real-time | [ ] |
| SQ.S6 | **Byzantine Fault Tolerant (BFT) Consensus** | ServQueue Raft | BFT extensions for multi-region cloud deployments with automated partition healing | [ ] |
| SQ.S7 | **eBPF Kernel Bypass & XDP Acceleration** | ServQueue Network | eBPF XDP network socket bypass for ultra-low latency packet ingestion (<10µs p99 delivery) | [ ] |
| SQ.S8 | **SIMD / AVX-512 Vectorized Filter Engine** | ServQueue Core | SIMD batch matching for processing 10M+ events/sec per CPU core | [ ] |

---




## Appendix C: Architectural Policy for OSS/EE Boundaries



All commercial enterprise features (**EE**) must have their core logic and implementations located exclusively inside the private `servverse-ee` repository. 

The open-source core repositories (such as `ServGate`, `ServStore`, etc.) must only expose clean interfaces, hooks, or config fields. The implementation of these hooks in the open-source code must use build-tagged placeholders (`//go:build !enterprise`), while the actual commercial code resides under the corresponding directories in `servverse-ee` and is built with `//go:build enterprise`.











