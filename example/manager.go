package main

import (
	tea "github.com/charmbracelet/bubbletea"
	overlay "github.com/rmhubbert/bubbletea-overlay"
)

type sessionState int

const (
	mainView sessionState = iota
	modalView
)

// Manager implements tea.Model, and manages the browser UI.
type Manager struct {
	state        sessionState
	windowWidth  int
	windowHeight int
	foreground   tea.Model
	background   tea.Model
	overlay      tea.Model
}

// Init initialises the Manager on program load. It partly implements the tea.Model interface.
func (m *Manager) Init() tea.Cmd {
	m.state = mainView
	m.foreground = &Foreground{}
	m.background = &Background{}
	m.overlay = overlay.New(
		m.foreground,
		m.background,
		overlay.Center,
		overlay.Center,
		0,
		0,
	)
	return nil
}

// Update handles event and manages internal state. It partly implements the tea.Model interface.
func (m *Manager) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m, tea.Quit

		case " ":
			if m.state == mainView {
				m.state = modalView
			} else {
				m.state = mainView
			}
			return m, nil
		}
	}

	fg, fgCmd := m.foreground.Update(message)
	m.foreground = fg

	bg, bgCmd := m.background.Update(message)
	m.background = bg

	cmds := []tea.Cmd{}
	cmds = append(cmds, fgCmd, bgCmd)

	return m, tea.Batch(cmds...)
}

// View applies and styling and handles rendering the view. It partly implements the tea.Model
// interface.
func (m *Manager) View() string {
	if m.state == modalView {
		return m.overlay.View()
	}
	return m.background.View()
}
