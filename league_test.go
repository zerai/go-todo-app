package main

import (
	_ "fmt"
	"net/http"
	"net/http/httptest"
	_ "net/url"
	"testing"

	"github.com/zerai/go-todo-app/todo"
)

func TestLeague(t *testing.T) {
	repository := todo.NewTodoRepositoryInMemory()
	server := &TodoServer{repository}

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}
