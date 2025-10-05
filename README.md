# Production-Ready Auth Service (Golang)

An authentication & authorization microservice built in Golang — designed for production usage (JWT, refresh tokens, API routes, middleware, etc.)

## Table of Contents

- [Features](#features)
- [Architecture](#architecture)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Clone & Setup](#clone--setup)
  - [Configuration](#configuration)
  - [Database Migrations](#database-migrations)
  - [Running the App](#running-the-app)
  - [Generating Swagger Docs](#generating-swagger-docs)
- [API Endpoints](#api-endpoints)
- [Environment Variables](#environment-variables)
- [Deployment](#deployment)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- User registration / login / logout
- JWT access tokens + refresh tokens
- Middleware for protected routes
- Role / permission support (if implemented)
- Structured logging, error handling
- Swagger / OpenAPI docs
- Configurable via environment variables
- Possibly database support (PostgreSQL, etc.)

> **Note:** Some of these are assumed features based on a typical production auth service. Adjust to match your actual implementation.

## Architecture

- `cmd/` — application entrypoint
- `internal/` — core business logic, domain, services, repositories
- `docs/` — swagger / API documentation
- `.env.sample` — sample environment variable file

## Prerequisites

- Go (version 1.18+ or whatever your project targets)
- A relational database (e.g. PostgreSQL, MySQL)
- (Optional) `swag` tool installed for generating Swagger docs

You can install swag via:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

## Getting Started

### Clone & Setup

```bash
git clone https://github.com/Kushan2k/production-ready-auth-service-using-golang
cd production-ready-auth-service-using-golang
go mod download
```

### Configuration

1. Copy `.env.sample` to `.env` and fill in your configuration values (DB connection, JWT secrets, etc.)

```bash
cp .env.sample .env
```

2. Update the `.env` file with your specific configuration values.

### Database Migrations

Not yet implemented. You can use a tool like `golang-migrate` or `GORM` auto-migrations.

```bash

```

### Running the App

```bash
go run cmd/main.go
```

### Generating Swagger Docs

```bash
swag init -g cmd/main.go
```

if you have `swag` installed, this will generate the Swagger docs in the `docs/` folder.

if you don't have `swag` installed, you can skip this step.

if you are in main.go file, you can run the following command to generate the Swagger docs:

```bash
swag init
```

## API Endpoints

- `POST /api/v1/auth/register` — Register a new user
- `POST /api/v1/auth/login` — Login and receive JWT tokens
- `POST /api/v1/auth/resent-otp` — Resend OTP for 2FA
- `POST /api/v1/auth/verify-otp` — Verify OTP for 2FA

## Environment Variables

The application uses the following environment variables (set in `.env`):

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=authdb

JWT_SECRET=supersecret
JWT_REFRESH_SECRET=anothersecret

MAIL_HOST=smtp.example.com
MAIL_PORT=587
MAIL_USERNAME=youruser
MAIL_PASSWORD=yourpassword

SERVER_PORT=8080
```

## Deployment

You can deploy this service using Docker, Kubernetes, or any cloud provider. Ensure your environment variables are set correctly in the deployment environment.

1. Build a binary or Dockerize your application

2. Make sure environment variables are set in your deployment environment

3. Migrate your database in the production environment

4. Run the binary / container

```dockerfile
FROM golang:1.20-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o auth-service cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/auth-service .
COPY .env ./
EXPOSE 8080
CMD ["./auth-service"]


```

You can also use Docker Compose if you want to spin up a DB + auth service together.

## Contributing

Thanks for your interest! Some guidelines:

1. Fork the repo

2. Create feature branches

3. Run tests, lint your code

4. Commit with meaningful messages

5. Submit PRs

Ensure Swagger docs are updated when endpoints change.

## License

MIT License. See `LICENSE` for details.
