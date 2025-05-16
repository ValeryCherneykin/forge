package tui

import "github.com/ValeryCherneykin/forge/internal/templates"

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
