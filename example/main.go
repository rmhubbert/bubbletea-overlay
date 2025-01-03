package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	InitWriter()

	tui := &Manager{}
	p := tea.NewProgram(tui, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
