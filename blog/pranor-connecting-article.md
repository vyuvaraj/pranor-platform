# Building the Pranor: A Unified, AI-Native Backend Ecosystem

*How a custom compiler, an S3 object store, a STOMP queue, an API gateway, and a glassmorphic console fit together to rethink microservices.*

---

## The Microservices Tax

Modern software engineering is dominated by "glue." 

If you want to build a simple event-driven microservice system that accepts user uploads, runs a filter, publishes an event, and saves metadata, you are immediately taxed. You have to configure an API Gateway (Kong/Nginx), a message broker (RabbitMQ/Kafka), an object storage container (MinIO/S3), an observability stack (OpenTelemetry, Collector, Jaeger), and write hundreds of lines of boilerplate setup in Go, Node.js, or Java to tie them together.

Then comes the Kubernetes YAML, the Dockerfiles, the IAM policies, and the connection string management. 

We spend 80% of our time setting up infrastructure and only 20% writing business logic. The infrastructure ceremony has swallowed the developer experience.

I wanted to change that. I wanted an ecosystem where the programming language, the database, the queues, the storage, the gateways, and the observability dashboards were all designed to work *together* out of the box.

I call this ecosystem the **Pranor**.

---

## The Pranor Architecture

The Pranor is structured into four cooperative layers. By designing each component to speak the same language and respect the same telemetry protocols, we eliminate the integration tax.

```
┌───────────────────────────────────────────────────────────────────────┐
│                           DEVELOPER LAYER                             │
│       Pranor (.pnr compiler) ➔ Defines HTTP, Pub/Sub, DB           │
├───────────────────────────────────────────────────────────────────────┤
│                           PLATFORM LAYER                              │
│       Pranor Gate (API Gateway) ➔ Edge filters, AI guards, Routing       │
├───────────────────────────────────────────────────────────────────────┤
│                        INFRASTRUCTURE LAYER                           │
│       Pranor Pulse (Broker)      │    Pranor Vault (S3 Storage)             │
│       STOMP + WASM Transforms │    Raft + Semantic Search             │
├───────────────────────────────────────────────────────────────────────┤
│                         OBSERVABILITY LAYER                           │
│       Pranor Console ➔ OTel waterfalls, Hash rings, Dynamic admin        │
└───────────────────────────────────────────────────────────────────────┘
```

Here is how the pieces fit together.

---

## 1. The Developer Layer: Pranor

Everything starts with **Pranor**, a domain-specific programming language that compiles to native Go binaries. Instead of writing connection pools and framework boilerplate, infrastructure is declared directly in syntax:

```pranor
// Declare infrastructure
server "8080"
broker "pranor-pulse://localhost:61613"

// Define a REST API route
route "POST" "/upload" (req) {
    let payload = req.body
    log.info("Processing upload...")
    
    // Publish directly using language primitives
    publish "raw-uploads" payload
    return { "status": "queued", "id": time.now() }
}
```

No imports. No frameworks. The compiler does the heavy lifting, generating optimized Go source code, wiring up HTTP servers, and initializing STOMP clients.

---

## 2. The Platform Layer: Pranor Gate

At the entry point of your system sits **Pranor Gate**, a programmable API Gateway. 

Because Pranor Gate is built for the Pranor, it auto-detects routes declared in Pranor binaries. Beyond reverse proxying, it introduces **WASM-powered inline middleware** and **AI-native guards**:
- **Dynamic WASM filters**: Upload Go/Rust code compiled to WebAssembly to modify requests or validate headers in isolated sandboxes at runtime.
- **AI Prompt Guard**: Inspects incoming JSON bodies to block jailbreak attempts and prompt injections before they reach your expensive LLM servers.
- **Semantic Caching**: Caches response vectors based on prompt similarity checks, avoiding duplicate LLM queries.
- **PII Redaction**: Dynamically scrubs credit card and social security numbers from response streams.

---

## 3. The Infrastructure Layer: Pranor Pulse & Pranor Vault

When your microservices need to communicate or persist data, they use Pranor's native infrastructure engines.

### Pranor Pulse (Message Broker)
Pranor Pulse is a STOMP-compliant message broker that features **Compute-in-Queue**. 

Instead of routing messages to an external worker service to transform them, Pranor Pulse runs WebAssembly transformers *inside* the broker. It also includes sliding-window deduplication and Dead Letter Queues (DLQs) to handle transformation errors gracefully without clogging the pipeline.

### Pranor Vault (Object Storage)
Pranor Vault is a distributed, Raft-consistent, S3-compatible object storage engine. 

While standard S3 buckets are black boxes, Pranor Vault is AI-native:
- **Semantic Search**: Text uploads are indexed using TF-IDF, allowing developers to query objects by meaning rather than file paths.
- **WASM Compute-near-Data**: Executes WASM transform binaries directly on storage nodes to process objects without moving bytes over the network.
- **Time Travel**: Retrieve the exact state of any bucket or key at a historical timestamp using native version history.

---

## 4. The Observability Layer: Pranor Console

Observability is usually an afterthought. In the Pranor, it is the glue. 

