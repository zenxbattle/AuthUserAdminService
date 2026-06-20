# auth user admin service

grpc service for user authentication, registration, profile management, and admin operations. uses postgresql for storage, redis for session/token caching, and supports google oauth, email verification, and 2fa.

## env vars

| var | default | description |
|-----|---------|-------------|
| environment | development | runtime environment |
| usergrpcport | 50051 | grpc listen port |
| postgresdsn | - | postgresql connection string |
| jwtsecretkey | - | hmac secret for jwt signing |
| appurl | http://localhost:7000 | public api base url |
| frontendurl | http://localhost:8080 | frontend url for cors |
| redisurl | localhost:6379 | redis address |
| googleclientid | - | google oauth client id |
| googleclientsecret | - | google oauth client secret |
| googleredirecturl | - | google oauth callback |
| smtphost | smtp.gmail.com | smtp server |
| smtpport | 587 | smtp port |
| smtpuser | - | smtp username |
| smtpappkey | - | smtp app password |
| adminpassword | admin | admin user password |
| adminusername | admin | admin username |
| resendapikey | - | resend email api key |

## grpc services

- user registration, login, token refresh, logout
- profile crud with avatar upload
- follow/unfollow system
- two-factor authentication setup and verify
- password management (change, forgot, reset)
- admin user management (create, ban, verify)
- google oauth login flow
