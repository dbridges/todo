package widgets

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dbridges/todo/models"
)

type AddTodo struct {
	todo    models.Todo
	field   string
	content tea.Model
	history string
}

type AddTodoMsg struct {
	Todo models.Todo
}

type AddTodoAbortMsg struct{}

func NewAddTodo() *AddTodo {
	return &AddTodo{
		field:   "title",
		content: NewInput("Title"),
	}
}

func (m *AddTodo) Init() tea.Cmd {
	return nil
}

func (m *AddTodo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case InputMsg:
		switch m.field {
		case "title":
			if msg.Value == "" {
				return m, func() tea.Msg { return AddTodoAbortMsg{} }
			}
			m.todo.Title = msg.Value
			m.history += m.content.View() + "\n"
			m.content = NewInput("Description")
			m.field = "description"
			return m, nil
		case "description":
			m.todo.Description = msg.Value
			m.history += m.content.View() + "\n"
			m.content = NewInput("Label")
			m.field = "label"
			return m, nil
		case "label":
			m.todo.Label = msg.Value
			m.history += m.content.View() + "\n"
			return m, func() tea.Msg { return AddTodoMsg{Todo: m.todo} }
		default:
			return m, nil
		}
	default:
		_, cmd := m.content.Update(msg)
		return m, cmd
	}
}

func (m *AddTodo) View() string {
	content := m.history
	content += m.content.View()
	return content
}
