package main

import (
	"fmt"
	"net/http"
	"strings"
)

// TodoServer currently returns 'TODO ONE' given _any_ request.
func TodoServer(w http.ResponseWriter, r *http.Request) {
	todoId := strings.TrimPrefix(r.URL.Path, "/todos/")

	fmt.Fprint(w, GetTodoData(todoId))
}

func GetTodoData(todoId string) string {
	if todoId == "123" {
		return "123 - first todo"
	}

	if todoId == "666" {
		return "666 - first todo"
	}

	return ""
}
