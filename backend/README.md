# Jobshare Backend API

Official backend service for the **Jobshare** platform, built with **Go**, **Fiber**, **PostgreSQL**, and structured following **Clean Architecture (DDD principles)**.

---

## 🚀 Tech Stack

- **Language:** Go (Golang)
- **Web Framework:** Fiber v2
- **Database:** PostgreSQL
- **Migration Tool:** golang-migrate
- **API Documentation:** Swagger (Swaggo)
- **Authentication:** Bcrypt (Password Hashing)

---

## 📂 Project Architecture

This project follows a strict Clean Architecture / Domain-Driven Design (DDD) layout:

```text
backend/
├── cmd/
│   └── api/                  # Main application entrypoint
├── internal/
│   ├── domain/               # Core entities and repository interfaces
│   ├── repository/           # Database implementations & raw SQL queries
│   ├── usecase/              # Business logic layer
│   └── delivery/             # HTTP handlers, DTOs, and routing
├── pkg/                      # Shared helper packages (response, pagination, etc.)
├── infrastructure/           # Database migrations and configurations
└── docs/                     # Auto-generated Swagger documentation
```

---

## 📋 Prerequisites

Make sure you have the following installed on your machine:

- [Go](https://golang.org/) (v1.20+)
- [Docker & Docker Compose](https://www.docker.com/)
- [golang-migrate CLI](https://github.com/golang-migrate/migrate) (for database migrations)

---

## ⚙️ Getting Started

### 1. Clone the Repository & Navigate to Backend

```bash
git clone https://github.com/ROGER898-Spec/JOB-SHARE.git
cd JOB-SHARE/backend
```

### 2. Run PostgreSQL via Docker

Ensure you have a PostgreSQL container running. Adjust the credentials based on your setup if necessary.

```bash
docker run --name jobshare-postgres \
  -e POSTGRES_USER=root \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=jobshare \
  -p 5432:5432 \
  -d postgres:alpine
```

### 3. Run Database Migrations

Apply the initial database schema using `golang-migrate`:

```bash
migrate \
  -path ./infrastructure/migrations \
  -database "postgres://root:password@localhost:5432/jobshare?sslmode=disable" \
  up
```

### 4. Run the Application

Start the Go Fiber server:

```bash
go run cmd/api/main.go
```

The server will run on:

**http://localhost:8080**

---

## 📖 API Documentation (Swagger)

Interactive API documentation is automatically generated and hosted locally.

Once the server is running, open your browser and navigate to:

**http://localhost:8080/swagger/index.html**

---

## 📌 Phase 1 Features (Auth & Core)

| Method | Endpoint                | Description                                               |
| ------ | ----------------------- | --------------------------------------------------------- |
| `POST` | `/api/v1/auth/register` | Register a new account (`admin`, `umkm`, or `freelancer`) |
| `POST` | `/api/v1/auth/login`    | Authenticate user and obtain session/profile data         |
| `GET`  | `/health`               | Server health check endpoint                              |
