package main

import (
	"bytes"
	"fmt"
	"html/template"
)

// Templates handles rendering message templates
type Templates interface {
	Render(templateName string, data map[string]string) (string, error)
}

// NewTemplates creates a new Templates
func NewTemplates(templateMapping map[string]*template.Template) Templates {
	return templates{
		templateMapping: templateMapping,
	}
}

type templates struct {
	templateMapping map[string]*template.Template
}

func (t templates) Render(templateName string, data map[string]string) (string, error) {
	tmpl, exists := t.templateMapping[templateName]
	if !exists {
		return "", fmt.Errorf("Template %s does not exist", templateName)
	}

	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, data)

	if err != nil {
		return "", fmt.Errorf("Error executing template %s with data %s: %s", templateName, data, err)
	}

	return buf.String(), nil
}
