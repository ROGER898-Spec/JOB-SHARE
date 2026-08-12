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
│   ├── repository/           # Database implementations & raw SQL queries (external .sql)
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
cd ./backend

```

### 2. Run PostgreSQL via Docker

```bash
docker run --name jobshare-postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -e POSTGRES_DB=jobshare -p 5432:5432 -d postgres:alpine

```

### 3. Run Database Migrations

```bash
migrate -path ./infrastructure/migrations -database "postgres://root:password@localhost:5432/jobshare?sslmode=disable" up

```

### 4. Run the Application

```bash
go run cmd/api/main.go

```

The server will run on: **http://localhost:8080**

---

## 📖 API Documentation (Swagger)

Interactive API documentation is automatically generated and hosted locally.
Once the server is running, open your browser and navigate to:
👉 **http://localhost:8080/swagger/index.html**

---

## 📌 API Endpoints Reference

### 🛠️ System & Documentation

| Method | Endpoint              | Description                  |
| :----- | :-------------------- | :--------------------------- |
| `GET`  | `/health`             | Server health check endpoint |
| `GET`  | `/swagger/index.html` | Swagger UI API Documentation |

---

### 🔐 Auth (Phase 1)

| Method | Endpoint                | Description                                               | Access |
| :----- | :---------------------- | :-------------------------------------------------------- | :----- |
| `POST` | `/api/v1/auth/register` | Register a new account (`admin`, `umkm`, or `freelancer`) | Public |
| `POST` | `/api/v1/auth/login`    | Authenticate user and obtain JWT token                    | Public |

---

### 👤 Profiles (Phase 1)

| Method | Endpoint                               | Description                    | Access       |
| :----- | :------------------------------------- | :----------------------------- | :----------- |
| `POST` | `/api/v1/umkm/profile`                 | Create UMKM Profile            | `UMKM`       |
| `GET`  | `/api/v1/umkm/profile/:user_id`        | Get UMKM Profile details       | `Protected`  |
| `POST` | `/api/v1/freelancers/profile`          | Create Freelancer Profile      | `Freelancer` |
| `GET`  | `/api/v1/freelancers/profile/:user_id` | Get Freelancer Profile details | `Protected`  |

---

### 🗂️ Master Data (Phase 2)

| Method | Endpoint                               | Description               | Access  |
| :----- | :------------------------------------- | :------------------------ | :------ |
| `GET`  | `/api/v1/categories`                   | Get all job categories    | Public  |
| `GET`  | `/api/v1/skills/category/:category_id` | Get skills by category ID | Public  |
| `POST` | `/api/v1/categories`                   | Create a new category     | `Admin` |
| `POST` | `/api/v1/skills`                       | Create a new skill        | `Admin` |

---

### 💼 Jobs / Vacancies (Phase 3)

| Method | Endpoint           | Description                                    | Access |
| :----- | :----------------- | :--------------------------------------------- | :----- |
| `GET`  | `/api/v1/jobs`     | Retrieve a list of all available job vacancies | Public |
| `GET`  | `/api/v1/jobs/:id` | Retrieve job vacancy details by ID             | Public |
| `POST` | `/api/v1/jobs`     | Create a new job post with required skills     | `UMKM` |

---

### 📝 Job Applications (Phase 4)

| Method  | Endpoint                              | Description                                    | Access       |
| :------ | :------------------------------------ | :--------------------------------------------- | :----------- |
| `POST`  | `/api/v1/applications`                | Freelancer applies for a job                   | `Freelancer` |
| `GET`   | `/api/v1/applications/job/:job_id`    | Get all applications for a specific job        | `UMKM`       |
| `GET`   | `/api/v1/applications/freelancer/:id` | Get all applications submitted by a freelancer | `Freelancer` |
| `PATCH` | `/api/v1/applications/:id/status`     | UMKM accepts/rejects an application            | `UMKM`       |

---

### 📋 Workspace & Kanban (Phase 5)

| Method  | Endpoint                            | Description                                        | Access      |
| :------ | :---------------------------------- | :------------------------------------------------- | :---------- |
| `POST`  | `/api/v1/kanban/tasks`              | Create a new task for a job                        | `Protected` |
| `GET`   | `/api/v1/kanban/jobs/:job_id/tasks` | Get all tasks for a specific job                   | `Protected` |
| `PATCH` | `/api/v1/kanban/tasks/:id/status`   | Update task status (`todo`, `in_progress`, `done`) | `Protected` |

---

### 💳 Transactions & Escrow (Phase 6)

| Method  | Endpoint                           | Description                        | Access      |
| :------ | :--------------------------------- | :--------------------------------- | :---------- |
| `POST`  | `/api/v1/transactions`             | Create escrow payment for a job    | `UMKM`      |
| `GET`   | `/api/v1/transactions/job/:job_id` | Get transaction details for a job  | `Protected` |
| `PATCH` | `/api/v1/transactions/:id/release` | Release escrow funds to freelancer | `UMKM`      |

---

### ⭐ Reviews & Ratings (Phase 7)

| Method | Endpoint                         | Description                                 | Access      |
| :----- | :------------------------------- | :------------------------------------------ | :---------- |
| `POST` | `/api/v1/reviews`                | Give rating & feedback after job completion | `UMKM`      |
| `GET`  | `/api/v1/reviews/job/:job_id`    | Get review for a specific job               | `Protected` |
| `GET`  | `/api/v1/reviews/freelancer/:id` | Get all reviews for a freelancer            | `Protected` |
