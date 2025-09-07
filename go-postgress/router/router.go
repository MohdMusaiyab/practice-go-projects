package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"go-postgress/handlers"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Home Page!"))
	}).Methods("GET")

	//Addtional Routes Like for Creating a To Do
	//POST API for Creating a To Do
	r.HandleFunc("/api/v1/todos", handlers.CreateTodoHandler).Methods("POST")
	//GET API for Getting All To Dos
	r.HandleFunc("/api/v1/todos", handlers.GetAllTodosHandler).Methods("GET")
	//GET API for Getting a single To Do by ID
	r.HandleFunc("/api/v1/todos/{id}", handlers.GetToDoById).Methods("GET")
	//PUT API for Updating a To Do by ID
	r.HandleFunc("/api/v1/todos/{id}", handlers.UpdateToDoByIdHandler).Methods("PUT")
	//DELETE API for Deleting a To Do by ID
	r.HandleFunc("/api/v1/todos/{id}", handlers.DeleteToDoByIdHandler).Methods("DELETE")
	//For Just Changing the Status of a To Do by ID
	r.HandleFunc("/api/v1/todos/{id}/status", handlers.UpdateToDoStatusByIdHandler).Methods("PATCH")
	return r
}
