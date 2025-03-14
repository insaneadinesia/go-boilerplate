# 🚀 Go-Boilerplate

A production-ready Go boilerplate implementing Clean Architecture principles. Built with modern tools and best practices to kickstart your next Golang project.

## ✨ Features

- **Clean Architecture** - Clear separation of concerns with Handler, Usecase, and Repository layers
- **Multiple Interfaces** - Support for REST, gRPC, and Worker implementations
- **Modern Stack** - Uses industry-standard tools and frameworks
- **Observability** - Built-in logging, tracing, and monitoring
- **Documentation** - Auto-generated Swagger/OpenAPI docs
- **Testing** - Ready-to-use testing setup with mocks
- **Configuration** - Environment-based configuration management
- **Database** - GORM integration with migration support
- **Developer Experience** - Hot reload, linting, and formatting tools included

## 🛠 Tech Stack

- **Framework:** [Echo](https://echo.labstack.com/) - High performance, minimalist Go web framework
- **ORM:** [GORM](https://gorm.io/) - The fantastic ORM library for Go
- **Logging:** [Zap](https://github.com/uber-go/zap) via [Gobang](https://github.com/insaneadinesia/gobang/tree/master/logger)
- **Tracing:** [OpenTelemetry](https://opentelemetry.io/) via [Gobang](https://github.com/insaneadinesia/gobang/tree/master/gotel)
- **Testing:** [Testify](https://github.com/stretchr/testify) & [Mockery](https://github.com/vektra/mockery)
- **Documentation:** [Swagger](https://swagger.io/)
- **Worker:** [Asynq](https://github.com/hibiken/asynq)

## 📁 Project Structure

```
.
├── cmd/                   # Application entry points
├── config/                # Configuration management
├── docs/                  # Documentation & Swagger files
├── internal/              # Private application code
│   ├── app/
│   │   ├── container/     # Dependency injection
│   │   ├── driver/        # Database & external service drivers
│   │   ├── entity/        # Domain entities
│   │   ├── handler/       # HTTP, gRPC & Worker handlers
│   │   ├── repository/    # Data access layer
│   │   ├── server/        # Server implementations
│   │   ├── usecase/       # Business logic
│   │   └── wrapper/       # External service clients
│   └── pkg/               # Internal shared packages
├── migration/             # Database migrations
└── main.go                # Application entry point
```

### 📂 Directory Overview

#### `/cmd`
Command-line tools and application entry points. Contains different service startup configurations:
- `server/rest.go` - REST API server
- `server/grpc.go` - gRPC server
- `server/worker.go` - Background worker
- `migrate.go` - Database migration tool

#### `/config`
Application configuration management:
- Environment variables handling
- Configuration structs and validation
- Different environment configs (dev, prod, etc.)

#### `/internal`
Private application code, not meant to be imported by other projects:

##### `/app/container`
Dependency injection container:
- Initializes and manages all application dependencies
- Provides clean dependency graph
- Makes testing and mocking easier

##### `/app/driver`
Initializes application dependencies, like database, redis, etc.

##### `/app/entity`
Domain entities and models:
- Database model definitions
- Shared types and interfaces
- Business domain objects

##### `/app/handler`
Request handlers for different interfaces:
- `/rest` - HTTP REST handlers
- `/grpc` - gRPC service implementations
- `/worker` - Background job handlers

##### `/app/repository`
Data access layer:
- Database operations
- Data persistence logic
- Query implementations

##### `/app/server`
Server implementations:
- HTTP server setup and middleware
- gRPC server configuration
- Worker server setup
- Routing and middleware configuration

##### `/app/usecase`
Business logic layer:
- Core business rules
- Service orchestration
- Data transformation
- Independent of external frameworks

##### `/app/wrapper`
External service clients:
- API client implementations
- Service integrations
- Third-party service wrappers

##### `/pkg`
Shared internal packages like common utilities, helper functions, etc.

#### `/migration`
Database migration files:
- Schema changes
- Data migrations
- Migration history

#### `/docs`
Project documentation:
- API documentation
- Swagger/OpenAPI specs
- Architecture diagrams
- Development guides

## 🚦 Getting Started

1. **Clone the repository**
   ```bash
   git clone https://github.com/insaneadinesia/go-boilerplate.git
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

4. **Run the application**
   ```bash
   # REST API Server
   go run main.go server rest

   # gRPC Server
   go run main.go server grpc

   # Worker
   go run main.go server worker

   # Database Create Migration
   go run main.go migrate create --name=xxx

   # Database Run Migration
   go run main.go migrate up

   # Database Rollback Migration
   go run main.go migrate down
   ```

## 🧪 Testing

```bash
# Generate mocks
make mock

# Run all tests with coverage
make test
```

## 📚 Documentation

- Generate Swagger documentation:
  ```bash
  make swag-init
  ```
- Access Swagger UI at: `http://localhost:9000/swagger/index.html`

## 🏗 Clean Architecture Layers

1. **Handler Layer** (Presentation)
   - Handles HTTP/gRPC requests and responses
   - Input validation
   - Request/Response transformation

2. **Usecase Layer** (Business Logic)
   - Implements business rules
   - Orchestrates data flow
   - Independent of external frameworks

3. **Repository Layer** (Data Access)
   - Database operations
   - External service interactions
   - Data persistence logic

## 📝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Support

For support, email rachmat.adi.p@gmail.com or create an issue in this repository.
