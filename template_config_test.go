package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTemplatesDir = filepath.Join(testDataDir, "templates")

func TestTemplateDirNotExist(t *testing.T) {
	_, err := ParseTemplates("not-a-dir")
	assert.NotNil(t, err)
}

func TestParseTemplates(t *testing.T) {
	tmpls, err := ParseTemplates(testTemplatesDir)
	assert.Nil(t, err)

	assert.NotNil(t, tmpls.Lookup("years-old"))
	assert.NotNil(t, tmpls.Lookup("dogs"))
}
