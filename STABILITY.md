# Pranor Ecosystem API Stability Policy

This document defines the stability guarantees and deprecation policies for the **Serv** ecosystem components.

---

## 1. Stability Tiers

### 🟢 Tier 1: Guaranteed Stable (v1.0.0+)
These endpoints, CLI commands, and libraries have strict backward compatibility guarantees. Breaking changes will only occur in major version bumps (e.g. v1.x to v2.0).

- **Pranor Vault S3 API**: Fully compatible with the standard AWS S3 REST API (SigV4, bucket creation, Put/Get/List/Delete/Multipart uploads).
- **Pranor Pulse STOMP API**: Complies with STOMP 1.1 and 1.2 specifications.
- **Pranor Auth OAuth2/OIDC API**: Implements standard `/oauth/token`, `/oauth/authorize`, `/userinfo`, and `.well-known/openid-configuration` endpoints.
- **Pranor Syntax**: Built-in keywords (`broker`, `store`, `cache`, `notify`, `workflow`) and compiler flags.

### 🟡 Tier 2: Stable, Subject to Evolution
These are stable APIs but may evolve as features are added. We guarantee at least a **3-month deprecation period** (via `Deprecation` and `Sunset` headers) before removal.

- **Pranor Console API**: `/api/v1/` endpoints for querying trace data, alerts, active topologies, and provision tasks.
- **Pranor Hub REST API**: Publishing, searching, and metadata endpoints.
- **Pranor Core Library**: Exported helper functions and middlewares.

### 🔴 Tier 3: Experimental / Internal APIs
These APIs are used strictly for inter-service cluster communications (e.g. Pranor Deploy node orchestration, internal mesh instances). They carry **no stability guarantees** and may change without notice.

---

## 2. Deprecation Process

When an API endpoint in Tier 2 is scheduled for decommissioning:
1. **Header Injections**: The API will immediately return `Deprecation: true` and `Sunset: <date>` HTTP headers.
2. **Logs**: A deprecation warning will be printed to the server logs.
3. **Ecosystem Notice**: The release notes for that version will detail the alternative endpoint.
