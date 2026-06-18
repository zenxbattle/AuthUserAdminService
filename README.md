# Auth User Admin Service

gRPC service for user authentication, registration, profiles, and admin management.

## Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `ENVIRONMENT` | No | `development` | Runtime environment |
| `USERGRPCPORT` | No | `50051` | gRPC listen port |
| `POSTGRESDSN` | Yes | — | PostgreSQL connection string |
| `JWTSECRETKEY` | Yes | — | HMAC secret for JWT signing |
| `APPURL` | No | `http://localhost:7000` | Public API base URL |
| `FRONTENDURL` | No | `http://localhost:8080` | Frontend URL for CORS |
| `REDISURL` | No | `localhost:6379` | Redis address |
| `GOOGLECLIENTID` | No | — | Google OAuth client ID |
| `GOOGLECLIENTSECRET` | No | — | Google OAuth client secret |
| `GOOGLEREDIRECTURL` | No | — | Google OAuth callback |
| `SMTPHOST` | No | `smtp.gmail.com` | SMTP server |
| `SMTPPORT` | No | `587` | SMTP port |
| `SMTPUSER` | No | — | SMTP username |
| `SMTPAPPKEY` | No | — | SMTP app password |
| `ADMINPASSWORD` | No | `admin` | Admin user password |
| `ADMINUSERNAME` | No | `admin` | Admin username |
| `RESENDAPIKEY` | No | — | Resend email API key |
