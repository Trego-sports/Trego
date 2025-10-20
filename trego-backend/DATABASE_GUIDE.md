# Database Setup & Access Guide

## Quick Setup 

TO USE IT, add import of "trego-backend/dbtest" in main.go

### 1. Start Database
```bash
# First, make sure Docker is running
open /Applications/Docker.app

# Wait for Docker to start, then:
docker run --name trego-postgres \
  -e POSTGRES_PASSWORD=password123 \
  -e POSTGRES_DB=trego \
  -p 5433:5432 \
  -d postgres:15
```

### Container name conflict
```bash
# If you get "container name already in use" error:
docker stop trego-postgres
docker rm trego-postgres
# Then run step 1 again
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Configure Environment
```bash
cp env.example .env
```

Edit `.env` file:
```env
DB_HOST=127.0.0.1
DB_PORT=5433
DB_USER=postgres
DB_PASSWORD=password123
DB_NAME=trego
DB_SSL_MODE=disable
PORT=8080
GIN_MODE=debug
```

### 4. Start Application
```bash
go run main.go
```

### Docker daemon not running
```bash
# If you get "Cannot connect to Docker daemon" error:
open /Applications/Docker.app
# Wait for Docker to start completely, then try again
```

### Database "trego" does not exist
```bash
# If you get "database trego does not exist" error:
docker exec trego-postgres psql -U postgres -c "CREATE DATABASE trego;"
```

### Tables don't exist
```bash
# If you get "relation does not exist" error:
# Stop the app (Ctrl+C) and restart it:
go run main.go
# This will run migrations and create tables automatically
```


## Access Database

### Command Line Access
```bash
# View all tables
docker exec trego-postgres psql -U postgres -d trego -c "\dt"

# View sports data
docker exec trego-postgres psql -U postgres -d trego -c "SELECT * FROM sports;"
```

### Web Interface Access

#### Start Adminer (Web Database Admin)
```bash
docker run --link trego-postgres:db -p 8081:8080 adminer
```

#### Connect via Browser
1. Open: `http://localhost:8081`
2. Login with:
   - **System:** PostgreSQL
   - **Server:** trego-postgres
   - **Username:** postgres
   - **Password:** password123
   - **Database:** trego

### Test API
```bash
# Health check
curl http://localhost:8080/health

# Ping test
curl http://localhost:8080/api/v1/ping
```

## Database Tables
- `users` - User profiles
- `sports` - Available sports (10 pre-loaded)
- `user_sports` - User-sport relationships
- `games` - Game events
- `game_players` - Game participation
- `schema_migrations` - Migration tracking


