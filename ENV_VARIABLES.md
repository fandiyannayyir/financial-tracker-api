# Environment Variables Configuration

This document explains the environment variables used in the Financial Tracker API.

## Setup

1. Copy the `.env.example` file to `.env`:
   ```bash
   cp .env.example .env
   ```

2. Update the values in `.env` according to your environment.

## Environment Variables

### Database Configuration

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `DB_HOST` | PostgreSQL database host | `localhost` | No |
| `DB_PORT` | PostgreSQL database port | `5432` | No |
| `DB_USER` | PostgreSQL database username | `postgres` | No |
| `DB_PASSWORD` | PostgreSQL database password | *(empty)* | Yes |
| `DB_NAME` | PostgreSQL database name | `financial_tracker` | No |
| `DB_SSL_MODE` | PostgreSQL SSL mode | `disable` | No |

### Server Configuration

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `SERVER_PORT` | Server port | `8080` | No |

### JWT Configuration

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `JWT_SECRET` | JWT secret key for authentication | *(none)* | Yes* |

*Required when implementing authentication features

### Application Configuration

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `APP_ENV` | Application environment (development/production) | `development` | No |

## Example Configuration

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_secure_password
DB_NAME=financial_tracker
DB_SSL_MODE=disable

# Server Configuration
SERVER_PORT=8080

# JWT Configuration
JWT_SECRET=your_very_secure_jwt_secret_key

# Application Environment
APP_ENV=development
```

## Notes

- The `.env` file is ignored by git to prevent sensitive information from being committed to the repository.
- Always use strong passwords and secret keys in production.
- The application will fall back to default values if environment variables are not set (except for required ones).
