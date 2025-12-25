# SHP-Learner E-Learning Platform

A modern e-learning platform built with Go (Pagoda framework) backend, Vue.js frontend, and PostgreSQL database.

## 🚀 Tech Stack

- **Backend**: Go 1.21+ with Pagoda framework
- **Frontend**: Vue.js 3 with Vite
- **Database**: PostgreSQL 14 with TimescaleDB extension
- **Containerization**: Docker & Docker Compose

## 📋 Features

- 🔐 JWT-based authentication
- 👥 User management and profiles
- 📚 Course management system
- 📝 Lesson and content delivery
- 📜 Certificate generation
- ❓ FAQ system
- 📊 Progress tracking
- 🎓 Enrollment management

## 🛠️ Setup Instructions

### Prerequisites

- Go 1.21 or higher
- Node.js 18+ and npm
- Docker and Docker Compose
- Git

### Development Setup (Local Backend & Frontend)

This setup runs only PostgreSQL in Docker while backend and frontend run on your host machine.

1. **Clone the repository**
   ```bash
   git clone https://github.com/SHP-Association/E-learningWeb.git
   cd E-learningWeb
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env and update values as needed
   ```

3. **Start PostgreSQL**
   ```bash
   docker-compose -f docker-compose.dev.yml up -d
   ```

4. **Run Go Backend**
   ```bash
   cd backend-go/cmd/server
   GO_DB_HOST=localhost go run . serve
   ```
   Backend will be available at `http://localhost:8002`

5. **Run Vue.js Frontend**
   ```bash
   cd web
   npm install
   npm run dev
   ```
   Frontend will be available at `http://localhost:5173`

### Production Setup (All Services in Docker)

This setup runs all services (PostgreSQL, Go backend, Vue frontend) in Docker containers.

1. **Clone and configure**
   ```bash
   git clone https://github.com/SHP-Association/E-learningWeb.git
   cd E-learningWeb
   cp .env.example .env
   # Edit .env and update production values
   ```

2. **Start all services**
   ```bash
   docker-compose -f docker-compose.prod.yml up -d
   ```

3. **Access the application**
   - Frontend: `http://localhost:5173`
   - Backend API: `http://localhost:8002`

### Stopping Services

**Development:**
```bash
docker-compose -f docker-compose.dev.yml down
```

**Production:**
```bash
docker-compose -f docker-compose.prod.yml down
```

## 📁 Project Structure

```
E-learningWeb/
├── backend-go/           # Go backend (Pagoda framework)
│   ├── cmd/
│   │   └── server/       # Main application entry point
│   ├── internal/
│   │   ├── api/          # API handlers and routes
│   │   ├── config/       # Configuration management
│   │   ├── database/     # Database connection
│   │   ├── models/       # Data models
│   │   └── store/        # Data access layer
│   └── Dockerfile
├── web/                  # Vue.js frontend
│   ├── src/
│   │   ├── components/   # Vue components
│   │   ├── views/        # Page views
│   │   ├── router/       # Vue Router
│   │   └── stores/       # Pinia stores
│   └── Dockerfile
├── docker-compose.dev.yml   # Development (postgres only)
├── docker-compose.prod.yml  # Production (all services)
├── .env.example             # Environment variables template
└── README.md
```

## 🔧 Environment Variables

Key environment variables (see `.env.example` for full list):

| Variable | Description | Default |
|----------|-------------|---------|
| `GO_DB_HOST` | Database host (`localhost` for local, `postgres` for Docker) | `postgres` |
| `GO_DB_NAME` | Database name | `shplearner` |
| `GO_DB_USER` | Database user | `sandesh` |
| `GO_DB_PASSWORD` | Database password | `password123` |
| `GO_PORT` | Backend server port | `8002` |
| `GO_JWT_SECRET` | JWT signing secret | Change in production! |
| `VITE_PORT` | Frontend dev server port | `5173` |
| `VITE_API_URL` | Backend API URL | `http://localhost:8002` |

## 🧪 Development

### Running Backend Tests
```bash
cd backend-go
go test ./...
```

### Building for Production
```bash
# Backend
cd backend-go
go build -o server ./cmd/server

# Frontend
cd web
npm run build
```

## 📝 API Documentation

The Go backend provides a RESTful API. Key endpoints:

- `POST /api/auth/register` - User registration
- `POST /api/auth/login` - User login
- `GET /api/courses` - List courses
- `GET /api/courses/:id` - Get course details
- `POST /api/enrollments` - Enroll in course

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📧 Contact

**Sandesh Patel**
- Email: sandeshpatel.sp.93@gmail.com
- Phone: +91 9399613606
- Twitter: [@SandeshPat007](https://x.com/SandeshPat007)
- LinkedIn: [Sandesh Patel](https://www.linkedin.com/in/sandesh-patel07)
- Instagram: [@sandesh_patel007](https://www.instagram.com/sandesh_patel007)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with [Pagoda](https://github.com/mikestefanello/pagoda) - Go web framework
- UI powered by [Vue.js](https://vuejs.org/)
- Database: [TimescaleDB](https://www.timescale.com/)
