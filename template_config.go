package main

import (
	"fmt"
	"path/filepath"
	"text/template"
)

// ParseTemplates parses templates
func ParseTemplates(templatesDir string) (*template.Template, error) {
	t, err := template.ParseGlob(filepath.Join(templatesDir, "*"))
	if err != nil {
		return nil, fmt.Errorf("Error parsing templates in dir %s: %s", templatesDir, err)
	}
	return t, nil
}
