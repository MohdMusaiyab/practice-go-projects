package handlers

import (
	"encoding/json"
	"go-postgress/config"
	"go-postgress/models"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode request body into a Todo struct (without ID & CreatedAt)
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Assign ID and CreatedAt
	todo.Id = uuid.New().String()
	todo.CreatedAt = time.Now().Format(time.RFC3339)

	// Default status if not provided
	if todo.Status == "" {
		todo.Status = models.StatusPending
	}

	// Insert into DB
	query := `INSERT INTO todos (id, title, description, status, created_at)
	          VALUES ($1, $2, $3, $4, $5)`

	_, err = config.DB.Exec(query, todo.Id, todo.Title, todo.Description, todo.Status, todo.CreatedAt)
	if err != nil {
		http.Error(w, "Failed to insert todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return created Todo as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func GetAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := config.DB.Query(`SELECT id, title, description, status, created_at FROM todos`)
	if err != nil {
		http.Error(w, "Failed to fetch todos: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt)
		if err != nil {
			http.Error(w, "Error scanning todos: "+err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	// Handle empty list case
	if len(todos) == 0 {
		todos = []models.Todo{} // return empty array, not null
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func GetToDoById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract 'id' from route variables
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	var todo models.Todo
	err := config.DB.QueryRow(`SELECT id, title, description, status, created_at FROM todos WHERE id=$1`, id).
		Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt)
	if err != nil {
		http.Error(w, "Todo not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
func UpdateToDoByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT method allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract 'id' from path parameters
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	// Decode updated fields from request body
	var updatedTodo models.Todo
	err := json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Update the todo in DB
	query := `UPDATE todos 
	          SET title=$1, description=$2, status=$3
	          WHERE id=$4`
	_, err = config.DB.Exec(query, updatedTodo.Title, updatedTodo.Description, updatedTodo.Status, id)
	if err != nil {
		http.Error(w, "Failed to update todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the updated todo
	updatedTodo.Id = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTodo)
}

func DeleteToDoByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE method allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract 'id' from path parameters
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	// Delete the todo from DB
	query := `DELETE FROM todos WHERE id=$1`
	result, err := config.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to delete todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error checking deletion: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

func UpdateToDoStatusByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Only PATCH method allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract 'id' from path parameters
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	// Decode new status from request body
	var payload struct {
		Status models.Status `json:"status"`
	}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Update the status in DB
	query := `UPDATE todos SET status=$1 WHERE id=$2`
	result, err := config.DB.Exec(query, payload.Status, id)
	if err != nil {
		http.Error(w, "Failed to update status: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error checking update: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Return updated todo
	var updatedTodo models.Todo
	err = config.DB.QueryRow(`SELECT id, title, description, status, created_at FROM todos WHERE id=$1`, id).
		Scan(&updatedTodo.Id, &updatedTodo.Title, &updatedTodo.Description, &updatedTodo.Status, &updatedTodo.CreatedAt)
	if err != nil {
		http.Error(w, "Error fetching updated todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTodo)
}
