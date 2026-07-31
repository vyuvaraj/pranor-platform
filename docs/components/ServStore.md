# Pranor Vault — AI-Native Object Storage

> **Status:** ✅ Production | **Port:** 8081 | **Repository:** [Pranor Vault](https://github.com/vyuvaraj/pranor/tree/main/packages/Pranor Vault)

## Overview

S3-compatible distributed object storage with WASM transforms, semantic search, time-travel queries, and content-addressed deduplication. Single binary with embedded web console.

## Key Features

- Full S3-compatible API (PUT/GET/DELETE/HEAD, multipart, presigned URLs)
- Object versioning (enabled, suspended, disabled) + delete markers
- AWS Signature V4 authentication
- WASM server-side transforms (wazero runtime)
- Transform pipeline DAG (multi-stage chaining)
- Semantic search (HNSW vector index + ONNX embeddings)
- Multi-modal embedding (text, images, PDFs, audio)
- Time-travel queries (any object at any timestamp)
- Content-addressed storage with reference-counted GC
- Raft consensus for distributed metadata
- Erasure coding (Reed-Solomon)
- Consistent hashing for data placement
- Geo-aware data placement and replication
- Cold storage tiering (transparent rehydration from S3/Glacier/B2)
- S3 Select (SQL queries on CSV/JSON in-place)
- Bucket event notifications (CloudEvents)
- Object tagging and metadata
- Lifecycle policies (auto-expire/transition)
- Object locking (WORM)
- Kubernetes operator + CSI plugin
- Embedded web console (glassmorphic UI)

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVSTORE_PORT` | Listen port | `8081` |
| `SERVSTORE_DATA_DIR` | Storage directory | `./data` |
| `SERVSTORE_ACCESS_KEY` | S3 access key | `admin` |
| `SERVSTORE_SECRET_KEY` | S3 secret key | `password` |
| `PRANOR_OTLP_ENDPOINT` | OTel collector | (disabled) |

## Endpoints

| Endpoint | Description |
|----------|-------------|
| `GET /healthz` | Liveness probe |
| `PUT /<bucket>/<key>` | Upload object |
| `GET /<bucket>/<key>` | Download object |
| `GET /<bucket>?query=semantic&q=<text>` | Semantic search |
| `POST /<bucket>/<wasm>?transform=true` | Execute WASM transform |
| `GET /console/` | Web UI dashboard |

## Pranor Integration

```srv
store "servstore://localhost:8081/my-bucket"

route "POST" "/upload" (req) {
    store.put(req.body.filename, req.body.data)
    return { "stored": true }
}
```

## Detailed API Reference

### 1. Put Object
Uploads an object payload to a specified bucket.
```bash
curl -X PUT http://localhost:8081/my-bucket/document.txt \
  -H "Authorization: Bearer gateway-secret-token" \
  -d "Hello Pranor Vault!"
```
**Response (201 Created):**
```json
{
  "etag": "e790a13df56aa...",
  "size": 17,
  "versionId": "v1"
}
```

### 2. Get Object
Downloads an object's binary contents.
```bash
curl -X GET http://localhost:8081/my-bucket/document.txt \
  -H "Authorization: Bearer gateway-secret-token"
```
**Response (200 OK):**
```text
Hello Pranor Vault!
```

### 3. Semantic Vector Search
Queries the bucket contents semantically.
```bash
curl -X GET "http://localhost:8081/my-bucket?query=semantic&q=distributed%20storage&limit=5" \
  -H "Authorization: Bearer gateway-secret-token"
```
**Response (200 OK):**
```json
[
  "document.txt",
  "architecture_overview.pdf"
]
```

---

## Operational Recovery Runbook

### Playbook 1: Metadata Cluster Sync / Split-Brain Recovery
If Raft consensus fails or a network partition splits the storage metadata ring:
1. **Identify Unhealthy Nodes**: Query `GET /api/v1/cluster/health` to find nodes marked `Down` or in a permanent state of reelection.
2. **Quorum Restoration**: If quorum is lost, restart the node with the highest last-applied index in force-bootstrap mode:
   ```bash
   servstore-server --force-bootstrap --data-dir /var/lib/servstore
   ```
3. **Re-add Stale Replicas**: Re-introduce remaining nodes into the cluster using:
   ```bash
   curl -X POST http://localhost:8081/console/cluster/join -d '{"node_addr":"10.0.0.5:8081"}'
   ```

### Playbook 2: Reed-Solomon Erasure Coding Data Recovery
When drive corruption or sector loss is detected:
1. **Detect Missing/Corrupted Shards**: Check logs for `RS Reconstruction Required` triggers.
2. **Trigger Reconstruction Job**: Send a POST request to rebuild missing data shards from surviving parity shards:
   ```bash
   curl -X POST http://localhost:8081/api/v1/maintenance/reconstruct \
     -H "Authorization: Bearer gateway-secret-token"
   ```
3. **Verify Integrity**: Validate the reconstructed objects using:
   ```bash
   curl -X GET http://localhost:8081/my-bucket/document.txt?verify-integrity=true
   ```

## Production & Operational Guidelines

### WASM Safety & Resource Limits
Pranor Vault allows executing user-defined WASM transformations near the storage boundary. To prevent resource exhaustion, the following rules apply:
* **Execution Timeout:** All WASM transforms are enforced with a strict `100ms` execution deadline context. Any transform exceeding this is aborted, and a `500 Internal Server Error` is logged.
* **Sandbox Memory Limits:** Instances run in isolated wazero sandboxes with a strict memory limit of `32MB` to safeguard cluster nodes against memory leak crashes.

### Persistence & Data Integrity Guarantees
Pranor Vault ensures high-durability persistence across failures:
* **Sync-to-Disk:** Object write operations (PUT) write payload shards to disk and issue `fsync()` system calls to guarantee block persistence before returning `201 Created`.
* **Metadata Consensus:** Cluster configurations and storage keys are recorded via Raft consensus into disk-backed log stores. Single-instance nodes run transactionally on SQLite databases.
* **Erasure Coding Parity:** Employs Reed-Solomon (data/parity) block distribution, guaranteeing zero data loss even if up to half of the physical storage nodes or block partitions fail.

### Standard Observability & Telemetry
* **Prometheus Metrics:** Exposes standard Prometheus-compatible telemetry at `GET /metrics`, reporting storage usage (bytes), active connection pools, read/write ops per second, and vector query response times.
* **OTel Integration:** Supports OpenTelemetry metrics and distributed tracing context propagation natively. Set `PRANOR_OTLP_ENDPOINT` to export traces to external monitoring infrastructure.

