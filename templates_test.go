package main

import (
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestTemplateNotExist(t *testing.T) {
	tmpl := template.Must(template.New("template").Parse(`Value of key is "{{index . "key"}}"`))
	templates := NewTemplates(tmpl)
	_, err := templates.Render("not a template", nil)
	assert.NotNil(t, err)
}

func TestTemplateRender(t *testing.T) {
	tmpl := template.Must(template.New("template").Parse(`Value of key is "{{index . "key"}}"`))
	data := map[string]string{"key": "value"}
	templates := NewTemplates(tmpl)

	msg, err := templates.Render("template", data)

	assert.Nil(t, err)
	assert.Equal(t, `Value of key is "value"`, msg)
}
