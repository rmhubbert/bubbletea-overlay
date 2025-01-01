package overlay

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Position represents a relative offset in the TUI.
type Position int

const (
	Top Position = iota + 1
	Right
	Bottom
	Left
	Center
)

// Model implements tea.Model, and manages the overlay UI.
type Model struct {
	Foreground tea.Model
	Background tea.Model
	XPosition  Position
	YPosition  Position
	XOffset    int
	YOffset    int
}

// New creates, instantiates, and returns a new overlay Model.
func New(fore tea.Model, back tea.Model, xPos Position, yPos Position, xOff int, yOff int) *Model {
	return &Model{
		Foreground: fore,
		Background: back,
		XPosition:  xPos,
		YPosition:  yPos,
		XOffset:    xOff,
		YOffset:    yOff,
	}
}

// Init initialises the Model on program load. It partly implements the tea.Model interface.
func (m *Model) Init() tea.Cmd {
	return nil
}

// Update handles event and manages internal state. It partly implements the tea.Model interface.
func (m *Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// View applies and styling and handles rendering the view. It partly implements the tea.Model
// interface.
func (m *Model) View() string {
	if m.Foreground == nil && m.Background == nil {
		return ""
	}
	if m.Foreground == nil && m.Background != nil {
		return m.Background.View()
	}
	if m.Foreground != nil && m.Background == nil {
		return m.Foreground.View()
	}

	return composite(
		m.Foreground.View(),
		m.Background.View(),
		m.XPosition,
		m.YPosition,
		m.XOffset,
		m.YOffset,
	)
}
