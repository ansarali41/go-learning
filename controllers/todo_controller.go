package controllers

import (
	"database/sql"
	"encoding/json"
	"go-learning/models"
	"net/http"
	"strconv"
)

// TodoController handles all todo-related requests
type TodoController struct {
	DB *sql.DB
}

// HandleTodos processes all /todos requests based on HTTP method
func (tc *TodoController) HandleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tc.GetAllTodos(w, r)
	case http.MethodPost:
		tc.CreateTodo(w, r)
	case http.MethodPut:
		tc.UpdateTodo(w, r)
	case http.MethodDelete:
		tc.DeleteTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// CreateTodo creates a new todo
func (tc *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert into database
	result, err := tc.DB.Exec("INSERT INTO todos (title, description, completed) VALUES (?, ?, ?)",
		todo.Title, todo.Description, todo.Completed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the ID of created todo
	id, _ := result.LastInsertId()
	todo.ID = int(id)

	// Send response
	json.NewEncoder(w).Encode(todo)
}

// GetTodo gets a single todo by ID
func (tc *TodoController) GetTodo(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL query
	id := r.URL.Query().Get("id")
	todoID, _ := strconv.Atoi(id)

	// Get todo from database
	var todo models.Todo
	err := tc.DB.QueryRow("SELECT id, title, description, completed FROM todos WHERE id = ?", todoID).
		Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)

	if err == sql.ErrNoRows {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	json.NewEncoder(w).Encode(todo)
}

// GetAllTodos gets all todos
func (tc *TodoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	// Get all todos from database
	rows, err := tc.DB.Query("SELECT id, title, description, completed FROM todos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Read all todos
	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}

	// Send response
	json.NewEncoder(w).Encode(todos)
}

// UpdateTodo updates a todo by ID
func (tc *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL query
	id := r.URL.Query().Get("id")
	todoID, _ := strconv.Atoi(id)

	// Read request body
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update in database
	_, err = tc.DB.Exec("UPDATE todos SET title = ?, description = ?, completed = ? WHERE id = ?",
		todo.Title, todo.Description, todo.Completed, todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todo.ID = todoID
	json.NewEncoder(w).Encode(todo)
}

// DeleteTodo deletes a todo by ID
func (tc *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL query
	id := r.URL.Query().Get("id")
	todoID, _ := strconv.Atoi(id)

	// Delete from database
	result, err := tc.DB.Exec("DELETE FROM todos WHERE id = ?", todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if todo was deleted
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Send success response
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo deleted successfully"})
}
