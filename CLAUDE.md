# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`opem-common` is a Go shared library (`github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-common`) consumed by other OPEM platform services. It provides HTTP clients, linked service integrations, and utilities.

## Common Commands

```bash
go build ./...          # Build all packages
go test ./...           # Run all tests
go test -v -run TestName ./path/to/pkg  # Run a single test
go vet ./...            # Static analysis
go fmt ./...            # Format code
go mod tidy             # Clean up dependencies
```

## Architecture

The module is organized into three main packages:

### `clients/`
Core HTTP client infrastructure used across all OPEM services.

- **`request-context.go`** — `ApiRequestContext` holds per-request metadata (domain, namespace, language, API key, headers). Auto-generates `Request-Id` UUIDs. Standard headers use the `X-R3ds9-*` prefix.
- **`response.go`** — `ApiResponse` is the standardized envelope for all API responses; implements the `error` interface. Factory methods: `NewSuccessResponse()`, `NewBadRequestError()`, `NewInternalServerError()`.
- **`apicms/`** — Client for the CMS API service. Uses URL path placeholders (`:hostDomain`, `:hostNamespace`, `:hostLang`, `:fileId`) for dynamic URL construction. Supports HAR tracing via `tpm-http-archive`.

### `linkedservices/`
Service registry and factory layer for initializing all external service clients.

- **`registry.go`** — `InitRegistry(cfg)` is the single entry point. Initializes MongoDB, S3, Kafka, and Hermodr clients. Provides factory methods like `NewHermodrClient()`.
- **`config.go`** — Aggregates configuration for all linked services. Call `PostProcess()` after unmarshalling.
- **`hermodr/`** — Client for the Hermodr authentication service. Wraps a REST HTTP client.

### `util/`
- **`range.go`** — `Range` (From/To) and `RangeSet` data structures. Key operations: `Contains()`, `Add()` with `Consecutive` or `MinMax` modes, `Defragment()` to consolidate adjacent ranges.

## Key Dependencies

| Package | Purpose |
|---|---|
| `tpm-http-client` | Base REST client used by all HTTP clients |
| `tpm-http-archive` | HAR tracing for HTTP requests |
| `tpm-mongo-common` | MongoDB driver wrapper |
| `tpm-kafka-common` | Kafka producer/consumer |
| `tpm-aws-common` | AWS S3 utilities |
| `zerolog` | Structured logging |
| `opentracing-go` / `jaeger-client-go` | Distributed tracing |

## Testing Notes

Integration tests in `clients/apicms/client_test.go` target `localhost:8082` by default and support Jaeger tracing via environment variables. These require a running CMS instance and are not pure unit tests.
