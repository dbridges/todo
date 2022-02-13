package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dbridges/todo/config"
	"github.com/dbridges/todo/scenes"
	"github.com/dbridges/todo/util"
)

func Run(cfg *config.Config) {
	scene, err := scenes.NewRunScene(cfg)
	util.CheckError(err)

	p := tea.NewProgram(scene)
	err = p.Start()
	util.CheckError(err)
}
