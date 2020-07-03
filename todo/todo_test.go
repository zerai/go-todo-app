package todo

import (
	"testing"
)

func TestTodo(t *testing.T) {

	anIdentifier := 23
	aLabel := "buy a car."
	aDescrition := "check a catalog."

	t.Run("create a todo", func(t *testing.T) {
		todo := NewTodo(anIdentifier, aLabel)

		if todo.Label() != aLabel {
			t.Error("unexpected error.")
		}
	})

	t.Run("Should returns the identifier", func(t *testing.T) {
		todo := NewTodo(anIdentifier, aLabel)

		if todo.ID() != anIdentifier {
			t.Error("unexpected error.")
		}
	})

	t.Run("Should returns the label", func(t *testing.T) {
		todo := NewTodo(anIdentifier, aLabel)

		if todo.Label() != aLabel {
			t.Error("unexpected error.")
		}
	})

	t.Run("Should returns the description", func(t *testing.T) {
		todo := Todo{
			description: aDescrition,
		}

		description, ok := todo.Description()
		if !ok {
			t.Error("unexpected error.")
		}
		if description != aDescrition {
			t.Errorf("unexpected error got %v want %v", description, aDescrition)
		}
	})

	t.Run("Should not returns description", func(t *testing.T) {
		todo := NewTodo(anIdentifier, aLabel)

		_, ok := todo.Description()
		if ok {
			t.Error("error should return false.")
		}
	})
}
