package cmd

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dbridges/todo/config"
	"github.com/dbridges/todo/models"
	"github.com/dbridges/todo/scenes"
	"github.com/dbridges/todo/util"
)

func Add(cfg *config.Config) {
	scene := scenes.NewAddScene(cfg)

	var todo = models.Todo{}

	if len(os.Args) == 3 {
		todo.Title = os.Args[2]
	} else {
		p := tea.NewProgram(scene)
		if err := p.Start(); err != nil {
			util.ExitError(err)
		}
		todo = scene.Todo
	}

	if todo.IsValid() {
		store, err := models.LoadTodos(cfg.Path)
		if err != nil {
			log.Fatalf("Error %v", err)
		}

		fmt.Println(len(store.Todos))

		store.Add(todo)
		if err := store.Save(); err != nil {
			util.ExitError(err)
		}

		fmt.Println("\nAdded 1 todo:")
		fmt.Println(todo.String())
	} else {
		fmt.Println("Aborted")
	}
}
