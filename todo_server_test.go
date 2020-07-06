package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zerai/go-todo-app/todo"
)

func TestGETTodos(t *testing.T) {
	repository := todo.NewTodoRepositoryInMemory()
	repository.Add(todo.NewTodoAsValue(123, "first todo"))
	repository.Add(todo.NewTodoAsValue(666, "second todo"))

	server := &TodoServer{repository}

	t.Run("returns todo data for ID 123", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/todos/123", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "123 - first todo")
	})

	t.Run("returns todo data for ID 666", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/todos/666", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "666 - second todo")
	})

	t.Run("returns error not found for unknow todo", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/todos/999", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
		//assertResponseBody(t, response.Body.String(), "666 - second todo")
	})
}

func newGetTodoRequest(todoID string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/todos/%s", todoID), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d want %d", got, want)
	}
}
