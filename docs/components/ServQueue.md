# Pranor Pulse — WASM-Enabled Multi-Protocol Message Broker

> **Status:** ✅ Production | **Ports:** `8082`/`9092` (HTTP & Admin UI), `61613` (STOMP), `1883` (MQTT v5.0), `9092` (Kafka) | **Repository:** [Pranor Pulse](https://github.com/vyuvaraj/pranor/tree/main/packages/Pranor Pulse)

---

## Overview

Pranor Pulse is a zero-dependency, high-performance message broker and event streaming engine written in Go as part of the unified **Serv monorepo**. It features a **Dual-Binary Architecture**:

- **`pranor-pulsed`**: Server daemon hosting core storage, inline WASM stream transforms, multi-protocol listeners, Prometheus metrics, and an embedded Web Admin UI.
- **`pranor-pulse`**: Dedicated administrative CLI for publishing, consuming, topic management, stream tailing, and point-in-time replay.

---

## Key Features

- **Dual-Binary Architecture**: Dedicated `pranor-pulsed` server daemon + lightweight `pranor-pulse` client CLI.
- **Multi-Protocol Support**:
  - **STOMP TCP Protocol** (`:61613`): `CONNECT`, `SUBSCRIBE`, `SEND`, `ACK`, `DISCONNECT`
  - **MQTT v5.0 IoT Gateway** (`:1883`): `CONNECT`, `PUBLISH`, `SUBSCRIBE`, `PINGREQ` for IoT device telemetry
  - **Kafka Wire Protocol Adapter** (`:9092`): `Produce`, `Fetch`, `ListOffsets`, `Metadata` binary request decoding
  - **HTTP REST Management API** (`:8082`/`:9092`): JSON publish, replay, topic admin, stats, and Web UI
- **Compute-in-Queue**: Wazero JIT-compiled WebAssembly (WASI) sandboxed stream transformers running inline.
- **Point-in-Time Event Replay (`seekToTime`)**: Seek consumer offsets to arbitrary past timestamps or durations (`pranor-pulse seek`).
- **Dead Letter Queues (DLQ)**: Automatic routing, poison-pill isolation, and backoff replay policies.
- **Cross-Cloud Active-Active Geo-Replication**: Multi-region cluster mirroring with Last-Write-Wins (LWW) CRDT conflict resolution.
- **Embedded Web Admin UI**: Embedded via `go:embed` at `http://localhost:8082/ui/` or `http://localhost:9092/ui/`.
- **Cloud-Native & Kubernetes Operations**:
  - **Kubernetes Operator (`Pranor PulseCluster`)**: Custom CRD controller managing cluster replica deployment and failover.
  - **KEDA Metrics Scaler**: Exposes topic consumer lag metrics to Kubernetes Event-driven Autoscaling.
  - **Automated Storage Tiering & Compaction**: TTL background eviction workers and cold segment offload to S3/Pranor Vault.
  - **Automated Chaos Injector**: Built-in testing harness (`pkg/testing/chaos_injector.go`) for latency, network partition, and disk corruption testing.
- **Multi-Language Client SDKs**: Standalone Go, TypeScript/Node.js, Python, and Browser WASM SDKs under `sdks/`.

---

## Dedicated Listener Ports

| Protocol | Default Port | Description |
|:---|:---|:---|
| **STOMP** | `:61613` | STOMP text-oriented messaging protocol |
| **MQTT v5.0** | `:1883` | Native IoT telemetry ingestion gateway |
| **Kafka Wire** | `:9092` | Binary Kafka protocol producer & consumer socket |
| **HTTP REST & UI** | `:8082` / `:9092` | REST API, Prometheus `/metrics`, and embedded Web Admin UI (`/ui/`) |

---

## Endpoints

| Endpoint | Method | Description |
|:---|:---|:---|
| `/healthz` | GET | Liveness and readiness probe |
| `/metrics` | GET | Prometheus metrics exporter |
| `/ui/` | GET | Embedded Web Admin Console UI |
| `/api/v1/publish` | POST | Publish message to a topic |
| `/api/v1/topics` | GET / POST | List or create topics |
| `/api/v1/seekToTime` | POST | Seek consumer offset to target timestamp |
| `/api/v1/replay/time` | POST | Replay event range from target timestamp |
| `/api/v1/dlq/:topic/replay` | POST | Replay DLQ messages back to primary queue |

---

## CLI Reference (`pranor-pulse`)

```bash
# Check node status
pranor-pulse status

# Topic Management
pranor-pulse topics list
pranor-pulse topics create orders.created

# Publish & Consume
pranor-pulse publish orders.created '{"order_id": "ord_1001", "total": 99.99}'
pranor-pulse consume orders.created
pranor-pulse tail orders.created --filter "total > 50"

# Point-in-Time Event Replay
pranor-pulse seek orders.created 15m
pranor-pulse seek orders.created 2026-07-26T05:00:00Z

# DLQ Management
pranor-pulse dlq replay orders.created
```

---

## Multi-Language SDK Usage

### Go SDK
```go
import "github.com/vyuvaraj/pranor/packages/Pranor Pulse/sdks/go"

client := pranor-pulse.NewClient("http://localhost:8082", "my-token")
err := client.Publish("orders.created", `{"order_id": "1001"}`)
offset, err := client.SeekToTime("orders.created", "15m")
```

### TypeScript / Node.js SDK
```javascript
const { Pranor PulseClient } = require('@pranor/queue-sdk');

const client = new Pranor PulseClient("http://localhost:8082", "my-token");
await client.publish("orders.created", JSON.stringify({ order_id: "1001" }));
```

### Python SDK
```python
from pranor-pulse import Pranor PulseClient

client = Pranor PulseClient(base_url="http://localhost:8082")
client.publish("orders.created", '{"order_id": "1001"}')
```

---

## Production & Operational Guidelines

### WASM Safety & Resource Limits
Pranor Pulse runs isolated WebAssembly stream transform filters inside message streams using the `wazero` engine:
* **Execution Timeout:** Each inline transform is bounded by a strict `50ms` timeout context.
* **Memory Capping:** Sandboxed with a `16MB` memory allocation ceiling.

### Persistence & Storage Tiering
* **Durability:** Every message is appended to the Write-Ahead Log (WAL) on disk before returning receipts.
* **Storage Tiering:** Closed WAL segments are automatically offloaded to S3 / Pranor Vault for infinite backlog retention, while TTL compaction workers purge expired records.

### Observability & Telemetry
* **Prometheus Metrics:** Pranor Pulse exposes metrics at `GET /metrics` and includes pre-built Grafana templates (`grafana_dashboard.json`).
* **OpenTelemetry Compatibility:** Propagates W3C Trace Context headers inside STOMP properties and OTLP spans across transform steps.
