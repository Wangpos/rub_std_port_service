# Go Student Management Service

This folder contains a minimal Go HTTP server implementation of the Student Management Service using Gin and GORM (SQLite).

Quick start (local)

1. Install Go 1.21+.

2. From this folder, download dependencies:

```bash
cd micro/std_man_ser/go-server
go mod tidy
```

3. Run the server (it will auto-migrate the schema and seed example data if DB is empty):

```bash
go run ./cmd/server
```

The server listens on port 4100 by default. APIs are under `/api` (e.g. `GET /api/students`).

Notes

- This is a minimal scaffold that keeps the same business logic as the TypeScript backend: unique email/student number, list/filter/paginate students, basic JSON envelope responses.
- Use the seeded SQLite file `go_dev.db` created in this folder for local testing.
- You can extend handlers and add middleware (auth, logging) as needed.
