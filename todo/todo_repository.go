package todo

import (
	"errors"
)

var ErrTodoNotFound = errors.New("Todo not found")

type TodoRepository interface {
	findByID(identifier int) Todo
}

type TodoRepositoryInMemory struct {
	todos map[int]Todo
}

func (r TodoRepositoryInMemory) FindByID(identifier int) (Todo, error) {
	todo, ok := r.todos[identifier]
	if !ok {
		return Todo{}, ErrTodoNotFound
	}
	return todo, nil
}
