package tui

import (
	"fmt"
	"github.com/ValeryCherneykin/forge/internal/icons"
	"github.com/ValeryCherneykin/forge/internal/templates"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

type Model struct {
	templates []string
	cursor    int
}

func NewModel() Model {
	return Model{
		templates: templates.GetTemplates(),
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
					os.Exit(1)
				}
				fmt.Println(m.templates[m.cursor])
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := "Forge: Шаблонизатор\n\n"
	if len(m.templates) == 0 {
		s += "None ~/.forge/templates/."
	} else {
		for i, template := range m.templates {
			cursor := "  "
			if m.cursor == i {
				cursor = "❯ "
			}
			icon := icons.GetIcon(template)
			style := lipgloss.NewStyle().
				Foreground(lipgloss.Color("#c0caf5"))
			iconStyle := lipgloss.NewStyle().
				Foreground(lipgloss.Color(icon.Color))
			if m.cursor == i {
				style = style.Bold(true).Foreground(lipgloss.Color("#7aa2f7"))
			}
			s += style.Render(cursor+iconStyle.Render(icon.Symbol)+" "+template) + "\n"
		}
	}
	s += "\nq — выход, j/k — навигация, enter/p — выбрать"

	style := lipgloss.NewStyle().
		Padding(2, 4).
		Foreground(lipgloss.Color("#c0caf5")).
		Background(lipgloss.Color("#1a1b26")).
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(lipgloss.Color("#7aa2f7"))

	return style.Render(s)
}
