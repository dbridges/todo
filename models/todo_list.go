package models

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type TodoList struct {
	Items []Todo
	path  string
}

func LoadTodos(path string) (*TodoList, error) {
	f, err := ioutil.ReadFile(path)

	if os.IsNotExist(err) {
		return &TodoList{}, nil
	} else if err != nil {
		return nil, err
	}

	items := strings.Split(strings.TrimSpace(string(f)), "\n\n")
	todos := make([]Todo, 0, len(items))

	for _, item := range items {
		if item == "" {
			continue
		}
		todo, err := ParseTodo(item)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return &TodoList{Items: todos, path: path}, nil
}

func (todos *TodoList) String() string {
	s := ""
	for _, todo := range todos.Items {
		s += todo.String() + "\n\n"
	}
	return s
}

func (todos *TodoList) Add(todo Todo) {
	todos.Items = append([]Todo{todo}, todos.Items...)
	todos.Sort()
}

func (todos *TodoList) Sort() {
	incomplete := make([]Todo, 0)
	complete := make([]Todo, 0)

	for _, todo := range todos.Items {
		if todo.Completed {
			complete = append(complete, todo)
		} else {
			incomplete = append(incomplete, todo)
		}
	}

	todos.Items = append(incomplete, complete...)
}

func (todos *TodoList) Save() error {
	dir := path.Dir(todos.path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	return os.WriteFile(todos.path, []byte(todos.String()), 0644)
}

func (todos *TodoList) ToggleCompleted(index int) {
	todos.Items[index].Completed = !todos.Items[index].Completed
	todos.Sort()
}

func (todos *TodoList) ClearCompleted() {
	items := make([]Todo, 0)
	for _, todo := range todos.Items {
		if !todo.Completed {
			items = append(items, todo)
		}
	}
	todos.Items = items
	todos.Sort()
}

func (todos *TodoList) Delete(index int) {
	items := make([]Todo, 0)
	for i, todo := range todos.Items {
		if i != index {
			items = append(items, todo)
		}
	}
	todos.Items = items
}

func (todos *TodoList) MoveUp(index int) {
	if index <= 0 {
		return
	}

	v1 := todos.Items[index]
	v2 := todos.Items[index-1]
	todos.Items[index-1] = v1
	todos.Items[index] = v2

	todos.Sort()
}

func (todos *TodoList) MoveDown(index int) {
	if index >= len(todos.Items)-1 {
		return
	}

	v1 := todos.Items[index]
	v2 := todos.Items[index+1]
	todos.Items[index+1] = v1
	todos.Items[index] = v2

	todos.Sort()
}
