# backend-api-repo

A reference Go backend service demonstrating dependency injection, repository pattern, and idiomatic Go practices with a clean layered architecture.

## Architecture

```
HTTP Request
    │
    ▼
┌──────────┐     ┌──────────┐     ┌──────────┐     ┌─────────┐
│ Handlers │ ──▶ │ Services │ ──▶ │  Store   │ ──▶ │ MongoDB │
│ (api/)   │     │          │     │(store/)  │     │         │
└──────────┘     └──────────┘     └──────────┘     └─────────┘
```

All layers communicate through **interfaces**, making it easy to swap implementations (e.g. replace MongoDB with PostgreSQL by adding a new store).

## Project Structure

```
internal/
├── api/             # HTTP handlers (one file per endpoint)
├── cmd/
│   ├── http/        # HTTP server entrypoint
│   │   └── dto/     # Request/Response types
│   └── grpc/        # gRPC server entrypoint (planned)
├── config/          # Environment-based config loading
├── models/          # Domain models
├── rpc/             # gRPC handlers (planned)
├── services/        # Business logic layer
│   └── user/
└── store/           # Data access layer
    └── mongodb/
external/
└── services/        # External service integrations
```

### Run with Docker Compose

```sh
docker compose up --build
```

This starts the API server on port `3000` and a MongoDB instance on port `27017`.

### Environment Variables

| Variable     | Required | Default | Description              |
|-------------|----------|---------|--------------------------|
| `MONGODB_URI` | Yes      | —       | MongoDB connection string |
| `HTTP_PORT`   | No       | 8080    | HTTP server port          |


## Testing

A [Bruno](https://www.usebruno.com/) API collection is included in `bruno-api-collection/` for manual testing.

```sh
go test ./...
```
