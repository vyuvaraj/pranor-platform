# Pranor Publication & Marketing Manifest

This document tracks the Medium article series, documentation layout, and marketing assets for the Pranor microservice ecosystem.

---

## 1. Hero Screenshots Manifest (`pranor run`)

To demonstrate the simplicity of the one-command developer loop, the following terminal outputs and dashboard visualization layouts have been compiled:

* **Hero Terminal Output (`pranor dev` / `pranor run`)**:
  ```bash
  $ pranor dev
  [12:00:00] Starting Pranor Stack...
  [12:00:01] Pranor Hub listening on http://localhost:8089 (Heartbeat TTL: 10s)
  [12:00:02] Pranor Gate listening on http://localhost:8080 (Auth Secret Loaded)
  [12:00:02] Pranor Trace listening on http://localhost:8090 (OTLP Endpoint active)
  [12:00:03] Pranor Console dashboard listening on http://localhost:3000
  [12:00:04] Pranor Tunnel mesh transport: mTLS handshake established successfully.
  [12:00:05] Pranor Deploy deploying services...
  [12:00:06] Service 'user-service' deployed successfully on port 8086.
  [12:00:07] [SYSTEM_ALERT] Circuit breaker for target 'payment-service' changed to CLOSED.
  [12:00:08] All systems operational. Go to console: http://localhost:3000
  ```

* **Interactive Distributed Tracing Waterfall Visual**: Gantt-chart rendering demonstrating bottleneck spans, critical paths, and slow database warnings (`⚠️ SLOW QUERY` badges).
* **Live Topology Mesh**: Graph canvas demonstrating traffic flow particles, latency gradients, and circuit-breaker warnings (tripped paths render red).

---

## 2. Medium Article Series (9 Components)

All nine articles have been prepared and cross-linked. The publication sequence is detailed below:

### Part 1: Pranor Architecture Overview
* **Title**: *Designing a Zero-Dependency Microservice Framework in Go*
* **Summary**: Explains the core philosophy of zero external dependencies, how `Pranor Core` acts as the utility wrapper, and how the modules connect via OTel.

### Part 2: Pranor Hub & Service Discovery
* **Title**: *Distributed Service Discovery: Building Pranor Hub from Scratch*
* **Summary**: Deep-dive into heartbeat TTL eviction loops, CA certificates distribution, and dynamic registry-lookup algorithms.

### Part 3: Pranor Gate & Dynamic Routing
* **Title**: *Creating a High-Performance API Gateway with Dynamic Hot-Reloading*
* **Summary**: Discusses reverse proxies, route tables stored on Pranor Vault (S3), and hot-reloading configurations without server restarts.

### Part 4: Pranor Mesh & Client-Side Resiliency
* **Title**: *Mesh Networks at the Client: Load Balancing and Circuit Breaking*
* **Summary**: Implementation details of mTLS transport tunnels, custom HTTP `RoundTrippers` with round-robin, and state transitions to `Open` (tripped).

### Part 5: Pranor Trace & Gantt Chart Visualizations
* **Title**: *Distributed Tracing: Implementing OTel Ingestion & Waterfall UIs*
* **Summary**: How Pranor Trace accepts OTLP trace payloads, builds hierarchical span trees, and renders interactive Gantt charts.

### Part 6: Pranor Vault & Cold Tier Storage
* **Title**: *Scalable Observability: Eviction Policies & S3 Cold Tier Archiving*
* **Summary**: How Pranor Trace handles high-scale memory eviction by archiving inactive trace payloads to S3-compatible cold tier stores.

### Part 7: Pranor Pulse & Message Brokerage
* **Title**: *High-Throughput Pub/Sub: Creating a Zero-Config Message Broker*
* **Summary**: Custom TCP-based messaging queue protocol with partitioned consumer offsets.

### Part 8: Pranor Chrono & Distributed Scheduling
* **Title**: *Distributed Crons: Reliable Scheduling Across Microservices*
* **Summary**: Cron execution loops using leader-election registries.

### Part 9: Pranor Deploy & In-Process Isolation
* **Title**: *Next-Gen Application Hosting: WASM and Docker Sandboxing*
* **Summary**: The future of Pranor Deploy, deploying services compiled to WASM in-process and orchestrating isolated Docker runtimes.

---

## 3. VS Code Extension & Links
* All component articles prominently link to the **Pranor Developer Tools** VS Code extension.
* VS Code extension provides:
  * Syntax highlighting for `.pnr` files.
  * Real-time telemetry feed from the console inside the editor sidebar.
  * Right-click deploy to local Pranor Deploy.
