# Genomic API Project Instructions

## 1. Prerequisites

- Docker and Docker Compose installed
- (Optional) Go installed for local development

## 2. Environment Setup

- Copy `.env.example` to `.env` and fill in your database credentials, or create a `.env` file in the project root with:
  ```env
  DB_USER=youruser
  DB_PASSWORD=yourpassword
  DB_NAME=yourdb
  DB_HOST=db
  DB_PORT=5432
  ```

## 3. Build and Run with Docker Compose

- Build and start all services:
  ```sh
  docker compose --env-file .env up --build
  ```
- Stop and remove containers, networks, and volumes:
  ```sh
  docker compose down -v
  ```

## 4. Local Development (without Docker)

- Start PostgreSQL locally and ensure your `.env` matches your local DB settings.
- Run the Go server:
  ```sh
  go run main.go
  ```

## 5. Database Initialization

- On first run, the database schema and initial data are loaded from `genomic_schema.dmbl.sql`.
- To reset the database, run:
  ```sh
  docker compose down -v
  docker compose --env-file .env up --build
  ```

## 6. API Usage

- The API runs at `http://localhost:8080`
- Use `/login` to obtain a JWT token.
- Use `/genomes` endpoints for CRUD operations (protected, requires JWT).

## 7. Useful Docker Commands

- View running containers:
  ```sh
  docker ps
  ```
- View logs:
  ```sh
  docker compose logs
  ```
- Enter a running container:
  ```sh
  docker exec -it genomic_api sh
  ```

---

For more, see the README.md.
