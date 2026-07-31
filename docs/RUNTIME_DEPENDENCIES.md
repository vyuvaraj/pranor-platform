# Pranor Runtime Dependencies & Integration Matrix

This document maps the complete runtime dependencies and network flow patterns across all 15 operational services of the Pranor ecosystem.

## Interaction Architecture Graph

```mermaid
graph TD
    %% Clients and Gateway Ingress
    Client[Browser / REST Client] -->|HTTP / WebSocket| Pranor Gate[Pranor Gate API Gateway]
    
    %% Gateway to Backend Services
    Pranor Gate -->|Routes Requests| Pranor Mesh[Pranor Mesh Service Discovery]
    Pranor Gate -->|Loads Config| Pranor Vault[Pranor Vault S3 Object Store]
    Pranor Gate -->|Authenticates| Pranor Auth[Pranor Auth Identity & JWT Provider]

    %% Service Mesh Routing Instance
    Pranor Mesh -->|Discovers Host| Pranor MeshInstances[Running Srv instances]
    
    %% Operational Core Services
    Pranor MeshInstances -->|Publishes Events| Pranor Pulse[Pranor Pulse Message Broker]
    Pranor MeshInstances -->|Schedules Workloads| Pranor Chrono[Pranor Chrono Scheduler]
    Pranor MeshInstances -->|Invokes Pipelines| Pranor Flow[Pranor Flow Workflow Engine]
    Pranor MeshInstances -->|Writes telemetry| Pranor Trace[Pranor Trace OTel Collector]
    Pranor MeshInstances -->|Queries Data| Pranor Pool[Pranor Pool SQL Proxy Manager]
    Pranor MeshInstances -->|Caches Responses| Pranor Cache[Pranor Cache Redis Wrapper]
    Pranor MeshInstances -->|Sends Emails| Pranor Notify[Pranor Notify SMTP Agent]
    
    %% Observability Control Center
    Pranor Console[Pranor Console Dashboard] -->|Polls Health| Pranor MeshInstances
    Pranor Console -->|Reads Logs/Spans| Pranor Trace
    Pranor Console -->|Exposes Tunneled Ports| Pranor Tunnel[Pranor Tunnel Local Ingress]
```

## Service Port Registry

| Port | Service Name | Protocol | Role |
|------|--------------|----------|------|
| `8080` | `Pranor Gate` | HTTP | Ingress API Gateway |
| `8081` | `Pranor Vault` | HTTP | S3 Storage Engine |
| `8082` | `Pranor Pulse` | HTTP/STOMP | Queue Broker |
| `8083` | `Pranor Console` | HTTP | Operational Dashboard |
| `8084` | `Pranor Cache` | RESP/HTTP | Redis Cache Proxy |
| `8085` | `Pranor Chrono` | HTTP | Scheduler Control plane |
| `8089` | `Pranor Mesh` | HTTP/UDP | Service Registry Node |
| `8090` | `Pranor Trace` | HTTP/gRPC | OpenTelemetry Collector |
| `8094` | `Pranor Notify` | HTTP | Transactional Mail Agent |
| `8096` | `Pranor Flow` | HTTP | DAG Workflow Engine |
| `8097` | `Pranor Pool` | HTTP | SQL Persistence Proxy |
| `8098` | `Pranor Auth` | HTTP | Identity and MFA provider |
| `8443` | `Pranor Tunnel` | HTTPS | Tunnel and Let's Encrypt Ingress |

## Interaction Flows

### 1. Ingress Request Authentication
1. **Client** hits `Pranor Gate` on `:8080/api/users`.
2. `Pranor Gate` extracts the token and validates against keys fetched from `Pranor Auth` OIDC configurations.
3. If valid, request is forwarded down to the corresponding `Pranor Mesh` registered target host.

### 2. Event-Driven Workflow Run
1. `Pranor Chrono` triggers a scheduled event on a timer payload.
2. The execution goes to `Pranor Pulse` topics.
3. A listening worker consumer picks up the task, executes a step, and writes artifacts to `Pranor Vault` S3 buckets.
