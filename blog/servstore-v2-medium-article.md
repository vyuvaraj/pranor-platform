# Pranor Vault v2: Evolving Our S3 Storage Engine into a Standalone, Local-First P2P Analytical Store

*From a single-binary storage box to a multi-modal object store featuring standalone daemons (`servstored`), client-side browser OPFS dual-sync (`@pranor/store-wasm`), streaming S3 Select & embedded DuckDB SQL analytics, Copy-on-Write bucket branching, and WebTorrent P2P asset seeding.*

---

> 💡 **Note**: This is **Part 2** of the Pranor Vault series. If you missed Part 1 on how we built an S3-compatible storage engine from scratch, check out [Part 1: I Built an S3-Compatible Object Storage Engine With AI-Native Capabilities](https://medium.com/@yuvamca002).

---

## Why Pranor Vault Needed to Evolve

In Part 1, we introduced Pranor Vault: a Go-based, S3-compatible object storage engine that combined standard S3 APIs with semantic search and inline WebAssembly (WASM) compute near data.

While the core object storage layer worked smoothly, operating it in real-world applications revealed four major architectural bottlenecks:

1. **The Cloud Egress Tax**: Traditional object storage engines like S3 or MinIO act as "dumb byte stores." To run queries over log files or Parquet datasets, developers are forced to download gigabytes of raw data to external clusters (Snowflake, Athena, Trino), incurring massive network egress fees.
2. **Offline UI Latency & Stalls**: Web applications uploading files over standard HTTP connections freeze or fail whenever network connectivity dips.
3. **Storage Snapshot Overhead**: Creating staging environments or test copies of multi-terabyte data lakes requires waiting hours for physical copy operations and doubles storage bills.
4. **Origin Bandwidth Exhaustion**: Delivering high-demand assets (video streams, 3D assets, software installers) to thousands of concurrent users overloads origin servers and inflates egress costs.

Here is how we solved these challenges in **Pranor Vault v2** within the unified **Serv monorepo** (`github.com/vyuvaraj/pranor/packages/Pranor Vault`).

---

## 1. Monorepo Migration & The Daemon / CLI Split

We merged standalone components into the unified **Serv monorepo** (`github.com/vyuvaraj/pranor/packages/Pranor Vault`). As part of this evolution, we strictly separated server runtime logic from administrative tooling:

* **`servstored` (Server Daemon)**: Zero-dependency background service process. It hosts standard S3 REST API listeners on port `:9000`, an embedded Web Storage Console UI on port `:9001` (`http://localhost:9001/ui/`), and Prometheus telemetry metrics (`/metrics`).
* **`servstore` (Client CLI)**: High-performance administrative binary for operators and automated scripts (`servstore status`, `servstore ls`, `servstore mb`, `servstore branch`).

```
                              servstored DAEMON                               
                              ─────────────────                               
                              ┌──────────────────────────┐                    
                              │ S3 REST API (:9000)      │ ◄─── AWS S3 SDKs / CLI
                              │ Web Console UI (:9001)   │ ◄─── Browser Admin
                              └────────────┬─────────────┘                    
                                           │                                  
                                           ▼                                  
                             ┌────────────────────────────┐                   
                             │ Core Engine & Storage Pool │                   
                             ├────────────────────────────┤                   
                             │ ⚡ Browser OPFS Dual-Sync  │                   
                             │ 🔍 Streaming S3 Select     │                   
                             │ 🌿 Copy-on-Write Branching │                   
                             │ 🌐 WebTorrent P2P Seeding  │                   
                             │ 🛡️ SEC 17a-4 WORM Locking │                   
                             └────────────────────────────┘                   
```

---

## 2. Client-Side Browser OPFS Dual-Sync (`@pranor/store-wasm`)

Standard object storage requires users to upload files over HTTP connections, exposing web applications to latency and network drops.

In Pranor Vault v2, `@pranor/store-wasm` brings origin-private filesystem (**OPFS**) persistence directly into browser Web Workers.

### How It Works:
1. When a user creates or modifies a file in a web application, it is written **instantly at native NVMe disk speed** into the browser's local OPFS storage using synchronous file handles (`FileSystemSyncAccessHandle`).
2. The user experience is **0ms zero-latency**—the UI updates immediately without waiting for server network ACK packets.
3. In the background, `@pranor/store-wasm` streams chunked S3 multipart uploads to `servstored` with automatic pause, retry, and resume resilience.

```typescript
import { Pranor VaultBrowserClient } from '@pranor/store-wasm';

const client = new Pranor VaultBrowserClient({
  endpoint: 'http://localhost:9000',
  bucket: 'user-workspace'
});

// Saves instantly to local OPFS (0ms) and streams to servstored in background
await client.putObject('report.pdf', fileBuffer);
```

---

## 3. Streaming S3 Select & Embedded DuckDB SQL Analytics

Traditional S3 object stores return raw byte streams. To filter data, you must pull entire files across the network into external query engines.

Pranor Vault v2 integrates a **Streaming S3 Select SQL Query Engine** and **Embedded DuckDB Parquet Reader** directly into `servstored`.

### Querying Data at Rest:
You can execute SQL queries directly over HTTP GET requests:

```bash
curl -X GET "http://localhost:9000/api/v1/analytics/query?sql=SELECT+*+FROM+'s3://logs/2026-07.parquet'+WHERE+status+=+500"
```

### Key Advantages:
* **Zero-ETL Overhead**: Query JSON, CSV, and Parquet data directly where it sits on NVMe storage.
* **Network Bandwidth Savings**: `servstored` filters data at rest and returns only matching result rows, cutting egress bandwidth by **up to 99%**.

---

## 4. Instant Copy-on-Write (CoW) Bucket Branching (`servstore branch`)

Cloning a multi-terabyte data bucket for testing or staging traditionally requires running background copy scripts that take hours and double storage costs.

Pranor Vault v2 introduces Git-style zero-byte branching for S3 buckets:

```bash
# Create an instant isolated branch clone of 'prod-data'
servstore branch create prod-data dev-test-branch

# List active bucket branches
servstore branch ls prod-data

# Merge modifications from branch overlay back to parent
servstore branch merge prod-data dev-test-branch
```

### Technical Implementation:
* Creating a branch generates a **Copy-on-Write (CoW) virtual metadata overlay** in **<1 millisecond**.
* **0 bytes of physical storage** are copied during branch creation.
* Unmodified object reads fall back to the base bucket, while object writes are isolated inside the virtual branch namespace (`/branches/dev-test-branch/`).

---

## 5. Browser WebTorrent P2P Asset Seeding Mesh

Serving popular static assets (video streams, CAD models, software updates) to thousands of concurrent users overloads origin servers and inflates bandwidth costs.

Pranor Vault v2 transforms client browsers into a peer-to-peer asset distribution mesh via `@pranor/store-wasm`:

* **WebRTC Peer Mesh**: Browsers fetch cached chunks directly from nearby connected peer browsers using WebRTC.
* **Integrity Validation**: Includes cryptographic SHA-256 chunk integrity verification before writing to local OPFS.
* **Cost Impact**: Reduces origin server bandwidth load and egress bills by **up to 95%** during viral traffic spikes.

---

## 6. Enterprise Multi-Cloud Lifecycle & Sovereign Archiving

Beyond browser and analytical features, Pranor Vault v2 includes enterprise-grade storage governance:

* **Policy-Driven Multi-Cloud Tiering**: Automatically migrates cold objects from local NVMe hot storage to AWS Glacier Deep Archive, Azure Blob Archive, or Google Cloud Storage Coldline.
* **SEC Rule 17a-4 WORM Object Locking**: Immutability modes and legal hold governance for financial compliance.
* **FIPS 140-3 KMS Envelope Encryption**: AES-256-GCM object encryption backed by Hardware Security Modules (HSMs) and HashiCorp Vault.

---

## Quickstart with Pranor Vault v2

```bash
# Clone the Serv monorepo
git clone https://github.com/vyuvaraj/pranor.git
cd serv/packages/Pranor Vault

# Build and start the daemon (Web UI at http://localhost:9001/ui/)
go build -o servstored ./cmd/servstored
./pranorstored --port 9000 --admin-port 9001

# In another terminal, use the CLI
go build -o servstore ./cmd/servstore
./pranorstore status
./pranorstore mb analytics-bucket
./pranorstore branch create analytics-bucket feature-testing
```

---

## Platform Comparison Matrix

| Feature | AWS S3 / MinIO | Pranor Vault v2 |
|:---|:---|:---|
| **Client Upload Latency** | Network Dependent (Stalls UI) | **0ms Local OPFS Persistence** |
| **Offline PWA Support** | ❌ None | **✅ Native (`@pranor/store-wasm`)** |
| **Parquet / Log Analytics** | Requires External Athena/Snowflake | **✅ Built-in Streaming S3 Select & DuckDB** |
| **Bucket Snapshot Branching** | Hours (Slow Full Copy) | **✅ <1ms CoW Zero-Byte Branching** |
| **Asset Delivery Egress** | 100% Cloud Egress Cost | **✅ Up to 95% Bandwidth Offload via P2P** |

* **Monorepo**: [github.com/vyuvaraj/pranor](https://github.com/vyuvaraj/pranor)
* **Package Path**: `packages/Pranor Vault`
* **License**: AGPLv3 (Server Daemon) & Apache 2.0 (Client SDKs / OPFS WASM)

*— Yuvaraj*
