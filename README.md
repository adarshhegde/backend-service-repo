# backend-api-repo

A reference Go backend service demonstrating dependency injection, repository pattern, and idiomatic Go practices with a clean layered architecture.

## Architecture

```
HTTP Request          gRPC Request
    │                     │
    ▼                     ▼
┌──────────┐     ┌──────────┐
│ Handlers │     │ Handlers │
│ (api/)   │     │ (rpc/)   │
└────┬─────┘     └────┬─────┘
     │                │
     ▼                ▼
   ┌────────────────────┐     ┌──────────┐     ┌─────────┐
   │     Services       │ ──▶ │  Store   │ ──▶ │ MongoDB │
   │                    │     │(store/)  │     │         │
   └────────────────────┘     └──────────┘     └─────────┘
```

Both HTTP and gRPC transport layers share the same services and store. All layers communicate through **interfaces**, making it easy to swap implementations (e.g. replace MongoDB with PostgreSQL by adding a new store).

## Project Structure

```
internal/
├── api/             # HTTP handlers (one file per endpoint)
├── cmd/
│   ├── http/        # HTTP server entrypoint
│   │   └── dto/     # Request/Response types
│   └── grpc/        # gRPC server entrypoint
├── config/          # Environment-based config loading
├── models/          # Domain models (with proto converters)
├── rpc/             # gRPC handlers (one file per RPC)
├── services/        # Business logic layer
│   └── user/
└── store/           # Data access layer
    └── mongodb/
proto-files/
├── backendservice/  # Proto definitions
└── generated-code/  # Generated Go code (do not edit)
external/
└── services/        # External service integrations
```

### Run with Docker Compose

```sh
docker compose up --build
```

This starts the HTTP server on port `3000`, the gRPC server on port `3001`, and a MongoDB instance on port `27017`.

### Makefile

```sh
make proto   # Regenerate protobuf code
make build   # Build all Docker images
make start   # Build and start all services
```

### Environment Variables

| Variable     | Required | Default | Description              |
|-------------|----------|---------|--------------------------|
| `MONGODB_URI` | Yes      | —       | MongoDB connection string |
| `HTTP_PORT`   | No       | 8080    | HTTP server port          |
| `GRPC_PORT`   | No       | 3001    | gRPC server port          |


## Testing

A [Bruno](https://www.usebruno.com/) API/RPC collection is included in `bruno-api-collection/` for manual testing.

```sh
go test ./...
```
