package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETTodos(t *testing.T) {

	t.Run("returns todo data for ID 123", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/todos/123", nil)
		response := httptest.NewRecorder()

		TodoServer(response, request)

		assertResponseBody(t, response.Body.String(), "123 - first todo")
	})

	t.Run("returns todo data for ID 666", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/todos/666", nil)
		response := httptest.NewRecorder()

		TodoServer(response, request)

		assertResponseBody(t, response.Body.String(), "666 - first todo")
	})

}

func newGetTodoRequest(todoId string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", todoId), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
