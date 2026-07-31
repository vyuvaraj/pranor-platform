# Serv-Lang Language Guide

> **Pranor** is a high-level, domain-specific language for building production-grade microservices. It compiles to native Go binaries with zero configuration.

---

## Table of Contents

1. [Project Setup](#1-project-setup)
2. [Unified Application Block](#unified-application-block)
3. [Server & Routes](#2-server--routes)
4. [Request Binding](#3-request-binding)
5. [HTML & Web Responses](#4-html--web-responses)
6. [Database](#5-database)
7. [Variables & Data Types](#6-variables--data-types)
8. [Type System](#7-type-system)
9. [Functions](#8-functions)
10. [Control Flow](#9-control-flow)
11. [Structs & Methods](#10-structs--methods)
12. [Interfaces](#11-interfaces)
13. [Enums](#12-enums)
14. [Generics](#13-generics)
15. [Imports & Modules](#14-imports--modules)
16. [Authentication](#15-authentication)
17. [Pub/Sub Messaging](#16-pubsub-messaging)
18. [Scheduling & Cron](#17-scheduling--cron)
19. [Object Store (S3)](#18-object-store-s3)
20. [WebSockets](#19-websockets)
21. [Middleware](#20-middleware)
22. [AI Integration](#21-ai-integration)
23. [MCP Tools & Agents](#22-mcp-tools--agents)
24. [Schema Migrations (table DSL)](#23-schema-migrations-table-dsl)
25. [Error Handling](#24-error-handling)
26. [Concurrency](#25-concurrency)
27. [Testing](#26-testing)
28. [External Functions (FFI)](#27-external-functions-ffi)
29. [Stream DSL WASM Transforms](#stream-dsl-wasm-transforms)
30. [Logic Configuration Policy Engine](#logic-configuration-policy-engine)
31. [Observability (OTel)](#28-observability-otel)
32. [Environment & Config](#29-environment--config)
33. [CLI Reference](#30-cli-reference)

---

## 1. Project Setup

### Single-file project
```python
# main.pnr
server "9000"

export route "GET" "/" (req) {
    return { "message": "Hello, world!" }
}
```

## Unified Application Block

You can wrap configuration nodes and routes inside an `app` block to create a clean logical boundary:

```python
app GatewayService {
    server "9000"
    database "sqlite://app.db"

    export route "GET" "/health" (req) {
        return { "status": "UP" }
    }
}
```

### Multi-file project (serv.toml)
```toml
# serv.toml
entry = "main.pnr"
name  = "my-service"
```

### Build & run
```bash
pranor build main.pnr          # compile to native binary
pranor run main.pnr            # compile + run
pranor run main.pnr --watch    # hot-reload on file changes
pranor run main.pnr --port 8080
```

---

## 2. Server & Routes

### Declare server port
```python
server "9000"
```

### Route declaration
```python
export route "METHOD" "/path" (req) {
    return { "key": "value" }
}
```

**Supported methods:** `GET`, `POST`, `PUT`, `PATCH`, `DELETE`, `HEAD`, `OPTIONS`

### Path parameters
```python
export route "GET" "/users/:id" (req) {
    let id = req.params["id"]
    return { "id": id }
}
```

### Query parameters
```python
export route "GET" "/search" (req) {
    let q = req.query["q"]
    return { "query": q }
}
```

### Rate limiting
```python
export route "POST" "/api/login" (req) @rate(5, "m") {
    # max 5 requests per minute per route
}
```

### CORS
```python
cors ["https://app.example.com", "https://admin.example.com"]
```

### Global IP rate limiting
```python
rate_limit 100 "m"   # 100 req/min per IP globally
```

---

## 3. Request Binding

### JSON body parsing
```python
export route "POST" "/api/users" (req) {
    let data = req.json()      # parse req.body as JSON
    let name  = data.name
    let email = data.email
}
```

### Form body parsing (application/x-www-form-urlencoded)
```python
export route "POST" "/contact" (req) {
    let form    = req.form()
    let message = form.message
}
```

### Safe param lookup (returns nil if missing)
```python
export route "GET" "/items/:id" (req) {
    let id = req.param("id")
    if id == nil {
        return { "error": "id required", "status": 400 }
    }
}
```

### Object destructuring
```python
let data = req.json()
let { name, email, age } = data
```

### Object shorthand (DX.S14)
```python
let name  = "Alice"
let email = "alice@example.com"
return { name, email }   # same as { name: name, email: email }
```

---

## 4. HTML & Web Responses

### Inline template
```python
export route "GET" "/" (req) {
    let tpl = `<!DOCTYPE html>
<html>
<head><title>{{.title}}</title></head>
<body><h1>Hello, {{.name}}!</h1></body>
</html>`
    return html.template(tpl, { "title": "Home", "name": "World" })
}
```

### File template
```python
export route "GET" "/" (req) {
    return html.render("views/index.html", { "user": user })
}
```

### Static file server
```python
html.static("/assets", "./public")    # serves ./public at /assets/
```

### Redirect
```python
export route "GET" "/old" (req) {
    return html.redirect("/new", 301)    # permanent
}

export route "GET" "/login-required" (req) {
    return html.redirect("/login", 302)  # temporary
}
```

### Implicit Content-Type inference (DX.S15)
When a route returns a plain string, Content-Type is automatically set:

| Return string starts with | Content-Type set |
|---|---|
| `<html`, `<!DOCTYPE` | `text/html; charset=utf-8` |
| `<?xml`, `<rss`, `<feed` | `application/xml; charset=utf-8` |
| `{...}` or `[...]` | `application/json` |
| anything else | `text/plain; charset=utf-8` |

---

## 5. Database

### Declare database
```python
database "sqlite://./app.db"
database "postgres://user:pass@localhost/mydb"
```

### Query
```python
let users = db.query("SELECT * FROM users WHERE active = ?", [true])
```

### Execute (insert/update/delete)
```python
db.exec("INSERT INTO users (name, email) VALUES (?, ?)", [name, email])
```

### Schema migrations
```python
migration "create_users_table" {
    db.exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT UNIQUE,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    )`)
}
```

---

## 6. Variables & Data Types

```python
let name  = "Alice"          # string
let age   = 30               # integer
let score = 9.5              # float
let active = true            # boolean
let items = [1, 2, 3]        # array
let user  = { "name": "Alice", "age": 30 }  # map/object

# Shorthand
let user  = { name, age }    # { name: name, age: age }

# Destructuring
let { name, age } = user

# String interpolation
let msg = f"Hello, {name}! You are {age} years old."

# Multi-line string
let html = `
<h1>Hello</h1>
<p>World</p>
`
```

---

## 7. Type System

### Type Annotations

```python
let name: string = "Alice"
let age: int = 30
let score: float = 9.5
let active: bool = true
```

### Type Aliases

```python
type UserID = int
type Email = string
type Handler = fn(Request) -> Response
```

### Null Safety (Optional Types)

Types suffixed with `?` allow nil values. Without `?`, assigning nil is a compile error.

```python
let name: string = "Alice"     # Cannot be nil
let email: string? = nil       # OK — optional type

fn findUser(id: int) -> User? {
    let row = db.query("SELECT * FROM users WHERE id = ?", id)
    if row == nil { return nil }
    return User { name: row.name }
}
```

### Union Types

```python
fn divide(a: int, b: int) -> int | error {
    if b == 0 { return "division by zero" }
    return a / b
}
```

### Optional Chaining

```python
let city = user?.address?.city    # nil if any part is nil
```

### Spread Operator

```python
let defaults = { "timeout": 30, "retries": 3 }
let config = { ...defaults, "timeout": 60 }
```

### Slice Expressions

```python
let items = [1, 2, 3, 4, 5]
let first3 = items[0:3]     # [1, 2, 3]
let rest = items[2:]         # [3, 4, 5]
let head = items[:2]         # [1, 2]
```

---

## 8. Functions

```python
fn greet(name) {
    return f"Hello, {name}!"
}

# Typed parameters and return
fn add(a: int, b: int) -> int {
    return a + b
}

export fn multiply(a, b) {      # exported = usable across files
    return a * b
}

# Anonymous function
let double = fn(x) { return x * 2 }

# Arrow functions (closures)
let triple = x => x * 3
let sum = (a, b) => a + b

# Higher-order functions
fn apply(val, transform) { return transform(val) }
let result = apply(10, x => x * x)
```

### Collection Methods (Arrow Functions)

```python
let users = [{ "name": "Alice", "active": true }, { "name": "Bob", "active": false }]
let active = users.filter(u => u.active).map(u => u.name)
# ["Alice"]

let items = [1, 2, 3, 4, 5]
items.filter(x => x > 2)          # [3, 4, 5]
items.map(x => x * 2)             # [2, 4, 6, 8, 10]
items.find(x => x == 3)           # 3
items.reduce(fn(a, b) { return a + b }, 0)  # 15
items.forEach(x => log.info(x))
items.contains(3)                  # true
```

### String Methods

```python
"hello world".split(" ")      # ["hello", "world"]
"  hi  ".trim()               # "hi"
"hello".replace("l", "L")     # "heLLo"
"hello".startsWith("he")      # true
"hello".includes("ell")       # true
"hello".toUpper()             # "HELLO"
"HELLO".toLower()             # "hello"
"hello".substring(1, 3)       # "el"
"ha".repeat(3)                # "hahaha"
```

---

## 9. Control Flow

### If / else
```python
if age >= 18 {
    return { "access": true }
} else {
    return { "access": false }
}
```

### For loop
```python
for i = 0; i < 10; i++ {
    log(i)
}
```

### For-in
```python
for item in items {
    log(item)
}

# Map iteration
for key, value in config {
    log.info(f"{key} = {value}")
}
```

### Break & Continue

```python
for item in items {
    if item == nil { continue }
    if item == "stop" { break }
    process(item)
}
```

### Match (Pattern Matching)
```python
match status {
    "active"   -> { return { "ok": true } }
    "inactive" -> { return { "ok": false } }
    _          -> { return { "error": "unknown" } }
}
```

---

## 10. Structs & Methods

```python
struct User {
    id: int
    name: string
    email: string
    active: bool
}

# Methods
fn User.fullName() -> string {
    return f"{self.name} ({self.email})"
}

fn User.greet() -> string {
    return f"Hi, I'm {self.name}"
}

# Instantiation
let u = User { id: 1, name: "Alice", email: "alice@test.com", active: true }
log.info(u.fullName())
```

---

## 11. Interfaces

Structural typing — if a struct has the methods, it satisfies the interface.

```python
interface Serializable {
    fn serialize() -> string
    fn deserialize(data: string)
}

# User satisfies Serializable if it has serialize() and deserialize()
fn User.serialize() -> string {
    return json.stringify(self)
}

fn User.deserialize(data: string) {
    let parsed = json.parse(data)
    self.name = parsed.name
}
```

---

## 12. Enums

```python
# Simple enum
enum Color { Red, Green, Blue }

# With explicit values
enum HttpStatus {
    OK = 200,
    NotFound = 404,
    ServerError = 500
}

# Usage
let status = HttpStatus.OK
match status {
    HttpStatus.OK -> { return { "success": true } }
    HttpStatus.NotFound -> { return { "error": "not found" } }
}
```

---

## 13. Generics

```python
# Generic function
fn filter[T](items: []T, pred: fn(T) -> bool) -> []T {
    let result: []T = []
    for item in items {
        if pred(item) { result.push(item) }
    }
    return result
}

fn map[T, U](items: []T, transform: fn(T) -> U) -> []U {
    let result: []U = []
    for item in items {
        result.push(transform(item))
    }
    return result
}

# Generic with constraints
fn max[T: Ordered](a: T, b: T) -> T {
    if a > b { return a }
    return b
}
```

### Constraints

| Constraint | Supports |
|-----------|----------|
| `Comparable` | `==`, `!=` |
| `Ordered` | `<`, `>`, `<=`, `>=` |
| `Numeric` | `+`, `-`, `*`, `/` |

---

## 14. Imports & Modules

```python
import "./handlers/users"          # imports users.pnr
import "./utils/validation.pnr"
import { validateEmail } from "./auth/utils"  # named import

# Wildcard directory import
import "./handlers/*"              # imports all .pnr in ./handlers/

# Stdlib
import "stdlib/auth"
import "stdlib/pagination"
import { ok, notFound } from "stdlib/response"

# Go package (requires .pnr.d declaration)
import uuid from "github.com/google/uuid"
let id = uuid.New()
```

---

## 15. Authentication

```python
auth "my-jwt-secret"               # enable JWT auth

# Register & login routes
export route "POST" "/auth/register" (req) {
    let data = req.json()
    return auth.register(data.username, data.password, data.email)
}

export route "POST" "/auth/login" (req) {
    let data = req.json()
    return auth.login(data.username, data.password)
}

# Protected route (JWT middleware auto-applied)
export route "GET" "/api/profile" (req) {
    let user = auth.currentUser(req)
    return { "user": user }
}

# Role-based access control
export route "DELETE" "/admin/users/:id" (req) @middleware("auth.role(\"admin\")") {
    # admin only
}
```

---

## 16. Pub/Sub Messaging

```python
broker "pranor-pulse://localhost:4222"
# or in-memory: broker "memory://"

publish "user.created" { "id": userId, "email": email }

subscribe "user.created" (event) {
    log(f"New user: {event.email}")
}
```

---

## 17. Scheduling & Cron

```python
every "5m" {
    # runs every 5 minutes
    let result = db.query("SELECT COUNT(*) as cnt FROM users")
    log(f"Total users: {result[0].cnt}")
}

cron "0 9 * * MON-FRI" {
    # 9:00 AM weekdays
    publish "reports.daily" { "type": "morning" }
}
```

---

## 18. Object Store (S3)

```python
store "s3://access:secret@localhost:9000/my-bucket"
# or: store "file://./data"

store.put("profile/alice.json", { "name": "Alice" })
let profile = store.get("profile/alice.json")
store.delete("profile/alice.json")
store.list("profile/")
```

---

## 19. WebSockets

```python
ws "/chat" (conn) {
    conn.send({ "msg": "Welcome!" })
    let msg = conn.receive()
    while msg != nil {
        conn.broadcast(msg)
        msg = conn.receive()
    }
}
```

---

## 20. Middleware

```python
middleware authMiddleware(req) {
    let token = req.headers["authorization"]
    if token == nil {
        return { "error": "Unauthorized", "status": 401 }
    }
    # return nil = pass through to handler
}

export route "GET" "/api/data" (req) use [authMiddleware] {
    return { "data": "secret" }
}

# Multiple middleware
export route "POST" "/admin" (req) use [authMiddleware, logging, rateLimit] {
    return { "admin": true }
}
```

---

## 21. AI Integration

```python
ai "openai://gpt-4o"              # OpenAI
# ai "anthropic://claude-3-5-sonnet"  # Anthropic
# ai "ollama://llama3"               # Local

# Text completion
let response = ai.complete("Summarize this article: " + text)

# Chat with message history
let reply = ai.chat([
    { "role": "system", "content": "You are a helpful assistant." },
    { "role": "user", "content": "What is Serv?" }
])

# Generate embeddings
let vector = ai.embed("distributed systems architecture")
```

---

## 22. MCP Tools & Agents

### Tool Declarations

```python
tool "calculator" "Performs math operations" (args) {
    let result = args.a + args.b
    return { "result": result }
}

tool "lookup_order" "Look up an order by ID" (args) {
    let row = db.query("SELECT * FROM orders WHERE id = ?", args.order_id)
    return row
}
```

### Agent Declarations

```python
agent SupportBot {
    system "You are a helpful customer support assistant."
    model  "openai://gpt-4o"
    tools  ["lookup_order", "calculator"]
}
```

**Supported model URI schemes:**
- `openai://gpt-4o` — OpenAI
- `anthropic://claude-3-5-sonnet` — Anthropic
- `google://gemini-2.0-flash` — Google Gemini
- `local://ollama/llama3` — Local Ollama

---

## 23. Schema Migrations (table DSL)

Declare database schema natively. The compiler generates SQL; `serv migrate` applies it.

```python
table users {
    id        int      @primary @autoincrement
    name      string   @required
    email     string   @unique
    role      string   @default(user)
    createdAt datetime @default(now)
}

table posts {
    id        int      @primary @autoincrement
    userId    int      @required
    title     string   @required
    body      string
    published bool     @default(0)
}
```

### Annotations

| Annotation | SQL equivalent |
|-----------|----------------|
| `@primary` | PRIMARY KEY |
| `@autoincrement` | AUTOINCREMENT |
| `@required` | NOT NULL |
| `@unique` | UNIQUE |
| `@default(value)` | DEFAULT value |

### Apply migrations

```bash
serv migrate                    # Apply to default db
serv migrate --db postgres://user:pass@host/db
```

### Raw migrations (advanced)

```python
migration "add_index" {
    db.exec("CREATE INDEX idx_users_email ON users (email)")
}
```

---

## 24. Error Handling

```python
# Try/catch
try {
    let data = db.query("SELECT * FROM users")
    return data
} catch (err) {
    return { "error": err, "status": 500 }
}

# Multi-return
let data, err = riskyCall()
if err != nil {
    log.error(err)
}

# ? operator — early return on nil/error
fn fetchUser(id) {
    let user = db.query("SELECT * FROM users WHERE id = ?", [id])?
    return user
}
```

The `?` operator: if the expression returns nil or error, returns nil from the enclosing function. Otherwise unwraps and continues.

---

## 25. Concurrency

```python
# Async (blocks until done)
let result = await fn() {
    return db.query("SELECT * FROM expensive_table")
}

# Parallel fan-out
let [users, orders] = await_all([
    fn() { return db.query("SELECT * FROM users") },
    fn() { return db.query("SELECT * FROM orders") }
])

# Fire-and-forget (inherits trace context)
spawn fn() {
    publish "notifications.send" { "to": email }
}

# Worker pool limit
spawn(5) heavyTask(data)
```

---

## 26. Testing

```python
test "math works" {
    let result = add(2, 3)
    assert result == 5
}

test "string methods" {
    assert "hello".toUpper() == "HELLO"
    assert "  hi  ".trim() == "hi"
}

test "database integration" {
    db.exec("INSERT INTO users (name) VALUES (?)", ["Test"])
    let rows = db.query("SELECT * FROM users WHERE name = ?", ["Test"])
    assert rows.length() > 0
}
```

Run with: `pranor test <file.pnr> [--cover] [--filter name]`

---

## 27. External Functions (FFI)

```python
# Go package binding
extern fn generateID() from "go:github.com/google/uuid:NewString"

# Python script binding
extern fn analyze(data) from "python:./scripts/analyzer.py:analyze"

# Usage
let id = generateID()
let result = analyze({ "text": "hello world" })
```

### Go Package Declarations (.pnr.d files)

```python
# uuid.pnr.d — generated by `serv add github.com/google/uuid`
declare module "github.com/google/uuid" {
    fn New() -> string
    fn NewString() -> string
}
```

Generate with: `serv add <go-package-path>`

---

## 28. Observability (OTel)

```python
otel "my-service-name"    # enable OpenTelemetry
```

Pranor automatically traces:
- Every HTTP request (with `traceparent` propagation)
- DB queries, cache ops, HTTP client calls, pub/sub, scheduler jobs

Built-in endpoints:
- `GET /metrics` — Prometheus metrics
- `GET /health` — liveness probe
- `GET /ready` — readiness probe

Environment: `PRANOR_OTLP_ENDPOINT=http://localhost:4318` to set collector.

---

## 29. Environment & Config

```python
let port = env("PORT")
let secret = env.secret("JWT_SECRET")  # masked in logs

# Config validation (fail-fast on startup)
validate {
    required "db.host",
    required "db.port",
    optional "log.level"
}
```

---

## Stream DSL WASM Transforms

Pranor provides native stream processing primitives to declare inline WASM message transforms in under 5 lines of code:

```python
# Declare a message transformation for a topic
transform "orders.raw" (msg) {
    let clean = msg
    # Return value is automatically re-routed or published
    return clean
}
```

---

## Logic Configuration Policy Engine

You can use Pranor as a high-performance configuration and routing policy engine:

```python
# Declare a policy routing rule evaluated at gateway speed
policy rate_limit_policy (ctx) {
    let path = ctx["path"]
    if path == "/api/admin" {
        return false
    }
    return true
}
```

---

## 33. CLI Reference

| Command | Description |
|---|---|
| `pranor build <file>` | Compile to native binary |
| `pranor run <file>` | Compile and run |
| `pranor run <file> --watch` | Run with hot-reload |
| `pranor dev <file>` | Hot-reload dev server with tests |
| `pranor test <file>` | Run .pnr tests |
| `pranor test --cover <file>` | Run tests with coverage |
| `serv fmt <file>` | Format source file |
| `pranor lint <file>` | Lint and static analysis |
| `serv migrate` | Apply table DSL migrations |
| `serv create "<prompt>"` | AI-powered scaffolding |
| `serv add <go-package>` | Generate .pnr.d declaration |
| `serv packages` | List installed declarations |
| `serv doctor` | Ecosystem health check |
| `pranor deploy --target <t>` | Deploy (fly/railway/render/docker) |
| `serv dockerize <file>` | Generate Dockerfile |
| `serv doc <file>` | Generate API docs |
| `serv repl` | Interactive REPL |
| `serv debug <file>` | Debug with Delve |
| `serv audit` | Audit dependencies for CVEs |
| `serv new <name>` | Scaffold new project |
| `pranor build --target wasm` | Compile to WebAssembly |

---

## Operators Reference

### Arithmetic
| Op | Description |
|----|-------------|
| `+` | Addition / string concat |
| `-` | Subtraction |
| `*` | Multiplication |
| `/` | Division |
| `%` | Modulo |

### Compound Assignment
`+=`, `-=`, `*=`, `/=`, `%=`

### Bitwise
| Op | Description |
|----|-------------|
| `&` | AND |
| `\|` | OR |
| `^` | XOR |
| `<<` | Left shift |
| `>>` | Right shift |

### Comparison
`==`, `!=`, `<`, `>`, `<=`, `>=`

### Logical
`and`, `or`, `!`

---

*This guide covers pranor v0.1.x. For changelog, see [RELEASE_NOTES.md](./RELEASE_NOTES.md).*
