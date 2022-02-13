package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/dbridges/todo/cmd"
	"github.com/dbridges/todo/config"
	"github.com/dbridges/todo/util"
)

var Version string

func initLog() error {
	cacheDir, err := config.UserCacheDir()
	if err != nil {
		log.Fatalf("Unable to open cache dir")
	}

	f, err := os.Create(path.Join(cacheDir, "todo.log"))

	if err != nil {
		return err
	}

	log.SetOutput(f)

	return nil
}

func Usage() {
	fmt.Printf("todo version %s\n", Version)
	fmt.Println("Todo is a tool for managing your todos.")
	fmt.Println("\nUsage:")
	fmt.Println("\n\ttodo <command> [arguments]")
	fmt.Println("\nThe commands are:")
	fmt.Println("\n\thelp\t\t display this help")
	fmt.Println("\trun\t\t start the todo TUI program")
	fmt.Println("\tadd, new \t add a new todo to the list")
	fmt.Println("\nIf no command is given it defaults to `run`")
	fmt.Println()
}

func main() {
	err := initLog()
	util.CheckError(err)

	cfg, err := config.Load()
	util.CheckError(err)
	if !cfg.IsValid() {
		util.CheckError(
			fmt.Errorf("Invalid config, must contain 'path' entry under '[core]'."),
		)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "add", "new":
			cmd.Add(cfg)
		default:
			Usage()
		}
	} else {
		cmd.Run(cfg)
	}
}
