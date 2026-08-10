# Serv Unified Ecosystem Roadmap & Architect Analysis


> Single source of truth for the **Serv** ecosystem: Pranor, Pranor Gate, Pranor Vault, Pranor Pulse, Pranor Console, Pranor Cache, Pranor Mesh, Pranor Chrono, Pranor Deploy, Pranor Trace, Pranor Tunnel, Pranor Auth, Pranor Pool, Pranor Notify, Pranor Flow, and the Pranor vision.  

> Last updated: August 10, 2026


---


## Ecosystem Completion Status


All items in Phases 1 through 87 have been fully implemented, verified, and pushed.


- For completed details of Phases 1 to 5: Refer to the git history and repository CHANGELOG.
- For completed details of Phases 6 to 10: See [UNIFIED_ROADMAP_COMPLETED_6_10.md](UNIFIED_ROADMAP_COMPLETED_6_10.md).
- For completed details of Phases 11 to 15: See [UNIFIED_ROADMAP_COMPLETED_11_15.md](UNIFIED_ROADMAP_COMPLETED_11_15.md).
- For completed details of Phase 16-20: See [UNIFIED_ROADMAP_COMPLETED_16_20.md](UNIFIED_ROADMAP_COMPLETED_16_20.md).
- For completed details of Phase 21-25: See [UNIFIED_ROADMAP_COMPLETED_21_25.md](UNIFIED_ROADMAP_COMPLETED_21_25.md).
- For completed details of Phase 26-30: See [UNIFIED_ROADMAP_COMPLETED_26_30.md](UNIFIED_ROADMAP_COMPLETED_26_30.md).
- For completed details of Phase 31-35: See [UNIFIED_ROADMAP_COMPLETED_31_35.md](UNIFIED_ROADMAP_COMPLETED_31_35.md).
- For completed details of Phase 36-40: See [UNIFIED_ROADMAP_COMPLETED_36_40.md](UNIFIED_ROADMAP_COMPLETED_36_40.md).
- For completed details of Phase 41-45: See [UNIFIED_ROADMAP_COMPLETED_41_45.md](UNIFIED_ROADMAP_COMPLETED_41_45.md).
- For completed details of Phase 46-50: See [UNIFIED_ROADMAP_COMPLETED_46_50.md](UNIFIED_ROADMAP_COMPLETED_46_50.md).
- For completed details of Phase 51-55: See [UNIFIED_ROADMAP_COMPLETED_51_55.md](UNIFIED_ROADMAP_COMPLETED_51_55.md).
- For completed details of Phase 56-60: See [UNIFIED_ROADMAP_COMPLETED_56_60.md](UNIFIED_ROADMAP_COMPLETED_56_60.md).
- For completed details of Phase 61-65: See [UNIFIED_ROADMAP_COMPLETED_61_65.md](UNIFIED_ROADMAP_COMPLETED_61_65.md).
- For completed details of Phase 66-70: See [UNIFIED_ROADMAP_COMPLETED_66_70.md](UNIFIED_ROADMAP_COMPLETED_66_70.md).
- For completed details of Phase 71-75: See [UNIFIED_ROADMAP_COMPLETED_71_75.md](UNIFIED_ROADMAP_COMPLETED_71_75.md).
- For completed details of Phase 76-80: See [UNIFIED_ROADMAP_COMPLETED_76_80.md](UNIFIED_ROADMAP_COMPLETED_76_80.md).
- For completed details of Phase 81-85: See [UNIFIED_ROADMAP_COMPLETED_81_85.md](UNIFIED_ROADMAP_COMPLETED_81_85.md).
- For completed details of Phase 86-87: See [UNIFIED_ROADMAP_COMPLETED_86_87.md](UNIFIED_ROADMAP_COMPLETED_86_87.md).

---

## v2.0 Branch Strategy

> **Decision (2026-08-10):** Pranor v1.0 is not yet released. All Phase 89/90 v2.0 feature work is developed on a dedicated `v2.0-dev` branch in `pranor` and `pranor-ee`. The `main` branch is frozen for v1.0 stabilization and release.

