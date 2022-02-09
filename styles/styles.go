package styles

import "github.com/charmbracelet/lipgloss"

// Dimensions
const (
	width              = 80
	contentSidePadding = 2
	todoWidth          = width - 2*contentSidePadding
	todoContentWidth   = todoWidth - 4
)

// Colors
const (
	secondary = lipgloss.Color("#586e75")
	tertiary  = lipgloss.Color("#3a494e")
	green     = lipgloss.Color("2")
	blue      = lipgloss.Color("4")
)

var Bold = lipgloss.NewStyle().Bold(true)

var Block = lipgloss.
	NewStyle().
	Width(width).
	Align(lipgloss.Center)

var Title = Block.Copy().Bold(true)

var Help = Block.Copy().Foreground(secondary).BorderStyle(lipgloss.NormalBorder()).BorderBottom(true)

var Content = Block.Copy().Padding(1, contentSidePadding).Align(lipgloss.Left)

var Message = Block.Copy().Padding(2)

var Secondary = lipgloss.NewStyle().Foreground(secondary)
var Green = lipgloss.NewStyle().Foreground(green)

var Todo = lipgloss.
	NewStyle().
	BorderStyle(lipgloss.HiddenBorder()).
	BorderLeft(true).
	MarginBottom(1)

var SelectedTodo = lipgloss.
	NewStyle().
	BorderStyle(lipgloss.ThickBorder()).
	BorderForeground(blue).
	MarginBottom(1).
	Inherit(Todo)

var TodoTitle = lipgloss.NewStyle().Width(todoContentWidth)
var TodoDescription = lipgloss.
	NewStyle().
	Foreground(secondary).
	Inherit(TodoTitle)

var CompletedTodoTitle = lipgloss.
	NewStyle().
	Strikethrough(true).
	Foreground(secondary).
	Inherit(TodoTitle)

var CompletedTodoDescription = lipgloss.
	NewStyle().
	Strikethrough(true).
	Foreground(tertiary).
	Inherit(TodoDescription)
