package projects

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Define a model to store the counter state
type model struct {
	count int
}

// Messages representing user input
type msg string

const (
	increment msg = "increment"
	decrement msg = "decrement"
	quit      msg = "quit"
)

// Init initializes the program (no initial commands needed here)
func (m model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model accordingly
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg: // Check if a key was pressed
		switch msg.String() {
		case "q": // Quit the program
			return m, tea.Quit
		case "up": // Increment counter
			m.count++
		case "down": // Decrement counter
			m.count--
		}
	}
	return m, nil
}

// View renders the UI
func (m model) View() string {
	return fmt.Sprintf("Counter: %d\n[↑] Increase  [↓] Decrease  [q] Quit", m.count)
}

func main() {
	// Start the Bubble Tea program
	p := tea.NewProgram(model{})
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting app: %v", err)
		os.Exit(1)
	}
}
