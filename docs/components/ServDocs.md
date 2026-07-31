# ServDocs — Documentation Generator

> **Status:** 🟢 Stable | **Port:** 8089 | **Repository:** [ServDocs](https://github.com/vyuvaraj/pranor/tree/main/packages/ServDocs)

## Overview

ServDocs is a CLI documentation generator that parses `.pnr` source files to extract routes, structs, functions, and middleware chains. It produces interactive HTML documentation sites, OpenAPI 3.0 JSON specifications, and typed client SDKs for TypeScript, Dart, and Swift — all from a single parse pass.

## Package Structure

```
pkg/parser/     — SrvDoc types + ParseSrvFile
pkg/generator/  — GenerateHtml (HTML site) + GenerateClientSDK (TS/Dart/Swift)
pkg/openapi/    — Generate (OpenAPI 3.0 JSON)
main.go         — CLI orchestration (<100 lines)
```

## Key Features

- `.pnr` file parsing: routes, structs, functions, middleware chains
- Interactive HTML documentation site (dark mode, live search, schema expand, versioned)
- OpenAPI 3.0 JSON spec generation with path parameter normalization
- Multi-language client SDK generation: TypeScript, Dart, Swift
- Multi-language code examples embedded in docs (cURL, Go, JavaScript)
- Versioned docs with in-page version selector
- Self-contained single-file HTML output

## CLI Commands

| Command | Description |
|---------|-------------|
| `servdocs generate` | Generate a static HTML documentation site |
| `servdocs openapi` | Generate an OpenAPI 3.0 JSON specification |
| `servdocs serve` | Serve docs locally via HTTP |
| `servdocs client` | Generate a typed client SDK |

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVDOCS_PORT` | HTTP listen port | `8089` |

## Endpoints

| Endpoint | Description |
|----------|-------------|
| `GET /healthz` | Liveness probe |
| `GET /` | Generated documentation UI |

## Pranor Integration

```srv
serv doc <file.pnr>
```

## Test Coverage

17 test functions across 4 packages (`pkg/parser`, `pkg/generator`, `pkg/openapi`, integration in `main`).

```bash
go test ./...   # runs all 4 packages
```
