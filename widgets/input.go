package widgets

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dbridges/todo/styles"
)

type Input struct {
	prompt    string
	value     string
	isFocused bool
}

type InputMsg struct {
	Value string
}

func NewInput(prompt string) *Input {
	return &Input{prompt: prompt, isFocused: true}
}

func (m *Input) Init() tea.Cmd {
	return nil
}

func (m *Input) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyCtrlU:
			m.value = ""
			return m, nil
		case tea.KeyEsc:
			return m, func() tea.Msg { return InputMsg{Value: ""} }
		case tea.KeyEnter:
			m.isFocused = false
			return m, func() tea.Msg { return InputMsg{Value: m.value} }
		case tea.KeyBackspace:
			if len(m.value) > 0 {
				m.value = m.value[:len(m.value)-1]
			}
		case tea.KeyRunes:
			m.value += msg.String()
		}
	}

	return m, nil
}

func (m *Input) View() string {
	s := styles.Bold.Render(m.prompt) + " " + m.value
	if m.isFocused {
		s += "█"
	}
	return s
}
