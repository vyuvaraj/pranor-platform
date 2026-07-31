# I Built a Programming Language for Backend Services — Here's Why

*What if building a microservice was as simple as describing what it should do?*

---

## The Problem

Every backend developer knows the ritual. You want a simple scheduled job that checks database health every 5 seconds. In Go, that's a `main.go`, a `time.Ticker`, a goroutine, signal handling, graceful shutdown, a `go.mod`, and about 50 lines before you've written any business logic.

Want to add an HTTP endpoint? Import `net/http`, set up a router, parse JSON manually, wire up middleware. A pub/sub consumer? Pick a client library, handle connection lifecycle, manage acknowledgments.

The infrastructure ceremony drowns out the intent.

I kept writing the same boilerplate across dozens of services. So I built a language where the infrastructure *is* the syntax.

---

## Meet Serv

Serv is a domain-specific language that compiles to native binaries via Go code generation. It's designed for one thing: making backend services trivial to write.

Here's a complete scheduled job:

```serv
every 5s {
    log.info("Running healthcheck...")
    log.info("Database connection is healthy.")
}

cron env("BACKUP_CRON") {
    log.warn("Starting database backup...")
    metric.inc("db_backups_total")
}
```

That's it. No main function. No imports. No boilerplate. Run `pranor build scheduler.pnr` and you get a native binary.

---

## A REST API in 15 Lines

```serv
server env("PORT")

route "GET" "/health" (req) {
    return { "status": "ok", "time": time.now() }
}

route "POST" "/orders" (req) {
    let body = req.body
    log.info("New order: ", body)
    let response = http.post("https://httpbin.org/post", body)
    metric.inc("orders_processed_total")
    return { "message": "Order processed", "upstream_status": response.status }
}
```

No framework to configure. No router to initialize. The `server` keyword starts an HTTP server. The `route` keyword defines endpoints. Return a map and it becomes JSON.

---
## Event-Driven by Default

Pub/sub and concurrency aren't afterthoughts — they're language primitives:

```serv
broker "nats://localhost:4222"

every 4s {
    publish "events.logs" "System status: OK"
}

subscribe "events.logs" (message) {
    log.info("Received: ", message)
    spawn processLogData(message)
}

fn processLogData(logLine) {
    log.info("[Worker] Indexing: ", logLine)
}
```

`broker` declares your messaging backend. `subscribe` and `publish` are keywords. `spawn` fires off a goroutine. Swap `"nats://..."` for `"kafka://..."` or `"rabbitmq://..."` and the code stays the same.

---

## Middleware and Auth — No Framework Required

```serv
server "8080"

middleware auth(req) {
    let token = req.params.authorization
    if token == nil {
        return { "status": 401, "error": "Unauthorized" }
    }
}

route "GET" "/admin" (req) use [auth] {
    return { "data": "admin panel" }
}
```

Middleware is a first-class concept. Attach it to routes with `use [...]`. If middleware returns a response, the request short-circuits. Simple.

---

## WebSockets in 8 Lines

```serv
server "8080"

ws "/chat" (conn) {
    for true {
        let msg = conn.receive()
        if msg == nil { break }
        conn.send(f"Echo: {msg}")
    }
}
```

---
## What Makes Serv Different

| Feature | Serv | Go | Node.js | Python |
|---------|------|-----|---------|--------|
| HTTP server | `server "8080"` | 15+ lines | Express setup | Flask/FastAPI setup |
| Scheduled job | `every 5s { }` | goroutine + ticker | node-cron package | APScheduler/Celery |
| Pub/Sub | `subscribe "topic" (msg) { }` | Client library + boilerplate | Library + callbacks | Library + boilerplate |
| Concurrency | `spawn task()` | `go task()` | Worker threads / async | asyncio / threading |
| Rate limiting | `limit 100/minute` on a route | Middleware library | express-rate-limit | Custom decorator |
| WebSocket | `ws "/path" (conn) { }` | gorilla/websocket setup | ws/socket.io | websockets library |

The difference isn't that Serv can do things others can't — it's that infrastructure concerns become *declarations* instead of *implementation details*.

---

## Under the Hood

Serv compiles through a classic pipeline:

```
.pnr source → Lexer → Parser (Pratt) → AST → Code Generator → Go source → Native binary
```

The generated Go code uses a runtime library that provides HTTP routing, pub/sub adapters, cron scheduling, database access, caching, metrics, and structured logging. You write 15 lines of Serv; the compiler generates the 200 lines of Go you'd have written by hand.

The compiler itself is written in Go (~21 source files across lexer, parser, AST, codegen, and semantic analysis). It includes:

- **Type checking** with optional types (`string?`), union types (`int | error`), and generics with constraints
- **Unused variable detection** and dead code warnings
- **Escape analysis** for performance optimization
- **Flow analysis** for unreachable code detection

---

## Real Features, Not a Toy

Serv isn't a weekend project. It has:

- **40+ working examples** covering REST APIs, schedulers, pub/sub, WebSockets, database migrations, generics, MCP tools, and more
- **Multiple database backends**: SQLite, PostgreSQL, MongoDB, Oracle
- **Multiple broker backends**: Kafka, NATS, RabbitMQ, MQTT, Redis Streams
- **A test framework**: `test "name" { assert expr }` blocks that generate Go tests
- **A formatter**: `serv fmt` with consistent 4-space indentation
- **A linter**: `serv lint` for static analysis without building
- **A REPL**: `serv repl` for interactive exploration
- **Python interop**: `extern fn` bindings to call Python scripts
- **Go package FFI**: Import and use any Go package via declaration files
- **A standard library**: Auth, JWT, validation, pagination, crypto modules written in Serv itself
- **Docker support**: `serv dockerize` generates production Dockerfiles
- **VS Code extension**: Syntax highlighting and language support

---
## Who Is This For?

If you're building:
- Background workers and scheduled jobs
- REST API microservices
- Event-driven processors (Kafka/NATS/RabbitMQ consumers)
- Internal tools and health-check services
- Rapid prototypes that need to become production services

...and you're tired of writing the same Go/Java/Python scaffolding every time, Serv might be what you're looking for.

---

## Try It

```bash
# Clone and build the compiler
git clone https://github.com/vyuvaraj/Pranor.git
cd Pranor
go build -o pranor.exe .

# Run an example
pranor run examples/02_rest_api.pnr

# Or build a native binary
pranor build examples/02_rest_api.pnr -o my-api.exe
./my-api.exe
```

---

## What's Next

Serv is open-source and actively developed. The roadmap includes:

- **serv-playground**: A web sandbox to try Serv without installing anything
- **Package registry**: `serv install` for community modules
- **More adapter backends**: MySQL, AWS SQS, GCP Pub/Sub, Meilisearch
- **Enhanced LSP**: Full autocomplete and go-to-definition
- **Deploy targets**: `pranor deploy --target fly` / `--target railway` / `--target k8s`

The goal isn't to replace Go — it's to make Go the *compilation target* while you focus on what your service actually does.

---

## Links

- **GitHub**: [github.com/vyuvaraj/Pranor](https://github.com/vyuvaraj/Pranor)
- **Language Reference**: Full docs in the repo under `docs/`
- **License**: Open source

---

*If you're building backend services and want to cut through the boilerplate, give Serv a try. Star the repo, open an issue, or just run an example and tell me what breaks.*

*— Yuvaraj*
