package templates

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

var Templates *template.Template

// InitializeTemplates recursively parses all templates in the "templates" directory
func InitializeTemplates() error {
	Templates = template.New("")

	err := filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			_, err := Templates.ParseFiles(path)
			if err != nil {
				return fmt.Errorf("error parsing template %s: %w", path, err)
			}
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("error initializing templates: %w", err)
	}

	fmt.Println("Templates initialized successfully...")
	return nil
}