| Repository | v1.0 Branch | v2.0 Branch | Notes |
|-----------|-------------|-------------|-------|
| `pranor` | `main` ← **freeze for v1.0** | `v2.0-dev` ✅ created | Sprint 2–7 feature work here |
| `pranor-ee` | `main` ← **freeze for v1.0** | `v2.0-dev` ✅ created | EE stubs for Graph, Decision, Flow |
| `pranor-platform` | `main` | `main` | Docs/roadmap — no feature code |

### Branch Rules
- **`main`**: v1.0 bug fixes, security patches, and release cuts only. No new Phase 89/90 modules.
- **`v2.0-dev`**: All Phase 89/90 Sprint 2–7 work (`graph/`, `decision/`, `trace/pkg/schema/`, `flow/pkg/agentstep/`, `learn/`).
- **Sprint 1 exception**: `scripts/check_cgo.sh`, `.github/workflows/cgo_check.yml`, and `_templates/` were committed to `main` (pure infrastructure benefiting v1.0 CI). They are inherited by `v2.0-dev` via branch creation.
- **Merge criteria**: `v2.0-dev → main` only after v1.0 is tagged (`v1.0.0`), all v2.0 tests pass, and a PR review is completed.

---

### Completion Tracker


| Initiative Area | Total Items | Completed | Pending | Progress | Status Bar |
|-----------------|-------------|-----------|---------|----------|------------|
| **Phase 81: Rebrand — Serv/Pranor → Pranor** | 87 | 87 | 0 | **100%** | ████████████████████ |
| **Phase 82: Documentation Consolidation** | 15 | 15 | 0 | **100%** | ████████████████████ |
| **Phase 83: Pranor LSP & IDE Intelligence Evolution** | 6 | 6 | 0 | **100%** | ████████████████████ |
| **Phase 84: Pranor VS Code Extension Ecosystem** | 6 | 6 | 0 | **100%** | ████████████████████ |
| **Phase 85: Enterprise Commercial Tier Expansion** | 21 | 21 | 0 | **100%** | ████████████████████ |
| **Phase 86: Advanced Enterprise Security & Operations** | 23 | 23 | 0 | **100%** | ████████████████████ |
| **Phase 87: Enterprise AI Sovereign Data & Resiliency** | 23 | 23 | 0 | **100%** | ████████████████████ |
| **Phase 88: Next-Gen AI Agent Security & Execution Governance** | 10 | 10 | 0 | **100%** | ████████████████████ |
| **Phase 89: Pranor v2.0 Core Governed Execution Fabric** | 7 | 5 | 2 | **71.4%** | ██████████████░░░░░░ |
| **Phase 90: Pranor v2.x Intelligence Extensions & ML Providers** | 5 | 1 | 4 | **20%** | ████░░░░░░░░░░░░░░░░ |

| **TOTAL ECOSYSTEM CODE WORK** | **787** | **780** | **7** | **99.1%** | ███████████████████░ |


---


## Phase 86: Advanced Enterprise Security, High-Availability & Operations Moats (Completed)

All 23 backlog items (EE.86.1 through EE.86.23) for Phase 86 have been fully completed, verified, and archived.
- For completed details of Phase 86: See [UNIFIED_ROADMAP_COMPLETED_86_87.md](UNIFIED_ROADMAP_COMPLETED_86_87.md).

---

## Phase 87: Enterprise AI Sovereign Data, Resiliency & Compliance Engineering (Completed)

All 23 backlog items (EE.87.1 through EE.87.23) for Phase 87 have been fully completed, verified, and archived.
- For completed details of Phase 87: See [UNIFIED_ROADMAP_COMPLETED_86_87.md](UNIFIED_ROADMAP_COMPLETED_86_87.md).

---

## Phase 88: Next-Gen AI Agent Security & Execution Governance (Completed)

This phase establishes Pranor Gate as a programmable, zero-trust execution control plane for AI agents, governing tool invocations, identity chains, execution risks, budgets, and HITL approvals.

