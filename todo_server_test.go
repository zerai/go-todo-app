package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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

func TestPostTodo(t *testing.T) {
	repository := todo.NewTodoRepositoryInMemory()
	server := &TodoServer{repository}

	t.Run("it records a Todo when POST", func(t *testing.T) {
		request := newPostTodoRequest("456")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		_, err := repository.FindByID(456)
		if err != nil {
			t.Fatal("should find added Todo:", err)
		}
	})

	t.Run("it not records a Todo if exist", func(t *testing.T) {
		identifier := 455
		identifierAsString := "455"
		aNewTodo := todo.NewTodoAsValue(identifier, "a label")
		err := repository.Add(aNewTodo)
		if err != nil {
			t.Fatal("should populate with fixture :", err)
		}

		request := newPostTodoRequest(identifierAsString)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusInternalServerError)

	})
}

func TestHandleDeleteTodo(t *testing.T) {
	identifier := 23
	identifierAsString := "23"
	repository := todo.NewTodoRepositoryInMemory()
	repository.Add(todo.NewTodoAsValue(identifier, "first todo"))
	server := &TodoServer{repository}

	t.Run("should remove a todo", func(t *testing.T) {
		request := newDeleteTodoRequest(identifierAsString)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)

		_, err := repository.FindByID(identifier)
		if err != todo.ErrTodoNotFound {
			t.Errorf("Expected todo with key %d to be deleted", identifier)
		}
	})
}

func newGetTodoRequest(todoID string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/todos/%s", todoID), nil)
	return req
}

func newPostTodoRequest(todoID string) *http.Request {
	data := url.Values{}
	data.Set("todo_id", todoID)
	data.Set("label", "a label")

	req, _ := http.NewRequest(http.MethodPost, "/todos/new", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	return req
}

func newDeleteTodoRequest(aTodoID string) *http.Request {
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/todos/%s", aTodoID), nil)
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
