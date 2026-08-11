# Pranor v2.0 — Governed AI Agent Execution Fabric

> **Published:** August 2026 | **Reading Time:** ~12 min | **Topic:** AI Agent Governance & Architecture

---

## Introduction

Building prototype AI agents with LLMs is easy. Running autonomous, multi-agent workflows in production microservices with strict cost limits, zero security breaches, and transactional rollback guarantees is remarkably hard.

**Pranor v2.0** extends the 17 core infrastructure modules of Pranor v1.0 with a **Governed AI Execution Fabric** written in pure, CGO-free Go (`CGO_ENABLED=0`).

In this post, we'll explore the architectural building blocks of Pranor v2.0.

---

## The 6-Level Priority Veto Ladder (`std/decision`)

Every tool invocation or capability request made by an AI agent passes through a 6-stage immutable veto ladder:

```
Level 1: Auth       (Hard DENY) ──> RBAC & tenant permissions
Level 2: Budget     (Hard DENY) ──> Daily token & USD cost limits
Level 3: Risk       (Soft/Hard) ──> Blast radius & security risk score
Level 4: Rules      (Transform) ──> Business policy rules
Level 5: Learn      (Advisory)  ──> ML predictor sidecar
Level 6: Default    (Fallback)  ──> Default ALLOW
```

If an agent attempts an unauthorized file deletion or exceeds its $0.50 per-request cost budget, **Level 1 or Level 2 immediately hard-denies the action** before any side-effect occurs.

---

## 3-Tier Entity Context Assembly (`std/graph`)

LLM performance degrades rapidly when inundated with low-relevance context. Pranor Graph (`std/graph`) assembles virtual entity context across three tiers:

- **Hot Tier (<2ms)**: Volatile in-session scratchpad memory.
- **Warm Tier (<15ms)**: Virtual SQL join across `pranor/pool` connection proxies.
- **Cold Tier (<100ms)**: S3 object storage fallback (`pranor/vault`).

Nodes are dynamically pruned based on `MaxTokenBudget` and semantic relevance, guaranteeing a **fail-closed contract** — if all tiers are exhausted, an explicit error is returned rather than sending corrupted data to an LLM.

---

## Transactional Agent Sagas (`std/flow`)

Multi-step agent workflows are managed by the `AgentStep` Saga Engine. If a step times out or fails (e.g., payment fails at Step 3), completed steps (Inventory Reservation at Step 1, Coupon Application at Step 2) are unwound in reverse order using `Compensate()`.

For policy boundary breaches, the workflow safely pauses and enqueues an `ApprovalRequest` to the **Human-In-The-Loop (HITL)** approval queue via REST or Slack webhooks.

---

## Summary of New v2.0 Modules

| Module | Scope |
|--------|-------|
| `std/graph` | Virtual 3-tier entity context assembly & token budget node pruning |
| `std/decision` | 6-level AI governance veto ladder & SIMULATION mode |
| `std/memory` | Working memory scratchpad & episodic Cosine Similarity vector recall |
| `std/eval` | Trajectory trace replay & 4 score evaluators (accuracy, latency, cost, safety) |
| `std/learn` | Pluggable ML inference provider (wazero WASM + gRPC sidecar) |
| `std/agent` | AgentSpec declarative registry & 7-state runtime state machine |
| `std/execctx` | Immutable execution context propagation across RPC & HTTP headers |
| `std/capability` | Governed tool registry with semver, risk classes, & rate limits |
| `agent/pkg/a2a` | Agent-to-Agent delegation protocol & escalation guard |
| `std/llm` | Model gateway router, HTTP/Echo drivers, & native token streaming |
| `gate/pkg/guardrails` | PII masking, prompt injection scanner, secret leak detector |
| `gate/pkg/shadow` | Simulation side-effect isolation & interceptor sandbox |
| `core/pkg/tenant` | Multi-tenant sandboxing & daily token/cost quota enforcer |
| `tools/agentctl` | Developer CLI for trace replay & policy dry-runs |

---

## Next Steps

- 📘 Explore the [Pranor v1.0 Documentation](https://vyuvaraj.github.io/pranor/v1.0/)
- 🚀 Explore the [Pranor v2.0 AI Execution Fabric Documentation](https://vyuvaraj.github.io/pranor/v2.0/)
- ⭐ Star the project on [GitHub](https://github.com/vyuvaraj/pranor)
