# Building a Production API Gateway with Pranor Gate

> **Published:** July 2026 | **Reading Time:** ~12 min | **Tags:** `pranor-gate`, `api-gateway`, `rate-limiting`, `routing`

---

Every serious backend needs an API gateway. It's the single entry point that handles routing, auth enforcement, rate limiting, and request transformation before traffic ever hits your services. **Pranor Gate** is the Pranor component that does exactly this — and it ships as a single Docker container with a file-based config.

---

## Why a Dedicated Gateway?

Without a gateway, each microservice handles:
- Auth token validation (duplicated 10×)
- Rate limiting (inconsistent across services)
- CORS headers (always forgotten somewhere)
- Request logging (siloed per service)

Pranor Gate centralizes all of this. Your upstream services focus on business logic. Pranor Gate handles the rest.

---

## Architecture Overview

```
Client
  │
  ▼
Pranor Gate :8081
  ├── /api/users/*     → UserService :3000
  ├── /api/orders/*    → OrderService :3001
  ├── /api/products/*  → ProductService :3002
  └── /api/payments/*  → PaymentService :3003
```

---

## Step 1: Install Pranor Gate

```bash
docker pull ghcr.io/vyuvaraj/pranor-gate:latest
```

Or download the binary from [releases](https://github.com/vyuvaraj/Pranor Gate/releases/latest).

---

## Step 2: Create the Config File

Create `pranor-gate.yaml`:

```yaml
server:
  port: 8081
  host: 0.0.0.0

# Upstream services
upstreams:
  user-service:
    url: http://user-service:3000
    health: /health
    timeout: 10s
  
  order-service:
    url: http://order-service:3001
    health: /health
    timeout: 15s
  
  payment-service:
    url: http://payment-service:3003
    health: /health
    timeout: 30s    # Longer timeout for payment processing

# Route definitions
routes:
  - path: /api/v1/users
    upstream: user-service
    methods: [GET, POST]
    auth: jwt
    rate_limit:
      requests: 100
      window: 1m

  - path: /api/v1/users/:id
    upstream: user-service
    methods: [GET, PUT, DELETE]
    auth: jwt
    rate_limit:
      requests: 50
      window: 1m

  - path: /api/v1/orders
    upstream: order-service
    methods: [GET, POST]
    auth: jwt
    rate_limit:
      requests: 200
      window: 1m

  - path: /api/v1/payments
    upstream: payment-service
    methods: [POST]
    auth: jwt
    rate_limit:
      requests: 10
      window: 1m      # Strict rate limit on payments

  # Public routes (no auth required)
  - path: /api/v1/products
    upstream: product-service
    methods: [GET]
    rate_limit:
      requests: 1000
      window: 1m

# Auth configuration
auth:
  jwt:
    secret: ${JWT_SECRET}
    # Or delegate to Pranor Auth:
    # provider: pranor-auth
    # endpoint: http://pranor-auth:8086

# Global middleware
middleware:
  cors:
    origins: ["https://yourdomain.com", "https://app.yourdomain.com"]
    methods: [GET, POST, PUT, DELETE, OPTIONS]
    headers: [Authorization, Content-Type]

  logging:
    format: json
    include_body: false     # Don't log request bodies (PII)

  tracing:
    enabled: true
    # endpoint: http://pranor-trace:4317   # Send to Pranor Trace

# Health check endpoint
health:
  path: /health
  include_upstreams: true
```

---

## Step 3: Run Pranor Gate

```bash
docker run -d \
  -p 8081:8081 \
  -v $(pwd)/pranor-gate.yaml:/app/pranor-gate.yaml \
  -e JWT_SECRET=your-secret-here \
  ghcr.io/vyuvaraj/pranor-gate:latest
```

Verify it's running:

```bash
curl http://localhost:8081/health
# {"status":"ok","upstreams":{"user-service":"healthy","order-service":"healthy"}}
```

---

## Step 4: Advanced Patterns

### Load Balancing

Pranor Gate can balance across multiple instances of the same service:

```yaml
upstreams:
  user-service:
    instances:
      - url: http://user-service-1:3000
      - url: http://user-service-2:3000
      - url: http://user-service-3:3000
    strategy: round-robin    # or: least-connections, ip-hash
    health: /health
```

### Circuit Breaking

Automatically remove unhealthy upstreams from rotation:

```yaml
upstreams:
  payment-service:
    url: http://payment-service:3003
    circuit_breaker:
      threshold: 5          # Open after 5 failures
      timeout: 30s          # Retry after 30s
      half_open_requests: 2 # Test with 2 requests before fully closing
```

### Request Transformation

Add, remove, or modify headers before forwarding:

```yaml
routes:
  - path: /api/v1/orders
    upstream: order-service
    transform:
      request:
        add_headers:
          X-Gateway: "Pranor Gate"
          X-Request-ID: "${request_id}"
        remove_headers:
          - Cookie    # Strip cookies before forwarding
```

### IP Allowlisting

Restrict sensitive routes to trusted IPs:

```yaml
routes:
  - path: /api/admin
    upstream: admin-service
    auth: jwt
    ip_allowlist:
      - 10.0.0.0/8
      - 172.16.0.0/12
```

---

## Step 5: Monitor with Pranor Console

If you're running Pranor Console, Pranor Gate will automatically report:
- Request rates per route
- Error rates and latency percentiles
- Upstream health status
- Rate limit hit counts

```bash
docker run -d -p 9000:9000 ghcr.io/vyuvaraj/pranor-console:latest \
  -e PRANOR_GATE_ENDPOINT=http://pranor-gate:8081
```

---

## Production Checklist

- [ ] Set `JWT_SECRET` via environment variable, not hardcoded in config
- [ ] Enable TLS termination (`tls.cert` and `tls.key` in config)
- [ ] Set strict rate limits on payment/auth routes
- [ ] Enable JSON logging for log aggregation
- [ ] Configure `ip_allowlist` for admin routes
- [ ] Set up health check monitoring against `/health`

---

## What's Next?

Now that traffic flows through your gateway, we need to make sure responses are fast. In the next post, we explore **Pranor Cache** and distributed caching strategies for high-traffic APIs.

➡️ [Distributed Caching Made Simple with Pranor Cache](blog.html?post=04-caching-with-pranor-cache)

---

*Issues with Pranor Gate? File a bug at [Pranor Gate/issues](https://github.com/vyuvaraj/Pranor Gate/issues).*
