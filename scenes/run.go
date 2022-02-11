package scenes

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dbridges/todo/config"
	"github.com/dbridges/todo/models"
	"github.com/dbridges/todo/styles"
	"github.com/dbridges/todo/widgets"
)

const (
	StateList = iota
	StateForm
)

type RunScene struct {
	config        *config.Config
	store         *models.Store
	content       tea.Model
	state         int
	isHelpVisible bool
}

func NewRunScene(cfg *config.Config) (*RunScene, error) {
	store, err := models.LoadTodos(cfg.Path)
	if err != nil {
		return nil, err
	}
	return &RunScene{
		config:  cfg,
		store:   store,
		content: widgets.NewTodoList(store),
		state:   StateList,
	}, nil
}

func (scene *RunScene) Init() tea.Cmd {
	return nil
}

func (scene *RunScene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if scene.state == StateForm {
			_, cmd := scene.content.Update(msg)
			return scene, cmd
		}

		switch msg.String() {
		case "esc":
			scene.isHelpVisible = false
			return scene, nil
		case "?":
			scene.isHelpVisible = !scene.isHelpVisible
			return scene, nil
		case "ctrl+c", "q":
			return scene, tea.Quit
		case "n":
			scene.state = StateForm
			scene.content = widgets.NewAddTodo()
			return scene, nil
		}
	case widgets.AddTodoAbortMsg:
		scene.state = StateList
		scene.content = widgets.NewTodoList(scene.store)
		return scene, nil
	case widgets.AddTodoMsg:
		if msg.Todo.IsValid() {
			scene.store.Add(msg.Todo)
			scene.store.Save()
		}
		scene.state = StateList
		scene.content = widgets.NewTodoList(scene.store)
		return scene, nil
	}

	_, cmd := scene.content.Update(msg)
	return scene, cmd
}

func (scene *RunScene) View() string {
	if scene.isHelpVisible {
		return scene.HeaderView() + styles.CenteredContent.Render(scene.HelpView())
	} else {
		return scene.HeaderView() + styles.Content.Render(scene.content.View())
	}
}

func (scene *RunScene) HeaderView() string {
	c := styles.Title.Render("TODO")
	c += "\n"
	c += styles.Header.Render(
		"By Dan Bridges. Press ? for help.",
	)

	return c
}

func (scene *RunScene) HelpView() string {
	s := "                   Help                   \n"
	s += "                                          \n"
	s += "  [j/k] Move Cursor             [n] New   \n"
	s += "[space] Toggle                [Del] Delete\n"
	s += "  [+/-] Swap up/down            [c] Clear \n"
	s += "    [?] Help                    [q] Quit  \n"

	return styles.Help.Render(s)
}
