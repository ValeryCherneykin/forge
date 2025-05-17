package templates

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetTemplates(templateDir string) ([]string, error) {
	if templateDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("failed to get home directory: %v", err)
		}
		templateDir = filepath.Join(home, ".forge", "templates")
	}

	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("template directory %s does not exist", templateDir)
	}

	entries, err := os.ReadDir(templateDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read template directory: %v", err)
	}

	var templates []string
	for _, entry := range entries {
		if !entry.IsDir() {
			templates = append(templates, entry.Name())
		}
	}
	return templates, nil
}

func CopyTemplates(fileName, templateDir string) error {
	if templateDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		templateDir = filepath.Join(home, ".forge", "templates")
	}
	src := filepath.Join(templateDir, fileName)
	dst := filepath.Join(".", fileName)

	if _, err := os.Stat(dst); !os.IsNotExist(err) {
		return fmt.Errorf("file %s already exists", fileName)
	}

	content, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, content, 0644)
}

func FilterTemplates(templates []string, query string) []string {
	var filtered []string
	query = strings.ToLower(query)
	for _, t := range templates {
		if strings.Contains(strings.ToLower(t), query) {
			filtered = append(filtered, t)
		}
	}
	return filtered
}
