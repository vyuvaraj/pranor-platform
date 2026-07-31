# Pranor 10-Minute Demo Video — Script & Storyboard

> **AG.4** | Record: install → write service → deploy → observe in console  
> Host on YouTube + embed in GitHub Pages / pranor-repo landing page  
> Total runtime: ~10 minutes

---

## Pre-recording Checklist

- [ ] Clean machine (or fresh VM / Docker Desktop) with no prior Serv installation
- [ ] Terminal: Warp or iTerm2 (large font, dark theme)
- [ ] VS Code open with the Serv LSP extension installed
- [ ] Pranor Console pre-seeded with a few traces from the showcase app
- [ ] OBS or Loom ready; resolution 1920x1080; 60fps

---

## Scene 1 — One-Line Install (0:00 – 1:00)

**Narration:**
> "Getting started with Pranor takes exactly one command. Let's install the entire ecosystem on a fresh machine."

**Screen:** Full-screen terminal

```bash
# macOS / Linux
curl -fsSL https://raw.githubusercontent.com/vyuvaraj/pranor-repo/main/scripts/install.sh | bash

# Windows (PowerShell)
irm https://raw.githubusercontent.com/vyuvaraj/pranor-repo/main/scripts/install.ps1 | iex
```

**Show:** Binary list printed after install — `pranor-gate`, `pranor-vault`, `pranor-pulse`, `pranor-console`, `serv` compiler, etc.

> "One command, 16 services, cross-platform. No Docker required, no JVM, no Node. Pure Go binaries."

---

## Scene 2 — Write a Service in Pranor (1:00 – 3:30)

**Narration:**
> "Now let's build a real REST API. We'll use Pranor — a compiled language where infrastructure is syntax, not an import."

**Screen:** VS Code, create `api.pnr`

```
service OrderAPI {

  store "orders" {
    backend: "pranor-vault://localhost:9000"
  }

  cache "order-cache" {
    ttl: 60s
  }

  broker "events" {
    backend: "pranor-pulse://localhost:8082"
  }

  route POST /orders -> Order {
    let order = json.decode<Order>(request.body)
    store.put("orders", order.id, order)
    broker.publish("events", "order.created", order)
    response.json(order)
  }

  route GET /orders/{id} -> Order {
    cached fn getOrder(id string) Order {
      return store.get("orders", id)
    }
    response.json(getOrder(params.id))
  }

}
```

**Show:** VS Code LSP in action — hover docs on `store`, `broker`, auto-complete, inline type errors.

> "Notice: `store`, `broker`, `cache` are keywords. Not imports. The compiler validates that these services exist at build time."

```bash
pranor build api.pnr
```

**Show:** Fast compile output, zero errors.

---

## Scene 3 — Run Locally (3:30 – 4:30)

**Screen:** Split terminal — start ecosystem, then the service

```bash
# Start the infrastructure (Pranor Vault, Pranor Pulse, Pranor Console)
docker-compose up -d

# Run the compiled service
./api
```

**Show:** Curl commands

```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{"id":"ord_001","item":"keyboard","qty":1}'

curl http://localhost:8080/orders/ord_001
```

> "Order created, stored in Pranor Vault, event published to Pranor Pulse — all from 15 lines of code."

---

## Scene 4 — Deploy to Pranor Deploy (4:30 – 5:30)

**Narration:**
> "Now let's deploy. Pranor Deploy is our orchestrator — think systemd meets Kubernetes, but a single binary."

```bash
pranor deploy api.pnr --cloud localhost:7070 --name order-api

pranor-deploy status
```

**Show:** Table listing `order-api -> RUNNING | 1 replica | p99: 2ms`

> "Blue-green deployments, autoscale rules, branch previews — all built in."

---

## Scene 5 — Observe in Pranor Console (5:30 – 8:00)

**Screen:** Browser -> http://localhost:9090

**Show in sequence:**

1. **Topology view** — live graph of services with health indicators
2. **Click `order-api`** — request waterfall, OTel traces from the 3 curl calls
3. **Expand a trace** — shows `route POST /orders` -> `store.put` -> `broker.publish` spans
4. **Metrics panel** — RPS, p99 latency, error rate
5. **Pranor Pulse panel** — `order.created` topic, message count, consumer lag = 0

> "Pranor Console is not just dashboards — it's a control plane. Every service auto-reports telemetry with zero configuration."

---

## Scene 6 — Ecosystem Overview Montage (8:00 – 9:30)

**Narration (voiceover):**
> "In the last 8 minutes you saw Pranor compile, Pranor Vault persist, Pranor Pulse route events, and Pranor Console observe everything. But the ecosystem goes deeper."

**Quick cuts (5-8 seconds each):**

| Cut | Shows |
|-----|-------|
| Pranor Gate config | AI prompt guard, WASM middleware, LLM routing |
| Pranor Auth login flow | JWT, MFA, OIDC, magic links |
| Pranor Flow DAG | Workflow definition and Mermaid visualization |
| Pranor Trace waterfall | Cross-service trace with .pnr source line mapping |
| VS Code extension | Sidebar health panel, CodeLens test runner |

---

## Scene 7 — Call to Action (9:30 – 10:00)

**Screen:** GitHub repo page

**Narration:**
> "Pranor is fully open source under Apache 2.0. Star the repo, try the showcase app, or join our Discord. Links in the description."

**Lower-third text:**
- github.com/vyuvaraj/pranor-repo
- star | docs | discord

---

## Production Notes

| Item | Detail |
|------|--------|
| **Music** | Royalty-free lo-fi (Pixabay or Uppbeat), -18 LUFS |
| **Captions** | Auto-generated via YouTube, reviewed for technical terms |
| **Chapters** | Add YouTube timestamps matching each scene above |
| **Thumbnail** | Dark background, "10 min" badge, Pranor logo, terminal window |
| **Description** | Include install commands, repo link, showcase app link, Discord |
| **Embed** | Add iframe to pranor-repo/index.html hero section |
