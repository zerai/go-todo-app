package todo

import (
	"testing"
)

func TestTodoRepository(t *testing.T) {

	todoRepositoryInMemory := TodoRepositoryInMemory{
		todos: map[int]Todo{
			123: {123, "a label", "a description"},
			321: {321, "a second label", "a second description"},
		},
	}

	t.Run("findByID should returns a Todo", func(t *testing.T) {
		got, err := todoRepositoryInMemory.FindByID(123)
		want := Todo{123, "a label", "a description"}

		if err != nil {
			t.Fatal("Not expected an error")
		}

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "123")
		}
	})

	t.Run("findByID should returns error", func(t *testing.T) {
		_, got := todoRepositoryInMemory.FindByID(555)

		assertError(t, got, ErrTodoNotFound)
	})

}

func TestAddTodo(t *testing.T) {
	t.Run("Should add a valid todo", func(t *testing.T) {
		emptytodoRepositoryInMemory := NewTodoRepositoryInMemory()
		aNewTodo := NewTodoAsValue(555, "a label")

		emptytodoRepositoryInMemory.Add(aNewTodo)

		got, err := emptytodoRepositoryInMemory.FindByID(555)
		want := aNewTodo
		if err != nil {
			t.Fatal("should find added Todo:", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Should returns error if Todo exist", func(t *testing.T) {
		repository := NewTodoRepositoryInMemory()
		aNewTodo := NewTodoAsValue(99, "a label")
		repository.Add(aNewTodo)

		err := repository.Add(aNewTodo)

		assertError(t, err, ErrAlreadyExist)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Shoud update existing Todo", func(t *testing.T) {
		repositoryInMemory := NewTodoRepositoryInMemory()
		todo := NewTodoAsValue(888, "a label")
		repositoryInMemory.Add(todo)
		updatedTodo := NewTodoAsValue(888, "an updated label")
		want := updatedTodo

		repositoryInMemory.Update(updatedTodo)

		got, err := repositoryInMemory.FindByID(888)
		if err != nil {
			t.Fatal("should find updated Todo:", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Should returns error for unknow Todo", func(t *testing.T) {
		repositoryInMemory := NewTodoRepositoryInMemory()
		todo := NewTodoAsValue(888, "a label")

		err := repositoryInMemory.Update(todo)

		assertError(t, err, ErrTodoNotFound)
	})
}

func TestRemove(t *testing.T) {
	identifier := 333
	repositoryInMemory := NewTodoRepositoryInMemory()
	todo := NewTodoAsValue(identifier, "a label")
	repositoryInMemory.Add(todo)

	repositoryInMemory.Remove(identifier)

	_, err := repositoryInMemory.FindByID(identifier)
	if err != ErrTodoNotFound {
		t.Errorf("Expected todo with key %d to be deleted", identifier)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
