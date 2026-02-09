# Go REST API

A RESTful API built with Go, Fiber, and GORM with PostgreSQL.

## Features

- User CRUD operations
- PostgreSQL database with GORM ORM
- Docker support
- Environment-based configuration

## Prerequisites

- Go 1.21+
- PostgreSQL
- Docker (optional)

## Getting Started

### Environment Variables

Create a `.env` file in the root directory:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=go_rest_api
```

### Run Locally

```bash
# Install dependencies
go mod tidy

# Run the application
go run main.go
```

### Run with Docker

```bash
# Build and run with docker-compose
docker-compose up --build

# Or build manually
docker build -t go-rest-api .
docker run -p 8081:8081 --env-file .env go-rest-api
```

## API Endpoints

### Users

| Method | Endpoint         | Description       |
|--------|------------------|-------------------|
| GET    | /api/users       | Get all users     |
| GET    | /api/users/:id   | Get user by ID    |
| POST   | /api/users       | Create new user   |
| PUT    | /api/users/:id   | Update user       |
| DELETE | /api/users/:id   | Delete user       |

### Example Requests

**Create User**
```bash
curl -X POST http://localhost:8081/api/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com", "password": "secret123"}'
```

**Get All Users**
```bash
curl http://localhost:8081/api/users
```

**Get User by ID**
```bash
curl http://localhost:8081/api/users/1
```

**Update User**
```bash
curl -X PUT http://localhost:8081/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Jane Doe"}'
```

**Delete User**
```bash
curl -X DELETE http://localhost:8081/api/users/1
```

## Project Structure

```
├── main.go
├── src/
│   ├── config/
│   │   ├── connect.config.go
│   │   └── env.config.go
│   ├── models/
│   │   └── user.model.go
│   ├── routes/
│   │   └── user.route.go
│   └── services/
│       └── user.service.go
├── Dockerfile
├── docker-compose.yaml
└── .env
```

## Deployment

### Staging
```bash
./deploy.staging.sh
```

### Production
```bash
./deploy.production.sh
```

## License

See [LICENCE](LICENCE) file.
