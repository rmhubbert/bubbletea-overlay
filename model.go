// Package Overlay is a component for Charm's Bubble Tea TUI framework that aims to simplify
// creating and managing overlays and modal windows in your TUI apps.
package overlay

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Position represents a relative offset in the TUI. There are five possible values; Top, Right,
// Bottom, Left, and Center.
type Position int

const (
	Top Position = iota + 1
	Right
	Bottom
	Left
	Center
)

type Viewable interface {
	View() string
}

// Model implements tea.Model, and manages calculating and compositing the overlay UI from the
// backbround and foreground models.
type Model struct {
	Foreground Viewable
	Background Viewable
	XPosition  Position
	YPosition  Position
	XOffset    int
	YOffset    int
}

// New creates, instantiates, and returns a pointer to a new overlay Model.
func New(fore Viewable, back Viewable, xPos Position, yPos Position, xOff int, yOff int) *Model {
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

// View applies the compositing and handles rendering the view. It partly implements the tea.Model
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

	return Composite(
		m.Foreground.View(),
		m.Background.View(),
		m.XPosition,
		m.YPosition,
		m.XOffset,
		m.YOffset,
	)
}
