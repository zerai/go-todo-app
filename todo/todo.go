package todo

// Todo describe a todo
type Todo struct {
	todoID      int
	label       string
	description string
}

// NewTodo factory method returns a Todo
func NewTodo(anIdentifier int, aLabel string) *Todo {
	return &Todo{
		todoID: anIdentifier,
		label:  aLabel,
	}
}

// ID returns the identifier
func (t Todo) ID() int {
	return t.todoID
}

// Label returns the label
func (t Todo) Label() string {
	return t.label
}

// Description returns the description
func (t Todo) Description() (description string, ok bool) {
	if len(t.description) > 0 {
		return t.description, true
	} else {
		return "", false
	}

}
