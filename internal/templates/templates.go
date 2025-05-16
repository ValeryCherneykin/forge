package templates

import (
	"fmt"
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

func CopyTemplates(fileName string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	src := filepath.Join(home, ".forge", "templates", fileName)
	dst := filepath.Join(".", fileName)

	if _, err = os.Stat(dst); !os.IsNotExist(err) {
		return fmt.Errorf("file %s is exist", fileName)
	}

	content, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, content, 0644)
}
