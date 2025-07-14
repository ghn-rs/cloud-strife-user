# Cloud Strife User Service

A gRPC microservice for user management in the Cloud Strife ecosystem.

## Features

- **User Management**: Complete CRUD operations for user entities
- **Authentication**: Password hashing using bcrypt
- **Database Integration**: MySQL with GORM ORM and auto-migrations
- **Logging**: Structured logging with Elasticsearch integration
- **Health Check**: Service health monitoring endpoint

## API Endpoints

The service provides the following gRPC methods:

- `CreateUser(UserRequest) -> UserResponse`
- `GetUser(IdRequest) -> UserResponse`
- `UpdateUser(UserUpdateRequest) -> UserResponse`
- `DeleteUser(IdRequest) -> BaseResponse`
- `PaginateUser(PaginationRequest) -> UsersResponse`
- `ListUsers(EmptyRequest) -> UsersResponse`
- `EditUser(UserUpdateRequest) -> UserResponse`
- `HealthCheck(EmptyRequest) -> BaseResponse`

## User Entity

The User entity includes:
- **EssentialEntity**: Base fields (ID, created_date, updated_date) from corelib
- **username**: Unique username (required)
- **password**: Hashed password (required, not returned in responses)
- **email**: Unique email address (required)
- **role**: User role (default: "user")
- **status**: User status (0=active, 1=inactive)
- **lastLogin**: Timestamp of last login

## Prerequisites

- Go 1.24.5+
- MySQL database
- Protocol Buffers compiler
- Buf CLI tool

## Setup

1. **Install dependencies**:
   ```bash
   make install_deps
   ```

2. **Configure the service**:
   Edit `config.json` with your database and service settings.

3. **Generate Protocol Buffers**:
   ```bash
   make proto
   ```

4. **Run the service**:
   ```bash
   make run
   ```

## Configuration

The service is configured via `config.json`:

```json
{
  "app": {
    "address": "0.0.0.0",
    "port": 5001,
    "name": "Cloud Strife User Service",
    "debug": true
  },
  "database": {
    "host": "localhost",
    "port": 3306,
    "database": "db_cloudstrife_user",
    "username": "your_username",
    "password": "your_password",
    "autoMigrate": true
  },
  "elastic": {
    "url": "http://localhost:9200/",
    "level": "debug"
  }
}
```

## Docker

Build and run using Docker:

```bash
docker build -t cloud-strife-user .
docker run -p 5001:5001 cloud-strife-user
```

## Development

- **Code formatting and linting**: `make tidy`
- **Clean build**: `make clean`
- **Full build**: `make all`

## Architecture

The service follows Clean Architecture principles:

- **Protocol Buffers**: Service contracts and message definitions
- **gRPC Layer**: Service handlers and server setup
- **Use Case Layer**: Business logic implementation
- **Repository Layer**: Database access and operations
- **Entity Layer**: Data models and domain objects
- **Dependencies**: Configuration, database, and external service connections