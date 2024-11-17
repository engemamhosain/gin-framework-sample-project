# Gin Framework Sample Project

A demonstration of best practices for building RESTful APIs using the **Gin Web Framework** and **GORM**.

---

## Project Overview

This project implements:
- A simple RESTful API.
- Integration with **MySQL** and **MongoDB** databases.
- Configuration management using YAML files.
- Basic authentication and user management.

---

## Features

- **Configuration Loading**: Supports environment-based configuration (e.g., `config.dev.yaml` for development).
- **Database Integration**:
  - MySQL: User management with predefined schema.
  - MongoDB: Additional NoSQL database support.
- **API Endpoints**:
  - User operations (`GET`, `POST`).
  - Authentication (`POST`).
- **Testing**: Sample tests with the `go test` command.

---

## Configuration

### Set Up Config

The configuration file is managed using **YAML** and supports environment-specific configurations.

1. Set `config.yaml`:
    ```yaml
    environment: dev
    ```
2. The system will load from `config.dev.yaml` based on the `environment` key.
3. Retrieve configuration values programmatically:
    ```go
    config.Config.GetString("app.url")
    ```

---

## Database Setup

### MySQL

Ensure the following table exists in your MySQL database:

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

## API Endpoints
This project provides the following RESTful API endpoints:

### 1. User API
- Retrieve User by ID
 
    Retrieve details of a specific user using their ID:
    ```
        curl http://localhost:8088/api/v1/user/{id}
    ```
    Example:

    ```
    curl http://localhost:8088/api/v1/user/1
    ```

- Create a New User
    ```
    curl -X POST -H 'Content-Type: application/json' -d '{
        "username": "john_doe",
        "password": "pass",
        "email": "john@example.com"
    }' http://localhost:8088/api/v1/user/
    
    ```

### 2. Authentication API

- Authenticate User
 
    Verify user credentials and generate an authentication token:

    ```
    curl -X POST -H 'Content-Type: application/json' -d '{
        "username": "john_doe",
        "password": "pass"
    }' http://localhost:8088/api/v1/auth/
    ```

## Testing


Run the test cases for this project using the following command:

```
go test test/*.go -v

```
# Project Structure

The following is the folder and file structure of the project:

```
├── config/ # Configuration files (e.g., config.yaml, config.dev.yaml) 
│ 
├── config.yaml                 # Main configuration file 
│ └── config.dev.yaml           # Development-specific configurations 
├── controllers/                # API controllers 
│ ├── user_controller.go        # Handles user-related API logic 
│ └── auth_controller.go        # Handles authentication API logic 
├── models/                     # Database models 
│ └── user.go                   # User model definition 
├── routes/                     # API route definitions 
│ ├── user_routes.go            # Routes for user endpoints 
│ └── auth_routes.go            # Routes for authentication endpoints 
├── services/                   # Service layer (business logic) 
│ ├── user_service.go           # Business logic for user management 
│ └── auth_service.go           # Business logic for authentication 
├── test/                       # Test files for unit and integration tests 
│ ├── user_test.go              # Tests for user API 
│ └── auth_test.go              # Tests for authentication API 
├── main.go                     # Application entry point 
└── README.md                   # Project documentation

```

