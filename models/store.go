package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/dbridges/todo/util"
)

const timeFormat = time.RFC3339

var columns = []string{"title", "description", "label", "completed_at"}
var columnCount = len(columns)

type Store struct {
	Todos          []Todo
	completedTodos []Todo
	path           string
}

func LoadTodos(path string) (*Store, error) {
	f, err := os.Open(path)
	defer f.Close()

	if os.IsNotExist(err) {
		return &Store{path: path}, nil
	} else if err != nil {
		return nil, err
	}

	r := csv.NewReader(f)

	records, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	completedTodos := make([]Todo, 0, len(records))
	todos := make([]Todo, 0)

	today := util.Today()

	for i, row := range records {
		// Don't parse header row
		if i == 0 {
			continue
		}
		if len(row) != columnCount {
			return nil, fmt.Errorf("Invalid number of columns in: '%v'", row)
		}
		todo := Todo{}
		todo.Title = row[0]
		todo.Description = row[1]
		todo.Label = row[2]
		todo.CompletedAt = row[3]

		if todo.Completed() && todo.CompletedAt < today {
			completedTodos = append(completedTodos, todo)
		} else {
			todos = append(todos, todo)
		}
	}

	s := &Store{Todos: todos, completedTodos: completedTodos, path: path}
	s.Sort()

	return s, nil
}

func (s *Store) Add(todo Todo) {
	s.Todos = append([]Todo{todo}, s.Todos...)
	s.Sort()
}

func (s *Store) Sort() {
	incomplete := make([]Todo, 0)
	complete := make([]Todo, 0)

	for _, todo := range s.Todos {
		if todo.Completed() {
			complete = append(complete, todo)
		} else {
			incomplete = append(incomplete, todo)
		}
	}

	s.Todos = append(incomplete, complete...)
}

func (s *Store) Save() error {
	dir := path.Dir(s.path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(s.path)

	if err != nil {
		return err
	}

	defer f.Close()

	w := csv.NewWriter(f)
	err = w.Write(columns)
	if err != nil {
		return err
	}

	for _, todo := range s.completedTodos {
		if err = w.Write(todo.Fields()); err != nil {
			return err
		}
	}

	for _, todo := range s.Todos {
		if err = w.Write(todo.Fields()); err != nil {
			return err
		}
	}

	w.Flush()
	return w.Error()
}

func (s *Store) ToggleCompleted(index int) {
	if s.Todos[index].CompletedAt == "" {
		s.Todos[index].CompletedAt = util.Today()
	} else {
		s.Todos[index].CompletedAt = ""
	}
	s.Sort()
}

func (s *Store) Delete(index int) {
	items := make([]Todo, 0)
	for i, todo := range s.Todos {
		if i != index {
			items = append(items, todo)
		}
	}
	s.Todos = items
}

func (todos *Store) MoveUp(index int) {
	if index <= 0 {
		return
	}

	v1 := todos.Todos[index]
	v2 := todos.Todos[index-1]
	todos.Todos[index-1] = v1
	todos.Todos[index] = v2

	todos.Sort()
}

func (todos *Store) MoveDown(index int) {
	if index >= len(todos.Todos)-1 {
		return
	}

	v1 := todos.Todos[index]
	v2 := todos.Todos[index+1]
	todos.Todos[index+1] = v1
	todos.Todos[index] = v2

	todos.Sort()
}
