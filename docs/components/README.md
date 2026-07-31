# Pranor Component Catalog

This catalog outlines all 15 operational services of the Pranor ecosystem and their current lifecycle statuses.

## Services Catalog

### 1. Ingress & Edge
* **Pranor Gate**: API Gateway. Handles routing, rate limiting, and transformations.
* **Pranor Tunnel**: Public secure tunnel endpoint for exposing local services.

### 2. Identity & Persistence
* **Pranor Auth**: OIDC token issuer, key-rotations, and TOTP MFA provider.
* **Pranor Pool**: SQL database proxy manager (SQLite, Postgres, Oracle).
* **Pranor Vault**: S3-compatible persistent object store.

### 3. Messaging & Workloads
* **Pranor Pulse**: WAL-backed message broker (STOMP / HTTP).
* **Pranor Chrono**: Cron scheduling control plane.
* **Pranor Flow**: Sagas and Schedulers DAG Workflow Engine.
* **Pranor Notify**: SMTPTransactional Mail Agent.

### 4. Service Mesh & Utilities
* **Pranor Mesh**: Service registry and client-side load balancer.
* **Pranor Cache**: Redis connection caching wrapper.
* **Pranor Core**: Shared middleware library.

## Core Component Reference APIs

### Pranor Gate
- **Endpoints**: Matches routes configured in `config.json`.
- **Admin**: `GET /api/routes` - lists current active paths.
- **Port**: `8080`

### Pranor Vault
- **Endpoints**:
  - `PUT /buckets/{name}/{object}` - uploads file payload.
  - `GET /buckets/{name}/{object}` - downloads file payload.
- **Port**: `8081`

### Pranor Auth
- **Endpoints**:
  - `POST /oauth/token` - authenticates credentials and returns JWT.
  - `GET /oauth/keys` - returns JWKS active rotation keys.
- **Port**: `8098`

### Pranor Mesh
- **Endpoints**:
  - `POST /api/register` - registers node instances.
  - `POST /api/heartbeat` - heartbeats keepalive.
- **Port**: `8089`

