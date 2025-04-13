package main

import (
	"log"
	"net/http"
	"skillistry-bff-app/handlers"
	"skillistry-bff-app/services"
)

func main() {
	mux := http.NewServeMux()
	services := &services.Container{}

	mux.HandleFunc("/health", handlers.HelloWorld)
	mux.HandleFunc("/skill/create", handlers.CreateSkill(services))
	mux.HandleFunc("/skill/list", handlers.ListSkill(services))
	mux.HandleFunc("/skill/tasks", handlers.ListTasks(services))

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", mux)
}
