package main

import (
	"bytes"
	"fmt"
	"text/template"
)

// Templates handles rendering message templates
type Templates interface {
	Render(templateName string, data map[string]string) (string, error)
}

// NewTemplates creates a new Templates
func NewTemplates(tmpls *template.Template) Templates {
	return templates{
		tmpls: tmpls,
	}
}

type templates struct {
	tmpls *template.Template
}

func (t templates) Render(templateName string, data map[string]string) (string, error) {
	tmpl := t.tmpls.Lookup(templateName)
	if tmpl == nil {
		return "", fmt.Errorf("Template %s does not exist", templateName)
	}

	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, data)

	if err != nil {
		return "", fmt.Errorf("Error executing template %s with data %s: %s", templateName, data, err)
	}

	return buf.String(), nil
}
