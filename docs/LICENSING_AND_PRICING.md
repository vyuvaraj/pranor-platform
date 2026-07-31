# Pranor Pulse Licensing & Commercial Pricing Strategy

This document outlines the official licensing strategy, dual-licensing policy, client SDK permissions, and commercial tiering model for **Pranor Pulse**.

---

## 1. Executive Summary & Licensing Recommendation

Pranor Pulse uses a **Dual-Licensing Open-Core Model** designed to maximize open-source developer adoption while building a defensible, high-margin enterprise business.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           SERVQUEUE ECOSYSTEM                               │
└─────────────────────────────────────────────────────────────────────────────┘
          │                                 │                               │
          ▼                                 ▼                               ▼
┌───────────────────┐             ┌───────────────────┐           ┌───────────────────┐
│ Client SDKs & OPFS│             │ Pranor Pulse Server  │           │ Pranor Pulse EE      │
│  (@pranor/...) │             │   (`servqueued`)  │           │    (`serv-ee`)    │
├───────────────────┤             ├───────────────────┤           ├───────────────────┤
│     Apache 2.0    │             │      AGPLv3       │           │    Commercial     │
│   (Frictionless)  │             │ (Copyleft Core)   │           │ (Proprietary SLA) │
└───────────────────┘             └───────────────────┘           └───────────────────┘
```

### Recommendation on AGPLv3 (GNU Affero General Public License)
**Recommendation: RETAIN AGPLv3 for Server Engine, Use Apache 2.0 for Client SDKs.**

- **Why keep AGPLv3 for `servqueued` server engine?**
  1. **Hyperscaler Protection**: AGPLv3 prevents AWS, GCP, Azure, or third-party cloud vendors from hosting Pranor Pulse as a managed cloud service without contributing modifications back to the open-source community.
  2. **Strong Enterprise Commercial Conversion**: Companies that wish to embed or modify Pranor Pulse within closed-source SaaS applications or multi-tenant platforms are required under AGPL to release their source code—or purchase a **Pranor Pulse Enterprise Commercial License**.
  3. **Industry Standard Precedent**: Successfully proven by infrastructure leaders such as **MinIO**, **Grafana**, **RabbitMQ**, and **MongoDB (originally)**.

- **Why use Apache 2.0 / MIT for Client SDKs (`sdks/go`, `@pranor/queue-wasm`)?**
  1. Frontend web apps, backend microservices, and mobile clients importing Pranor Pulse libraries must **never** be subject to copyleft restrictions.
  2. Enables 100% frictionless integration into any proprietary enterprise application stack.

---

## 2. Licensing Matrix by Component

| Component | Repository Path | License | Commercial Exemption Option |
|---|---|---|---|
| **Pranor Pulse Core Server Daemon (`servqueued`)** | `serv/packages/Pranor Pulse` | **GNU AGPLv3** | Yes (Commercial License) |
| **Pranor Pulse Dual-CLI (`servqueue`)** | `serv/packages/Pranor Pulse/cmd/servqueue` | **GNU AGPLv3** | Yes (Commercial License) |
| **Local Browser OPFS WASM Engine (`@pranor/queue-wasm`)** | `serv/packages/Pranor Pulse/pkg/opfs` | **Apache 2.0** | Included in Apache 2.0 |
| **Go & Multi-Language Client SDKs** | `serv/packages/Pranor Pulse/sdks/*` | **Apache 2.0** | Included in Apache 2.0 |
| **Pranor Console Web Inspector & Admin UI** | `pranor-repo/servconsole` | **GNU AGPLv3** | Yes (Commercial License) |
| **Pranor Pulse Enterprise Commercial Engine (`serv-ee`)** | `serv-ee/src/Pranor Pulse` | **Commercial Proprietary** | Requires License Key |

---

## 3. Commercial Tiers & Feature Matrix

Pranor Pulse is structured into three clear commercial tiers:

| Feature / Module | Community (Free / AGPLv3) | Enterprise Tier ($30/core/mo) | Sovereign / Financial Tier ($60/core/mo) |
|---|:---:|:---:|:---:|
| **Core Broker Engine & STOMP / MQTT 5.0 Protocols** | ✅ | ✅ | ✅ |
| **Local-First Browser OPFS WASM Queue (`@pranor/queue-wasm`)** | ✅ | ✅ | ✅ |
| **Point-in-Time Event Replay & Poison-Pill DLQ Engine** | ✅ | ✅ | ✅ |
| **Prometheus `/metrics` & Basic Grafana Templates** | ✅ | ✅ | ✅ |
| **Cross-Cloud Active-Active Geo-Replication (WAN Sync)** (`SQ.E15`) | ❌ | ✅ | ✅ |
| **Kafka Wire Protocol Compatibility Adapter** (`SQ.E16`) | ❌ | ✅ | ✅ |
| **Multi-Cloud S3 / Pranor Vault Cold Tier Compaction** (`SQ.E20`) | ❌ | ✅ | ✅ |
| **AWS EventBridge & Enterprise Signed Webhooks (HMAC)** (`SQ.E21`) | ❌ | ✅ | ✅ |
| **FIPS 140-3 PKCS#11 HSM & Merkle Audit Ledger** (`SQ.E17`) | ❌ | ❌ | ✅ |
| **Post-Quantum Cryptography (NIST Kyber768/Dilithium)** (`SQ.E17`) | ❌ | ❌ | ✅ |
| **Inline WASM AI Guardrails & Interceptor (ONNX/WASM PII)** (`SQ.E18`) | ❌ | ❌ | ✅ |
| **eBPF Kernel Bypass & XDP Socket Offload (<10µs Latency)** (`SQ.E19`) | ❌ | ❌ | ✅ |
| **Multi-Cluster K8s Federation Operator & KEDA Auto-scaler** (`SQ.E22`) | ❌ | ❌ | ✅ |
| **Dedicated 24/7 SLA Support & Architecture Review** | ❌ | 8x5 Email | 24/7 Phone & Dedicated AM |

---

## 4. Commercial Pricing Models

### Model A: Per-Core CPU Subscription (Self-Hosted / On-Prem / Kubernetes)

Calculated based on the total number of vCPUs / CPU cores assigned to the `servqueued-ee` instances.

- **Community Tier**: **$0** (Free, AGPLv3 Open Source, Unlimited Cores).
- **Enterprise Tier**: **$30 / vCPU Core / Month** (billed annually at **$360 / core / year**).
  - *Example:* A 3-node cluster with 4 vCPUs per node (12 cores total) = **$4,320 / year**.
- **Sovereign & Defense Tier**: **$60 / vCPU Core / Month** (billed annually at **$720 / core / year**).
  - *Example:* A high-security 5-node cluster with 8 vCPUs per node (40 cores total) = **$28,800 / year**.

### Model B: Pranor Pulse Cloud (Managed Serverless SaaS)

For organizations seeking a fully managed cloud service without infrastructure overhead:

- **Data Ingestion & Egress**: $0.04 per GB transferred.
- **Hot Storage Buffer (SSD Log)**: $0.025 per GB / month.
- **Cold Storage Tier (S3 Archiving)**: $0.005 per GB / month.
- **WASM AI Guardrail Executions**: $0.001 per 1,000 payload checks.

---

## 5. Technical License Key Enforcement & Verification

In `serv-ee`, commercial feature modules are compiled behind the `//go:build enterprise` build tag.

### License Key Validation Flow

1. Pranor Pulse Enterprise daemon startup:
   ```bash
   servqueued-ee --config=/etc/servqueue/config.yaml --license-key=/etc/servqueue/license.lic
   ```
2. The daemon validates the cryptographically signed JWT / RSA license file:
   - **Payload Check**: Organization Name, Target Tier (`enterprise` vs `sovereign`), Max Core Limit, Expiration Date.
   - **Signature Verification**: Validated via offline public RSA key (no phone-home requirement for air-gapped sovereign environments).
3. If valid, `serv-ee` features (`GeoReplication`, `KafkaAdapter`, `HSMUnsealer`, `eBPFXDP`) activate seamlessly.

---

## 6. Summary Comparison: Licensing Choices

| Option | Open Source Engine | Client Libraries | Cloud Hyperscaler Risk | Enterprise Monetization Potential |
|---|---|---|---|---|
| **Recommended Strategy** | **AGPLv3** | **Apache 2.0** | **Low** (Hyperscalers must share SaaS code or buy license) | **Very High** (Direct conversion path to Commercial EE) |
| Pure Apache 2.0 | Apache 2.0 | Apache 2.0 | **High** (AWS can re-sell without contributing back) | Medium |
| BSL 1.1 (Business Source) | BSL 1.1 | Apache 2.0 | Low | High |

---

*Document updated: July 2026 | Version 2.0*
