package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETTodos(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	TodoServer(response, request)

	t.Run("returns a todo data", func(t *testing.T) {
		got := response.Body.String()
		want := "TODO ONE"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}