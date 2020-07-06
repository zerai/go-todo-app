package main

import (
	"github.com/zerai/go-todo-app/todo"
	"log"
	"net/http"
)

func main() {
	server := &TodoServer{todo.NewTodoRepositoryInMemory()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
