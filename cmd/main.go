package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ValeryCherneykin/forge/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	templateDir := flag.String("dir", "", "Custom template directory")
	flag.Parse()

	model := tui.NewModel(*templateDir)
	p := tea.NewProgram(model)
	result, err := p.StartReturningModel()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if m, ok := result.(tui.Model); ok && m.SelectedFile != "" {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error getting current directory: %v\n", err)
			os.Exit(1)
		}
		fullPath := filepath.Join(currentDir, m.SelectedFile)
		fmt.Print("\033[2J\033[1;1H")
		fmt.Printf("\033[32mCreate file %s\033[0m\n", fullPath)
	}
}
