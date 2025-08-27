# Environment Variables Setup

Create a `.env` file in the backend directory with the following variables:

```bash
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=your_database_name
DB_SSLMODE=disable

# Redis Configuration
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# Server Configuration
PORT=8080

# CORS Configuration (optional)
# CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001,http://localhost:5173
```

## Required Environment Variables

### Database
- **DB_USER**: Your PostgreSQL username
- **DB_PASSWORD**: Your PostgreSQL password
- **DB_NAME**: Your PostgreSQL database name

### Redis
- No required variables (Redis runs without authentication by default)

## Optional Environment Variables (with defaults)

### Database
- **DB_HOST**: Database host (default: localhost)
- **DB_PORT**: Database port (default: 5432)
- **DB_SSLMODE**: SSL mode (default: disable)

### Redis
- **REDIS_HOST**: Redis host (default: localhost)
- **REDIS_PORT**: Redis port (default: 6379)
- **REDIS_PASSWORD**: Redis password (default: empty - no auth)
- **REDIS_DB**: Redis database number (default: 0)

### Server
- **PORT**: Server port (default: 8080)

### CORS
- **CORS_ALLOWED_ORIGINS**: Comma-separated list of allowed origins for CORS (default: http://localhost:3000,http://localhost:3001,http://localhost:3002,http://localhost:5173)

## Setup Instructions

1. Copy the environment variables above into a new `.env` file in the backend directory
2. Replace the placeholder values with your actual PostgreSQL credentials
3. Make sure your PostgreSQL server is running
4. Make sure your Redis server is running (default: `redis-server` on port 6379)
5. Run the application with `go run cmd/api/main.go`

