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

	switch r.Method {
	case http.MethodPost:
		p.addTodo(w, r)
	case http.MethodGet:
		p.showTodo(w, r)
	case http.MethodDelete:
		p.deleteTodo(w, r)
	}
}

func (p *TodoServer) showTodo(w http.ResponseWriter, r *http.Request) {
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

func (p *TodoServer) addTodo(w http.ResponseWriter, r *http.Request) {

	todoID, err := strconv.Atoi(r.FormValue("todo_id"))

	if err != nil {
		fmt.Fprint(w, "Error 500 input param conversion error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	label := r.FormValue("label")

	todo := todo.NewTodoAsValue(todoID, label)

	// store
	errRepository := p.repository.Add(todo)
	if errRepository != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, errRepository.Error())
		return
	}
	w.WriteHeader(http.StatusAccepted)
	return
}

func (p *TodoServer) deleteTodo(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, "/todos/")
	todoID, err := strconv.Atoi(param)
	if err != nil {
		fmt.Fprint(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = p.repository.FindByID(todoID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err.Error())
		return
	}

	p.repository.Remove(todoID)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
