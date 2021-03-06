package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"todo-app-go.com/v1/controller"
	"todo-app-go.com/v1/database"
	"todo-app-go.com/v1/error_handler"
	"todo-app-go.com/v1/model"
	"todo-app-go.com/v1/util"
)

func TestGetAll(t *testing.T) {
	wantedTodos := []model.Todo{
		model.NewTodo(1, "FirstTodo", false),
		model.NewTodo(2, "SecondTodo", false),
	}
	store := database.NewStubTodoStore(
		&map[int]model.Todo{
			1: wantedTodos[0],
			2: wantedTodos[1],
		},
	)
	server := controller.NewTodoServer(&store)

	t.Run("it todos as JSON", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/todos", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		var got []model.Todo

		err := json.NewDecoder(res.Body).Decode(&got)

		assert.Nil(t, err)
		assert.Equal(t, res.Code, http.StatusOK)

		util.SortTodoSliceById(got)

		assert.Equal(t, wantedTodos, got)
		assert.Equal(t, util.JsonContentType, res.Result().Header.Get("content-type"))
	})
}

func TestAdd(t *testing.T) {
	todo := "invalid todo"
	store := database.NewStubTodoStore(
		&map[int]model.Todo{},
	)
	server := controller.NewTodoServer(&store)

	t.Run("invalid json format should return 400", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/api/todos", bytes.NewBuffer([]byte(todo)))
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		var parseError error_handler.TodoAppError
		json.NewDecoder(res.Body).Decode(&parseError)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, error_handler.BAD_REQUEST_CODE, parseError.Code)
	})
}
