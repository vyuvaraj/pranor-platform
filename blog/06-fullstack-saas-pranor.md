# Full-Stack SaaS in Under an Hour with Pranor

> **Published:** July 2026 | **Reading Time:** ~20 min | **Tags:** `tutorial`, `saas`, `full-stack`, `docker-compose`

---

This is the capstone tutorial of the Pranor blog series. We'll build a complete SaaS application — a project management tool — using 8 Pranor components, from scratch to a running Docker Compose stack.

**What we're building:** A multi-tenant project management API (think simplified Linear or Jira backend) with:
- User authentication and multi-tenancy
- Projects, tasks, and comments
- Real-time notifications via events
- File attachments
- Scheduled reminders
- Full observability

---

## Services We'll Use

| Service | Role |
|---------|------|
| **Pranor Gate** | API gateway — routing, auth, rate limiting |
| **Pranor Auth** | Authentication — JWT, signup, login, org management |
| **Pranor Cache** | Caching — task lists, user sessions |
| **Pranor Pulse** | Events — task created, comment added, reminders |
| **Pranor Vault** | File storage — task attachments |
| **Pranor Notify** | Transactional email — invites, reminders |
| **Pranor Chrono** | Scheduler — daily digest, overdue task reminders |
| **Pranor Trace** | Observability — distributed traces, latency |

Plus our custom **AppService** written in Pranor.

---

## Project Structure

```
taskflow/
├── app/
│   ├── services/
│   │   ├── task.pnr
│   │   ├── project.pnr
│   │   ├── comment.pnr
│   │   └── notification.pnr
│   ├── models/
│   │   ├── user.pnr
│   │   ├── project.pnr
│   │   └── task.pnr
│   └── pranor.yaml
├── config/
│   ├── pranor-gate.yaml
│   └── pranor-auth.yaml
└── docker-compose.yml
```

---

## Step 1: Define the Data Models

`app/models/task.pnr`:
```pranor
model Project {
  id:          string   @id
  org_id:      string   @required @index
  name:        string   @required
  description: string   @optional
  status:      enum     [active, archived]   @default=active
  created_by:  string   @required
  created_at:  time     @auto
  updated_at:  time     @auto
}

model Task {
  id:          string   @id
  project_id:  string   @required @index
  org_id:      string   @required @index
  title:       string   @required
  description: string   @optional
  status:      enum     [todo, in_progress, done, cancelled]   @default=todo
  priority:    enum     [low, medium, high, critical]          @default=medium
  assignee_id: string   @optional @index
  due_date:    time     @optional
  created_by:  string   @required
  created_at:  time     @auto
  updated_at:  time     @auto
}

model Comment {
  id:         string   @id
  task_id:    string   @required @index
  org_id:     string   @required @index
  author_id:  string   @required
  content:    string   @required
  created_at: time     @auto
}
```

---

## Step 2: Define the Services

`app/services/task.pnr`:
```pranor
import store
import cache
import queue
import store as files

service TaskService {
  base /api/v1/tasks

  # List tasks for a project
  route GET / {
    auth   required
    params project_id: string @required
    cache  ttl=30s  key="tasks:{auth.org_id}:{project_id}"
    return store.list(Task, filter={org_id: auth.org_id, project_id: params.project_id})
  }

  # Get a single task
  route GET /:id {
    auth   required
    cache  ttl=60s  key="task:{id}"
    task = store.get(Task, id, filter={org_id: auth.org_id})
    return task
  }

  # Create a task
  route POST / {
    auth     required
    validate body(Task)
    task = store.create(Task, {
      ...body,
      org_id:     auth.org_id,
      created_by: auth.user_id
    })
    # Invalidate project task list cache
    cache.del(key="tasks:{auth.org_id}:{body.project_id}")
    # Publish event
    queue.publish("tasks.created", {
      task_id:    task.id,
      org_id:     task.org_id,
      project_id: task.project_id,
      created_by: task.created_by,
      assignee_id: task.assignee_id
    })
    return task
  }

  # Update a task
  route PUT /:id {
    auth     required
    validate body(Task) @partial
    old_task = store.get(Task, id, filter={org_id: auth.org_id})
    task = store.update(Task, id, body)
    # Clear task and list cache
    cache.del(key="task:{id}")
    cache.del(key="tasks:{auth.org_id}:{task.project_id}")
    # Emit status change event if status changed
    if body.status && body.status != old_task.status {
      queue.publish("tasks.status_changed", {
        task_id:    id,
        old_status: old_task.status,
        new_status: body.status,
        changed_by: auth.user_id
      })
    }
    return task
  }

  # Upload attachment
  route POST /:id/attachments {
    auth   required
    file   upload @max=25MB
    url = files.upload(file, bucket="task-attachments", path="{id}/{file.name}")
    attachment = store.create(Attachment, {task_id: id, url: url, name: file.name})
    return attachment
  }
}
```

