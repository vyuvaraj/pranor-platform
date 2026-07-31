# Pranor Hub — Package Registry Hub

> **Status:** ✅ Production | **Port:** 8088 | **Repository:** [Pranor Hub](https://github.com/vyuvaraj/pranor/tree/main/packages/Pranor Hub)

## Overview

Pranor Hub is a package registry hub for `.pnr` modules. It provides S3-backed storage via Pranor Vault, semver resolution, BFS dependency tree computation, cryptographic package signing, JWT-authenticated publishing, and an embedded dashboard for browsing packages.

## Key Features

- Package publishing and versioning for `.pnr` modules
- S3-backed storage via Pranor Vault
- Semantic versioning resolution with range constraints
- BFS dependency tree computation
- Cryptographic package signing and verification
- JWT-authenticated publish operations
- Embedded web dashboard for package browsing
- Dependency vulnerability advisory integration

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVREGISTRY_PORT` | HTTP listen port | `8088` |
| `SERVSTORE_URL` | Pranor Vault URL for package storage | (required) |
| `PRANOR_JWT_SECRET` | JWT verification secret | (required) |

## Endpoints

| Endpoint | Description |
|----------|-------------|
| `GET /healthz` | Liveness probe |
| `POST /api/v1/publish` | Publish a package version |
| `GET /api/v1/packages/{name}` | Get package metadata |
| `GET /api/v1/resolve/{name}/{version}` | Resolve version and dependencies |

## Pranor Integration

```srv
serv add <package>
serv install
serv publish
```
