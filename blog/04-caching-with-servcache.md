# Distributed Caching Made Simple with Pranor Cache

> **Published:** July 2026 | **Reading Time:** ~10 min | **Tags:** `pranor-cache`, `caching`, `performance`, `redis-alternative`

---

Caching is one of the highest-leverage performance improvements available to any backend. A well-placed cache can reduce database load by 90% and cut p99 latency from seconds to milliseconds. **Pranor Cache** is the Pranor cache server — a Redis-compatible key-value store with clustering, TTL, and namespacing built in.

---

## What Makes Pranor Cache Different From Redis?

| Feature | Redis | Pranor Cache |
|---------|-------|-----------|
| Protocol | RESP | REST + gRPC |
| Auth | ACL files | JWT / API key (same as rest of Pranor) |
| Namespacing | Conventions only | First-class namespaces |
| Dashboard | External tool | Built into Pranor Console |
| Clustering | Redis Cluster (complex) | Raft-based, auto-discovery |
| Binary | External install | Single Docker image |

---

## Step 1: Start Pranor Cache

```bash
# Standalone (single node)
docker run -d \
  -p 8082:8082 \
  -v pranor-cache_data:/data \
  --name pranor-cache \
  ghcr.io/vyuvaraj/pranor-cache:latest
```

Verify:
```bash
curl http://localhost:8082/health
# {"status":"ok","version":"0.2.0","nodes":1}
```

---

## Step 2: Basic Operations via REST API

> 💡 **Note:** While we showcase the **REST API** using standard `curl` commands below for ease of demonstration, Pranor Cache also exposes a high-performance **gRPC** interface for low-overhead client access within your service grid.

### Set a Value

```bash
curl -X PUT http://localhost:8082/cache/my-key \
  -H "Content-Type: application/json" \
  -d '{"value": "hello world", "ttl": 300}'
```

Response:
```json
{"ok": true, "expires_at": "2026-07-10T17:00:00Z"}
```

### Get a Value

```bash
curl http://localhost:8082/cache/my-key
```

Response:
```json
{"key": "my-key", "value": "hello world", "ttl": 295}
```

### Delete a Value

```bash
curl -X DELETE http://localhost:8082/cache/my-key
```

### Bulk Get

```bash
curl -X POST http://localhost:8082/cache/bulk \
  -H "Content-Type: application/json" \
  -d '{"keys": ["user:1", "user:2", "user:3"]}'
```

---

## Step 3: Namespaces

Namespaces allow you to logically partition your cache and flush entire sections at once — a pattern Redis struggles with.

```bash
# Set values in the "products" namespace
curl -X PUT http://localhost:8082/ns/products/item:42 \
  -d '{"value": {"name":"Widget","price":9.99}, "ttl": 600}'

curl -X PUT http://localhost:8082/ns/products/item:43 \
  -d '{"value": {"name":"Gadget","price":19.99}, "ttl": 600}'

# Flush the entire "products" namespace (e.g., after a bulk update)
curl -X DELETE http://localhost:8082/ns/products
# Deletes all keys under products/*
```

This is extremely useful for multi-tenant systems where you need to clear one tenant's cache without affecting others.

---

## Step 4: Use with Pranor

When your `serv.yaml` sets the cache driver to `pranor-cache`, Pranor handles all cache calls transparently:

```yaml
# serv.yaml
cache:
  driver: pranor-cache
  endpoint: http://pranor-cache:8082
  namespace: my-api       # All keys are scoped to this namespace
  default_ttl: 60s
```

In your service file:

```serv
service ProductService {
  route GET /products/:id {
    cache ttl=5m key="product:{id}"    # Pranor Cache handles this
    return store.get(Product, id)
  }
}
```

On first request: DB query, result stored in Pranor Cache for 5 minutes.
On subsequent requests: Served from cache in <1ms.

---

## Step 5: Clustering for High Availability

For production, run 3 nodes. Pranor Cache uses Raft consensus for leader election:

```yaml
# docker-compose.yml
version: "3.9"
services:
  pranor-cache-1:
    image: ghcr.io/vyuvaraj/pranor-cache:latest
    environment:
      NODE_ID: "node-1"
      CLUSTER_PEERS: "pranor-cache-2:8082,pranor-cache-3:8082"
      DATA_DIR: /data
    volumes:
      - cache1_data:/data
    ports:
      - "8082:8082"

  pranor-cache-2:
    image: ghcr.io/vyuvaraj/pranor-cache:latest
    environment:
      NODE_ID: "node-2"
      CLUSTER_PEERS: "pranor-cache-1:8082,pranor-cache-3:8082"
      DATA_DIR: /data
    volumes:
      - cache2_data:/data

  pranor-cache-3:
    image: ghcr.io/vyuvaraj/pranor-cache:latest
    environment:
      NODE_ID: "node-3"
      CLUSTER_PEERS: "pranor-cache-1:8082,pranor-cache-2:8082"
      DATA_DIR: /data
    volumes:
      - cache3_data:/data

volumes:
  cache1_data:
  cache2_data:
  cache3_data:
```

```bash
docker compose up -d

# Check cluster health
curl http://localhost:8082/cluster
# {"leader":"node-1","nodes":["node-1","node-2","node-3"],"quorum":"healthy"}
```

With 3 nodes, Pranor Cache can survive one node failure without any downtime.

---

## Step 6: Cache Invalidation Patterns

### Pattern 1: Time-Based (TTL)

The simplest approach — let items expire naturally.

```bash
# Set with 5-minute TTL
curl -X PUT http://localhost:8082/cache/user:1 \
  -d '{"value": {...}, "ttl": 300}'
```

Best for: Frequently read data with acceptable staleness (e.g., product catalog, user profiles).

### Pattern 2: Event-Driven Invalidation

Invalidate on write events via Pranor Pulse:

```serv
service OrderService {
  route POST /orders {
    result = store.create(Order, body)
    # Publish invalidation event
    queue.publish("cache.invalidate", {
      namespace: "orders",
      key: "user:{body.user_id}:orders"
    })
    return result
  }
}
```

A cache consumer subscribes and flushes the key immediately when an order is created.

### Pattern 3: Write-Through

Update the cache on every write:

```serv
route PUT /users/:id {
  result = store.update(User, id, body)
  cache.set(key="user:{id}", value=result, ttl=5m)
  return result
}
```

Best for: Hot data that must always be fresh (e.g., user session, auth tokens).

---

## Performance Numbers

In benchmarks on a t3.medium (2 vCPU, 4 GB RAM):

| Operation | Latency (p50) | Latency (p99) | Throughput |
|-----------|--------------|--------------|-----------|
| GET (hit) | 0.3ms | 1.1ms | 45,000 req/s |
| SET | 0.4ms | 1.4ms | 38,000 req/s |
| Namespace flush | 2.1ms | 8.3ms | N/A |

---

## What's Next?

Caching helps with read performance. But for async workflows and service decoupling, you need a message queue. In the next post, we build an event-driven microservice architecture with **Pranor Pulse**.

➡️ [Event-Driven Microservices with Pranor Pulse](blog.html?post=05-event-driven-pranor-pulse)

---

*Spotted a bug or have a cache pattern to share? Open a [discussion](https://github.com/vyuvaraj/Pranor Cache/discussions).*
