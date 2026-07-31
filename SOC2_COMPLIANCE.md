# SOC2 Security Controls & Compliance Documentation

This document describes the security controls and compliance mechanisms implemented across the Pranor ecosystem, providing evidence for audit and security reviews.

---

## 1. Encryption Controls

### Encryption in Transit
- **TLS Enforced API Endpoints**: All service communication through `Pranor Gate` is encrypted using TLS 1.3 / HTTPS.
- **Inter-Service Mesh mTLS**: Mutual TLS (mTLS) with client certificate verification is enforced inside `Pranor Mesh` to prevent man-in-the-middle attacks and verify identity.
- **Database Proxy mTLS**: `Pranor Pool` requires mutual TLS verification for database client connections.
- **Message Broker mTLS**: `Pranor Pulse` restricts enterprise topic publishers and subscribers to mTLS verified clients.

### Encryption at Rest
- **Envelope Encryption**: `Pranor Vault` S3 storage layers implement envelope encryption (data encryption keys wrapped by master keys) to secure object payloads at rest.
- **Token Cryptography**: `Pranor Auth` secures refresh and access tokens using cryptographically signed HMAC/SHA-256 tokens before database persistence.

---

## 2. Access Control & Authorization (RBAC)

- **Token Revocation**: `Pranor Auth` records refresh tokens in a database to enable remote token revocation and active session audits.
- **Role-Based Access Control (RBAC)**: Fine-grained RBAC is enforced on administrative endpoints (e.g., creating crons in `Pranor Chrono`, changing routing configurations).
- **Access Privilege Audits**: `Pranor Console` provides dashboards to monitor active admin sessions, API token lifetimes, and assigned user roles.

---

## 3. Security Auditing & Monitoring

- **Structured JSON Audit Trails**: `Pranor Auth` outputs structured JSON audit logs for login attempts, MFA adjustments, and security level alterations.
- **Replication History**: `Pranor Pool` records query histories to track data schema changes and database transactions.
- **Centralized Instrumentation**: OpenTelemetry (OTel) instrumentation across all components monitors latency anomaly patterns, database queries, and queue consumer lag.

---

## 4. Data Retention & Incident Recovery

- **Mail Queue Retention**: `Pranor Notify` provides automated disk queue retention settings, purging non-pending items older than configured thresholds.
- **Dead Letter Queues (DLQ)**: `Pranor Notify` and `Pranor Pulse` isolate failing or corrupted messages in DLQs to prevent data loss and support incident inspection.
- **Checkpoint Saga States**: `Pranor Flow` persists workflow execution step state machine checkpoints in database storage layers, supporting safe state recovery during node failovers.
