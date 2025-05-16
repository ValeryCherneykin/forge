package tui

import (
	"fmt"
	"os"

	"github.com/ValeryCherneykin/forge/internal/templates"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	templates []string
	cursor    int
}

func NewModel() Model {
	templates := templates.GetTemplates()
	return Model{
		templates: templates,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.templates)-1 {
				m.cursor++
			}
		case "enter", "p":
			if len(m.templates) > 0 {
				if err := templates.CopyTemplates(m.templates[m.cursor]); err != nil {
					fmt.Fprintf(os.Stderr, "error: %v\n", err)
				} else {
					fmt.Printf("Шаблон %s скопирован\n", m.templates[m.cursor])
				}
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := "Forge \n\n"
	if len(m.templates) == 0 {
		s += "None found in ~/.forge/templates/."
	} else {
		for i, template := range m.templates {
			cursor := "  "
			if m.cursor == i {
				cursor = "➜ "
			}
			s += cursor + template + "\n"
		}
	}
	s += "\nq — exit, j/k — up/down"
	return s
}
