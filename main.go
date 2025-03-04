package main

import (
	"fmt"
	"go-learning/config"
	"go-learning/controllers"
	"log"
	"net/http"
)

func main() {
	// Step 1: Connect to MySQL database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	// Step 2: Create a new todo controller
	todoController := &controllers.TodoController{DB: db}

	// Step 3: Setup API routes
	// POST /todos - Create a new todo
	// GET /todos - Get all todos
	// GET /todo?id=1 - Get a single todo
	// PUT /todos?id=1 - Update a todo
	// DELETE /todos?id=1 - Delete a todo
	http.HandleFunc("/todos", todoController.HandleTodos)
	http.HandleFunc("/todo", todoController.GetTodo)

	// Step 4: Start the server
	fmt.Println("Server starting at http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
