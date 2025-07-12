package overlay

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// debug prints its input to a debug file, but only when a DEBUG environment variable has been set.
func debug(lines ...string) {
	if len(os.Getenv("DEBUG")) > 0 {
		ts := strconv.Itoa(int(time.Now().UnixMilli())) + "\n"
		s := strings.Join(lines, "\n")

		f, err := tea.LogToFile("debug.log", ts+s)
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer func() {
			err := f.Close()
			if err != nil {
				fmt.Printf("could not close debug file: %v", err)
			}
		}()
	}
}