| # | Item | Component | Open-Source (OSS) Scope | Enterprise (EE) Scope | Selected / Priority |
|---|------|-----------|-------------------------|------------------------|---------------------|
| **EE.88.1** | **AI Agent Security Firewall** | Pranor Gate | Single-node rule engine inspecting tool calls/intents for `ALLOW / DENY / TRANSFORM` decisions. | AI Self-defending WAF integration & global policy distribution across multi-cloud clusters. | ✅ **Completed** |
| **EE.88.2** | **Agent Security Chain & Identity Propagation** | Pranor Gate / Auth | Local JWT & SPIFFE `Agent ID → User ID → Tenant ID → Capability` context propagation. | Multi-tenant directory sync, enterprise IAM mapping & automated credential rotation. | ✅ **Completed** |
| **EE.88.3** | **Human-in-the-Loop (HITL) Execution Engine** | Pranor Gate | Basic REST API & CLI webhook approval endpoints for manual intervention. | Slack/Teams interactive approvals, RBAC escalation paths, SLA timers & audit vault. | ✅ **Completed** |
| **EE.88.4** | **Agent Trajectory Replay & Simulation Engine** | Pranor Gate | Local CLI `pranor-gate replay` tool for trajectory diffs against candidate models/tools. | Parallel batch simulation across test datasets & automated regression reporting in Console UI. | ✅ **Completed** |
| **EE.88.5** | **AI Capability Risk & Trust Engine** | Pranor Gate | Static tool risk scoring rules (LOW/MEDIUM/HIGH per route/tool). | Real-time ML risk scoring evaluating behavioral anomalies, user trust metrics & context. | ✅ **Completed** |
| **EE.88.6** | **Agent Blast-Radius & Tool Budget Enforcer** | Pranor Gate | Per-instance tool count & session rate limits. | Distributed Redis/Cluster-wide budget enforcement across multi-cloud deployments. | ✅ **Completed** |
| **EE.88.7** | **Agent Memory & Context Governance** | Pranor Gate | Standard PII regex detection, data masking & parameter redaction. | ML-driven data classification, sensitivity tagging & enterprise DLP policies. | ✅ **Completed** |
| **EE.88.8** | **Agent-to-Agent Delegation Governance** | Pranor Gate | Point-to-point agent-to-agent authorization & token passing. | Multi-agent mesh control plane, delegation lineage tracking & root-cause audit logs. | ✅ **Completed** |
| **EE.88.9** | **Protocol-Agnostic Capability Exposer** | Pranor Gate | Native MCP, gRPC, and HTTP/REST adapters for single nodes. | GraphQL federation + automatic multi-protocol schema translation & registration. | ✅ **Completed** |
| **EE.88.10** | **End-to-End Agent Execution Trace Visualization** | Pranor Gate / Trace | Standard OpenTelemetry spans emitted to local OTLP collectors. | Full visual multi-step agent execution tree UI in Pranor Console dashboard. | ✅ **Completed** |

---

## Phase 89: Pranor v2.0 Core Governed Execution Fabric (Planned / Active Phase)

Guided by `requirements_definitive.md` (§2.1–§2.6, §3) and the Architectural Litmus Test, Phase 89 establishes Pranor's core AI-native execution runtime — all modules are pure-Go, CGO-free, and split by OSS / EE build tags.

**Selected for active implementation:** V2.89.7, V2.89.4, V2.89.1, V2.89.3 (Sprint 1–4). V2.89.2, V2.89.5, V2.89.6 follow as unblocked dependents.

