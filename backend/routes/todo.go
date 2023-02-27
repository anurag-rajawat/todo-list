package routes

import (
	"github.com/anurag-rajawat/todo-list/backend/handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/todo-completed", handlers.GetCompletedTodos).Methods("GET")
	router.HandleFunc("/todo-incomplete", handlers.GetInCompletedTodos).Methods("GET")
	router.HandleFunc("/todo", handlers.CreateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", handlers.UpdateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", handlers.DeleteItem).Methods("DELETE")
	return router
}
