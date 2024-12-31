package overlay

import (
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
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

// Model implements tea.Model, and manages the browser UI.
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

// offsets calculates the actual vertical and horizontal offsets used to position the foreground
// tea.Model relative to the background tea.Model.
func (m *Model) offsets() (int, int) {
	var x, y int
	switch m.XPosition {
	case Center:
		halfBackgroundWidth := (lipgloss.Width(m.Background.View()) + 1) / 2
		halfForegroundWidth := (lipgloss.Width(m.Foreground.View()) + 1) / 2
		x = halfBackgroundWidth - halfForegroundWidth
	case Right:
		x = lipgloss.Width(m.Background.View()) - lipgloss.Width(m.Foreground.View())
	}

	switch m.YPosition {
	case Center:
		halfBackgroundHeight := (lipgloss.Height(m.Background.View()) + 1) / 2
		halfForegroundHeight := (lipgloss.Height(m.Foreground.View()) + 1) / 2
		y = halfBackgroundHeight - halfForegroundHeight
	case Bottom:
		y = lipgloss.Height(m.Background.View()) - lipgloss.Height(m.Foreground.View())
	}

	debug(
		"X position: "+strconv.Itoa(int(m.XPosition)),
		"Y position: "+strconv.Itoa(int(m.YPosition)),
		"X offset: "+strconv.Itoa(x+m.XOffset),
		"Y offset: "+strconv.Itoa(y+m.YOffset),
		"Background width: "+strconv.Itoa(lipgloss.Width(m.Background.View())),
		"Foreground width: "+strconv.Itoa(lipgloss.Width(m.Foreground.View())),
		"Background height: "+strconv.Itoa(lipgloss.Height(m.Background.View())),
		"Foreground height: "+strconv.Itoa(lipgloss.Height(m.Foreground.View())),
	)

	return x + m.XOffset, y + m.YOffset
}

// composite merges and flattens the background and foreground views into a single view.
// This implementation is based off of the one used by Superfile -
// https://github.com/yorukot/superfile/blob/main/src/pkg/string_function/overplace.go
func (m *Model) composite() string {
	fg := m.Foreground.View()
	bg := m.Background.View()
	fgWidth, fgHeight := lipgloss.Size(fg)
	bgWidth, bgHeight := lipgloss.Size(bg)

	if fgWidth >= bgWidth && fgHeight >= bgHeight {
		return fg
	}

	x, y := m.offsets()
	x = clamp(x, 0, bgWidth-fgWidth)
	y = clamp(y, 0, bgHeight-fgHeight)

	fgLines := lines(fg)
	bgLines := lines(bg)
	var sb strings.Builder

	for i, bgLine := range bgLines {
		if i > 0 {
			sb.WriteByte('\n')
		}
		if i < y || i >= y+fgHeight {
			sb.WriteString(bgLine)
			continue
		}

		pos := 0
		if x > 0 {
			left := ansi.Truncate(bgLine, x, "")
			pos = ansi.StringWidth(left)
			sb.WriteString(left)
			if pos < x {
				sb.WriteString(whitespace(x - pos))
				pos = x
			}
		}

		fgLine := fgLines[i-y]
		sb.WriteString(fgLine)
		pos += ansi.StringWidth(fgLine)

		right := ansi.TruncateLeft(bgLine, pos, "")
		bgWidth := ansi.StringWidth(bgLine)
		rightWidth := ansi.StringWidth(right)
		if rightWidth <= bgWidth-pos {
			sb.WriteString(whitespace(bgWidth - rightWidth - pos))
		}
		sb.WriteString(right)
	}
	return sb.String()
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
	return m.composite()
}