---

## Step 3: Configure Ecosystem Services

`config/pranor-gate.yaml`:
```yaml
server:
  port: 8081

upstreams:
  app-service:
    url: http://app:3000
    health: /health

routes:
  - path: /api/v1
    upstream: app-service
    auth: jwt
    rate_limit:
      requests: 300
      window: 1m

  # Public auth routes (no JWT required)
  - path: /auth
    upstream: pranor-auth
    rate_limit:
      requests: 20
      window: 1m

auth:
  jwt:
    provider: pranor-auth
    endpoint: http://pranor-auth:8086

middleware:
  cors:
    origins: ["https://app.taskflow.io"]
  tracing:
    endpoint: http://pranor-trace:4317
```

`config/pranor-auth.yaml`:
```yaml
server:
  port: 8086

jwt:
  secret: ${JWT_SECRET}
  expiry: 24h
  refresh_expiry: 30d

multi_tenancy:
  enabled: true
  org_claim: "org_id"

providers:
  - type: password
    enabled: true
  - type: google
    client_id: ${GOOGLE_CLIENT_ID}
    client_secret: ${GOOGLE_CLIENT_SECRET}

mail:
  provider: pranor-notify
  endpoint: http://pranor-notify:8088
```

---

## Step 4: Set Up Scheduled Jobs

`app/services/notification.pnr`:
```pranor
import cron
import queue
import store

# Daily digest at 8am UTC
cron.schedule("0 8 * * *", "daily-digest") {
  # Find all tasks due today
  today = time.now().truncate(day)
  due_tasks = store.list(Task, filter={
    due_date: today,
    status:   ["todo", "in_progress"]
  })
  
  # Group by assignee and publish digest events
  by_assignee = due_tasks.group_by(task => task.assignee_id)
  for assignee_id, tasks in by_assignee {
    queue.publish("notifications.daily_digest", {
      assignee_id: assignee_id,
      tasks: tasks
    })
  }
}

# Overdue task check every hour
cron.schedule("0 * * * *", "overdue-check") {
  overdue = store.list(Task, filter={
    due_date: { lt: time.now() },
    status:   ["todo", "in_progress"]
  })
  for task in overdue {
    queue.publish("notifications.task_overdue", {
      task_id:     task.id,
      assignee_id: task.assignee_id,
      days_overdue: time.since(task.due_date).days()
    })
  }
}
```

---

## Step 5: Wire Up Email Notifications

`app/services/notification.pnr` (continued):
```pranor
import mail

# Handle daily digest
consumer "notifications.daily_digest" {
  group "mail-consumers"

  handler(msg) {
    assignee = store.get(User, msg.assignee_id)
    mail.send({
      to:       assignee.email,
      subject:  "Your tasks for today",
      template: "daily-digest",
      data:     { user: assignee, tasks: msg.tasks }
    })
  }
}

# Handle task assignment
consumer "tasks.created" {
  group "mail-consumers"

  handler(msg) {
    if msg.assignee_id {
      assignee = store.get(User, msg.assignee_id)
      assigner = store.get(User, msg.created_by)
      task     = store.get(Task, msg.task_id)
      mail.send({
        to:       assignee.email,
        subject:  "New task assigned: ${task.title}",
        template: "task-assigned",
        data:     { task: task, assigner: assigner }
      })
    }
  }
}
```

---

## Step 6: Docker Compose Stack

