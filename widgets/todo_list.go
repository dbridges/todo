package widgets

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dbridges/todo/models"
	"github.com/dbridges/todo/styles"
	"github.com/dbridges/todo/util"
)

type TodoList struct {
	todos  *models.TodoList
	cursor int
}

func NewTodoList(todos *models.TodoList) *TodoList {
	return &TodoList{todos: todos}
}

func (m *TodoList) Init() tea.Cmd {
	return nil
}

func (m *TodoList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.cursor = util.Clamp(m.cursor+1, 0, len(m.todos.Items)-1)
		case "k", "up":
			m.cursor = util.Clamp(m.cursor-1, 0, len(m.todos.Items)-1)
		case " ":
			m.todos.ToggleCompleted(m.cursor)
			m.todos.Save()
		case "=", "+":
			m.todos.MoveUp(m.cursor)
			m.cursor = util.Clamp(m.cursor-1, 0, len(m.todos.Items)-1)
			m.todos.Save()
		case "-", "_":
			m.todos.MoveDown(m.cursor)
			m.cursor = util.Clamp(m.cursor+1, 0, len(m.todos.Items)-1)
			m.todos.Save()
		case "c":
			m.todos.ClearCompleted()
			m.cursor = util.Clamp(m.cursor, 0, len(m.todos.Items)-1)
			m.todos.Save()
		case "backspace":
			m.todos.Delete(m.cursor)
			m.cursor = util.Clamp(m.cursor, 0, len(m.todos.Items)-1)
			m.todos.Save()
		}
	}

	return m, nil
}

func (m *TodoList) View() string {
	if len(m.todos.Items) == 0 {
		return m.EmptyView()
	}
	return m.ListView()
}

func (m *TodoList) EmptyView() string {
	return styles.Message.Render("No todos yet!")
}

func (m *TodoList) ListView() string {
	items := make([]string, len(m.todos.Items))

	for i, todo := range m.todos.Items {
		items[i] = m.RenderTodo(todo, m.cursor == i)
	}

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

func (m *TodoList) RenderTodo(todo models.Todo, selected bool) string {
	check := ""
	var titleStyle lipgloss.Style
	var descriptionStyle lipgloss.Style

	if todo.Completed {
		check = styles.Green.Render(" ✓ ")
		titleStyle = styles.CompletedTodoTitle
		descriptionStyle = styles.CompletedTodoDescription
	} else {
		check = styles.Secondary.Render(" ☐ ")
		titleStyle = styles.TodoTitle
		descriptionStyle = styles.TodoDescription
	}

	bodyItems := make([]string, 0)
	bodyItems = append(bodyItems, titleStyle.Render(todo.Title))
	if len(todo.Description) > 0 {
		bodyItems = append(bodyItems, descriptionStyle.Render(todo.Description))
	}

	body := lipgloss.JoinVertical(lipgloss.Left, bodyItems...)

	c := lipgloss.JoinHorizontal(lipgloss.Top, check, body)

	if selected {
		return styles.SelectedTodo.Render(c)
	}
	return styles.Todo.Render(c)
}
