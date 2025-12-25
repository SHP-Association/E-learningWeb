# E-Learning Backend (Go)

Go backend implementation for the E-Learning platform using modern Go practices and the Pagoda framework pattern.

## Features

- ✅ RESTful API with JWT authentication
- ✅ PostgreSQL database integration with GORM
- ✅ Shares database with Django backend
- ✅ CORS support for frontend integration
- ✅ Structured logging
- ✅ Role-based access control
- ✅ Pagination and filtering
- ✅ Docker support

## Project Structure

```
backend-go/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── models/
│   │   ├── user.go              # User model
│   │   └── models.go            # Other models
│   ├── store/
│   │   ├── store.go             # Store interface
│   │   └── postgres/
│   │       └── postgres.go      # PostgreSQL implementation
│   └── api/
│       ├── api.go               # API setup and routes
│       ├── auth.go              # Authentication handlers
│       ├── middleware.go        # Middleware functions
│       ├── courses.go           # Course handlers
│       └── handlers.go          # Other handlers
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 15
- Docker and Docker Compose (optional)

### Local Development

1. **Install dependencies:**
   ```bash
   go mod download
   ```

2. **Set up environment variables:**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Run the server:**
   ```bash
   go run cmd/server/main.go
   ```

The server will start on `http://localhost:8002`

### Docker Development

1. **Build and run with Docker Compose:**
   ```bash
   # From the project root
   docker compose -f docker-compose.dev.yml up backend-go
   ```

2. **Access the API:**
   - Health check: http://localhost:8002/api/health
   - API endpoints: http://localhost:8002/api/

## API Endpoints

### Authentication
- `POST /api/login` - User login
- `POST /api/register` - User registration
- `POST /api/logout` - User logout

### Users
- `GET /api/users` - List users
- `POST /api/users` - Create user
- `GET /api/users/{id}` - Get user
- `PUT /api/users/{id}` - Update user
- `DELETE /api/users/{id}` - Delete user

### Categories
- `GET /api/categories` - List categories
- `POST /api/categories` - Create category
- `GET /api/categories/{id}` - Get category
- `PUT /api/categories/{id}` - Update category
- `DELETE /api/categories/{id}` - Delete category

### Courses
- `GET /api/courses` - List courses
- `POST /api/courses` - Create course
- `GET /api/courses/{id}` - Get course
- `PUT /api/courses/{id}` - Update course
- `DELETE /api/courses/{id}` - Delete course

### Lessons, Enrollments, Quizzes, FAQs
Similar CRUD endpoints available for each resource.

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `GO_PORT` | Server port | `8002` |
| `GO_ENV` | Environment (development/production) | `development` |
| `GO_JWT_SECRET` | JWT signing secret | - |
| `GO_JWT_EXPIRY` | JWT token expiry duration | `24h` |
| `GO_DB_HOST` | Database host | `localhost` |
| `GO_DB_PORT` | Database port | `5432` |
| `GO_DB_NAME` | Database name | `shplearner` |
| `GO_DB_USER` | Database user | - |
| `GO_DB_PASSWORD` | Database password | - |
| `GO_CORS_ALLOWED_ORIGINS` | Allowed CORS origins | `http://localhost:5173` |

## Database

The Go backend uses the same PostgreSQL database as the Django backend. It connects to existing tables created by Django migrations.

### Table Mapping

| Django Model | Go Model | Table Name |
|--------------|----------|------------|
| CustomUser | User | `Account_customuser` |
| Category | Category | `Category_category` |
| Course | Course | `courses_course` |
| Lesson | Lesson | `Lesson_lesson` |
| Enrollment | Enrollment | `Enrollment_enrollment` |
| Quiz | Quiz | `Quiz_quiz` |
| FAQ | FAQ | `FAQ_faq` |

## Development

### Running Tests
```bash
go test ./...
```

### Building
```bash
go build -o bin/server cmd/server/main.go
```

### Code Formatting
```bash
go fmt ./...
```

## Deployment

### Docker Build
```bash
docker build -t e-learning-backend-go .
```

### Docker Run
```bash
docker run -p 8002:8002 --env-file .env e-learning-backend-go
```

## Architecture

This backend follows the faptap project structure:
- **cmd/**: Application entry points
- **internal/**: Private application code
  - **config/**: Configuration management
  - **models/**: Data models
  - **store/**: Data access layer (interface + implementations)
  - **api/**: HTTP handlers and routing

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.
