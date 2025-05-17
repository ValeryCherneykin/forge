package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ValeryCherneykin/forge/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	templateDir := flag.String("dir", "", "Custom template directory")
	flag.Parse()

	model := tui.NewModel(*templateDir)
	p := tea.NewProgram(model, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
