package models

import (
	"fmt"
	"strings"
)

type Todo struct {
	Title       string
	Description string
	Completed   bool
}

func ParseTodo(txt string) (Todo, error) {
	todo := Todo{}
	lines := strings.SplitN(strings.TrimSpace(txt), "\n", 2)

	if len(lines) < 1 {
		return todo, fmt.Errorf("Empty todo detected")
	}

	if !(strings.HasPrefix(lines[0], "- [ ] ") ||
		strings.HasPrefix(lines[0], "- [x] ")) {
		return todo, fmt.Errorf("Unable to parse '%s'", txt)
	}

	todo.Title = lines[0][6:]

	if strings.HasPrefix(lines[0], "- [x] ") {
		todo.Completed = true
	}

	if len(lines) == 2 {
		todo.Description = lines[1][6:]
	}

	return todo, nil
}

func (todo Todo) String() string {
	s := ""
	if todo.Completed {
		s += "- [x] "
	} else {
		s += "- [ ] "
	}

	s += todo.Title

	if len(todo.Description) > 0 {
		s += "\n      " + todo.Description
	}

	return s
}

func (todo Todo) IsValid() bool {
	return todo.Title != ""
}
