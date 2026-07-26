# Event-Driven Microservices with ServQueue

> **Published:** July 2026 | **Reading Time:** ~12 min | **Tags:** `servqueue`, `messaging`, `event-driven`, `pub-sub`, `wasm`, `monorepo`

---

Synchronous request-response works fine until it doesn't. When a user places an order, you don't want them waiting while your service sends a confirmation email, updates inventory, triggers a fulfillment workflow, and notifies your analytics pipeline — all in sequence. **ServQueue** (part of the unified [Serv monorepo](https://github.com/vyuvaraj/serv)) decouples these operations so each one runs independently, securely, and reliably.

---

## What Is ServQueue?

ServQueue is the Servverse message broker. It supports:

- **Dual-Binary Architecture** — `servqueued` server daemon + `servqueue` management CLI
- **Multi-Protocol Support** — Native STOMP (`:61613`), MQTT v5.0 (`:1883`), Kafka Wire Protocol (`:9092`), and HTTP REST (`:8082`/`:9092`)
- **Compute-in-Queue** — Inline WASM transform execution on topics
- **Point-in-Time Event Replay** — Seek consumer offsets to arbitrary timestamps (`seekToTime`)
- **Dead Letter Queues (DLQ)** — Poison-pill isolation and automatic retry policies
- **Active-Active Geo-Replication** — Multi-region cluster mirroring with CRDT conflict resolution
- **Embedded Web Admin UI** — Built-in visual management UI at `http://localhost:9092/ui/`
- **Cloud-Native Integration** — Kubernetes Operator (`ServQueueCluster`) and KEDA auto-scaler

---

## Step 1: Run the ServQueue Daemon (`servqueued`)

Build or run `servqueued` directly from the monorepo:

```bash
cd serv/packages/ServQueue
go run ./cmd/servqueued --port 9092
```

The server outputs:
```
Starting ServQueue Standalone Daemon (servqueued) on port 9092...
servqueued Web Admin UI available at http://localhost:9092/ui/
```

---

## Step 2: CLI Operations (`servqueue`)

Use the dedicated `servqueue` CLI tool to inspect and manage topics:

```bash
# Check status
servqueue status

# Create a topic
servqueue topics create orders.created

# Publish a message
servqueue publish orders.created '{"order_id": "ord_1001", "total": 99.99}'

# Consume messages
servqueue consume orders.created

# Stream live topic messages
servqueue tail orders.created

# Seek consumer offset to 15 minutes ago
servqueue seek orders.created 15m
```

---

## Step 3: Multi-Language Client SDKs

ServQueue provides official client SDKs:

### Go SDK
```go
import "github.com/vyuvaraj/serv/packages/ServQueue/sdks/go"

client := servqueue.NewClient("http://localhost:8082", "my-token")
err := client.Publish("orders.created", `{"order_id": "1001"}`)
offset, err := client.SeekToTime("orders.created", "15m")
```

### Node.js / TypeScript SDK
```javascript
const { ServQueueClient } = require('@servverse/queue-sdk');

const client = new ServQueueClient("http://localhost:8082", "my-token");
await client.publish("orders.created", JSON.stringify({ order_id: "1001" }));
```

### Python SDK
```python
from servqueue import ServQueueClient

client = ServQueueClient(base_url="http://localhost:8082")
client.publish("orders.created", '{"order_id": "1001"}')
```

---

## Step 4: Observability & Monitoring

ServQueue exposes Prometheus metrics natively at `/metrics`:
* `servqueue_messages_published_total`
* `servqueue_queue_depth`
* `servqueue_consumer_lag`
* `servqueue_wasm_executions_total`

Pre-built Grafana dashboard templates are included in `packages/ServQueue/grafana_dashboard.json`.

---

## Summary

ServQueue simplifies event-driven architectures by combining multi-protocol compatibility, inline WASM stream processing, point-in-time replay, and native Kubernetes scaling into a single zero-dependency platform.
