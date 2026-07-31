# Completed Initiatives (Phases 31-35 Archive)

This document archives completed initiatives for reference, keeping the main roadmap lean.

## Phase 32: Pranor Lock & Pranor Secret Standalone & Hardening (Completed)

> **Goal:** Support zero-dependency standalone execution modes for both components and implement enterprise-grade security and transport hardening.

### Pranor Lock Standalone & Hardening

| # | Item | Description | Status |
|---|------|-------------|--------|
| SL.6 | **Zero-Dependency Standalone Mode** | Support loading standalone server configuration from `servlock.yaml` (ports, backends) without requiring mesh/shared auth | [x] |
| SL.7 | **API Key Authentication** | Implement API Key token header authorization for standalone clients to access the locking APIs securely | [x] |
| SL.8 | **Lease Event Pub/Sub** | Implement SSE (Server-Sent Events) or WebSocket channels for lock release notifications to eliminate polling | [x] |
| SL.9 | **TLS & mTLS Transport Hardening** | Support native TLS and mutual TLS server configs inside the binary for secure client connection tunnels | [x] |

### Pranor Secret Standalone & Hardening

| # | Item | Description | Status |
|---|------|-------------|--------|
| SS.6 | **Zero-Dependency Standalone Mode** | Support loading config from `servsecret.yaml` (storage path, encryption schemes, cache rules, auth keys) | [x] |
| SS.7 | **Automated Backup & Recovery** | Configure automated scheduled encrypted backups of the secrets database to local storage or S3/MinIO objects | [x] |
| SS.8 | **Dynamic Environment Injector** | Create helper command `servsecret env run --cmd "app"` to fetch secrets and inject them directly to child processes | [x] |
| SS.9 | **Key Rotation Schedules** | Support background rotation of the master key on a configurable period (e.g. 90 days) with backup keys retention | [x] |

## Phase 33: Pranor Lock & Pranor Secret Advanced Capabilities (Completed)

> **Goal:** Enhance `Pranor Lock` and `Pranor Secret` with advanced operational, architectural, and security capabilities.

### Pranor Lock Advanced Capabilities

| # | Item | Description | Status |
|---|------|-------------|--------|
| SL.10 | **Active Heartbeat-Based Lease Extensions** | Automatic lease renewals based on client connection heartbeat pings | [x] |
| SL.11 | **Consensus-Based Clustering with Raft** | support local distributed clustering using Raft consensus mechanism | [x] |
| SL.12 | **Dynamic Waiter Priority Queues** | Allow lock waiters to specify priority levels to bypass standard FIFO queues | [x] |
| SL.13 | **Lease Hold-Time Alerting** | Monitor lease durations and alert on zombie locks held past threshold | [x] |
| SL.14 | **Read-Only Shared Locks** | Implement shared (read) and exclusive (write) multi-access locking | [x] |

### Pranor Secret Advanced Capabilities

| # | Item | Description | Status |
|---|------|-------------|--------|
| SS.10 | **Secret Versioning and Rollbacks** | Maintain a timeline of secret revisions allowing rollout rollback | [x] |
| SS.11 | **Dynamic Database Credentials** | Just-In-Time generation of short-lived database user logins | [x] |
| SS.12 | **Client-Side Zero-Knowledge Decryption** | Secure mode keeping the master decryption key strictly client-side | [x] |
| SS.13 | **IP CIDR & Geofencing Policies** | Restrict secret lookups by network subnets or geofenced zones | [x] |
| SS.14 | **Secret Leak Detection & Scanning** | Scan payloads/logs to flag plaintext exposure of managed secrets | [x] |

## Phase 31: Pranor Lock & Pranor Secret Ecosystem Integration (Completed)

> **Goal:** Integrate `Pranor Lock` and `Pranor Secret` as first-class primitives throughout other core components in the Pranor ecosystem.

### Ecosystem Integration Backlog

| # | Item | Target Component | Description | Status |
|---|------|------------------|-------------|--------|
| EI.1 | **Pranor Gate Secret Integration** | Pranor Gate | Fetch SSL/TLS certificates and JWT verification keys dynamically from `Pranor Secret` instead of static local files | [x] |
| EI.2 | **Pranor Flow Distributed Locking** | Pranor Flow | Integrate `Pranor Lock` into the task execution engine to prevent duplicate transaction runs in multi-instance clusters | [x] |
| EI.3 | **Pranor Console Management Dashboard** | Pranor Console | Add UI views to list active locks (via `Pranor Lock` observability) and manage keys/secret rotations (via `Pranor Secret` APIs) | [x] |
| EI.4 | **Pranor Built-in Lock/Secret Operators** | Pranor | Introduce native runtime standard library operators (e.g., `secret("db.pass")` or `lock("resource") { ... }`) | [x] |
| EI.5 | **Pranor Chrono Scheduler Lock Gating** | Pranor Chrono | Use `Pranor Lock` in the job execution cycle to prevent scheduler drift and duplicate scheduler execution | [x] |

