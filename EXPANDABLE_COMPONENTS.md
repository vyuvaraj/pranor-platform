# Pranor Expandable Components

> Components that could be added to the ecosystem in the future. None are committed — this is a living ideas document for when the time is right.

---

## Tier 1: High-Value Gaps (Missing from a complete backend platform)

### Pranor Secret — Secret & Credential Management
**Problem:** Every service needs secrets (DB passwords, API keys, JWT keys). Currently scattered via env vars with no rotation, no audit trail, no encryption-at-rest for sensitive config.

| Aspect | Detail |
|--------|--------|
| What it does | Centralized vault with encryption, rotation, access audit |
| Language keyword | `let dbPass = env.secret("DB_PASSWORD")` (already exists — needs backend) |
| Adapters | HashiCorp Vault, AWS Secrets Manager, Doppler, encrypted local store |
| CLI | `serv secret set KEY`, `serv secret rotate KEY`, `serv secret list` |
| Integrates with | Pranor Deploy (inject at deploy), Pranor Auth (key material), pranor.yaml |
| Standalone value | High — any app needs secrets management |
| Effort | Medium (2-3 weeks) |

---

### ServConfig — Live Configuration Service
**Problem:** Each service reads config independently at startup. No live reload, no environment scoping, no type validation, no centralized management.

| Aspect | Detail |
|--------|--------|
| What it does | Centralized config server with live push, environment scoping (dev/staging/prod), type validation, change audit |
| Language keyword | `let host = config("db.host")` (already exists — needs backend) |
| Features | Live reload (no restart), config inheritance, JSON Schema validation, diff preview |
| Integrates with | Pranor Console (config editor UI), Pranor Deploy (deploy-time injection), `pranor.yaml` |
| Standalone value | Medium — most useful within ecosystem |
| Effort | Medium (2-3 weeks) |

---

### Pranor Lock — Distributed Locking
**Problem:** Multi-instance services need mutual exclusion (payment processing, order fulfillment). Currently no ecosystem solution — teams roll their own Redis SETNX.

| Aspect | Detail |
|--------|--------|
| What it does | Cross-service mutex with TTL, deadlock detection, queueing, fencing tokens |
| Language keyword | `lock("order-123", 30s) { ... }` |
| Backends | Redis, etcd, Pranor Vault (CAS-based), in-memory (single node) |
| Features | Reentrant locks, read/write locks, lock queuing, automatic release on crash |
| Integrates with | Pranor Mesh (service-aware locking), Pranor Trace (lock contention spans) |
| Standalone value | High — any distributed system needs this |
| Effort | Small (1-2 weeks) |

---

### ServSearch — Full-Text & Semantic Search
**Problem:** Pranor Vault has vector search on objects, but there's no general-purpose search for application data (indexing DB records, structured queries with facets).

| Aspect | Detail |
|--------|--------|
| What it does | Search index with full-text + semantic (vector) queries, facets, filters |
| Language keyword | `search "meilisearch://host/index"` |
| Adapters | Meilisearch, Typesense, Elasticsearch, Pranor Vault (built-in), Zinc |
| Features | Auto-index on DB write, typo tolerance, faceted search, geo-search |
| Integrates with | Pranor Pool (index on write trigger), Pranor Vault (semantic layer) |
| Standalone value | High — any app with search needs this |
| Effort | Medium (3-4 weeks) |

---

### ServMetrics — Time-Series Metrics Storage
**Problem:** Pranor Trace handles traces. Pranor Console displays metrics. But there's no dedicated metrics storage for long-term retention, custom dashboards, and alerting queries.

| Aspect | Detail |
|--------|--------|
| What it does | TSDB for counters, gauges, histograms. PromQL-compatible query engine |
| Ingest | Prometheus scrape + OTLP metrics push |
| Storage | Compressed columnar storage with configurable retention |
| Features | Recording rules, downsampling, multi-tenant isolation |
| Integrates with | Pranor Console (chart data source), Pranor Trace (metrics from spans) |
| Standalone value | High — lightweight Prometheus/VictoriaMetrics alternative |
| Effort | Large (6-8 weeks) |

