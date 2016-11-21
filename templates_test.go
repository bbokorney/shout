package main

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateNotExist(t *testing.T) {
	templates := NewTemplates(make(map[string]*template.Template))
	_, err := templates.Render("not a template", nil)
	assert.NotNil(t, err)
}

func TestTemplateRender(t *testing.T) {
	tmpl := template.Must(template.New("tmpl").Parse(`Value of key is "{{index . "key"}}"`))
	data := map[string]string{"key": "value"}
	tmplMap := make(map[string]*template.Template)
	tmplMap["template"] = tmpl
	templates := NewTemplates(tmplMap)

	msg, err := templates.Render("template", data)

	assert.Nil(t, err)
	assert.Equal(t, `Value of key is "value"`, msg)
}
