package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

	router := http.NewServeMux()
	router.Handle("/api/todos", http.HandlerFunc(server.todosHandler))

	server.Handler = router

	return server
}

func (t *TodoServer) todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		t.todoCreation(w, r)
	case http.MethodGet:
		t.getAllTodos(w, r)
	}
}

// func (t *TodoServer) showDescription(w http.ResponseWriter, todoId int) {
// 	todo, err := t.Store.GetTodoById(todoId)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 	}

// 	fmt.Fprintf(w, "%v", todo)
// }

func (t *TodoServer) getAllTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "getAllTodos")
}

func (t *TodoServer) todoCreation(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var todo model.Todo
	json.Unmarshal(reqBody, &todo)
	t.todoStore.AddTodo(todo)
}
