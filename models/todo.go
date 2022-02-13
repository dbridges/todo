package models

type Todo struct {
	Title       string
	Description string
	Label       string
	CompletedAt string
}

func (todo Todo) String() string {
	s := todo.Title

	if len(todo.Description) > 0 {
		s += "\n" + todo.Description
	}

	return s
}

func (todo Todo) IsValid() bool {
	return todo.Title != ""
}

func (todo Todo) Completed() bool {
	return todo.CompletedAt != ""
}

func (todo Todo) Fields() []string {
	return []string{todo.Title, todo.Description, todo.Label, todo.CompletedAt}
}
