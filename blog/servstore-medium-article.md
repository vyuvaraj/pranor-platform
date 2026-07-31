# I Built an S3-Compatible Object Storage Engine With AI-Native Capabilities — From Scratch

*What happens when you take MinIO's premise and push it toward semantic search, time travel, and compute-near-data?*

---

## Why Build Another Object Storage?

MinIO proved that self-hosted S3-compatible storage is valuable. Thousands of companies use it to avoid AWS lock-in while keeping the S3 API their tools already speak.

But MinIO stops at being a storage box. You put bytes in, you get bytes out. If you want to search content, you pipe it into Elasticsearch. If you want to transform data, you pull it out, process it, push it back. If you want historical snapshots, you're managing versioning policies manually.

I wanted an object storage engine that *understands* its contents. One where you can:

- **Search semantically** — ask "find documents about consensus algorithms" and get ranked results
- **Time travel** — get the state of any object at any historical timestamp
- **Compute near data** — run WASM transforms server-side without moving bytes across the network

So I built Pranor Vault.

---

## What Pranor Vault Is

Pranor Vault is a distributed, S3-compatible object storage engine written in Go. It speaks the AWS S3 REST API (Signature V4), so any S3 SDK or CLI works out of the box. But beyond the standard S3 feature set, it adds an AI-native storage layer.