| # | Item | Component | Open-Source (OSS) Scope | Enterprise (EE) Scope | Sprint / Priority | Status |
|---|------|-----------|-------------------------|------------------------|-------------------:|--------|
| **V2.89.7** | **Zero-CGO Invariant CI Pipeline (`scripts/check_cgo.sh`)** | Pranor Core | Automated CI script: `CGO_ENABLED=0` cross-compile matrix (`linux/amd64`, `arm64`, `darwin`, `windows`); `ldd` static-link verification; `go list` CGO source scan. | FIPS 140-3 compliant HSM crypto verification build stage. | 🥇 **Sprint 1 — P0** | ✅ **Completed** |
| **V2.89.4** | **Pranor Trace — OTLP Span Schema (`std/trace`)** | Pranor Trace | Canonical `pranor.agent_execution` root span hierarchy across Gate, Graph, Decision, Flow, and Learn modules; mandatory span attributes (`agent_id`, `user_id`, `tenant_id`, `request_id`, `module`, `outcome`); best-effort emission (never on critical path). | Tail-based sampling; long-term trace archive; `Eval` replay query API. | 🥈 **Sprint 2 — P0** | ✅ **Completed** |
| **V2.89.1** | **Pranor Graph (`std/graph`)** | Pranor Core | Virtual entity context layer linking `Pool`, `Vault`, and `Pulse`; Hot materialized cache (`< 2ms`); fail-closed on backend unavailability (`ErrGraphContextUnavailable`). | Cross-datacenter entity graph synchronization & RBAC tenant isolation. | 🥉 **Sprint 3 — P0** | ✅ **Completed** |
| **V2.89.3** | **Pranor Flow — AgentStep & Saga Engine (`std/flow`)** | Pranor Flow | Idempotent, compensatable `AgentStep` primitive; `SagaConfig` with `MaxSteps`, `StepTimeout`, `TotalTimeout`; `COMPENSATE` and `PAUSE_FOR_HITL` limit policies; reverse compensation unwind on failure. | Distributed saga state replication (Raft); multi-region saga recovery & failover. | 4️⃣ **Sprint 4 — P0** | ✅ **Completed** |
| **V2.89.5** | **Pranor Graph — Fault Tolerance Contract** | Pranor Graph | Transparent Hot→Warm fallback with `graph.cache_miss` span; fail-closed `ErrGraphContextUnavailable` on all-tier exhaustion; no fail-open permitted. | Same contract applied across multi-datacenter replicas. | 5️⃣ Sprint 5 — P1 (unblocks after V2.89.1) | Planned |
| **V2.89.2** | **Pranor Decision Engine (`std/decision`)** | Pranor Core | 6-level priority veto ladder: Authorization > Budget > Risk > Rules > Learn > Default. Hard `DENY` on Auth/Budget; `APPROVE`/`DENY` on Risk; soft advisory on Learn. | Enterprise rule distribution & real-time risk model synchronization. | 6️⃣ Sprint 6 — P1 (unblocks after V2.89.1) | Planned |
| **V2.89.6** | **Pranor Decision — Fault Tolerance Contract** | Pranor Decision | `DENY` with `ERR_CONTEXT_UNAVAILABLE` when Graph unavailable; skip Learn (Priority 5) when sidecar down; `SIMULATION` mode records decisions without committing side-effects. | Same contract enforced across distributed rule clusters. | 7️⃣ Sprint 7 — P1 (unblocks after V2.89.2) | Planned |

---

## Active Sprint Execution Order (Phase 89 + 90 Selected Items)

> Items are sequenced by hard dependency order and architectural impact. P0 items block all downstream work; P1 items unblock once their dependency sprint is complete.

```
┌─────────────────────────────────────────────────────────────────────────┐
│  SPRINT 1 — Guardrails (P0)                                             │
│  V2.89.7  Zero-CGO CI Pipeline      ← Safety net for all future work   │
│  V2.90.5  OSS/EE Build-Tag Scaffold  ← Pattern set before any modules  │
├─────────────────────────────────────────────────────────────────────────┤
│  SPRINT 2 — Shared Instrumentation (P0)                                 │
│  V2.89.4  Pranor Trace OTLP Schema   ← All modules wire spans here     │
├─────────────────────────────────────────────────────────────────────────┤
│  SPRINT 3 — Core Dependency (P0)                                        │
│  V2.89.1  Pranor Graph               ← Decision, Flow, Gate all depend  │
├─────────────────────────────────────────────────────────────────────────┤
│  SPRINT 4 — Differentiating Capability (P0)                             │
│  V2.89.3  Pranor Flow (AgentStep)    ← Execution fabric differentiator  │
├─────────────────────────────────────────────────────────────────────────┤
│  SPRINT 5–7 — Dependent Modules (P1, unblock after Sprint 3)            │
│  V2.89.5  Graph Fault Tolerance      ← Completes V2.89.1               │
│  V2.89.2  Pranor Decision Engine     ← Requires Graph (V2.89.1)        │
│  V2.89.6  Decision Fault Tolerance   ← Completes V2.89.2               │
└─────────────────────────────────────────────────────────────────────────┘

  Phase 90 items begin after Sprint 4 core is stable:
  V2.90.1  Pranor Learn   (v2.x — after Trace + Decision complete)
  V2.90.2  Pranor Eval    (v2.x — after Learn + Trace complete)
  V2.90.3  Simulation     (v2.x — after Decision complete)
  V2.90.4  HITL Queue     (v2.x — after Flow complete)
```