---

## Tier 2: Differentiation & Modern Use Cases

### ServEvents — CloudEvents-Compatible Event Bus
**Problem:** Pranor Pulse is a message broker (pub/sub). But modern event-driven architectures need a higher-level event bus with schemas, event sourcing, replay, and CloudEvents compatibility.

| Aspect | Detail |
|--------|--------|
| What it does | Event bus with CloudEvents spec, event sourcing, projections, replay from any point |
| Difference from Pranor Pulse | Pranor Pulse = messaging transport. ServEvents = event semantics layer on top |
| Features | Event schemas with evolution, projections (materialize views from events), event replay, dead letter |
| Integrates with | Pranor Pulse (transport layer), Pranor Vault (event archive), Pranor Flow (event triggers) |
| Standalone value | Medium — needs Pranor Pulse underneath |
| Effort | Large (4-6 weeks) |

---

### ServNotify — Push Notifications & Webhook Delivery
**Problem:** Pranor Notify handles email/Slack/SMS. But mobile push (FCM/APNs), browser push (Web Push API), and outbound webhook delivery with retry/signing are missing.

| Aspect | Detail |
|--------|--------|
| What it does | Fan-out notification delivery: mobile push, browser push, webhook relay |
| Channels | Firebase Cloud Messaging, Apple Push, Web Push API, outbound webhooks |
| Features | Delivery tracking, retry with backoff, webhook signature (HMAC), payload templates |
| Integrates with | Pranor Notify (unified notification orchestrator), Pranor Pulse (async delivery) |
| Standalone value | Medium-High — webhook delivery is universally needed |
| Effort | Medium (3-4 weeks) |

---

### ServJobs — Background Job Queue
**Problem:** Pranor Chrono handles periodic tasks. Pranor Pulse handles async messaging. But there's no dedicated background job queue with worker pools, priorities, progress tracking, and retry (like Sidekiq/BullMQ/Celery).

| Aspect | Detail |
|--------|--------|
| What it does | Background job processing: enqueue, prioritize, retry, track progress |
| Difference from Pranor Chrono | Pranor Chrono = scheduled (time-based). ServJobs = on-demand (event-triggered) |
| Difference from Pranor Pulse | Pranor Pulse = message delivery. ServJobs = task execution with state |
| Features | Priority queues, progress tracking, rate limiting, job chaining, dead jobs |
| Language keyword | `enqueue "process_image" { file: uploaded.key, size: "thumbnail" }` |
| Integrates with | Pranor Pulse (transport), Pranor Console (job dashboard), Pranor Trace (job spans) |
| Standalone value | High — every web app needs background processing |
| Effort | Medium (3-4 weeks) |

---

### ServEdge — Edge/WASM Execution Runtime
**Problem:** Pranor Deploy deploys to central infrastructure. For latency-sensitive workloads (personalization, auth checks, geo-routing), code needs to run at the edge.

| Aspect | Detail |
|--------|--------|
| What it does | Deploy Serv functions as WASM to edge locations. Geo-routed, low-latency |
| Execution | Wazero runtime at edge PoPs. Cold start <5ms |
| Features | Geo-routing, KV store at edge, offline sync, auto-failover to origin |
| Language keyword | `deploy "edge" { regions: ["us-east", "eu-west", "ap-south"] }` |
| Integrates with | Pranor Gate (edge routing), Pranor Cache (edge KV), Pranor Vault (origin data) |
| Standalone value | Medium — needs hosting infrastructure |
| Effort | Very Large (8-12 weeks) |

---

### ServAI — Dedicated AI/ML Gateway
**Problem:** Pranor Gate has cost-aware LLM routing and prompt guard. But a dedicated AI service could handle model registry, prompt versioning, RAG orchestration, fine-tuning jobs, and token budget management as a product.

