package cmd

import (
	"flag"
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
	os.Args = os.Args[1:]

	labelFlag := flag.String("l", "", "Set the new todo label")
	descriptionFlag := flag.String("d", "", "Set the new todo description")
	flag.Parse()

	scene := scenes.NewAddScene(cfg)

	var todo = models.Todo{
		Label:       *labelFlag,
		Description: *descriptionFlag,
	}

	if flag.Arg(0) != "" {
		todo.Title = flag.Arg(0)
	} else {
		p := tea.NewProgram(scene)
		err := p.Start()
		util.CheckError(err)
		todo = scene.Todo
	}

	if todo.IsValid() {
		store, err := models.LoadTodos(cfg.Path)
		if err != nil {
			log.Fatalf("Error %v", err)
		}

		fmt.Println(len(store.Todos))

		store.Add(todo)
		err = store.Save()
		util.CheckError(err)

		fmt.Println("\nAdded 1 todo:")
		fmt.Println(todo.String())
	} else {
		fmt.Println("Aborted")
	}
}