---

## Phase 90: Pranor v2.x Intelligence Extensions & ML Providers (Planned / Future Phase)

Establishes the extensible provider framework, OSS/EE build-tag discipline, and evaluation suite for external machine learning runtimes.

**Selected for active implementation:** V2.90.5 (Sprint 1 alongside Phase 89). V2.90.1–V2.90.4 begin after Phase 89 Sprint 4 core is stable.

| # | Item | Component | Open-Source (OSS) Scope | Enterprise (EE) Scope | Sprint / Priority | Status |
|---|------|-----------|-------------------------|------------------------|-------------------:|--------|
| **V2.90.5** | **OSS / EE Build-Tag Contract (All v2.x Modules)** | Pranor Core | Enforced `//go:build !enterprise` stubs returning `ERR_EE_REQUIRED`; shared interface files carry no build tags; OSS stubs API-compatible with EE implementations. | EE build path (`//go:build enterprise`) in `pranor-ee` repo across all v2.x modules. | 🥇 **Sprint 1 — P0** | ✅ **Completed** |
| **V2.90.1** | **Pranor Learn Provider Architecture (`std/learn`)** | Pranor Learn | `Predictor` interface (`Predict`, `HealthCheck`); typed `PredictInput` / `PredictOutput` structs with schema validation; pure-Go WASM (`wazero`) and gRPC sidecar drivers; `ErrSidecarTimeout` / `ErrModelBudgetExceeded` fault contracts. | GPU-accelerated PyTorch/TabPFN sidecar pool & multi-cluster model routing. | Post-Phase-89 — P1 | Planned |
| **V2.90.3** | **Decision Simulation Engine ("What-If" Analysis)** | Pranor Decision | Counterfactual decision evaluation in `SIMULATION` mode without committing backend state to any module. | Multi-variant decision A/B testing & simulation analytics dashboard. | Post-Phase-89 — P1 | Planned |
| **V2.90.4** | **Interactive HITL Approval Queue (`Pranor Console`)** | Pranor Console | Basic Webhook & REST approval endpoints for manual gate interventions. | Interactive Slack, Microsoft Teams, and Email approval workflows with SLA timers. | Post-Phase-89 — P1 | Planned |
| **V2.90.2** | **Pranor Eval Framework (`std/eval`)** | Pranor Eval | Trajectory trace replay engine and score evaluators (accuracy, latency, cost, safety). | Automated CI/CD quality gate blocking regressions before deployment. | Post-Phase-89 — P2 (requires Learn) | Planned |

---

## Deferred Community / Non-Code Items

| # | Item | Component | Description | Status |
|---|------|-----------|-------------|--------|
| AG.4 | **10-minute demo video** | pranor-repo | Screen recording: install → write service → deploy → observe in console. Hosted on YouTube + embedded in GitHub Pages | [Deferred] |
| AG.5 | **Discord/community server** | - | Developer community for questions, showcases, and contributors | [Deferred] |
| AG.12 | **Customer pilot program** | - | Find 2-3 teams to run in staging. Gather real feedback on DX, performance, gaps | [Deferred] |
| V1.9 | **API freeze period** | - | 4 weeks with zero breaking changes after all V1.1-V1.6 are done | [Deferred] |