| Aspect | Detail |
|--------|--------|
| What it does | AI model gateway: routing, versioning, evaluation, budget, RAG pipeline management |
| Difference from Pranor Gate AI | Pranor Gate = proxy-level (route requests). ServAI = platform-level (manage models) |
| Features | Model registry, prompt A/B testing, eval pipelines, fine-tune job queue, embedding cache |
| Integrates with | Pranor Gate (request routing), Pranor Vault (model artifacts), Pranor Pulse (fine-tune jobs) |
| Standalone value | High — dedicated AI gateway is a product category |
| Effort | Large (6-8 weeks) |

---

## Tier 3: Specialized / Niche

### ServForms — Form Builder & Submission Store
| What | Backend for form submissions — schema validation, file uploads, webhooks on submit |
| Use case | Contact forms, surveys, waitlists without building a custom backend |
| Effort | Small (1-2 weeks) |

### ServFiles — CDN & Image Processing
| What | File serving with on-the-fly transforms: resize, format conversion, watermark, thumbnails |
| Difference from Pranor Vault | Pranor Vault = raw object storage. ServFiles = delivery-optimized with transforms |
| Effort | Medium (3-4 weeks) |

### ServScheduler — Visual DAG Scheduler (Airflow-like)
| What | Complex multi-step job pipelines with visual editor, dependencies, retries |
| Difference from Pranor Chrono | Pranor Chrono = single jobs on schedule. ServScheduler = multi-step DAGs with data passing |
| Note | Overlaps significantly with Pranor Flow. May be better as a Pranor Flow extension |
| Effort | Large (6+ weeks) |

### ServProxy — TCP/Database Tunnel
| What | TCP proxy for tunneling database connections, Redis, gRPC through firewalls |
| Difference from Pranor Tunnel | Pranor Tunnel = HTTP only. ServProxy = raw TCP (postgres, redis, mysql protocols) |
| Effort | Medium (2-3 weeks) |

### ServAnalytics — Product Analytics
| What | Self-hosted product analytics: page views, events, funnels, cohorts, retention |
| Similar to | Plausible, PostHog, Amplitude (but self-hosted, single binary) |
| Effort | Very Large (8+ weeks) |

### ServBilling — Subscription & Usage Billing
| What | Metered billing: track usage per tenant, generate invoices, integrate with Stripe |
| Use case | SaaS products built on Pranor that need usage-based pricing |
| Effort | Large (6+ weeks) |

### ServIdentity — Decentralized Identity (DID/Verifiable Credentials)
| What | W3C DID and Verifiable Credential issuance/verification |
| Difference from Pranor Auth | Pranor Auth = traditional OAuth/OIDC. ServIdentity = decentralized/blockchain-anchored |
| Effort | Large (6+ weeks) |

---

## Prioritization Framework

When deciding which to build next, score each on:

| Criterion | Weight | Question |
|-----------|--------|----------|
| **Daily pain** | 30% | Do developers hit this gap every day? |
| **Standalone value** | 25% | Can someone use this without the rest of Pranor? |
| **Ecosystem leverage** | 20% | Does it make existing components more valuable? |
| **Differentiation** | 15% | Does it create something competitors don't have? |
| **Effort** | 10% | Can it ship in <3 weeks? |

### Recommended Build Order

| Priority | Component | Score | Rationale |
|----------|-----------|-------|-----------|
| 1 | **Pranor Secret** | 92 | Every service needs it, standalone usable, small effort |
| 2 | **Pranor Lock** | 88 | Universal need, small effort, enables production patterns |
| 3 | **ServSearch** | 82 | High standalone value, many adapters, differentiating |
| 4 | **ServJobs** | 78 | Every web app needs background jobs, clear gap |
| 5 | **ServNotify** | 72 | Webhook delivery is universal, complements Pranor Notify |
| 6 | **ServConfig** | 68 | Makes ecosystem better but less standalone value |
| 7 | **ServAI** | 65 | Differentiating but large effort |
| 8 | **ServMetrics** | 60 | Valuable but large effort, competing with Prometheus |
| 9 | **ServEdge** | 55 | Differentiating but very large effort + needs infra |
| 10 | **ServEvents** | 50 | Architecturally elegant but needs Pranor Pulse underneath |

---

*This document is reviewed quarterly. Items move to `UNIFIED_ROADMAP.md` when a build decision is made.*
