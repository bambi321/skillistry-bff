package main

import (
	"log"
	"net/http"
	"skillistry-bff-app/src/handlers"
	"skillistry-bff-app/src/services"
)

func main() {
	services := &services.Container{}

	http.HandleFunc("/", handlers.HelloWorld)
	http.HandleFunc("skill/create", handlers.CreateSkill(services))
	http.HandleFunc("skill/list", handlers.ListSkill(services))
	http.HandleFunc("skill/tasks", handlers.ListTasks(services))

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
