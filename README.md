# xcode Auth-User-Admin Service

A robust gRPC-based user authentication and management service built with Go, featuring comprehensive user operations, OAuth integration, and advanced security features.

## Features

### Authentication & Authorization
- **Email/Password Authentication** - Secure user registration and login
- **Google OAuth Integration** - Seamless social login
- **JWT Token Management** - Access and refresh token handling
- **Two-Factor Authentication (2FA)** - TOTP-based security
- **Admin Authentication** - Separate admin login system

### User Management
- **User Registration** - Email verification with OTP
- **Profile Management** - Update user information and avatar
- **Password Recovery** - Forgot password with email tokens
- **Account Verification** - Email-based account activation
- **Ban System** - User banning with expiration support

### Social Features
- **Follow/Unfollow System** - User relationship management
- **Social Links** - GitHub, Twitter, LinkedIn integration
- **User Profiles** - Comprehensive user information display

### Security & Monitoring
- **Password Hashing** - bcrypt with salt
- **Input Validation** - Email and password format validation
- **Rate Limiting** - Redis-based caching
- **Structured Logging** - Zap logger with BetterStack integration
- **Error Handling** - Comprehensive error types and gRPC status codes

## Tech Stack

- **Language**: Go 1.24.1
- **Framework**: gRPC
- **Database**: PostgreSQL with GORM
- **Cache**: Redis
- **Authentication**: JWT, Google OAuth2
- **Logging**: Zap with BetterStack
- **Email**: Resend API
- **2FA**: TOTP (Time-based One-Time Password)
- **Containerization**: Docker & Docker Compose

## Prerequisites

- Go 1.24.1 or higher
- PostgreSQL 15+
- Redis 7+
- Docker & Docker Compose (optional)

## Quick Start

### 1. Clone the Repository
```bash
git clone <repository-url>
cd UserService
```

### 2. Environment Setup
Create a `.env` file in the root directory:
```env
# Server Configuration
ENVIRONMENT=development
USERGRPCPORT=50051

# Database
POSTGRESDSN=host=localhost port=5432 user=admin password=password dbname=xcodedev sslmode=disable

# Authentication
JWTSECRETKEY=your-jwt-secret-key
ADMINPASSWORD=your-admin-password
ADMINUSERNAME=admin

# URLs
APPURL=http://localhost:7000
FRONTENDURL=http://localhost:8080

# Google OAuth
GOOGLECLIENTID=your-google-client-id
GOOGLECLIENTSECRET=your-google-client-secret
GOOGLEREDIRECTURL=http://localhost:8080/auth/google/callback

# Redis
REDISURL=localhost:6379

# Email Service
RESENDAPIKEY=your-resend-api-key
SMTPHOST=smtp.gmail.com
SMTPPORT=587
SMTPUSER=your-email@gmail.com
SMTPAPPKEY=your-smtp-app-key

# Logging
BETTERSTACKSOURCETOKEN=your-betterstack-token
BETTERSTACKUPLOADURL=https://in.logs.betterstack.com
```

### 3. Start Dependencies
```bash
# Using Docker Compose
docker-compose up -d

# Or start services individually
# PostgreSQL, Redis, etc.
```

### 4. Run the Service
```bash
# Install dependencies
go mod tidy

# Run the service
go run cmd/main.go
```

The gRPC server will start on port `50051` (or your configured port).

## Project Structure

```
.
├── cache/                  # Redis cache implementation
├── cmd/                    # Application entry point
├── configs/                # Configuration management
├── customerrors/           # Custom error
├── db/                     # Database models and connection
├── logger/                 # Logging utilities
├── reposit           # Data access layer
├── service/                # Business logic layer
├── utils/                  # Utility functions (JWT, SMTP)
├── docker-compose.yml      # Docker services
├── Dockerfile             # Container configuration
└── README.md              # This file
```

## Configuration

### Envirt Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `ENVIRONMENT` | Application environment | `development` |
| `USERGRPCPORT` | gRPC server port | `50051` |
| `POSTGRESDSN` | PostgreSQL connection string | - |
| `JWTSECRETKEY` | JWT signing secret | - |
| `REDISURL` | Redis connection URL | `localhost:6379` |
| `GOOGLECLIENTID` | Google OAuth client ID | - |
| `RESENDAPIKEY` | Resend email service API key | - |

## API Endpoints

### Authentication
- `RegisterUser` - User registration with email verification
- `LoginUser` - Email/password authentication
- `LoginWithGoogle` - Google OAuth authentication
- `LoginAdmin` - Admin authentication
- `TokenRefresh` - Refresh access tokens
- `LogoutUser` - User logout

### User Management
- `VerifyUser` - Email verification with OTP
- `ResendEmailVerification` - Resend verification email
- `ForgotPassword` - Initiate password recovery
- `FinishForgotPassword` - Complete password reset
- `UpdateProfile` - Update user profile
- `UpdateProfileImage` - Update user avatar
- `GetUserProfile` - Retrieve user profile

### Social Features
- `FollowUser` - Follow another user
- `UnfollowUser` - Unfollow a user
- `GetFollowing` - Get users being followed
- `GetFollowers` - Get user's followers

### Security & Admin
- `CheckBanStatus` - Check if user is banned
- `EnableTwoFactorAuth` - Enable 2FA
- `DisableTwoFactorAuth` - Disable 2FA
- `ValidateTwoFactorAuth` - Validate 2FA token


## Docker Deployment

### Using Docker Compose
```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Building Docker Image
```bash
# Build image
docker build -t user-service .

# Run container
docker run -p 50051:50051 --env-file .env user-service
```

## Monitoring & Logging

The service includes comprehensive logging with:
- **Structured Logging**: JSON-formatted logs with Zap
- **BetterStack Integration**: Centralized log management
- **Request Tracing**: Trace ID for request correlation
- **Error Tracking**: Detailed error logging with context

## Security Features

- **Password Security**: bcrypt hashing with unique salts
- **JWT Security**: Secure token generation and validation
- **Input Validation**: Email format and password strength validation
- **Rate Limiting**: Redis-based request throttling
- **2FA Support**: TOTP-based two-factor authentication
- **Account Security**: Ban system and verification requirements

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -m 'Add some new feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For support and questions:
- Create an issue in the repository
- Check the documentation
- Review the error logs for troubleshooting
---