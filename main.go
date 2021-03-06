package main

import (
	"log"
	"net/http"

	"todo-app-go.com/v1/controller"
	"todo-app-go.com/v1/database"
)

func main() {
	// server := &controller.TodoServer{database.NewInMemoryTodoStore()}
	store := database.NewInMemoryTodoStore()
	server := controller.NewTodoServer(&store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
