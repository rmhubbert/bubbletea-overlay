package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Manager implements tea.Manager, and manages the browser UI.
type Background struct {
	windowWidth  int
	windowHeight int
}

// Init initialises the Manager on program load. It partly implements the tea.Manager interface.
func (m *Background) Init() tea.Cmd {
	return nil
}

// Update handles event and manages internal state. It partly implements the tea.Manager interface.
func (m *Background) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
	}

	return m, cmd
}

// View applies and styling and handles rendering the view. It partly implements the tea.Manager
// interface.
func (m *Background) View() string {
	backStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(lipgloss.Color("6")).
		Width(m.windowWidth - 2).
		Height(m.windowHeight - 5).
		Foreground(lipgloss.Color("8"))

	var row, mainContent string
	for range (m.windowWidth / 2) - 1 {
		row += "XO"
	}
	for x := range m.windowHeight - 5 {
		if x > 0 {
			mainContent += "\n"
		}
		mainContent += row
	}

	footerStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(lipgloss.Color("6")).
		Align(lipgloss.Center).
		Width(m.windowWidth - 2)
	footerContent := "Press <space> to toggle the modal window. Press q or <esc> to quit."

	content := backStyle.Render(mainContent)
	footer := footerStyle.Render(footerContent)

	return lipgloss.JoinVertical(lipgloss.Left, content, footer)
}
