package scenes

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dbridges/todo/config"
	"github.com/dbridges/todo/models"
	"github.com/dbridges/todo/styles"
	"github.com/dbridges/todo/widgets"
)

type RunScene struct {
	config        *config.Config
	todos         *models.TodoList
	content       tea.Model
	isHelpVisible bool
}

func NewRunScene(cfg *config.Config) (*RunScene, error) {
	todos, err := models.LoadTodos(cfg.Path)
	if err != nil {
		return nil, err
	}
	return &RunScene{
		config:  cfg,
		todos:   todos,
		content: widgets.NewTodoList(todos),
	}, nil
}

func (scene *RunScene) Init() tea.Cmd {
	return nil
}

func (scene *RunScene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
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
			scene.content = widgets.NewAddTodo()
			return scene, nil
		}
	case widgets.AddTodoAbortMsg:
		scene.content = widgets.NewTodoList(scene.todos)
		return scene, nil
	case widgets.AddTodoMsg:
		if msg.Todo.IsValid() {
			scene.todos.Add(msg.Todo)
			scene.todos.Save()
		}
		scene.content = widgets.NewTodoList(scene.todos)
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
