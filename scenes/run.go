package scenes

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dbridges/todo/config"
	"github.com/dbridges/todo/models"
	"github.com/dbridges/todo/styles"
	"github.com/dbridges/todo/widgets"
)

type RunScene struct {
	Config  *config.Config
	todos   *models.TodoList
	state   string
	content tea.Model
}

func NewRunScene(cfg *config.Config) (*RunScene, error) {
	todos, err := models.LoadTodos(cfg.Path)
	if err != nil {
		return nil, err
	}
	return &RunScene{
		Config:  cfg,
		todos:   todos,
		state:   "list",
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
		case "ctrl+c", "q":
			return scene, tea.Quit
		case "+":
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
	return scene.HeaderView() + styles.Content.Render(scene.content.View())
}

func (scene *RunScene) HeaderView() string {
	c := styles.Title.Render("TODO")
	c += "\n"
	c += styles.Help.Render(
		"[j/k] Up/Down  [+] New  [space] Toggle  [Del] Delete  [c] Clear  [q] Quit",
	)

	return c
}
