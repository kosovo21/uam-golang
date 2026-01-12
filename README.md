# User Access Management (UAM) - Go JWT Authentication Service

A secure, production-ready REST API for user authentication and management built with Go (Golang), Gin, and JWT.

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Technology Stack](#technology-stack)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [API Documentation (Swagger)](#api-documentation-swagger)
- [Configuration](#configuration)
- [API Endpoints](#api-endpoints)
- [Security Features](#security-features)
- [Project Structure](#project-structure)
- [Usage Examples](#usage-examples)
- [Development](#development)

## 🎯 Overview

UAM is a comprehensive user authentication and management service that provides secure JWT-based authentication. It is rewritten in Go to provide high performance and low resource usage while maintaining industry-standard security measures including Argon2 password hashing and robust token validation.

## ✨ Features

- **JWT Authentication**: Secure token-based authentication
- **User Registration**: Create new user accounts with email validation
- **User Management**: Retrieve user information (password-protected)
- **Password Security**: Argon2id for secure password hashing
- **Input Validation**: Request validation using validator library
- **Exception Handling**: Structured error responses
- **CORS Support**: Configured for frontend integration
- **Audit Trail**: Automatic timestamp tracking (CreatedAPI, UpdatedAt)
- **Role-Based**: User roles support
- **API Documentation**: Interactive Swagger/OpenAPI documentation

## 🛠 Technology Stack

- **Language**: Go 1.22+
- **Framework**: [Gin Web Framework](https://github.com/gin-gonic/gin)
- **Database**: PostgreSQL
- **ORM**: [GORM](https://gorm.io/)
- **Configuration**: [Viper](https://github.com/spf13/viper)
- **Authentication**: [golang-jwt](https://github.com/golang-jwt/jwt)
- **Password Hashing**: [Argon2](https://github.com/alexedwards/argon2)
- **API Documentation**: [Swaggo](https://github.com/swaggo/swag)
- **Build Tool**: Go Modules

## 📦 Prerequisites

- Go 1.22 or higher
- PostgreSQL 12+ (or Docker)
- Make (optional, for convenience commands)

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone <repository-url>
cd uam-golang
```

### 2. Database Setup

Create a PostgreSQL database:

```sql
CREATE DATABASE uam;
CREATE USER myuser WITH PASSWORD 'mypassword';
GRANT ALL PRIVILEGES ON DATABASE uam TO myuser;
```

Or use Docker Compose:

```bash
docker-compose up -d
```

### 3. Configuration

Create a `.env` file or `config.yaml` in the root directory:

```properties
DB_HOST=localhost
DB_USER=myuser
DB_PASSWORD=mypassword
DB_NAME=uam
DB_PORT=5432
JWT_SECRET=your-very-long-and-secure-secret-key-at-least-64-characters
JWT_EXPIRATION=24h
SERVER_PORT=8080
```

### 4. Build and Run

Install dependencies:

```bash
go mod download
```

Run the application:

```bash
go run cmd/api/main.go
```

Or build the binary:

```bash
go build -o bin/uam cmd/api/main.go
./bin/uam
```

The service will start on `http://localhost:8080`.

## 📚 API Documentation (Swagger)

The API documentation is generated using Swaggo.

### Access Swagger UI

Once running, navigate to:

**Swagger UI:** `http://localhost:8080/swagger/index.html`

### Features

- **Interactive API Testing**: Test endpoints directly
- **JWT Auth**: Use the "Authorize" button (Bearer scheme)
- **Schemas**: View request/response models

## ⚙️ Configuration

The application uses Viper for configuration, supporting `.env`, `config.yaml`, or environment variables.

| Variable | Description | Default |
|----------|-------------|---------|
| `JWT_SECRET` | JWT signing secret | (Required) |
| `JWT_EXPIRATION` | Token duration | 24h |
| `DB_DSN` | Database connection string | (Constructed from vars) |
| `GIN_MODE` | Gin mode (debug/release) | debug |

## 📡 API Endpoints

### Authentication

#### Register User
```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

#### Login
```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### User Management

#### Get Profile
```http
GET /api/v1/users/me
Authorization: Bearer <token>
```

#### Get All Users (Admin)
```http
GET /api/v1/users
Authorization: Bearer <token>
```

## 🔒 Security Features

### Authentication
- **JWT**: Stateless Bearer token authentication.
- **Middleware**: Custom Gin middleware for protecting routes.

### Password Security
- **Argon2id**: High-security password hashing using the specialized Go library.
- **Salt**: Automatic random salt generation.

### Input Validation
- **Struct Tags**: Uses `go-playground/validator` for email format and required field validation.

## 📁 Project Structure

Standard Go project layout:

```
.
├── cmd/
│   └── api/
│       └── main.go           # Application entrypoint
├── internal/
│   ├── config/               # Configuration loading
│   ├── handlers/             # HTTP Controllers
│   ├── middleware/           # Auth and other middleware
│   ├── models/               # GORM entities & DTOs
│   ├── repository/           # Database access
│   └── service/              # Business logic
├── pkg/
│   └── utils/                # Shared utilities (JWT, Crypto)
├── docs/                     # Swagger generated docs
├── go.mod                    # Dependencies
└── README.md
```

## 💡 Usage Examples

### Using cURL

#### Register
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"securepass123"}'
```

#### Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"securepass123"}'
```

## 🔧 Development

### Running Tests
```bash
go test ./...
```

### Generating Swagger Docs
```bash
swag init -g cmd/api/main.go
```

## 📝 Notes

- **Production**: Set `GIN_MODE=release`.
- **Database**: Ensure PostgreSQL is running and accessible.

## 📄 License

[Specify your license here]
