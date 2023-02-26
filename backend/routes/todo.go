package routes

import (
	"github.com/anurag-rajawat/todo-list/backend/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/todo-completed", middleware.GetCompletedTodos).Methods("GET")
	router.HandleFunc("/todo-incomplete", middleware.GetInCompletedTodos).Methods("GET")
	router.HandleFunc("/todo", middleware.CreateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", middleware.UpdateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", middleware.DeleteItem).Methods("DELETE")
	return router
}
