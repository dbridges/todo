package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dbridges/todo/config"
	"github.com/dbridges/todo/scenes"
	"github.com/dbridges/todo/util"
)

func Run(cfg *config.Config) {
	scene, err := scenes.NewRunScene(cfg)
	if err != nil {
		util.ExitError(err)
	}

	p := tea.NewProgram(scene)
	if err := p.Start(); err != nil {
		util.ExitError(err)
	}
}
