package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model implements tea.Model, and manages the browser UI.
type Foreground struct {
	windowWidth  int
	windowHeight int
}

// Init initialises the Model on program load. It partly implements the tea.Model interface.
func (m *Foreground) Init() tea.Cmd {
	return nil
}

// Update handles event and manages internal state. It partly implements the tea.Model interface.
func (m *Foreground) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
	}

	return m, cmd
}

// View applies and styling and handles rendering the view. It partly implements the tea.Model
// interface.
func (m *Foreground) View() string {
	foreStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(lipgloss.Color("6")).
		Padding(0, 1)

	boldStyle := lipgloss.NewStyle().Bold(true)
	title := boldStyle.Render("Bubble Tea Overlay")
	content := "Hello! I'm in a modal window.\n\nPress <space> to close the window."
	layout := lipgloss.JoinVertical(lipgloss.Left, title, content)

	return foreStyle.Render(layout)
}
