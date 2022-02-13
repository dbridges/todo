package widgets

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dbridges/todo/models"
	"github.com/dbridges/todo/styles"
	"github.com/dbridges/todo/util"
)

type TodoList struct {
	store       *models.Store
	cursor      int
	labelColors map[string]string
}

func NewTodoList(store *models.Store, labelColors map[string]string) *TodoList {
	return &TodoList{store: store, labelColors: labelColors}
}

func (m *TodoList) Init() tea.Cmd {
	return nil
}

func (m *TodoList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.cursor = util.Clamp(m.cursor+1, 0, len(m.store.Todos)-1)
		case "k", "up":
			m.cursor = util.Clamp(m.cursor-1, 0, len(m.store.Todos)-1)
		case " ":
			m.store.ToggleCompleted(m.cursor)
			m.store.Save()
		case "=", "+":
			m.store.MoveUp(m.cursor)
			m.cursor = util.Clamp(m.cursor-1, 0, len(m.store.Todos)-1)
			m.store.Save()
		case "-", "_":
			m.store.MoveDown(m.cursor)
			m.cursor = util.Clamp(m.cursor+1, 0, len(m.store.Todos)-1)
			m.store.Save()
		case "backspace":
			m.store.Delete(m.cursor)
			m.cursor = util.Clamp(m.cursor, 0, len(m.store.Todos)-1)
			m.store.Save()
		}
	}

	return m, nil
}

func (m *TodoList) View() string {
	if len(m.store.Todos) == 0 {
		return m.EmptyView()
	}
	return m.ListView()
}

func (m *TodoList) EmptyView() string {
	return styles.Message.Render("No todos yet!")
}

func (m *TodoList) ListView() string {
	items := make([]string, len(m.store.Todos))

	for i, todo := range m.store.Todos {
		items[i] = m.RenderTodo(todo, m.cursor == i)
	}

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

func (m *TodoList) RenderTodo(todo models.Todo, selected bool) string {
	check := ""
	var titleStyle lipgloss.Style
	var descriptionStyle lipgloss.Style
	labelColor := "#1074e6"

	color, ok := m.labelColors[todo.Label]
	if ok {
		labelColor = color
	}

	labelStyle := lipgloss.
		NewStyle().
		Foreground(styles.ForegroundTextColor(labelColor)).
		Background(lipgloss.Color(labelColor))

	if todo.Completed() {
		check = styles.Green.Render(" ✓ ")
		titleStyle = styles.CompletedTodoTitle
		descriptionStyle = styles.CompletedTodoDescription
	} else {
		check = styles.Secondary.Render(" ☐ ")
		titleStyle = styles.TodoTitle
		descriptionStyle = styles.TodoDescription
	}

	var title string
	if todo.Label != "" {
		title = labelStyle.Render(" "+todo.Label+" ") + " " + titleStyle.Render(todo.Title)
	} else {
		title = titleStyle.Render(todo.Title)
	}

	bodyItems := make([]string, 0)
	bodyItems = append(bodyItems, title)
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
