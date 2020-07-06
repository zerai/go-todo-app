package main

import (
	"fmt"
	"github.com/zerai/go-todo-app/todo"
	"net/http"
	"strconv"
	"strings"
)

// TodoServer ...
type TodoServer struct {
	repository todo.TodoRepository
}

func (p *TodoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	param := strings.TrimPrefix(r.URL.Path, "/todos/")
	todoID, conversionErr := strconv.Atoi(param)

	if conversionErr != nil {
		fmt.Fprint(w, "Error 500 input param conversion error")
		return
	}

	todo, err := p.repository.FindByID(todoID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err.Error())
		return
	}
	resultString := fmt.Sprintf("%v - ", todo.ID()) + todo.Label()
	fmt.Fprint(w, resultString)
}
