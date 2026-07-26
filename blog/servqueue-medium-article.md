# ServQueue v2: Evolving Our WASM Message Broker into a Monorepo Stream Engine

*From a single-binary WASM broker to a multi-protocol streaming engine featuring dual CLI binaries, MQTT v5.0, Kafka compatibility, point-in-time replay, and CRDT active-active geo-replication.*

---

> 💡 **Note**: This is **Part 2** of the ServQueue series. If you missed Part 1 on how we built inline WASM stream processing from scratch, check out [Part 1: I Built a Message Broker With Inline WASM Stream Processing From Scratch](https://medium.com/@yuvamca002/i-built-a-message-broker-with-inline-wasm-stream-processing-from-scratch-2196a391cfac).

---

## Why ServQueue Needed to Evolve

In Part 1, we introduced ServQueue: a Go-based message broker executing inline WebAssembly (WASM) filters inside the dispatch loop.

While the core WASM stream processing engine worked seamlessly, running it in real-world infrastructure revealed key operational challenges:

1. **Mixed Binary Responsibilities**: The client CLI commands and broker server daemon were bundled together, making daemon deployment in Docker/K8s heavier than necessary.
2. **Protocol Barriers**: Developers wanted IoT devices to stream telemetry via MQTT and existing microservices to produce via Kafka without needing external bridge proxy containers.
3. **Disaster Recovery**: Consumers needed the ability to rewind and replay stream history back to exact timestamps (`seekToTime`) rather than just raw offset numbers.
4. **Cross-Cloud Active-Active Sync**: Distributed clusters across multiple cloud regions required conflict-free multi-primary event mirroring.

Here is how we addressed these challenges in **ServQueue v2** and migrated to the unified **Serv monorepo** (`github.com/vyuvaraj/serv`).

---

## 1. Monorepo Migration & The Daemon / CLI Split

We merged standalone services into the unified **Serv monorepo** (`github.com/vyuvaraj/serv/packages/ServQueue`). As part of this migration, we strictly separated server runtime logic from client administration tooling:

* **`servqueued` (Server Daemon)**: Zero-dependency background service hosting the WAL log engine, WASM execution runner, protocol adapters, Prometheus metrics (`/metrics`), and an embedded Web Admin UI served at `http://localhost:9092/ui/` via Go `embed`.
* **`servqueue` (Client CLI)**: Fast-booting administrative binary for operators and scripts (`status`, `topics`, `publish`, `consume`, `tail`, `seek`).

```
                              ┌────────────────────────────────────────┐
                              │               servqueued               │
┌──────────────┐              │ ┌──────────┐  ┌──────────┐ ┌─────────┐ │              ┌──────────────┐
│  STOMP /     │  publish     │ │  STOMP   │  │   MQTT   │ │  Kafka  │ │   dispatch   │  Consumers   │
│  MQTT /      ├─────────────►│ │  :61613  │  │  :1883   │ │  :9092  │ ├─────────────►│  (Subscribers│
│  Kafka Client│              │ └────┬─────┘  └────┬─────┘ └────┬────┘ │              └──────────────┘
└──────────────┘              │      └───────────┼────────────┘      │
                              │                  ▼                   │
                              │     [ WASM Stream Processor ]        │
                              │                  │                   │
                              │     [ Point-in-Time Replay ]         │
                              │                  │                   │
                              │    [ CRDT Active-Active Sync ]       │
                              └────────────────────────────────────────┘
```

---

## 2. Multi-Protocol Compatibility: Speaking MQTT & Kafka Natively

Rather than requiring users to deploy protocol translation bridges, `servqueued` implements native binary socket decoders for multiple industry wire protocols:

* **STOMP (`:61613`)**: Traditional text-oriented pub/sub frame parser.
* **MQTT v5.0 Gateway (`:1883`)**: Decodes MQTT `CONNECT`, `PUBLISH`, `SUBSCRIBE`, and `PINGREQ` frames directly into ServQueue topics — allowing low-power IoT devices to publish directly into the broker.
* **Kafka Compatibility Adapter (`:9092`)**: Decodes binary Kafka request headers (`Produce`, `Fetch`, `Metadata`), allowing standard Kafka client SDKs to publish directly to ServQueue without changing code.

---

## 3. Point-in-Time Event Replay (`seekToTime`)

When downstream consumers crash or data corruptions occur, seeking by relative offset count is error-prone. 

ServQueue v2 introduces timestamp-based stream seeking (`seekToTime`). The log engine indexes entry timestamps, enabling instant seeks via the CLI or HTTP API:

```bash
# Seek consumer offset to 15 minutes ago
servqueue seek orders.events 15m

# Seek to an explicit ISO-8601 timestamp
servqueue seek orders.events 2026-07-26T05:00:00Z
```

The server returns the exact target offset and timestamp, allowing consumers to resume processing safely.

---

## 4. Cross-Cloud Active-Active Geo-Replication (CRDTs)

For multi-region deployments, ServQueue v2 introduces background cluster mirroring powered by **Last-Write-Wins (LWW) CRDTs**.

When nodes in `us-east-1` and `eu-west-1` operate concurrently, the mirror engine resolves concurrent state updates across cluster boundaries without central locks, ensuring eventual consistency even across temporary network partitions.

---

## 5. Cloud-Native Operations & K8s Ecosystem

To operationalize ServQueue in cloud environments, we built native Kubernetes controllers and monitoring exporters:

* **Kubernetes Operator (`ServQueueCluster`)**: A custom CRD controller managing cluster replica deployment, state reconciliation, and automated pod failover.
* **KEDA Metrics Scaler**: Exposes topic consumer lag metrics to Kubernetes Event-driven Autoscaling (KEDA) to scale consumer pod replicas dynamically.
* **Prometheus & Grafana**: Native `/metrics` endpoint with pre-built Grafana dashboard templates (`grafana_dashboard.json`) tracking throughput, queue depth, consumer lag, and WASM transform execution latency.
* **Automated Chaos Injector**: Built-in testing harness (`pkg/testing/chaos_injector.go`) for injecting network latency, partition failures, and disk corruption during integration testing.

---

## 6. Multi-Language SDKs

ServQueue v2 provides official client packages under `packages/ServQueue/sdks/`:

* **Go**: `import "github.com/vyuvaraj/serv/packages/ServQueue/sdks/go"`
* **TypeScript / Node.js**: `const { ServQueueClient } = require('@servverse/queue-sdk');`
* **Python**: `from servqueue import ServQueueClient`
* **Browser WASM**: `@servverse/queue-wasm` for embedded Web Worker event streaming using OPFS (`FileSystemSyncAccessHandle`).

---

## Quickstart with ServQueue v2

```bash
# Clone the unified Serv monorepo
git clone https://github.com/vyuvaraj/serv.git
cd serv/packages/ServQueue

# Build and start the daemon (Web UI at http://localhost:9092/ui/)
go build -o servqueued ./cmd/servqueued
./servqueued --port 9092

# In another terminal, use the CLI
go build -o servqueue ./cmd/servqueue
./servqueue status
./servqueue topics create orders.created
./servqueue publish orders.created '{"order_id": 1001, "amount": 99.99}'
./servqueue seek orders.created 10m
```

---

## Summary & What's Next

ServQueue has grown from an experimental WASM pub/sub broker into a multi-protocol stream processing engine. By combining inline WASM transforms with native MQTT/Kafka protocol support, timestamp-based replay, and Kubernetes operator tooling, developers get the power of a modern event streaming platform with zero external dependencies.

* **Monorepo**: [github.com/vyuvaraj/serv](https://github.com/vyuvaraj/serv)
* **Package Path**: `packages/ServQueue`
* **License**: Apache 2.0

*— Yuvaraj*
