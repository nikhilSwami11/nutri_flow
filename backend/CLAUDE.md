# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run the server
go run ./cmd/api

# Build binary
go build -o bin/api ./cmd/api

# Tidy dependencies
go mod tidy

# Start infrastructure (MongoDB + Redis)
docker compose up -d   # run from the parent directory (nutriflow/)
```

No test files exist yet. When adding tests, run them with:
```bash
go test ./...
go test ./internal/pantry/...  # single package
```

## Architecture

Go REST API using **chi** router with a 3-tier layered architecture per domain:

- `cmd/api/main.go` — entry point: loads env, connects DB, registers routes, starts server
- `pkg/db/db.go` — MongoDB connection setup (returns `*mongo.Database`)
- `internal/<domain>/` — each domain has four files:
  - `model.go` — structs with BSON/JSON tags
  - `handler.go` — HTTP handlers wired to chi routes
  - `service.go` — business logic, cross-domain coordination, external API calls
  - `repository.go` — all MongoDB operations (10s context timeout pattern)

Request flow: **Handler → Service → Repository → MongoDB**

**Current domains:** `pantry`, `profile`
**Placeholder domains (empty):** `recipes`, `photo`, `sessions`
**Placeholder package (empty):** `pkg/ai`

Routes are registered in `main.go`:
```
GET  /health
GET|POST      /pantry
PUT|DELETE    /pantry/{id}
GET|POST|PUT  /profile
```

All endpoints are scoped per user via `user_id` (query param or request body). There is no auth middleware yet.

## Environment

Copy `.env example` to `.env`. Required variables:

```
MONGODB_URI=mongodb://nutriflow:nutriflow123@localhost:27017/nutriflow_db?authSource=admin
REDIS_URL=redis://localhost:6379
PORT=8080
```

MongoDB and Redis are defined in `docker-compose.yml` at the repo root (`nutriflow/`), not inside `backend/`.

Redis is configured but **not yet used** in application code.

## Conventions

- MongoDB collections: `pantry_items`, `user_profiles`, `recipes`
- Profile documents use `user_id` as the MongoDB `_id` field
- Recipe documents use `ObjectID` as `_id`
- Partial updates in profile use `map[string]interface{}` with `$set`
- New domains go in `internal/` and must include `handler.go`, `service.go`, `repository.go`, and `model.go`

claude --resume 932bf3b1-dd9d-4dd4-8e59-26282becfaa6 