**The standard stuff** (you'd expect from any serious object storage):
- Full S3 REST API: PUT, GET, DELETE, HEAD, list, multipart uploads
- Object versioning (Enabled / Suspended / Disabled)
- Object locking (WORM — write-once-read-many)
- Lifecycle policies (auto-expire after N days)
- AES-256-GCM encryption at rest
- TLS 1.3 enforcement
- AWS Signature V4 authentication

**The distributed stuff** (for running at scale):
- Raft consensus for strong consistency
- Gossip-based membership and failure detection
- Consistent hash ring for balanced data placement
- Reed-Solomon erasure coding
- Cross-region replication
- Auto-healing when nodes go down
- BLAKE3 checksums with bit-rot detection

**The AI-native stuff** (what makes it different):
- Semantic search via TF-IDF indexing
- Time travel queries
- Serverless WASM transforms (compute-near-data)
- Content-addressed storage with deduplication
- Cold storage tiering to any S3-compatible backend

---
## Semantic Search — Built Into Storage

Most object storage is a dumb key-value store. You know the key, you get the object. You don't know the key? Good luck.

Pranor Vault automatically indexes text content on ingest. Upload a Markdown file, a text document, or anything with a `text/*` content type, and it gets TF-IDF indexed using cosine similarity ranking.

```bash
# Upload documents (indexed automatically)
aws s3api put-object --bucket docs --key raft.txt --body raft.txt \
  --content-type text/plain --endpoint-url http://localhost:9000

# Search semantically
curl "http://localhost:9000/docs?query=semantic&q=consensus+metadata+replication&max-results=3"
```

The response is S3-compatible XML with ranked results. No Elasticsearch. No external indexing pipeline. The storage engine knows what's inside its objects.

---

## Time Travel — Get Any Object at Any Point in History

Versioning in S3 gives you version IDs. You need to know which version you want. Pranor Vault lets you ask a different question: *what did this object look like on January 1st?*

```bash
# Get object state at a specific timestamp
curl "http://localhost:9000/mybucket/config.json?at=2025-06-01T12:00:00Z"
```

It resolves against version `LastModified` metadata. No extra storage overhead — it uses the existing version history. Simple API, powerful capability. Think Git for objects, but with timestamp-based checkout.

---

## WASM Compute-Near-Data — Process Without Moving

The traditional pattern: download 500MB file → transform locally → upload result. Three network trips, latency, bandwidth cost.

Pranor Vault flips this. Upload a WASI-compatible `.wasm` binary, then execute it *server-side* against any object:

```bash
# Build a WASM transform
GOOS=wasip1 GOARCH=wasm go build -o uppercase.wasm ./transforms/uppercase/

# Upload the transform and the data
aws s3api put-object --bucket transforms --key uppercase.wasm --body uppercase.wasm \
  --endpoint-url http://localhost:9000
aws s3api put-object --bucket transforms --key hello.txt --body hello.txt \
  --endpoint-url http://localhost:9000

# Execute server-side — data never leaves the storage node
curl -X POST "http://localhost:9000/transforms/uppercase.wasm?transform=true&target-key=hello.txt&mem-limit=64&timeout=30"
```

Each WASM invocation gets a fresh, isolated `wazero` runtime — pure Go, no CGO, no host filesystem access. Configurable memory limits and timeouts per call. The data stays where it is; the compute goes to the data.

---
## Content-Addressed Storage & Cold Tiering

Enable CAS on a bucket and Pranor Vault deduplicates automatically. Objects are stored by their BLAKE3 hash with reference counting — data is only deleted when the last reference is removed.

For cost optimization, cold storage tiering archives inactive CAS blocks to any S3-compatible backend (AWS Glacier, Backblaze B2, another MinIO cluster):

```bash
# Configure cold tiering
curl -X PUT "http://localhost:9000/mybucket?cold-tier" \
  -H "Content-Type: application/json" \
  -d '{
    "endpoint": "https://s3.amazonaws.com",
    "remote_bucket": "cold-archive",
    "min_age_days": 30,
    "scan_interval_min": 60
  }'

# GetObject transparently re-hydrates archived blocks
aws s3api get-object --bucket mybucket --key archived.bin /tmp/out.bin \
  --endpoint-url http://localhost:9000
```

Transparent re-hydration. No API changes. Objects archived cold are fetched seamlessly when accessed.

---

## Distributed — Not Just Replicated

Pranor Vault isn't a single-node toy with "distributed" in the README. The clustering layer includes:

- **Raft consensus** for all metadata mutations — strong consistency, not eventual
- **Gossip protocol** for node discovery and failure detection — nodes find each other and detect failures within seconds
- **Consistent hash ring** (CRUSH-style) with virtual nodes for balanced placement — add a node and data rebalances automatically
- **Reed-Solomon erasure coding** (configurable data/parity ratio) — survive shard loss without full replication overhead
- **Cross-region replication** — async replication across geographic regions with loop prevention
- **Auto-healing** — detects under-replicated objects and rebuilds replicas in the background

Add a node with a single API call:
```bash
curl -X POST "http://localhost:9000/console/cluster/join" \
  -d '{"address": "node3:9000"}'
```

Background rebalancing redistributes existing objects without downtime.

---

## Cloud-Native: Kubernetes Operator & Helm

Pranor Vault includes a full Kubernetes deployment story:

```yaml
apiVersion: storage.servstore.io/v1alpha1
kind: Pranor VaultCluster
metadata:
  name: my-storage
spec:
  replicas: 3
  image: ghcr.io/vyuvaraj/servstore:latest
  erasureCoding:
    enabled: true
    dataShards: 2
    parityShards: 1
  storage:
    size: 50Gi
```

The operator manages StatefulSets, rolling upgrades, bucket configuration, and credential-to-policy mapping. Helm chart included.

---
## Performance & Hardening

Some details that matter in production:

- **Direct I/O bypass** for objects >16MB — skips the OS page cache for direct disk throughput
- **Parallel BLAKE3 hashing** for objects >8MB — concurrent chunk hashing across CPU cores
- **Tenant-isolated rate limiting** — token-bucket per namespace, 429 responses with `Retry-After` headers
- **Chaos testing manifests** included — pod failures, network partitions, and disk I/O errors

---

## Observability Out of the Box

- **OpenTelemetry tracing** on all HTTP routes and storage I/O — zero external dependencies
- **Prometheus metrics** at `/metrics` — request rate, latency histograms, storage utilization, cluster state
- **Structured JSON logging** via Go's `slog` with trace IDs, method, path, status, and duration

No sidecar agents. No configuration. Start the binary and your observability stack picks it up.

---

## How It Compares

| Feature | AWS S3 | MinIO | Pranor Vault |
|---------|--------|-------|-----------|
| S3 API compatible | ✅ (native) | ✅ | ✅ |
| Self-hostable | ❌ | ✅ | ✅ |
| Semantic search | ❌ | ❌ | ✅ (built-in) |
| Time travel queries | ❌ | ❌ | ✅ |
| WASM compute-near-data | ❌ | ❌ | ✅ |
| Content-addressed dedup | ❌ | ❌ | ✅ |
| Cold storage tiering | ✅ (Glacier) | ✅ (tier) | ✅ (any S3 backend) |
| Erasure coding | ✅ | ✅ | ✅ (Reed-Solomon) |
| K8s operator | ❌ | ✅ | ✅ |
| Single binary | ❌ | ✅ | ✅ |
| Open source | ❌ | ✅ (AGPL) | ✅ (Apache 2.0) |

Pranor Vault isn't trying to replace AWS S3 for everyone. It's for teams who want self-hosted object storage that does more than store bytes — storage that understands, searches, and transforms its contents.

---

## Getting Started

```bash
# Clone and build
git clone https://github.com/vyuvaraj/Pranor Vault.git
cd Pranor Vault
go build -o servstore ./cmd/servstore

# Run (no auth, port 9000)
./servstore

# Or with full security
./servstore --auth --access-key "mykey" --secret-key "mysecret" \
  --encryption-key "my-passphrase" \
  --tls-cert ./server.crt --tls-key ./server.key

# Open web console
# http://localhost:9000
```

The web console gives you drag-and-drop uploads, bucket management, versioning controls, and object version history — embedded directly in the binary.

---

## What's Next

- More embedding models for semantic search (currently TF-IDF, exploring dense vector support)
- WASM transform chaining (pipe output of one transform into another)
- S3 Select-style query API for structured data (JSON/CSV/Parquet)
- Expanded RBAC with policy language
- Performance benchmarks against MinIO (coming soon)

---

## Open Source vs. Enterprise Storage Features

Pranor Vault offers both free community-supported and commercial tiers:

* **Open Source (OSS)**: Pure S3 REST API compliance, object versioning, WORM locking, TF-IDF semantic text indexing, server-side WASM transforms (compute-near-data), and local storage pooling.
* **Enterprise Edition (EE)**: Adds multi-region active-active database replication, distributed Raft clustering with cross-region consistency, Reed-Solomon erasure coding, auto-healing background scrubbers, and 24/7 SLA support.

For enterprise scale deployments, reach out to the core team at **pranor@gmail.com**.

---

## Links

- **GitHub**: [github.com/vyuvaraj/Pranor Vault](https://github.com/vyuvaraj/Pranor Vault)
- **License**: Apache 2.0
- **Language**: Pure Go, zero CGO dependencies

---

*If you need object storage that goes beyond put-and-get — search, time travel, server-side compute — give Pranor Vault a try. It's a single binary that runs anywhere Go compiles.*

*— Yuvaraj*
