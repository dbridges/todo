package styles

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Dimensions
const (
	width              = 80
	contentSidePadding = 2
	todoWidth          = width - 2*contentSidePadding
	todoContentWidth   = todoWidth - 4
	helpWidth          = todoWidth - 4
)

// Colors
const (
	secondary = lipgloss.Color("#586e75")
	tertiary  = lipgloss.Color("#3a494e")
	green     = lipgloss.Color("2")
	blue      = lipgloss.Color("4")
	TextLight = lipgloss.Color("#f2f2f2")
	TextDark  = lipgloss.Color("#4a4a4a")
)

var Bold = lipgloss.NewStyle().Bold(true)

var Block = lipgloss.
	NewStyle().
	Width(width).
	Align(lipgloss.Center)

var Title = Block.Copy().Bold(true)

var Header = Block.
	Copy().
	Foreground(secondary).
	BorderStyle(lipgloss.NormalBorder()).
	BorderBottom(true)

var Content = Block.
	Copy().
	Padding(1, contentSidePadding).
	Align(lipgloss.Left)

var CenteredContent = Block.
	Copy().
	Padding(1, contentSidePadding).
	Align(lipgloss.Center)

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

var Help = lipgloss.
	NewStyle().
	Padding(0, 8).
	Align(lipgloss.Center).
	BorderForeground(secondary).
	BorderStyle(lipgloss.RoundedBorder())

func ForegroundTextColor(c string) lipgloss.Color {
	hex := strings.TrimPrefix(c, "#")
	if len(hex) != 6 {
		return TextLight
	}
	rgb := make([]float64, 3)
	for i, v := range []string{hex[0:2], hex[2:4], hex[4:6]} {
		h, err := strconv.ParseInt(v, 16, 64)
		if err != nil {
			return TextLight
		}
		rgb[i] = float64(h)
	}

	if rgb[0]*0.299+rgb[1]*0.587+rgb[2]*0.114 > 186 {
		return TextDark
	}
	return TextLight

}
