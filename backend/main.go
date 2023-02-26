package main

import (
	"github.com/anurag-rajawat/todo-list/backend/routes"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	router := routes.Router()
	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
	}).Handler(router)
	log.Info("server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
