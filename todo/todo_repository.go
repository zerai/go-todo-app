package todo

import (
	"errors"
)

var (
	// ErrTodoNotFound ...
	ErrTodoNotFound = errors.New("Todo not found")

	// ErrAlreadyExist ...
	ErrAlreadyExist = errors.New("Todo already exist")
)

// TodoRepository interface for TodoRepository
type TodoRepository interface {
	findByID(identifier int) Todo
}

// TodoRepositoryInMemory in memory storage for Todo
type TodoRepositoryInMemory struct {
	todos map[int]Todo
}

// NewTodoRepositoryInMemory factory method
func NewTodoRepositoryInMemory() *TodoRepositoryInMemory {

	return &TodoRepositoryInMemory{
		todos: make(map[int]Todo),
	}
}

// FindByID find Todo by identifier, returns Todo or error
func (r TodoRepositoryInMemory) FindByID(identifier int) (Todo, error) {
	todo, ok := r.todos[identifier]
	if !ok {
		return Todo{}, ErrTodoNotFound
	}
	return todo, nil
}

// Add add a Todo to repository
func (r *TodoRepositoryInMemory) Add(aTodo Todo) error {
	key := aTodo.ID()
	_, err := r.FindByID(key)

	switch err {
	case ErrTodoNotFound:
		r.todos[key] = aTodo
	case nil:
		return ErrAlreadyExist
	default:
		return err
	}

	return nil
}
