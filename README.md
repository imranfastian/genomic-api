# 🧬 Genomic API — Golang Microservice

A secure, containerized REST API for genomic data, built with Go, Gin, and PostgreSQL.

## Features

- RESTful CRUD endpoints for genomes and related resources
- JWT authentication
- PostgreSQL integration
- Docker & Docker Compose support
- Automatic DB schema initialization

## Getting Started

### Prerequisites

- Docker & Docker Compose
- (Optional) Go for local development

### Setup

1. Copy `.env.example` to `.env` and update credentials, or create a `.env` file in the project root.
2. Build and run with Docker Compose:
   ```sh
   docker compose --env-file .env up --build
   ```
3. The API will be available at `http://localhost:8080`

### API Usage

- `POST /login` — obtain JWT token
- `GET /genomes` — list genomes (JWT required)
- `POST /genomes` — create genome (JWT required)
- `GET /genomes/:id` — get genome by ID (JWT required)
- `PUT /genomes/:id` — update genome (JWT required)
- `DELETE /genomes/:id` — delete genome (JWT required)

### Database

- Schema is initialized from `genomic_schema.dmbl.sql` on first run.
- To reset the DB, run:
  ```sh
  docker compose down -v
  docker compose --env-file .env up --build
  ```

### Development

- Run locally with:
  ```sh
  go run main.go
  ```

---

See `INSTRUCTIONS.md` for more details and commands.
