# Component Maturity Analysis & Architectural Roadmap

This document analyzes the external feedback regarding the maturity of **Pranor Gate**, **Pranor Pulse**, and **Pranor Vault**, details the associated production risks, and proposes architectural mitigations divided into Open Source (OSS) and Enterprise (EE) domains.

---

## 1. Pranor Gate (API Gateway)

### Gaps Identified & Detailed Feedback
*   **Dynamic Upstream Discovery:** Currently relies on hardcoded JSON route maps. Production gateways require dynamic integration with service discovery registries (Consul, Kubernetes CoreDNS) to automatically detect when downstream services scale or crash.
*   **Distributed Rate Limiting:** The current localized rate limiter fails behind a round-robin load balancer. It needs a shared back-end state adapter (such as a Redis Sentinel cluster) using a sliding-window token bucket algorithm to enforce global API thresholds.
*   **Circuit Breaking & Outage Isolation:** Lacks automatic circuit breaking when a downstream service or queue stalls. Without this, pending connections back up, exhausting file descriptors and triggering cascading cluster failure.
*   **Security & Interconnect:** Needs robust mutual TLS (mTLS) with dynamic validation and multi-tenant certificate authority integration.

### Production Risk
*   Concurrent request storms will trigger memory/CPU spikes, and localized limiters will fail under load-balanced topologies.
*   A downstream failure or stall will cascade back to the gateway, causing file descriptor starvation and crashing the edge.

### Mitigation Plan
*   **[OSS] Rate Limiting:** Implement sliding-window rate limiting using local memory or Redis backend.
*   **[OSS] Circuit Breaker:** Add a basic circuit-breaker proxy state-machine (Closed, Open, Half-Open).
*   **[EE] Dynamic Discovery:** Integrate with Consul and Kubernetes CoreDNS for dynamic upstream registration.
*   **[EE] Distributed Rate Limiting:** Implement Redis Sentinel integration for shared token-bucket rate limiting.
*   **[EE] Advanced mTLS:** Dynamic certificate handshake and exchange with tenant-based validations.

---

## 2. Pranor Pulse (Message Queue)

### Gaps Identified & Detailed Feedback
*   **WASM Resource Sandboxing & Throttling:** Running WebAssembly via Wazero is fast, but a faulty user script with an infinite loop or high memory allocation will drain CPU cores and crash the host broker process. Needs strict runtime limiters to terminate slow WASM execution cycles.
*   **Split-Brain Prevention:** For multi-AZ clusters, the broker requires a replication coordinator. A network split will cause partition drift and duplicate message offset consumption without strict consensus.
*   **Dead Letter Queue (DLQ) Eviction Policies:** If a WASM data filter throws an exception or a consumer fails to acknowledge payloads repeatedly, messages must automatically offload to a DLQ with contextual metadata headers describing the failure.
*   **Memory Safety:** The WASM engine relies on raw `unsafe.Pointer` mappings, creating high vulnerability to crashes (segfaults) during out-of-bounds allocation or uncaught panics.
*   **Unbounded Memory Queues:** Internal buffers lack backpressure constraints. If producers flood the queue faster than consumers or filters process them, the broker consumes memory indefinitely until terminated by the OS OOM (Out-Of-Memory) killer.

### Production Risk
*   A single faulty WASM filter script or incorrect memory address calculation (via `unsafe.Pointer`) can trigger a native segmentation fault (SIGSEGV) and instantly crash the primary broker process.
*   High-throughput producer spikes will exhaust host memory (OOM) if consumers fall behind.
*   Network partition events will corrupt message logs or cause duplicate offset commits without a partition coordinator.

### Mitigation Plan
*   **[OSS] Safe WASM Runner:** Replace `unsafe.Pointer` memory mappings in the WASM runner with safe, bounds-checked slices and explicit memory copies.
*   **[OSS] WASM Execution Limits:** Add configurable execution timeouts (e.g., terminate filter if it takes longer than 50ms) to Wazero configuration.
*   **[OSS] Backpressure & Bounds:** Implement strict buffer limits on memory queues to block or throttle producers when consumer limits are reached.
*   **[OSS] Dead Letter Queue:** Implement a secondary DLQ eviction system with failure context headers.
*   **[EE] Distributed Consensus:** Implement Raft-based message replication across broker node clusters.
*   **[EE] Partition Resilience:** Add split-brain prevention and automated broker failover logic.

---

## 3. Pranor Vault (State Store)

### Gaps Identified & Detailed Feedback
*   **Formal Raft Consensus Verification:** Managing configuration tables requires linearizable consistency. Pranor Vault needs a verified consensus library (such as `hashicorp/raft`) to manage state mutations safely and prevent silent database corruption during server restarts.
*   **RBAC & TLS Interconnect:** To run securely in shared environments, all service-to-service communication paths must enforce mandatory mutual TLS (mTLS) certificate handshakes, paired with distinct write/read permissions for separate network keys.

### Production Risk
*   State synchronization bugs can silently corrupt metadata records, leading to systemic failures across downstream applications.
*   Lacking TLS interconnect and RBAC exposes sensitive configuration metadata to unauthorized internal nodes.

### Mitigation Plan
*   **[OSS] Lock Backend Stability:** Standardize interface abstraction layer for SQL/key-value storage backends.
*   **[EE] Audited Raft Integration:** Integrate an audited, industry-standard Raft implementation (`hashicorp/raft`) for state replication.
*   **[EE] TLS Interconnect & RBAC:** Enforce mutual TLS handshakes and RBAC permissions per cluster access token.

---

## 4. Ecosystem & Shared Middleware (Pranor Core)

### Gaps Identified & Detailed Feedback
*   **Naïve Error Propagation:** Critical error paths (such as database handshakes or network calls) are often handled by printing the error or immediately triggering a hard panic/exit, rather than using structured, resilient retry policies. 

### Production Risk
*   Momentary network blips, database restarts, or transient timeouts will cause downstream microservices to crash completely instead of gracefully waiting and reconnecting.

### Mitigation Plan
*   **[OSS] Resilient Retries:** Refactor `Pranor Core` database and HTTP client middleware to use standard retry adapters (e.g., exponential backoff) to recover from transient outages.
*   **[OSS] Structured Panic Recovery:** Enforce standard panic-recovery handlers in all HTTP and queue listeners to avoid dropping the process on individual request errors.

---

## Architecture Verification Checklist

To gauge if the infrastructure is production-ready, verify the following capabilities:

- [ ] **State Resiliency:** Can I pull the power cord on 1 out of 3 running Pranor Vault nodes without corrupting active configurations?
- [ ] **Edge Protection:** Does Pranor Gate reject traffic smoothly with an HTTP 429 error when hit by a simulated DDoS attack?
- [ ] **WASM Isolation:** Does Pranor Pulse terminate a WASM data filter if it takes longer than 50ms to run?
- [ ] **Ecosystem Resilience:** Does a momentary network split or database connection timeout trigger an automatic retry (with backoff) rather than a hard crash/panic?

