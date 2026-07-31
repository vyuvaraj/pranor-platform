# Serv Runtime Cross-Service Dependencies & Tracing

This document defines the runtime dependency graph and OTel tracing flows across the core Pranor ecosystem.

## Runtime Architecture Diagram

The diagram below illustrates how components interact dynamically at runtime:

```mermaid
graph TD
    subgraph Client Space
        Console[Pranor Console]
    end

    subgraph Service Layer
        Auth[Pranor Auth]
        DB[Pranor Pool]
        Mail[Pranor Notify]
        Flow[Pranor Flow]
    end

    subgraph Infrastructure
        Store[Pranor Vault]
        Queue[Pranor Pulse]
    end

    subgraph Observability
        Trace[Pranor Trace]
    end

    %% Routing / Proxy flows
    Console -->|Auth Verification| Auth
    Console -->|SQL Commands| DB
    Console -->|Delivery Requests| Mail
    Console -->|Workflow Definitions| Flow

    %% Storage persistence flows
    Auth -->|S3 load/save users| Store
    DB -->|S3 load/save migrations| Store
    Mail -->|S3 load/save templates| Store
    Flow -->|S3 load/save checkpoints| Store

    %% Queue flows
    Mail -->|DLQ publish| Queue
    Flow -->|Event queues| Queue

    %% Tracing / OTel flows
    Auth -.->|OTel Tracing| Trace
    DB -.->|OTel Tracing| Trace
    Mail -.->|OTel Tracing| Trace
    Flow -.->|OTel Tracing| Trace
    Console -.->|OTel Tracing| Trace
```

## Dependency Descriptions

1. **State Persistence**: On startup and write mutations, all 4 services query the **Pranor Vault** S3 gateway bucket (or fail back gracefully to mock storage) to ensure dynamic state reload.
2. **Asynchronous Messaging**: **Pranor Notify** pushes failed deliveries to the dead-letter-queue (DLQ) in **Pranor Pulse**. **Pranor Flow** publishes execution steps to **Pranor Pulse** for durable orchestrations.
3. **Observability Pipeline**: All services wrap handlers in OTel tracing middleware, sending span updates to **Pranor Trace** to monitor distributed transactions.
