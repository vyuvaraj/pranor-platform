# I Built a Message Broker With Inline WASM Stream Processing — From Scratch

*How ServQueue evolved into a monorepo distribution featuring dual-CLI binaries, STOMP & MQTT & Kafka protocol compatibility, point-in-time event replay, embedded Web Admin UI, and CRDT active-active geo-replication.*

---

## The Stream Processing & Ecosystem Evolution

Message brokers like RabbitMQ, NATS, or Kafka do an excellent job of delivering bytes from producer to consumer. But traditional brokers are fundamentally passive pipelines: transforming messages or routing across protocol boundaries requires deploying complex stream processors (Apache Flink, Kafka Streams) or external bridge containers.

**ServQueue** was built to solve this natively within the **Servverse monorepo** ecosystem (`github.com/vyuvaraj/serv/packages/ServQueue`). It combines **Compute-in-Queue** (inline WASM transform execution), zero-dependency deployment, protocol adapters, and automated cloud-native operations into a single platform.

---

## What is ServQueue?

ServQueue is a high-performance message broker written in Go as part of the unified **Servverse** monorepo. It features a clean **Daemon/Client separation**:

* **`servqueued`**: Single zero-dependency server daemon hosting the log storage engine, STOMP server (`:61613`), HTTP REST API (`:8082`/`:9092`), MQTT v5.0 gateway (`:1883`), Kafka wire protocol adapter (`:9092`), and an embedded Web Admin UI (`http://localhost:9092/ui`).
* **`servqueue`**: Lightweight client CLI for administrative tasks (`status`, `topics`, `publish`, `consume`, `tail`, `seek`).

```
Producer ➔ [ Publish (STOMP / MQTT / Kafka / REST) ]
                  │
                  ▼
         ( ServQueue Topic ) ➔ [ WASM Transform Sandbox ] ➔ [ Dispatch ] ➔ Consumer
                  │                        │
                  │               (Failed / Dropped)
                  │                        │
                  ▼                        ▼
       ( CRDT Geo-Mirror )     ( Dead Letter Queue )
```

---

## Key Differentiators & Advanced Features

### 1. Compute-in-Queue (Inline WASM Transforms)
Compile transform filters written in Go or Rust into WebAssembly (WASI) and upload them directly to topics using the HTTP management API:

```bash
# Upload a compiled .wasm transform to the 'orders' topic
curl -X POST http://localhost:8082/api/v1/topics/orders/transform \
  --data-binary @my_transform.wasm
```

Messages pass through the sandboxed `wazero` WASM runner inline during dispatch, eliminating extra network hops.

### 2. Point-in-Time Event Replay (`seekToTime`)
Seek consumer offsets to arbitrary past timestamps for disaster recovery or historical event replay:

```bash
# Seek consumer offset to 15 minutes ago
servqueue seek orders 15m

# Seek to explicit RFC3339 timestamp
servqueue seek orders 2026-07-26T05:00:00Z
```

### 3. Protocol Adapters: Wire-level STOMP, MQTT v5.0, & Kafka
Connect using standard existing client SDKs without changing client code:
* **STOMP TCP Server** (`tcp://localhost:61613`): Native pub/sub subscription frames.
* **MQTT v5.0 IoT Gateway** (`tcp://localhost:1883`): Native IoT device telemetry ingestion with `CONNECT`, `PUBLISH`, `SUBSCRIBE`, and `PUBACK`.
* **Kafka Compatibility Adapter** (`tcp://localhost:9092`): Decodes Kafka binary protocol requests (`Produce`, `Fetch`, `Metadata`).

### 4. Embedded Web Admin UI & ServConsole Queue Inspector
* **Embedded UI**: Accessible directly at `http://localhost:9092/ui/` via Go `embed`, providing real-time stats, active topics, queue depth, consumer lag, and live stream tailing.
* **ServConsole Integration**: Real-time consumer lag monitoring, outbox relay status, and stream inspection integrated directly into the central dashboard.

### 5. Cross-Cloud Active-Active Geo-Replication (CRDT)
Multi-region active-active cluster mirroring across cloud providers with Last-Write-Wins (LWW) CRDT conflict resolution.

### 6. Cloud-Native & K8s Operations
* **Kubernetes Operator**: Custom `ServQueueCluster` CRD for automated cluster provisioning and replica failover.
* **KEDA Metrics Adapter**: Auto-scale consumer pods based on real-time topic lag.
* **Storage Tiering & Auto-Compaction**: Automatic TTL background eviction and cold segment offloading to S3/ServStore.
* **Prometheus Metrics**: Exposes native `/metrics` endpoint with ready-to-use Grafana dashboard templates (`grafana_dashboard.json`).

---

## Multi-Language Client SDKs

ServQueue ships with standalone client SDKs in `packages/ServQueue/sdks/`:

* **Go**: `import "github.com/vyuvaraj/serv/packages/ServQueue/sdks/go"`
* **TypeScript / Node.js**: `const { ServQueueClient } = require('@servverse/queue-sdk');`
* **Python**: `from servqueue import ServQueueClient`
* **Browser WASM**: `@servverse/queue-wasm` for OPFS-backed embedded Web Worker event logs.

---

## Architecture Comparison

| Feature | RabbitMQ | NATS JetStream | Kafka | ServQueue |
|:---|:---|:---|:---|:---|
| **Repository** | Separate | Separate | Separate | **Serv Monorepo** |
| **Daemon / CLI Split** | Partial | `nats` CLI | `kafka-tools` | **`servqueued` / `servqueue`** |
| **Compute-in-Queue** | ❌ (Plugins only) | ❌ | ❌ | **✅ (WASM WASI)** |
| **Multi-Protocol** | AMQP/STOMP | NATS | Kafka | **STOMP, MQTT, Kafka, REST** |
| **Point-in-Time Seek** | ❌ | Offset only | Offset/Timestamp | **`seekToTime` (CLI & REST)** |
| **Embedded Admin UI** | Plugin | ❌ | External | **Built-in (`http://localhost:9092/ui`)** |
| **CRDT Geo-Mirroring** | ❌ | ❌ | MirrorMaker 2 | **Native Active-Active CRDT** |
| **K8s Operator & KEDA** | Third-party | Third-party | Strimzi | **Built-in Operator & KEDA** |

---

## Quickstart (Monorepo Setup)

```bash
# Clone the unified Serv monorepo
git clone https://github.com/vyuvaraj/serv.git
cd serv/packages/ServQueue

# Build and run the daemon
go build -o servqueued ./cmd/servqueued
./servqueued --port 9092

# In another terminal, interact using the CLI
go build -o servqueue ./cmd/servqueue
./servqueue status
./servqueue topics create orders
./servqueue publish orders '{"order_id": 1001, "total": 49.99}'
./servqueue consume orders
./servqueue seek orders 5m
```

---

## Links & Ecosystem

- **Monorepo**: [github.com/vyuvaraj/serv](https://github.com/vyuvaraj/serv)
- **ServQueue Package**: [packages/ServQueue](https://github.com/vyuvaraj/serv/tree/main/packages/ServQueue)
- **Ecosystem Specs**: Check `UNIFIED_ROADMAP.md` in `servverse-repo`.
- **License**: Apache 2.0

---

*Streamline your infrastructure. Run transformations near the data inside the broker, speak standard STOMP/MQTT/Kafka protocols, and scale seamlessly from embedded local-first PWAs to active-active cloud clusters.*

*— Yuvaraj*
