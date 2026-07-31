# Pranor Pulse v2: Evolving Our WASM Message Broker into a Monorepo Stream Engine

*From a single-binary WASM broker to a multi-protocol streaming engine featuring dual CLI binaries, MQTT v5.0, Kafka compatibility, point-in-time replay, and CRDT active-active geo-replication.*

---

> 💡 **Note**: This is **Part 2** of the Pranor Pulse series. If you missed Part 1 on how we built inline WASM stream processing from scratch, check out [Part 1: I Built a Message Broker With Inline WASM Stream Processing From Scratch](https://medium.com/@yuvamca002/i-built-a-message-broker-with-inline-wasm-stream-processing-from-scratch-2196a391cfac).

---

## Why Pranor Pulse Needed to Evolve

In Part 1, we introduced Pranor Pulse: a Go-based message broker executing inline WebAssembly (WASM) filters inside the dispatch loop.

While the core WASM stream processing engine worked seamlessly, running it in real-world infrastructure revealed key operational challenges:

1. **Mixed Binary Responsibilities**: The client CLI commands and broker server daemon were bundled together, making daemon deployment in Docker/K8s heavier than necessary.
2. **Protocol Barriers**: Developers wanted IoT devices to stream telemetry via MQTT and existing microservices to produce via Kafka without needing external bridge proxy containers.
3. **Disaster Recovery**: Consumers needed the ability to rewind and replay stream history back to exact timestamps (`seekToTime`) rather than just raw offset numbers.
4. **Cross-Cloud Active-Active Sync**: Distributed clusters across multiple cloud regions required conflict-free multi-primary event mirroring.

Here is how we addressed these challenges in **Pranor Pulse v2** and migrated to the unified **Serv monorepo** (`github.com/vyuvaraj/pranor`).

---

## 1. Monorepo Migration & The Daemon / CLI Split

We merged standalone services into the unified **Serv monorepo** (`github.com/vyuvaraj/pranor/packages/Pranor Pulse`). As part of this migration, we strictly separated server runtime logic from client administration tooling:

* **`pranor-pulsed` (Server Daemon)**: Zero-dependency background service process. It hosts dedicated listeners for STOMP (`:61613`), MQTT v5.0 (`:1883`), Kafka (`:9092`), HTTP REST (`:8082`), Prometheus metrics (`/metrics`), and an embedded Web Admin UI (`http://localhost:8082/ui/`) served via Go `embed`.
* **`pranor-pulse` (Client CLI)**: Fast-booting administrative binary for operators and scripts (`status`, `topics`, `publish`, `consume`, `tail`, `seek`).

---

## 2. Multi-Protocol Architecture & Dedicated Ports

A common point of confusion with multi-protocol brokers is whether all clients connect to the same port. 

In `pranor-pulsed`, **each protocol listens on its standard dedicated TCP port**, but all protocol handlers funnel into the **same underlying core message broker engine**. This enables seamless **Cross-Protocol Fan-Out**: an IoT sensor publishing via MQTT on port `1883` can be read by a STOMP subscriber on port `61613`, a Kafka client on port `9092`, or inspected via the HTTP Web UI on port `8082`.

```
PRODUCER CLIENTS                                pranor-pulsed DAEMON                                CONSUMER CLIENTS
────────────────                               ─────────────────                                ────────────────
[ STOMP Client ]   ─── tcp://localhost:61613 ──► ┌──────────────────────────┐ ─── tcp://localhost:61613 ──► [ STOMP Subscriber ]
                                                 │ STOMP Server (:61613)   │
[ MQTT Device ]    ─── tcp://localhost:1883  ──► │ MQTT Gateway (:1883)    │ ─── tcp://localhost:1883  ──► [ MQTT Subscriber ]
                                                 │ Kafka Adapter (:9092)   │
[ Kafka Producer ] ─── tcp://localhost:9092  ──► │ HTTP REST API (:8082)   │ ─── tcp://localhost:9092  ──► [ Kafka Consumer ]
                                                 └────────────┬─────────────┘
[ REST / CLI ]     ─── http://localhost:8082 ──►              │
                                                              ▼
                                               ┌────────────────────────────┐
                                               │ Core Engine & Topic Router │
                                               ├────────────────────────────┤
                                               │ ⚙️ WASM Stream Processor   │
                                               │ ⏪ Point-in-Time Replay    │
                                               │ 🌐 CRDT Geo-Replication     │
                                               │ ☠️ Dead Letter Queue (DLQ) │
                                               └────────────────────────────┘
```

### Port Mapping Summary

| Protocol | Default Port | Primary Use Case |
|:---|:---|:---|
| **STOMP** | `:61613` | Lightweight pub/sub messaging across languages |
| **MQTT v5.0** | `:1883` | Low-power IoT telemetry & edge sensor publishing |
| **Kafka Wire** | `:9092` | Direct producer/consumer integration with Kafka client SDKs |
| **HTTP REST / UI** | `:8082` | Management API, Web Admin UI (`/ui/`), & Prometheus `/metrics` |

---

## 3. Point-in-Time Event Replay (`seekToTime`)

When downstream consumers crash or data corruptions occur, seeking by relative offset count is error-prone. 

Pranor Pulse v2 introduces timestamp-based stream seeking (`seekToTime`). The log engine indexes entry timestamps, enabling instant seeks via the CLI or HTTP API:

```bash
# Seek consumer offset to 15 minutes ago
pranor-pulse seek orders.events 15m

# Seek to an explicit ISO-8601 timestamp
pranor-pulse seek orders.events 2026-07-26T05:00:00Z
```

The server returns the exact target offset and timestamp, allowing consumers to resume processing safely.

---

## 4. Cross-Cloud Active-Active Geo-Replication (CRDTs)

For multi-region deployments, Pranor Pulse v2 introduces background cluster mirroring powered by **Last-Write-Wins (LWW) CRDTs**.

When nodes in `us-east-1` and `eu-west-1` operate concurrently, the mirror engine resolves concurrent state updates across cluster boundaries without central locks, ensuring eventual consistency even across temporary network partitions.

## 5. Local-First Browser OPFS Queue (`@pranor/queue-wasm`)

Beyond server-side streaming, Pranor Pulse extends event brokerage directly into the browser for local-first and offline-first Progressive Web Apps (PWAs):

* **Embedded Web Worker Storage**: Powered by `@pranor/queue-wasm`, Pranor Pulse runs an embedded event log inside browser Web Workers leveraging the Origin Private File System (`FileSystemSyncAccessHandle`) for high-throughput local disk storage.
* **Offline Outbox & Reconnect Relay**: When a browser client loses network connectivity, published events are stored locally in the OPFS WAL. Upon network reconnection, an outbox relay automatically streams unacknowledged event ranges in exact sequence to the remote `pranor-pulsed` cluster.

🔗 **Interactive Live Demo**: Test local-first browser queueing and outbox replay live in your browser: [Launch Pranor Pulse OPFS Interactive Demo](https://vyuvaraj.github.io/pranor/playground/opfs_demo.html).

---

## 6. Cloud-Native Operations & K8s Ecosystem

To operationalize Pranor Pulse in cloud environments, we built native Kubernetes controllers and monitoring exporters:

* **Kubernetes Operator (`Pranor PulseCluster`)**: A custom CRD controller managing cluster replica deployment, state reconciliation, and automated pod failover.
* **KEDA Metrics Scaler**: Exposes topic consumer lag metrics to Kubernetes Event-driven Autoscaling (KEDA) to scale consumer pod replicas dynamically.
* **Prometheus & Grafana**: Native `/metrics` endpoint with pre-built Grafana dashboard templates (`grafana_dashboard.json`) tracking throughput, queue depth, consumer lag, and WASM transform execution latency.
* **Automated Chaos Injector**: Built-in testing harness (`pkg/testing/chaos_injector.go`) for injecting network latency, partition failures, and disk corruption during integration testing.

---

## 7. Multi-Language SDKs & Standard STOMP Clients

Pranor Pulse v2 provides official client packages under `packages/Pranor Pulse/sdks/`:

* **Go**: `import "github.com/vyuvaraj/pranor/packages/Pranor Pulse/sdks/go"`
* **TypeScript / Node.js**: `const { Pranor PulseClient } = require('@pranor/queue-sdk');`
* **Python**: `from pranor-pulse import Pranor PulseClient`
* **Browser WASM**: `@pranor/queue-wasm` for embedded Web Worker event streaming using OPFS (`FileSystemSyncAccessHandle`).

### Zero-Lock-In: Any Standard STOMP Client Works Out-of-the-Box
Because Pranor Pulse listens natively on STOMP port `:61613`, you are never locked into custom SDKs. Any standard open-source STOMP library across language ecosystems connects seamlessly:
* **Java / Spring**: `spring-messaging` / `StompClient`
* **Python**: `stomp.py`
* **Node.js**: `@stomp/stompjs`
* **C# / .NET**: `Apache.NMS.Stomp`
* **Ruby**: `stomp` gem

---

---

## 8. Platform Differentiating Factors: Pranor Pulse vs. Pranor Vault vs. Pranor Gateway

A frequent question from developers evaluating the **Pranor Platform** (`github.com/vyuvaraj/pranor`) is what makes each engine uniquely disruptive compared to traditional cloud infrastructure:

```
┌──────────────────────────────────────────────────────────────────────────────────────────────────┐
│                                    PRANOR PLATFORM ECOSYSTEM                                  │
├──────────────────────────────┬───────────────────────────────────┬───────────────────────────────┤
│    ⚡ Pranor Pulse (Messaging)   │     📦 Pranor Vault (Storage)        │    🌐 Pranor Gateway (Edge AI)   │
├──────────────────────────────┼───────────────────────────────────┼───────────────────────────────┤
│ • Browser OPFS WAL Outbox    │ • Browser OPFS Dual-Sync (`.wasm`)│ • Dual Server/Browser WASM    │
│ • STOMP, MQTT v5, Kafka Wire │ • Native Embedded DuckDB/Parquet  │ • Edge AI Token (TPM) Limits  │
│ • Point-in-Time Replay       │ • K+M Reed-Solomon Erasure Code   │ • Semantic Prompt Caching     │
│ • CRDT Geo-Replication       │ • SEC Rule 17a-4 WORM Object Lock │ • eBPF XDP Kernel DDoS Bypass │
└──────────────────────────────┴───────────────────────────────────┴───────────────────────────────┘
```

### 📦 Pranor Vault: Differentiating Factors (vs. S3 / MinIO / Cloudflare R2)

1. **Client-Side OPFS Browser Dual-Sync (`@pranor/store-wasm`)**:
   - *The Problem with MinIO / S3*: Traditional object stores require continuous network connection for direct uploads. Poor mobile connectivity drops uploads and blocks the UI.
   - *The Pranor Vault Solution*: `@pranor/store-wasm` leverages origin-private filesystem (**OPFS**) inside browser Web Workers. Web apps save gigabytes of media/files **instantly at native NVMe disk speed** into browser OPFS. In the background, Pranor Vault streams chunked S3 multipart uploads to `pranor-vaultd` with zero UI freezing and automatic retry on reconnect.
2. **Inline Parquet & DuckDB Zero-ETL Analytics**:
   - *The Problem with MinIO / S3*: Traditional storage engines are "dumb byte stores" requiring external, expensive query clusters (AWS Athena, Snowflake, Trino).
   - *The Pranor Vault Solution*: `pranor-vaultd` embeds an analytical SQL engine directly inside the storage process. Execute ANSI SQL (`SELECT * FROM 's3://bucket/data.parquet' WHERE status = 500`) directly via HTTP. Pranor Vault filters data at rest, returning only requested result rows and cutting network egress by **up to 99%**.
3. **Hybrid Reed-Solomon Erasure Coding**:
   - *The Problem with Traditional Storage*: Forced 3x replication (200% storage cost overhead).
   - *The Pranor Vault Solution*: Configurable $K+M$ erasure coding (e.g., $4+2$ parity = 50% storage overhead for 11 9s of durability).

---

### 🌐 Pranor Gateway: Differentiating Factors (vs. Kong / Envoy / Cloudflare Workers)

1. **Universal WASM Edge Filter Engine (Server & Browser Service Worker)**:
   - *The Problem with Kong / Envoy*: Kong relies on Lua (single-threaded CPU bottleneck), while Envoy WASM has high IPC overhead. Neither can run inside a user's web browser.
   - *The Pranor Gateway Solution*: Compiled WebAssembly filters execute in-process with sub-10 microsecond latency. The **exact same WASM filter rules** running in `pranor-gatewayd` on the server can also run inside the browser as a Service Worker (`@pranor/gateway-wasm`), delivering offline-first edge mock APIs and zero-latency client-side request validation.
2. **Native Edge AI Proxy, Semantic Prompt Caching & Token Throttling**:
   - *The Problem with Generic Gateways*: Traditional gateways only count HTTP requests (RPM), ignoring actual LLM token consumption.
   - *The Pranor Gateway Solution*: Parses OpenAI, Anthropic, and Ollama streams natively:
     - **Token-per-Minute (TPM) Limits**: Enforces real-time rate limits on prompt + completion tokens.
     - **Semantic Prompt Caching**: Hashes prompt embeddings to return cached LLM responses in <1ms, slashing LLM API costs by up to 80%.
     - **Automatic PII Redaction**: Strips sensitive data (SSNs, credit cards, keys) before prompts leave the edge.
3. **Kernel-Level eBPF XDP DDoS Bypass (<5µs Latency)**:
   - *The Problem with User-Space Gateways*: Traditional gateways inspect malicious traffic after the TCP handshake, remaining vulnerable to SYN floods.
   - *The Pranor Gateway Solution*: `pranor-gatewayd` attaches eBPF XDP programs directly to the NIC driver layer, dropping SYN floods and malicious IP ranges in kernel space (<5µs latency) before packets hit TCP sockets.

---

## Quickstart with Pranor Pulse v2

```bash
# Clone the unified Serv monorepo
git clone https://github.com/vyuvaraj/pranor.git
cd serv/packages/Pranor Pulse

# Build and start the daemon (Web UI at http://localhost:8082/ui/)
go build -o pranor-pulsed ./cmd/pranor-pulsed
./pranorqueued --port 8082

# In another terminal, use the CLI
go build -o pranor-pulse ./cmd/pranor-pulse
./pranorqueue status
./pranorqueue topics create orders.created
./pranorqueue publish orders.created '{"order_id": 1001, "amount": 99.99}'
./pranorqueue seek orders.created 10m
```

---

## Summary & What's Next

The **Pranor Platform** (`Pranor Pulse`, `Pranor Vault`, `Pranor Gateway`) transforms traditional backend infrastructure into ultra-fast, zero-dependency engines. By combining inline WASM processing, client-side OPFS persistence, multi-protocol native compatibility, and kernel-level eBPF DDoS protection, developers can build local-first, financial-grade distributed systems without cloud lock-in.

* **Monorepo**: [github.com/vyuvaraj/pranor](https://github.com/vyuvaraj/pranor)
* **Packages**: `packages/Pranor Pulse`, `packages/Pranor Vault`, `packages/Pranor Gate`
* **License**: AGPLv3 (Server Engines) & Apache 2.0 (Client SDKs / OPFS WASM Drivers)

*— Yuvaraj*

