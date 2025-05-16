package templates

import (
	"os"
	"path/filepath"
)

func GetTemplates() []string {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	dir := filepath.Join(home, ".forge", "templates")
	writes, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}

	var templates []string

	for _, write := range writes {
		if !write.IsDir() {
			templates = append(templates, write.Name())
		}
	}

	return templates
}
