# GAuth Web Application

A comprehensive web application for GAuth (Gimel Authentication) RFC-0150 implementation, providing a modern interface for authentication and authorization management.

## Features

- 🔐 **Authentication Management**: Login, logout, token management
- 👥 **User Management**: User registration, profile management, role assignment
- 🛡️ **Authorization Policies**: Create, manage, and test authorization policies
- 📊 **Dashboard**: Real-time monitoring and analytics
- 🔧 **Admin Panel**: System administration and configuration
- 📱 **Responsive Design**: Mobile-friendly interface
- 🚀 **Modern Tech Stack**: React/Next.js frontend with Go backend

## Architecture

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────────┐
│   Frontend      │    │   Backend API    │    │   GAuth Core        │
│   (Next.js)     │◄──►│   (Go/Gin)       │◄──►│   (RFC-0150)        │
│                 │    │                  │    │                     │
│ • Dashboard     │    │ • REST API       │    │ • Authentication    │
│ • User Mgmt     │    │ • WebSocket      │    │ • Authorization     │
│ • Policy Mgmt   │    │ • Middleware     │    │ • Token Management  │
│ • Analytics     │    │ • Auth Guard     │    │ • Policy Engine     │
└─────────────────┘    └──────────────────┘    └─────────────────────┘
```

## Tech Stack

### Frontend
- **Framework**: Next.js 14 with App Router
- **Language**: TypeScript
- **Styling**: Tailwind CSS
- **Components**: Radix UI + shadcn/ui
- **State Management**: Zustand
- **Forms**: React Hook Form + Zod validation
- **HTTP Client**: Axios
- **Real-time**: Socket.IO client

### Backend
- **Framework**: Go with Gin
- **Database**: PostgreSQL
- **Cache**: Redis
- **Authentication**: JWT tokens
- **Real-time**: WebSocket/Socket.IO
- **Documentation**: Swagger/OpenAPI
- **Testing**: Go testing + testify

### Infrastructure
- **Containerization**: Docker & Docker Compose
- **Reverse Proxy**: Nginx
- **Monitoring**: Prometheus + Grafana
- **Logging**: Structured logging with logrus

## Quick Start

### Prerequisites
- Node.js 18+ and npm/yarn
- Go 1.21+
- Docker & Docker Compose
- PostgreSQL (or use Docker)
- Redis (or use Docker)

### Development Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/mauriciomferz/Gauth_web_app_001.git
   cd Gauth_web_app_001
   ```

2. **Start with Docker Compose** (recommended):
   ```bash
   docker-compose up -d
   ```
   This starts all services: frontend, backend, database, and cache.

3. **Or run manually**:
   
   **Backend**:
   ```bash
   cd backend
   cp .env.example .env  # Configure your environment
   go mod download
   go run cmd/server/main.go
   ```
   
   **Frontend**:
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

4. **Access the application**:
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080
   - API Documentation: http://localhost:8080/swagger

## Project Structure

```
├── frontend/                # Next.js frontend application
│   ├── app/                # App router pages
│   ├── components/         # React components
│   ├── lib/               # Utilities and configurations
│   ├── hooks/             # Custom React hooks
│   └── types/             # TypeScript type definitions
├── backend/               # Go backend application
│   ├── cmd/               # Application entry points
│   ├── internal/          # Private application code
│   ├── pkg/               # Public packages
│   ├── api/               # API handlers and routes
│   └── docs/              # API documentation
├── shared/                # Shared resources
│   ├── types/             # Shared type definitions
│   └── docs/              # Documentation
├── infrastructure/        # Infrastructure as code
│   ├── docker/            # Docker configurations
│   └── k8s/               # Kubernetes manifests
└── docs/                  # Project documentation
```

## API Documentation

The API is fully documented using OpenAPI/Swagger. Once the backend is running, you can access the interactive documentation at:
- **Swagger UI**: http://localhost:8080/swagger
- **OpenAPI Spec**: http://localhost:8080/swagger.json

## Environment Configuration

### Backend (.env)
```env
# Server
PORT=8080
GIN_MODE=release

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=gauth
DB_PASSWORD=gauth_password
DB_NAME=gauth_db

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# JWT
JWT_SECRET=your_jwt_secret_key
JWT_EXPIRY=24h

# GAuth Core
GAUTH_SERVER_URL=http://localhost:9090
```

### Frontend (.env.local)
```env
# API
NEXT_PUBLIC_API_URL=http://localhost:8080
NEXT_PUBLIC_WS_URL=ws://localhost:8080

# Authentication
NEXTAUTH_URL=http://localhost:3000
NEXTAUTH_SECRET=your_nextauth_secret
```

## Features Overview

### Authentication
- User registration and login
- JWT token management
- Multi-factor authentication (MFA)
- Password reset and recovery
- Session management

### Authorization
- Role-based access control (RBAC)
- Policy-based authorization
- Dynamic permission management
- Resource-level access control

### User Management
- User profiles and preferences
- Role assignment and management
- User activity tracking
- Bulk operations

### Policy Management
- Visual policy builder
- Policy testing and validation
- Policy templates and examples
- Import/export functionality

### Analytics & Monitoring
- Real-time dashboards
- Authentication metrics
- Authorization analytics
- System health monitoring
- Audit logs and compliance

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Testing

### Backend Testing
```bash
cd backend
go test -v ./...
go test -race ./...
go test -cover ./...
```

### Frontend Testing
```bash
cd frontend
npm test
npm run test:e2e
npm run test:coverage
```

## Deployment

### Docker Deployment
```bash
docker-compose -f docker-compose.prod.yml up -d
```

### Kubernetes Deployment
```bash
kubectl apply -f infrastructure/k8s/
```

## Security

- JWT token authentication
- HTTPS/TLS encryption
- Input validation and sanitization
- SQL injection prevention
- XSS protection
- CORS configuration
- Rate limiting
- Security headers

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For support and questions:
- Create an issue in this repository
- Email: support@gimel.foundation
- Documentation: [docs/](docs/)

## Roadmap

- [ ] Multi-tenant support
- [ ] SSO integration (SAML, OAuth2, OIDC)
- [ ] Mobile applications
- [ ] Advanced analytics
- [ ] Plugin system
- [ ] API rate limiting dashboard
- [ ] Compliance reporting