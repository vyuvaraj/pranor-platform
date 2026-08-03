# Event-Driven Microservices with Pranor Pulse

> **Published:** July 2026 | **Reading Time:** ~12 min | **Tags:** `pranor-pulse`, `messaging`, `event-driven`, `pub-sub`, `wasm`, `monorepo`

---

Synchronous request-response works fine until it doesn't. When a user places an order, you don't want them waiting while your service sends a confirmation email, updates inventory, triggers a fulfillment workflow, and notifies your analytics pipeline — all in sequence. **Pranor Pulse** (part of the unified [Pranor monorepo](https://github.com/vyuvaraj/pranor)) decouples these operations so each one runs independently, securely, and reliably.

---

## What Is Pranor Pulse?

Pranor Pulse is the Pranor message broker. It supports:

- **Dual-Binary Architecture** — `pranor-pulsed` server daemon + `pranor-pulse` management CLI
- **Multi-Protocol Support** — Native STOMP (`:61613`), MQTT v5.0 (`:1883`), Kafka Wire Protocol (`:9092`), and HTTP REST (`:8082`/`:9092`)
- **Compute-in-Queue** — Inline WASM transform execution on topics
- **Point-in-Time Event Replay** — Seek consumer offsets to arbitrary timestamps (`seekToTime`)
- **Dead Letter Queues (DLQ)** — Poison-pill isolation and automatic retry policies
- **Active-Active Geo-Replication** — Multi-region cluster mirroring with CRDT conflict resolution
- **Embedded Web Admin UI** — Built-in visual management UI at `http://localhost:9092/ui/`
- **Cloud-Native Integration** — Kubernetes Operator (`Pranor PulseCluster`) and KEDA auto-scaler

---

## Step 1: Run the Pranor Pulse Daemon (`pranor-pulsed`)

Build or run `pranor-pulsed` directly from the monorepo:

```bash
cd pranor/packages/Pranor Pulse
go run ./cmd/pranor-pulsed --port 9092
```

The server outputs:
```
Starting Pranor Pulse Standalone Daemon (pranor-pulsed) on port 9092...
pranor-pulsed Web Admin UI available at http://localhost:9092/ui/
```

---

## Step 2: CLI Operations (`pranor-pulse`)

Use the dedicated `pranor-pulse` CLI tool to inspect and manage topics:

```bash
# Check status
pranor-pulse status

# Create a topic
pranor-pulse topics create orders.created

# Publish a message
pranor-pulse publish orders.created '{"order_id": "ord_1001", "total": 99.99}'

# Consume messages
pranor-pulse consume orders.created

# Stream live topic messages
pranor-pulse tail orders.created

# Seek consumer offset to 15 minutes ago
pranor-pulse seek orders.created 15m
```

---

## Step 3: Multi-Language Client SDKs

Pranor Pulse provides official client SDKs:

### Go SDK
```go
import "github.com/vyuvaraj/pranor/packages/Pranor Pulse/sdks/go"

client := pranor-pulse.NewClient("http://localhost:8082", "my-token")
err := client.Publish("orders.created", `{"order_id": "1001"}`)
offset, err := client.SeekToTime("orders.created", "15m")
```

### Node.js / TypeScript SDK
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

## Step 4: Observability & Monitoring

Pranor Pulse exposes Prometheus metrics natively at `/metrics`:
* `pranor-pulse_messages_published_total`
* `pranor-pulse_queue_depth`
* `pranor-pulse_consumer_lag`
* `pranor-pulse_wasm_executions_total`

Pre-built Grafana dashboard templates are included in `packages/Pranor Pulse/grafana_dashboard.json`.

---

## Summary

Pranor Pulse simplifies event-driven architectures by combining multi-protocol compatibility, inline WASM stream processing, point-in-time replay, and native Kubernetes scaling into a single zero-dependency platform.
