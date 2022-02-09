package models

import "testing"

func TestTodoString(t *testing.T) {
	todo := Todo{Title: "Do stuff"}
	expected := "- [ ] Do stuff"
	if todo.String() != expected {
		t.Fatalf("Wanted '%s' got '%s'", expected, todo.String())
	}

	todo.Description = "Some extra info."
	expected =
		`- [ ] Do stuff
      Some extra info.`
	if todo.String() != expected {
		t.Fatalf("Wanted '%s' got '%s'", expected, todo.String())
	}

	todo.Completed = true
	expected =
		`- [x] Do stuff
      Some extra info.`
	if todo.String() != expected {
		t.Fatalf("Wanted '%s' got '%s'", expected, todo.String())
	}
}
