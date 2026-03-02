# Car Management System API

A professional RESTful API built with Go (Golang) for managing a car dealership system with comprehensive RBAC (Role-Based Access Control), user management, and inventory tracking.

## 🚀 Features

- **Authentication & Authorization**
  - JWT-based authentication
  - Role-Based Access Control (RBAC)
  - Permission-based endpoint protection

- **User Management**
  - User CRUD operations
  - Role assignment
  - Password hashing with bcrypt

- **Role & Permission System**
  - Dynamic role creation
  - Granular permissions
  - Role-to-user and permission-to-role mapping

- **Car Inventory Management**
  - Complete car CRUD operations
  - Car models, makes, and grades
  - Photo and document management
  - Stock tracking

- **Order & Payment System**
  - Order management
  - Purchase history tracking
  - Payment history and installments
  - Shopping cart functionality

- **API Documentation**
  - Interactive Swagger/OpenAPI documentation
  - Complete request/response schemas
  - Try-it-out functionality

## 📋 Table of Contents

- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Database Setup](#database-setup)
- [Environment Variables](#environment-variables)
- [API Endpoints](#api-endpoints)
- [Authentication](#authentication)
- [Swagger Documentation](#swagger-documentation)
- [Development](#development)

## 🛠 Tech Stack

- **Language:** Go 1.24
- **Web Framework:** Gin
- **Database:** PostgreSQL
- **ORM/Query Builder:** sqlx
- **Authentication:** JWT (golang-jwt/jwt)
- **Password Hashing:** bcrypt
- **Validation:** go-playground/validator
- **API Documentation:** Swaggo
- **Environment Management:** godotenv

## 📁 Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── db/
│   │   └── db.go                # Database connection
│   ├── dto/
│   │   ├── user_dto.go          # User data transfer objects
│   │   ├── role_dto.go          # Role DTOs
│   │   └── permission_dto.go    # Permission DTOs
│   ├── handlers/
│   │   ├── user_handler.go      # User HTTP handlers
│   │   ├── role_handler.go      # Role HTTP handlers
│   │   ├── permission_handler.go # Permission HTTP handlers
│   │   └── car_handler.go       # Car HTTP handlers
│   ├── middleware/
│   │   ├── auth.go              # JWT authentication middleware
│   │   └── permission.go        # Permission checking middleware
│   ├── models/
│   │   ├── user.go              # User model
│   │   ├── role.go              # Role model
│   │   ├── car.go               # Car model
│   │   └── ...                  # Other models
│   ├── repository/
│   │   ├── user_repository.go   # User database operations
│   │   ├── role_repository.go   # Role database operations
│   │   ├── permission_repository.go
│   │   └── car_repository.go    # Car database operations
│   ├── routes/
│   │   └── routes.go            # Route definitions
│   ├── service/
│   │   ├── user_service.go      # User business logic
│   │   ├── role_service.go      # Role business logic
│   │   ├── permission_service.go
│   │   └── car_service.go       # Car business logic
│   └── utils/
│       ├── jwt.go               # JWT utilities
│       ├── password.go          # Password hashing utilities
│       ├── response.go          # Standardized API responses
│       ├── logger.go            # Logging utilities
│       └── errors.go            # Custom error definitions
├── schema/
│   └── schema.sql               # Database schema
├── docs/
│   ├── docs.go                  # Swagger documentation (generated)
│   ├── swagger.json             # Swagger JSON (generated)
│   └── swagger.yaml             # Swagger YAML (generated)
├── .env                         # Environment variables
├── go.mod                       # Go module dependencies
└── README.md                    # This file
```

## 🚦 Getting Started

### Prerequisites

- Go 1.24 or higher
- PostgreSQL 12 or higher
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd first
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Set up the database**
   ```bash
   # Create database
   createdb car_db

   # Run schema
   psql -d car_db -f schema/schema.sql
   ```

5. **Install Swagger CLI (for documentation generation)**
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

6. **Generate Swagger documentation**
   ```bash
   swag init -g cmd/api/main.go
   ```

7. **Run the application**
   ```bash
   go run cmd/api/main.go
   ```

The server will start on `http://localhost:8080`

## 🗄 Database Setup

### Create Database

```sql
CREATE DATABASE car_db;
```

### Run Schema

```bash
psql -U postgres -d car_db -f schema/schema.sql
```

### Database Tables

The system includes the following main tables:

- **Authentication & Authorization**
  - `users` - User accounts
  - `roles` - User roles
  - `permissions` - System permissions
  - `role_user` - User-role mapping
  - `permission_role` - Permission-role mapping

- **Car Management**
  - `car_makes` - Car manufacturers
  - `car_models` - Car models
  - `cars` - Car inventory
  - `car_photos` - Car images
  - `car_grades` - Car condition grades
  - `car_details` - Detailed car information
  - `documents` - Car documents

- **Business Operations**
  - `stocks` - Inventory stock
  - `carts` - Shopping carts
  - `orders` - Customer orders
  - `order_items` - Order line items
  - `purchase_history` - Purchase records
  - `payment_history` - Payment records
  - `installments` - Payment installments

## ⚙️ Environment Variables

Create a `.env` file in the root directory:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=car_db
DB_SSLMODE=disable

# JWT Configuration
JWT_SECRET=your-super-secret-key-change-this-in-production
JWT_EXPIRY_HOURS=24

# Server Configuration
PORT=8080
```

## 📡 API Endpoints

### Base URL
```
http://localhost:8080
```

### Public Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Health check |
| `POST` | `/api/v1/login` | User login (returns JWT token) |
| `GET` | `/swagger/*any` | Swagger documentation UI |

### Protected Endpoints (Require Authentication)

All endpoints below require `Authorization: Bearer <token>` header.

#### User Management (`/api/v1/users`)

| Method | Endpoint | Description | Permission Required |
|--------|----------|-------------|---------------------|
| `POST` | `/api/v1/users` | Create new user | - |
| `GET` | `/api/v1/users` | List all users | - |
| `GET` | `/api/v1/users/:id` | Get user by ID | - |
| `PUT` | `/api/v1/users/:id` | Update user | - |
| `DELETE` | `/api/v1/users/:id` | Delete user | - |

#### Role Management (`/api/v1/roles`)

| Method | Endpoint | Description | Permission Required |
|--------|----------|-------------|---------------------|
| `POST` | `/api/v1/roles` | Create new role | - |
| `GET` | `/api/v1/roles` | List all roles | - |
| `GET` | `/api/v1/roles/:id` | Get role by ID | - |
| `PUT` | `/api/v1/roles/:id` | Update role | - |
| `DELETE` | `/api/v1/roles/:id` | Delete role | - |
| `POST` | `/api/v1/roles/assign` | Assign role to user | - |

#### Permission Management (`/api/v1/permissions`)

| Method | Endpoint | Description | Permission Required |
|--------|----------|-------------|---------------------|
| `POST` | `/api/v1/permissions` | Create new permission | - |
| `GET` | `/api/v1/permissions` | List all permissions | - |
| `GET` | `/api/v1/permissions/:id` | Get permission by ID | - |
| `PUT` | `/api/v1/permissions/:id` | Update permission | - |
| `DELETE` | `/api/v1/permissions/:id` | Delete permission | - |

#### Car Management (`/api/v1/cars`)

| Method | Endpoint | Description | Permission Required |
|--------|----------|-------------|---------------------|
| `POST` | `/api/v1/cars` | Create new car | `car-create` |
| `GET` | `/api/v1/cars` | List all cars | `car-read` |
| `GET` | `/api/v1/cars/:id` | Get car by ID | `car-read` |
| `PUT` | `/api/v1/cars/:id` | Update car | `car-update` |
| `DELETE` | `/api/v1/cars/:id` | Delete car | `car-delete` |

## 🔐 Authentication

### Login

**Request:**
```bash
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "password123"
  }'
```

**Response:**
```json
{
  "message": "login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "name": "Admin User",
      "username": "admin",
      "email": "admin@example.com",
      "is_active": true,
      "last_login_at": "2026-02-08T12:00:00Z",
      "created_at": "2026-02-01T09:00:00Z",
      "updated_at": "2026-02-08T12:00:00Z"
    }
  },
  "track_id": "d6c7f2a4-29c3-4f5e-a8a0-49f8ed8a9a10"
}
```

### Using the Token

Include the token in the `Authorization` header for all protected endpoints:

```bash
curl -X GET http://localhost:8080/api/v1/users \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

## 📚 Swagger Documentation

Interactive API documentation is available via Swagger UI.

### Access Swagger UI

Navigate to: **http://localhost:8080/swagger/index.html**

### Features

- Browse all available endpoints
- View request/response schemas
- Test endpoints directly from the browser
- Download OpenAPI specification

### Regenerate Documentation

After making changes to API annotations:

```bash
swag init -g cmd/api/main.go
```

## 🔧 Development

### Project Architecture

This project follows a **layered architecture**:

1. **Handler Layer** - HTTP request handling and response formatting
2. **Service Layer** - Business logic and validation
3. **Repository Layer** - Database operations
4. **Model Layer** - Data structures

### Adding a New Endpoint

1. **Define the model** in `internal/models/`
2. **Create DTOs** in `internal/dto/`
3. **Implement repository** in `internal/repository/`
4. **Implement service** in `internal/service/`
5. **Create handler** in `internal/handlers/`
6. **Add Swagger annotations** to the handler
7. **Register routes** in `internal/routes/routes.go`
8. **Regenerate Swagger docs**

### Code Style

- Follow Go best practices and conventions
- Use `gofmt` for code formatting
- Add comments for exported functions
- Include Swagger annotations for all endpoints

### Running Tests

```bash
go test ./...
```

## 📝 API Request/Response Examples

### Create User

**Request:**
```bash
POST /api/v1/users
Content-Type: application/json
Authorization: Bearer <token>

{
  "name": "John Doe",
  "username": "johndoe",
  "email": "john@example.com",
  "password": "securepassword123"
}
```

**Response:**
```json
{
  "message": "User created successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "username": "johndoe",
    "email": "john@example.com",
    "is_active": true,
    "created_at": "2026-02-08T12:00:00Z",
    "updated_at": "2026-02-08T12:00:00Z"
  },
  "track_id": "0b9cf5f0-54f7-4f57-95a3-38836f26c512"
}
```

### Create Car

**Request:**
```bash
POST /api/v1/cars
Content-Type: application/json
Authorization: Bearer <token>

{
  "model_id": 1,
  "ref_no": "CAR-2026-001",
  "package": "Premium",
  "body_type": "Sedan",
  "year": 2024,
  "color": "Black",
  "mileage_km": 15000,
  "fuel": "Petrol",
  "transmission": "Automatic",
  "status": "available"
}
```

**Response:**
```json
{
  "message": "Car created successfully",
  "data": {
    "id": 1,
    "model_id": 1,
    "ref_no": "CAR-2026-001",
    "package": "Premium",
    "body_type": "Sedan",
    "year": 2024,
    "color": "Black",
    "reg_year_month": "2024-01",
    "mileage_km": 15000,
    "engine_cc": 1800,
    "fuel": "Petrol",
    "transmission": "Automatic",
    "drive": "FWD",
    "seats": 5,
    "status": "available",
    "created_at": "2026-02-08T12:00:00Z",
    "updated_at": "2026-02-08T12:00:00Z"
  },
  "track_id": "3c1f5f0f-6f47-4c31-9cb5-bf6c6ac0f0e7"
}
```

### List Cars

**Request:**
```bash
GET /api/v1/cars
Authorization: Bearer <token>
```

**Response:**
```json
{
  "message": "Cars fetched successfully",
  "data": [
    {
      "id": 1,
      "model_id": 1,
      "ref_no": "CAR-2026-001",
      "package": "Premium",
      "body_type": "Sedan",
      "year": 2024,
      "color": "Black",
      "reg_year_month": "2024-01",
      "mileage_km": 15000,
      "engine_cc": 1800,
      "fuel": "Petrol",
      "transmission": "Automatic",
      "drive": "FWD",
      "seats": 5,
      "status": "available",
      "created_at": "2026-02-08T12:00:00Z",
      "updated_at": "2026-02-08T12:00:00Z"
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "limit": 1,
    "total_items": 1,
    "links": {
      "self": "/api/v1/cars"
    }
  },
  "limit": 1,
  "track_id": "3b3d0a8f-bcd3-4b0c-b0ef-8d8e3df9a103"
}
```

### Get Car by ID

**Request:**
```bash
GET /api/v1/cars/1
Authorization: Bearer <token>
```

**Response:**
```json
{
  "message": "Car fetched successfully",
  "data": {
    "id": 1,
    "model_id": 1,
    "ref_no": "CAR-2026-001",
    "package": "Premium",
    "body_type": "Sedan",
    "year": 2024,
    "color": "Black",
    "reg_year_month": "2024-01",
    "mileage_km": 15000,
    "engine_cc": 1800,
    "fuel": "Petrol",
    "transmission": "Automatic",
    "drive": "FWD",
    "seats": 5,
    "status": "available",
    "created_at": "2026-02-08T12:00:00Z",
    "updated_at": "2026-02-08T12:00:00Z"
  },
  "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "limit": 1,
    "total_items": 1,
    "links": {
      "self": "/api/v1/cars/1"
    }
  },
  "limit": 1,
  "track_id": "1a0e3b0c-5b7a-4f55-9c7c-0ebfd1d3d244"
}
```

### Update Car

**Request:**
```bash
PUT /api/v1/cars/1
Content-Type: application/json
Authorization: Bearer <token>

{
  "model_id": 1,
  "ref_no": "CAR-2026-001",
  "package": "Premium",
  "body_type": "Sedan",
  "year": 2024,
  "color": "Black",
  "reg_year_month": "2024-01",
  "mileage_km": 15000,
  "engine_cc": 1800,
  "fuel": "Petrol",
  "transmission": "Automatic",
  "drive": "FWD",
  "seats": 5,
  "status": "available"
}
```

**Response:**
```json
{
  "message": "Car updated successfully",
  "data": null,
  "track_id": "f40d4f0a-4ad8-44c1-9c62-4e1d2484cd83"
}
```

### Delete Car

**Request:**
```bash
DELETE /api/v1/cars/1
Authorization: Bearer <token>
```

**Response:**
```json
{
  "message": "Car deleted successfully",
  "data": null,
  "track_id": "b75db21d-c8c2-41b5-9a77-9c1f9d3c00f3"
}
```

### Assign Role to User

**Request:**
```bash
POST /api/v1/roles/assign
Content-Type: application/json
Authorization: Bearer <token>

{
  "user_id": 1,
  "role_id": 2
}
```

**Response:**
```json
{
  "message": "Role assigned successfully",
  "data": null,
  "track_id": "6f0523e1-3a88-4b72-8b97-2e9501f8b9ad"
}
```

## 🔒 Security Features

- **Password Hashing:** bcrypt with cost factor 10
- **JWT Authentication:** Secure token-based authentication
- **RBAC:** Fine-grained access control
- **SQL Injection Protection:** Parameterized queries via sqlx
- **Input Validation:** Request validation using validator/v10
- **CORS:** Configurable cross-origin resource sharing

## 🐛 Error Handling

The API uses standardized error responses:

```json
{
  "message": "User not found",
  "error": "record not found",
  "track_id": "a01e71c4-3fe8-4f71-9c53-c2427fb6756b"
}
```

**Common HTTP Status Codes:**
- `200` - Success
- `201` - Created
- `400` - Bad Request (validation errors)
- `401` - Unauthorized (missing/invalid token)
- `403` - Forbidden (insufficient permissions)
- `404` - Not Found
- `409` - Conflict (duplicate resource)
- `500` - Internal Server Error

## 📦 Dependencies

Key dependencies (see `go.mod` for complete list):

- `github.com/gin-gonic/gin` - Web framework
- `github.com/jmoiron/sqlx` - SQL toolkit
- `github.com/lib/pq` - PostgreSQL driver
- `github.com/golang-jwt/jwt/v5` - JWT implementation
- `github.com/go-playground/validator/v10` - Input validation
- `github.com/swaggo/gin-swagger` - Swagger integration
- `github.com/joho/godotenv` - Environment variable management
- `golang.org/x/crypto` - Cryptography (bcrypt)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License.

## 👥 Authors

- Your Name - Initial work

## 🙏 Acknowledgments

- Gin framework documentation
- Go community
- PostgreSQL documentation
- Swagger/OpenAPI specification

---

**Built with ❤️ using Go**
