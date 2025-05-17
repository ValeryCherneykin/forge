package tui

import (
	"fmt"
	"strings"

	"github.com/ValeryCherneykin/forge/internal/icons"
	"github.com/ValeryCherneykin/forge/internal/templates"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const maxVisibleItems = 4

type Model struct {
	templates    []string
	filtered     []string
	cursor       int
	startIndex   int
	filterInput  string
	filtering    bool
	message      string
	templateDir  string
	SelectedFile string
}

func NewModel(templateDir string) Model {
	templates, err := templates.GetTemplates(templateDir)
	if err != nil {
		return Model{
			message:     fmt.Sprintf("Error: %v", err),
			templateDir: templateDir,
		}
	}
	return Model{
		templates:   templates,
		filtered:    templates,
		templateDir: templateDir,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.filtering {
			switch msg.String() {
			case "enter":
				m.filtering = false
				m.templates = m.filtered
				m.cursor = 0
				m.startIndex = 0
				m.message = "Filter applied."
			case "esc":
				m.filtering = false
				m.filterInput = ""
				m.filtered = m.templates
				m.message = "Filter cleared."
			default:
				if len(msg.String()) == 1 || msg.String() == "backspace" {
					if msg.String() == "backspace" && len(m.filterInput) > 0 {
						m.filterInput = m.filterInput[:len(m.filterInput)-1]
					} else if msg.String() != "backspace" {
						m.filterInput += msg.String()
					}
					m.filtered = templates.FilterTemplates(m.templates, m.filterInput)
					m.cursor = 0
					m.startIndex = 0
				}
			}
			return m, nil
		}
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
				if m.cursor < m.startIndex {
					m.startIndex--
				}
			}
		case "down", "j":
			if m.cursor < len(m.filtered)-1 {
				m.cursor++
				if m.cursor >= m.startIndex+maxVisibleItems {
					m.startIndex++
				}
			}
		case "enter", "p":
			if len(m.filtered) > 0 {
				if err := templates.CopyTemplates(m.filtered[m.cursor], m.templateDir); err != nil {
					m.message = fmt.Sprintf("Error: %v", err)
					return m, nil
				}
				m.SelectedFile = m.filtered[m.cursor]
				return m, tea.Quit
			}
		case ":":
			m.filtering = true
			m.filterInput = ""
			m.message = "Enter filter query (esc to cancel):"
			return m, nil
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := ""

	if m.message != "" {
		messageStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5555"))
		if strings.Contains(m.message, "successfully") {
			messageStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#50fa7b"))
		}
		s += messageStyle.Render(m.message) + "\n"
	}

	if m.filtering {
		s += fmt.Sprintf("Filter: %s\n", m.filterInput)
	}

	if len(m.filtered) == 0 {
		s += "No templates found.\n"
	} else {
		endIndex := m.startIndex + maxVisibleItems
		if endIndex > len(m.filtered) {
			endIndex = len(m.filtered)
		}

		for i := m.startIndex; i < endIndex; i++ {
			template := m.filtered[i]
			cursor := "  "
			if m.cursor == i {
				cursor = "❯ "
			}
			icon := icons.GetIcon(template)
			style := lipgloss.NewStyle().Foreground(lipgloss.Color("#c0caf5"))
			iconStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(icon.Color))
			if m.cursor == i {
				style = style.
					Bold(true).
					Foreground(lipgloss.Color("#7aa2f7")).
					Background(lipgloss.Color("#3b4261"))
			}
			s += style.Render(cursor+iconStyle.Render(icon.Symbol)+" "+template) + "\n"
		}

		highlightStyle := lipgloss.NewStyle().Background(lipgloss.Color("#3b4261")).Height(1)
		if len(m.filtered) > 0 {
			s += highlightStyle.Render(strings.Repeat(" ", lipgloss.Width(m.filtered[m.cursor])+2)) + "\n"
		}
	}

	style := lipgloss.NewStyle().
		Padding(0, 4).
		Foreground(lipgloss.Color("#c0caf5")).
		Background(lipgloss.Color("#1a1b26"))

	return style.Render(s)
}
