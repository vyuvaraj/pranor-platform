# Upgrading and Migration Guide

This document provides a template and guidelines for performing major and minor upgrades across the **Pranor** ecosystem.

---

## 1. Ecosystem Upgrade Order

When upgrading an entire Pranor installation (e.g. from `v1.0.x` to `v1.1.0`), services should be updated in the following dependency order to avoid service disruptions:

```mermaid
graph TD
    Pranor Core --> Pranor Auth
    Pranor Auth --> Pranor Mesh
    Pranor Mesh --> Pranor Pool
    Pranor Pool --> Pranor Vault
    Pranor Vault --> Pranor Pulse
    Pranor Pulse --> Pranor Console
```

1. **Ecosystem Shared library** (`Pranor Core`)
2. **Identity & Auth Control Plane** (`Pranor Auth`)
3. **Service Mesh Discovery** (`Pranor Mesh`)
4. **Database Proxies & Data Stores** (`Pranor Pool`, `Pranor Vault`)
5. **Message Queuing & Orchestration** (`Pranor Pulse`, `Pranor Flow`)
6. **Dashboard & Dashboards** (`Pranor Console`)

---

## 2. Database Schema Migrations

Every state-persisting service (like `Pranor Auth`, `Pranor Console`, `Pranor Notify`) handles database migrations automatically during startup.
- **Migration Policy**: We only support additive migrations (e.g., `ALTER TABLE ... ADD COLUMN`) in minor/patch releases.
- **Offline / Rollback**: Destructive changes (e.g. dropping columns, renaming tables) are strictly reserved for major releases and must be documented below.

---

## 3. Version Migration Log

### Upgrading to v1.0.0
- **Prefix Changes**: Ensure all HTTP clients target `/api/v1/...` instead of `/api/...`. Backward compatibility routes will remain active during the deprecation window, but are scheduled for removal in `v2.0.0`.
- **Payload Limits**: Request body size limits are now enforced at 10MB by default across all services. Large file uploads should utilize multipart S3 flows in `Pranor Vault`.