## Phase 34: Pranor Lock & Pranor Secret Enterprise & UI (Completed)

> **Goal:** Build advanced Enterprise Edition (EE) capabilities and full integration with Pranor Console dashboard controls.

### Pranor Lock & Pranor Secret Enterprise & UI Roadmap

| # | Item | Type | Description | Status |
|---|------|------|-------------|--------|
| EE.1 | **Multi-Tenant Cryptographic Isolation** | EE | Segment storage namespaces and encrypt each tenant vault with dedicated KMS/HSM keys | [x] |
| EE.2 | **Shamir's Secret Sharing Bootstrapping** | OSS/EE | Require multiple operator key shards to unseal the master key on startup | [x] |
| EE.3 | **Pranor Console Locks Dashboard** | UI | Real-time monitoring UI in Pranor Console for active locks, priority waiter lists, and SSE events | [x] |
| EE.4 | **Pranor Console Vault Explorer** | UI | Field-masked secrets manager GUI for secret CRUD operations and rollback actions | [x] |
| EE.5 | **Automatic Dynamic Credential Rotation** | EE | Out-of-the-box cron rotation connectors for dynamic SQL/NoSQL logins | [x] |
| EE.6 | **Standalone CLI Administration Client** | OSS | Unified CLI `servlockctl` / `servsecretctl` for administration from terminals | [x] |

## Phase 35: Pranor Language Ergonomics (Completed)

> **Goal:** Reduce reliance on Python daemon wrappers and external utilities by introducing native built-in namespaces and language sugar that make common day-to-day operations expressive.

| # | Item | Effort | Component | Description | Status |
|---|------|--------|-----------|-------------|--------|
| LE.1 | **`exec` namespace — Native Shell/Script Execution** | Small | Pranor | Add a built-in `exec` namespace: `exec.run("powershell -File ./deploy.ps1")`. Returns `{ stdout, stderr, exitCode }`. | [x] |
| LE.2 | **`csv` built-in namespace** | Small | Pranor | Promote `stdlib/csv.pnr` to a first-class compiler built-in: `csv.parse()` and `csv.stringify()`. | [x] |
| LE.3 | **`xml` namespace — Native XML Parsing** | Small | Pranor | Add a built-in `xml` namespace: `xml.parse()` and `xml.stringify()`. | [x] |
| LE.4 | **`yaml` namespace — YAML Parsing & Emit** | Small | Pranor | Add a built-in `yaml` namespace: `yaml.parse()` and `yaml.stringify()`. | [x] |
| LE.5 | **`file` namespace — Direct File I/O** | Small | Pranor | Add a built-in `file` namespace: `file.read()`, `file.write()`, `file.exists()`, `file.list()`. | [x] |
| LE.6 | **`path` namespace — File Path Utilities** | Small | Pranor | Add a built-in `path` namespace wrapping `path/filepath`: `path.join()`, `path.dirname()`, `path.basename()`. | [x] |
| LE.7 | **`regex` namespace — Regular Expression Support** | Small | Pranor | Add a built-in `regex` namespace: `regex.match()`, `regex.find()`, `regex.replace()`. | [x] |
| LE.8 | **`math` namespace — Mathematical Functions** | Small | Pranor | Add a built-in `math` namespace: `math.floor()`, `math.ceil()`, `math.round()`, `math.abs()`, `math.pow()`. | [x] |
| LE.9 | **`encoding` namespace — Base64 & Hex** | Small | Pranor | Add a built-in `encoding` namespace: `encoding.base64.encode()`, `encoding.hex.encode()`. | [x] |
| LE.10 | **`hash` namespace — Cryptographic Hashing** | Small | Pranor | Add a built-in `hash` namespace: `hash.md5()`, `hash.sha256()`, `hash.sha512()`, `hash.hmac()`. | [x] |
| LE.11 | **`uuid` namespace — Unique ID Generation** | Small | Pranor | Add a built-in `uuid` namespace: `uuid.v4()`, `uuid.v7()`. | [x] |
| LE.12 | **`rand` namespace — Random Value Generation** | Small | Pranor | Add a built-in `rand` namespace: `rand.int()`, `rand.float()`, `rand.string()`, `rand.bool()`. | [x] |
| LE.13 | **`url` namespace — URL Parsing & Encoding** | Small | Pranor | Add a built-in `url` namespace: `url.parse()`, `url.encode()`, `url.decode()`. | [x] |
| LE.14 | **`env` namespace — Typed Environment Variables** | Small | Pranor | Extend the current `env` helper to `env.get()`, `env.require()`, `env.int()`, `env.bool()`. | [x] |
| LE.15 | **Optional chaining (`?.`)** | Medium | Pranor | Add null-safe member access: `user?.address?.city`. | [x] |
| LE.16 | **Spread operator (`...`) for arrays and maps** | Medium | Pranor | Add spread syntax: `[...arr1, newItem]` and `{...baseConfig, timeout: 30}`. | [x] |
| LE.17 | **`time` namespace — Full Date/Time Support** | Medium | Pranor | Overhaul the `time` built-in into a comprehensive date/time namespace. | [x] |
| LE.18 | **Multiline string dedentation** | Small | Pranor | Strip common leading whitespace from indented raw string literals. | [x] |
| LE.19 | **`jwt` namespace — User-Accessible Token Signing** | Small | Pranor | Expose a first-class `jwt` namespace: `jwt.sign()`, `jwt.verify()`, `jwt.decode()`. | [x] |
| LE.20 | **`compress` namespace — Gzip & Deflate** | Small | Pranor | Add a built-in `compress` namespace: `compress.gzip()`, `compress.ungzip()`, `compress.deflate()`. | [x] |
| LE.21 | **`semver` namespace — Semantic Version Parsing** | Small | Pranor | Add a built-in `semver` namespace: `semver.parse()`, `semver.compare()`, `semver.satisfies()`. | [x] |
| LE.22 | **`duration` namespace — Human-Readable Time Spans** | Small | Pranor | Add a built-in `duration` namespace: `duration.parse()`, `duration.format()`, `duration.since()`. | [x] |
| LE.23 | **`format` namespace — Human-Readable Value Formatting** | Small | Pranor | Add a built-in `format` namespace: `format.bytes()`, `format.number()`, `format.percent()`. | [x] |
| LE.24 | **`ip` namespace — IP Address Utilities** | Small | Pranor | Add a built-in `ip` namespace: `ip.parse()`, `ip.isPrivate()`, `ip.inCIDR()`, `ip.version()`. | [x] |
| LE.25 | **`dns` namespace — DNS Lookup Utilities** | Small | Pranor | Add a built-in `dns` namespace: `dns.lookup()`, `dns.txt()`, `dns.pnr()`. | [x] |
| LE.26 | **`multipart` namespace — File Upload Parsing** | Small | Pranor | Add a built-in `multipart` namespace: `multipart.parse()`. | [x] |
| LE.27 | **`diff` namespace — Structural Diff & Patch** | Small | Pranor | Add a built-in `diff` namespace: `diff.text(a, b)` and `diff.json(objA, objB)`. | [x] |
| LE.28 | **`proto` namespace — Protocol Buffer Support** | Medium | Pranor | Add a built-in `proto` namespace: `proto.encode(obj, schema)` and `proto.decode(bytes, schema)`. | [x] |