`docker-compose.yml`:
```yaml
version: "3.9"

services:
  # Infrastructure
  postgres:
    image: postgres:16-alpine
    environment:
      POSTGRES_DB: taskflow
      POSTGRES_USER: taskflow
      POSTGRES_PASSWORD: ${DB_PASSWORD}
    volumes:
      - pg_data:/var/lib/postgresql/data

  # Pranor components
  pranor-gate:
    image: ghcr.io/vyuvaraj/pranor-gate:latest
    ports: ["80:8081"]
    volumes:
      - ./config/pranor-gate.yaml:/app/pranor-gate.yaml
    environment:
      JWT_SECRET: ${JWT_SECRET}
    depends_on: [pranor-auth, app]

  pranor-auth:
    image: ghcr.io/vyuvaraj/pranor-auth:latest
    volumes:
      - ./config/pranor-auth.yaml:/app/pranor-auth.yaml
    environment:
      JWT_SECRET: ${JWT_SECRET}
      GOOGLE_CLIENT_ID: ${GOOGLE_CLIENT_ID}
      GOOGLE_CLIENT_SECRET: ${GOOGLE_CLIENT_SECRET}
    depends_on: [postgres]

  pranor-cache:
    image: ghcr.io/vyuvaraj/pranor-cache:latest
    volumes:
      - cache_data:/data

  pranor-pulse:
    image: ghcr.io/vyuvaraj/pranor-pulse:latest
    volumes:
      - queue_data:/data

  pranor-vault:
    image: ghcr.io/vyuvaraj/pranor-vault:latest
    environment:
      STORAGE_DRIVER: local
      STORAGE_PATH: /data
    volumes:
      - store_data:/data

  pranor-notify:
    image: ghcr.io/vyuvaraj/pranor-notify:latest
    environment:
      SMTP_HOST: ${SMTP_HOST}
      SMTP_PORT: 587
      SMTP_USER: ${SMTP_USER}
      SMTP_PASS: ${SMTP_PASS}
      FROM_ADDRESS: noreply@taskflow.io

  pranor-trace:
    image: ghcr.io/vyuvaraj/pranor-trace:latest
    ports: ["16686:16686"]    # Jaeger-compatible UI
    volumes:
      - trace_data:/data

  pranor-console:
    image: ghcr.io/vyuvaraj/pranor-console:latest
    ports: ["9000:9000"]
    environment:
      PRANOR_GATE_URL: http://pranor-gate:8081
      PRANOR_CACHE_URL: http://pranor-cache:8082
      PRANOR_PULSE_URL: http://pranor-pulse:8083
      PRANOR_TRACE_URL: http://pranor-trace:4317

  # Application
  app:
    build: ./app
    environment:
      DATABASE_URL: postgres://taskflow:${DB_PASSWORD}@postgres:5432/taskflow
      PRANOR_CACHE_URL: http://pranor-cache:8082
      PRANOR_PULSE_URL: http://pranor-pulse:8083
      PRANOR_VAULT_URL: http://pranor-vault:8084
      PRANOR_AUTH_URL: http://pranor-auth:8086
      PRANOR_CHRONO_URL: http://pranor-chrono:8089
    depends_on:
      - postgres
      - pranor-cache
      - pranor-pulse
      - pranor-vault

volumes:
  pg_data:
  cache_data:
  queue_data:
  store_data:
  trace_data:
```

---

## Step 7: Launch the Stack

```bash
# Set environment variables
cp .env.example .env
# Edit .env with your secrets

# Build and start
docker compose up -d

# Watch logs
docker compose logs -f app pranor-gate

# Check everything is healthy
curl http://localhost/health
```

Open Pranor Console at `http://localhost:9000` to see:
- Live request throughput via Pranor Gate
- Cache hit rates in Pranor Cache
- Queue depths and consumer lag in Pranor Pulse
- Distributed traces in Pranor Trace

---

## What You Built

In under an hour, you have:

- ✅ **Multi-tenant REST API** with JWT auth and org isolation
- ✅ **Caching** on all read-heavy endpoints
- ✅ **Event-driven notifications** via Pranor Pulse → Pranor Notify
- ✅ **File uploads** via Pranor Vault
- ✅ **Scheduled jobs** (daily digest, overdue checks) via Pranor Chrono
- ✅ **Full observability** — traces, metrics, logs via Pranor Trace + Pranor Console
- ✅ **One-command deploy** with Docker Compose

The complete source code is available in [pranor-repo/examples/taskflow](https://github.com/vyuvaraj/pranor-repo/tree/main/examples/taskflow).

---

## Series Wrap-Up

This series covered the entire Pranor developer journey:

| Post | Topic |
|------|-------|
| [1 — Introducing Pranor](blog.html?post=01-introducing-pranor) | Ecosystem overview |
| [2 — Pranor in 10 min](blog.html?post=02-getting-started-pranor) | First service |
| [3 — Pranor Gate](blog.html?post=03-api-gateway-pranor-gate) | API gateway deep dive |
| [4 — Pranor Cache](blog.html?post=04-caching-with-pranor-cache) | Distributed caching |
| [5 — Pranor Pulse](blog.html?post=05-event-driven-pranor-pulse) | Event-driven architecture |
| **6 — Full-Stack SaaS** | **This post** |

---

*Want more? Star [pranor-repo](https://github.com/vyuvaraj/pranor-repo) and watch for new posts. Feedback? Open a [discussion](https://github.com/vyuvaraj/pranor-repo/discussions).*
