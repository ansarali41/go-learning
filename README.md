# Go Todo Application

## Overview

A simple, lightweight Todo application built with Go and MySQL, designed to demonstrate web development fundamentals and RESTful API design.

## 🚀 Features

-   Create new todo items
-   Read all todos
-   Read a specific todo by ID
-   Update existing todos
-   Delete todos
-   Persistent storage with MySQL

## 📋 Prerequisites

-   Go (version 1.20+)
-   MySQL (version 8.0+)
-   Postman (optional, for API testing)

## 🔧 Technology Stack

-   Backend: Go (Golang)
-   Database: MySQL
-   Architecture: RESTful API, MVC Pattern

## 📦 Project Structure

```
go-learning/
│
├── config/
│   └── db.go          # Database configuration
│
├── controllers/
│   └── todo_controller.go  # HTTP request handlers
│
├── models/
│   └── todo.go        # Data structures
│
├── main.go            # Application entry point
├── schema.sql         # Database schema
└── go.mod             # Dependency management
```

## 🛠️ Setup Instructions

### 1. Database Setup

1. Install MySQL
2. Create the database:
    ```bash
    mysql -u root -p < schema.sql
    ```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run the Application

```bash
go run main.go
```

## 🌐 API Endpoints

-   `POST /todos`: Create a new todo
-   `GET /todos`: List all todos
-   `GET /todo?id=X`: Get a specific todo
-   `PUT /todos?id=X`: Update a todo
-   `DELETE /todos?id=X`: Delete a todo

## 📝 Example Payload

```json
{
    "title": "Learn Go Programming",
    "description": "Study Go fundamentals and build projects",
    "completed": false
}
```

## 🧪 Testing

Use the included Postman collection (`Todo_App.postman_collection.json`) to test the API.

## 🔒 Configuration

Update database credentials in `config/db.go`

## 🚧 Future Improvements

-   Add user authentication
-   Implement input validation
-   Create more robust error handling
-   Add logging
-   Implement database migrations

## 📄 License

MIT License

## 👨‍💻 Author

Ansar Ali Sarkar
