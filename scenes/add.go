package scenes

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dbridges/todo/config"
	"github.com/dbridges/todo/models"
	"github.com/dbridges/todo/widgets"
)

type AddScene struct {
	Config  *config.Config
	Todo    models.Todo
	content tea.Model
	history string
}

func NewAddScene(cfg *config.Config) *AddScene {
	return &AddScene{
		Config:  cfg,
		content: widgets.NewAddTodo(),
	}
}

func (scene *AddScene) Init() tea.Cmd {
	return nil
}

func (scene *AddScene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case widgets.AddTodoAbortMsg:
		return scene, tea.Quit
	case widgets.AddTodoMsg:
		scene.Todo = msg.Todo
		return scene, tea.Quit
	default:
		_, cmd := scene.content.Update(msg)
		return scene, cmd
	}
}

func (scene *AddScene) View() string {
	content := scene.history
	content += scene.content.View()
	return content
}