### `time` Namespace Full Breakdown (LE.17 detail)

| # | Item | Effort | Component | Description | Status |
|---|------|--------|-----------|-------------|--------|
| LE.29 | **`time.parse(str, layout)`** | Small | Pranor | Parse a date/time string into a time value. | [x] |
| LE.30 | **`time.format(t, layout)`** | Small | Pranor | Format a time value into a custom string. | [x] |
| LE.31 | **`time.inZone(t, tz)` — Timezone Conversion** | Small | Pranor | Convert a time value to a named timezone. | [x] |
| LE.32 | **`time.utc(t)` and `time.local(t)`** | Small | Pranor | Convenience shorthands for UTC/Local conversion. | [x] |
| LE.33 | **`time.add(t, duration)` — Time Arithmetic** | Small | Pranor | Add a duration to a time value. | [x] |
| LE.34 | **`time.sub(t1, t2)` — Time Difference** | Small | Pranor | Compute the difference between two time values in seconds. | [x] |
| LE.35 | **`time.before(t1, t2)` and `time.after(t1, t2)`** | Small | Pranor | Boolean time comparison. | [x] |
| LE.36 | **`time.fromUnix(seconds)` — Unix to Time** | Small | Pranor | Convert a Unix epoch integer back to a formatted string. | [x] |
| LE.37 | **`time.components(t)` — Date Part Extraction** | Small | Pranor | Destructure a time value into its parts. | [x] |
| LE.38 | **Predefined layout constants** | Small | Pranor | Expose common format constants (RFC3339, DATE, etc.). | [x] |

### Go Interop Escape Hatches

| # | Item | Effort | Component | Description | Status |
|---|------|--------|-----------|-------------|--------|
| LE.39 | **`@inline go` blocks — Raw Go Code Embedding** | Medium | Pranor | Add `@inline go fn` declarations. | [x] |
| LE.40 | **`extern fn` → `go:` binding improvements** | Small | Pranor | Support extern fn declarations in block layouts, method receivers, and alias auto-inference. | [x] |