**Pranor Console** is a premium glassmorphic dashboard that connects to all downstream engines using the `PRANOR_DISCOVERY` protocol. It provides:
- **Distributed OTel Span Waterfalls**: Tracks requests across network boundaries (Gateway ➔ Custom Service ➔ Broker ➔ Storage) and visualizes execution latency in a cascaded chart.
- **Consistent Hash Ring Visualizer**: Displays how data shards partition across storage nodes in real-time.
- **Administrative Hub**: Drag-and-drop WASM filters, view topic queues, and publish test payloads directly from the UI.

---

## A Lifecycle of a Request in the Pranor

To see the ecosystem in action, let's follow a single user request:

1. **Client Request**: A client uploads an image containing personal information.
2. **Pranor Gate Intercept**: Pranor Gate intercepts the upload. A registered **WASM filter** validates the client's JWT auth token (synced across the ecosystem).
3. **Pranor Service**: The request is routed to a custom binary compiled from `Pranor`. It issues a `publish "images.raw"` command.
4. **Pranor Pulse Processing**: Pranor Pulse receives the message. It executes a serverless **WASM transform** that redacts PII strings from the image metadata.
5. **Pranor Vault Persist**: The processed image is stored in **Pranor Vault**. Because Pranor Vault is configured with content-addressed storage (CAS), it deduplicates the payload and writes it to a distributed hash ring.
6. **Pranor Console Trace**: The developer opens Pranor Console, views the OTel Waterfall Chart, sees the exact 140ms path, audits the WASM execution time, and verifies the consistent hash node where the image was stored.

---

## The Power of the Shared Protocol

What makes this possible is zero-config wiring. We enforce three shared environment variables across all components:
1. `PRANOR_DISCOVERY`: A JSON string or file path containing the addresses of all services, allowing the gateway, storage, broker, and console to find each other instantly.
2. `PRANOR_JWT_SECRET`: A shared token signing secret that allows unified, secure authorization across all components.
3. `PRANOR_OTLP_ENDPOINT`: A shared trace exporter endpoint, ensuring every component reports spans to the same observability pool.

By agreeing on these three primitives, we eliminate configuration friction.

---

## Getting Started

To spin up the entire Pranor locally:

```bash
# 1. Start Pranor Vault (Storage) on :8081
./pranorstore.exe

# 2. Start Pranor Pulse (Broker) on :8082
./pranorqueue.exe

# 3. Start Pranor Gate (Gateway) on :8080
./pranorgate.exe --config=config.json

# 4. Launch Pranor Console on :8083
export PRANOR_DISCOVERY='{
  "gate":          "http://localhost:8080",
  "store":         "http://localhost:8081",
  "queue":         "http://localhost:8082"
}'
./pranorconsole.exe
```

Open `http://localhost:8083` in your browser. You are now running a fully integrated, observable, programmable microservice platform.

---

## What's Next

The Pranor is designed to grow. The platform has expanded to include major new core components:
- **Pranor Auth**: An OIDC identity and token validation provider with lockout and recovery reset flows.
- **ServDB**: A high-efficiency connection pooler and read/write splitting proxy database driver.
- **Pranor Notify**: A transactional message delivery hub supporting SMTP, Slack, and SMS channels with DLQ retries.
- **Pranor Flow**: A stateful DAG-based workflow orchestrator and Saga compensation engine.

---

## Open Source vs. Enterprise Editions

To support standard self-hostable development while addressing enterprise scale and security requirements, the Pranor ecosystem is split into two models:

* **Open Source (OSS)**: Under the Apache 2.0 license, providing the full core DSL compiler (`Pranor`), single-tenant reverse proxy (`Pranor Gate`), object storage (`Pranor Vault`), queue (`Pranor Pulse`), and local dashboard observability (`Pranor Console`).
* **Enterprise Edition (EE)**: Provides production-scale extensions like multi-region active-active database replication, strict multi-tenant JWT context isolation, high-throughput prompt injection firewalls, federated GraphQL gateways with LLM semantic routing, and 24/7 SLA-backed support.

For commercial inquiries or SLA support, contact the core team directly at **pranor@gmail.com**.

---

## Links

- **Pranor**: [github.com/vyuvaraj/Pranor](https://github.com/vyuvaraj/Pranor)
- **Pranor Vault**: [github.com/vyuvaraj/Pranor Vault](https://github.com/vyuvaraj/Pranor Vault)
- **Pranor Pulse**: [github.com/vyuvaraj/Pranor Pulse](https://github.com/vyuvaraj/Pranor Pulse)
- **Pranor Gate**: [github.com/vyuvaraj/Pranor Gate](https://github.com/vyuvaraj/Pranor Gate)
- **Pranor Console**: [github.com/vyuvaraj/Pranor Console](https://github.com/vyuvaraj/Pranor Console)
- **Pranor Auth**: [github.com/vyuvaraj/Pranor Auth](https://github.com/vyuvaraj/Pranor Auth)
- **ServDB**: [github.com/vyuvaraj/ServDB](https://github.com/vyuvaraj/ServDB)
- **Pranor Notify**: [github.com/vyuvaraj/Pranor Notify](https://github.com/vyuvaraj/Pranor Notify)
- **Pranor Flow**: [github.com/vyuvaraj/Pranor Flow](https://github.com/vyuvaraj/Pranor Flow)

---

*Backend development doesn't have to be fragmented. Welcome to a world where language, storage, queues, and gateways speak the same language.*

*— Yuvaraj*
