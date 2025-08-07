# 🧬 Genomic API — Golang Microservice

A secure, containerized REST API for genomic data, built with Go, Gin, and PostgreSQL.

## Features

- RESTful CRUD endpoints for genomes, samples, sequence files, variant files, and users
- JWT authentication
- PostgreSQL integration
- Docker & Docker Compose support
- Automatic DB schema and sample data initialization
- Live code reload in development (via Docker volume mount)
- Swagger UI documentation (`/swagger/index.html`)

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
4. Swagger UI available at `http://localhost:8080/swagger/index.html`

### API Usage

- `POST /api/login` — obtain JWT token
- `GET /api/genomes` — list genomes
- `POST /api/genomes` — create genome
- `GET /api/genomes/:id` — get genome by ID
- `PUT /api/genomes/:id` — update genome
- `DELETE /api/genomes/:id` — delete genome

- `GET /api/samples` — list samples
- `POST /api/samples` — create sample
- `GET /api/samples/:id` — get sample by ID
- `PUT /api/samples/:id` — update sample
- `DELETE /api/samples/:id` — delete sample

- `GET /api/sequence` — list sequence files
- `POST /api/sequence` — create sequence file
- `GET /api/sequence/:id` — get sequence file by ID
- `PUT /api/sequence/:id` — update sequence file
- `DELETE /api/sequence/:id` — delete sequence file

- `GET /api/variants` — list variant files
- `POST /api/variants` — create variant file
- `GET /api/samples/:id/variants` — get variants for a sample
- `DELETE /api/variants/:id` — delete variant file

- `GET /api/users` — list users
- `POST /api/users` — create user
- `GET /api/users/:id` — get user by ID
- `PUT /api/users/:id` — update user
- `DELETE /api/users/:id` — delete user

### Database

- Schema and sample data are initialized from `genomic_schema.dmbl.sql` on first run.
- To reset the DB, run:
  ```sh
  docker compose down -v
  docker compose --env-file .env up --build
  ```

### Development

- Code changes are reflected live in the running container (no rebuild needed).
- Run locally with:
  ```sh
  go run main.go
  ```
- To update Swagger docs after changing handler annotations:
  ```sh
  swag init
  ```

---

See `INSTRUCTIONS.md` for more details and commands.
