package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"todo-app-go.com/v1/database"
	"todo-app-go.com/v1/model"
)

type TodoServer struct {
	todoStore database.TodoStore
	http.Handler
}

func NewTodoServer(todoStore database.TodoStore) *TodoServer {
	server := new(TodoServer)

	server.todoStore = todoStore

	router := mux.NewRouter()

	router.Handle("/api/todos", http.HandlerFunc(server.todosHandler))
	router.Handle("/api/todos/{id:[0-9]+}", http.HandlerFunc(server.todoByIdHandler))

	server.Handler = router

	return server
}

func (t *TodoServer) todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(t.todoStore.GetAllTodos())
		w.WriteHeader(http.StatusOK)
	case http.MethodPost:
		var todo model.Todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			panic(fmt.Sprintf("invalid json data, err: %v", err))
		}
		t.todoStore.AddTodo(todo)
		w.WriteHeader(http.StatusCreated)
	}
}

func (t *TodoServer) todoByIdHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	id, _ := strconv.Atoi(idParam)
	w.Header().Set("content-type", "application/json")
	todo, _ := t.todoStore.GetTodoById(id)
	json.NewEncoder(w).Encode(todo)
	w.WriteHeader(http.StatusOK)
}
