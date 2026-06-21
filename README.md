# ZenXBattle Auth & User Admin Service

gRPC service for user authentication, authorization, and admin operations. Handles JWT token management, user CRUD, and role-based access control.

## Architecture

```
ApiGateway → gRPC (:50051) → AuthUserAdminService
                                 ├── PostgreSQL (GORM)
                                 ├── Redis (session cache)
                                 └── JWT token store
```

## Tech Stack

- **Go** + gRPC
- **GORM** (PostgreSQL)
- **Redis** for session/cache
- **JWT** tokens
- **Zap** structured logging

## gRPC Endpoints

| Method | Description |
|--------|-------------|
| `Register` | Create new user account |
| `Login` | Authenticate, return JWT |
| `ValidateToken` | Verify JWT validity |
| `RefreshToken` | Rotate tokens |
| `GetUser` | Fetch user profile |
| `UpdateUser` | Modify user details |
| `ListUsers` | Admin: list all users |
| `DeleteUser` | Admin: delete user |
| `BanUser` | Admin: ban user |
| `GetAdminStats` | Admin dashboard metrics |

## Quick Start

```bash
export DB_HOST=localhost DB_PORT=5432 DB_USER=zenx DB_PASS=zenx123 DB_NAME=zenxbattle
export REDIS_ADDR=localhost:6379
export JWT_SECRET=your-secret-key

go run cmd/main.go
# → gRPC server on :50051
```

## Database

Uses PostgreSQL with GORM auto-migration. Schema managed by service startup.

## Docker

```bash
docker build -t zenxbattle-auth .
docker run -p 50051:50051 zenxbattle-auth
```

## Related Services

- [ApiGateway](https://github.com/zenxbattle/ApiGateway) — REST proxy to this service
- [Frontend](https://github.com/zenxbattle/Frontend) — login/register UI
- [CommonProto](https://github.com/zenxbattle/CommonProto) — protobuf definitions

## Deploy (K3s)

```bash
kubectl apply -k https://github.com/zenxbattle/infrastructure/tree/k3s/k3s/services/auth
```
