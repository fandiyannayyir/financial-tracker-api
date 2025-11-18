# Financial Tracker API

A RESTful API built with Go and Echo framework for managing personal financial tracking. This API provides user management functionality and serves as the backend for a financial tracking application.

## Features

- **User Management**: Complete CRUD operations for user accounts
- **RESTful API**: Clean and intuitive API endpoints
- **Database Integration**: PostgreSQL database with connection pooling
- **Password Security**: Secure password hashing
- **Modular Architecture**: Well-organized code structure with separation of concerns

## Technology Stack

- **Language**: Go 1.24
- **Web Framework**: Echo v4
- **Database**: PostgreSQL
- **Database Driver**: lib/pq
- **Password Hashing**: bcrypt

## Project Structure

```
financial-tracker-api/
├── config/
│   └── db.go              # Database configuration and connection
├── modules/
│   └── users/
│       ├── user_dto.go    # Data Transfer Objects
│       ├── user_handler.go # HTTP request handlers
│       ├── user_model.go   # User data model
│       ├── user_repository.go # Database operations
│       ├── user_route.go   # Route definitions
│       └── user_service.go # Business logic
├── routes/
│   └── route.go           # Main route registration
├── utils/
│   └── hash.go            # Password hashing utilities
├── main.go                # Application entry point
├── go.mod                 # Go module dependencies
└── go.sum                 # Dependency checksums
```

## Prerequisites

Before running this application, make sure you have:

- Go 1.24 or higher installed
- PostgreSQL database server running
- Git for cloning the repository

## Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd financial-tracker-api
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up PostgreSQL database**
   ```sql
   CREATE DATABASE financial_tracker;
   ```

4. **Configure environment variables**
   
   Copy the environment template file:
   ```bash
   cp .env.example .env
   ```
   
   Update the values in `.env` file:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=financial_tracker
   DB_SSL_MODE=disable
   SERVER_PORT=8080
   JWT_SECRET=your_jwt_secret_key
   APP_ENV=development
   ```

5. **Create database tables**
   
   Create the users table:
   ```sql
   CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       name VARCHAR(255) NOT NULL,
       email VARCHAR(255) UNIQUE NOT NULL,
       password VARCHAR(255) NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
   ```

## Running the Application

1. **Start the server**
   ```bash
   go run main.go
   ```

2. **Server will start on port 8080**
   ```
   Server running at http://localhost:8080
   ```

## API Endpoints

### Users

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | `/api/users` | Get all users |
| GET    | `/api/users/:id` | Get user by ID |
| POST   | `/api/users` | Create new user |
| PUT    | `/api/users/:id` | Update user |
| DELETE | `/api/users/:id` | Delete user |

### Example Requests

**Create User**
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepassword"
  }'
```

**Get All Users**
```bash
curl -X GET http://localhost:8080/api/users
```

**Get User by ID**
```bash
curl -X GET http://localhost:8080/api/users/1
```

**Update User**
```bash
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Smith",
    "email": "johnsmith@example.com"
  }'
```

**Delete User**
```bash
curl -X DELETE http://localhost:8080/api/users/1
```

## Development

### Project Architecture

This project follows a clean architecture pattern with clear separation of concerns:

- **Handlers**: Handle HTTP requests and responses
- **Services**: Contain business logic
- **Repositories**: Handle database operations
- **Models**: Define data structures
- **DTOs**: Data transfer objects for API communication

### Adding New Features

1. Create new modules in the `modules/` directory
2. Follow the existing pattern: model, repository, service, handler, routes
3. Register new routes in `routes/route.go`

### Code Style

- Follow Go naming conventions
- Use meaningful variable and function names
- Add comments for exported functions
- Keep functions small and focused

## Environment Variables

This application uses environment variables for configuration. See `ENV_VARIABLES.md` for detailed documentation.

### Quick Setup

1. Copy the template:
   ```bash
   cp .env.example .env
   ```

2. Edit `.env` with your configuration:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_secure_password
   DB_NAME=financial_tracker
   SERVER_PORT=8080
   JWT_SECRET=your_jwt_secret_key
   ```

### Production Environment

For production deployment, set environment variables directly in your deployment environment instead of using the `.env` file.

## Testing

Run tests with:
```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Future Enhancements

- [ ] JWT authentication and authorization
- [ ] Financial transactions management
- [ ] Categories and budgets
- [ ] Income and expense tracking
- [ ] Financial reports and analytics
- [ ] API documentation with Swagger
- [ ] Unit and integration tests
- [ ] Docker containerization
- [ ] CI/CD pipeline

## Support

If you have any questions or issues, please open an issue on GitHub or contact the development team.
