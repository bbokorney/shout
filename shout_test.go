package main

import (
	"fmt"
	"testing"

	"github.com/bbokorney/shout/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUsersExist(t *testing.T) {
	users := NewUsers(map[string]string{"user_id": "username"})
	mockTemplates := new(mocks.Templates)

	recipients := []string{"user_id"}
	templateName := "template-name"
	data := make(map[string]string)

	mockTemplates.On("Render", templateName, data).Return("", nil)

	shouter := NewShouter(users, mockTemplates)

	err := shouter.Send(recipients, templateName, data)

	assert.Nil(t, err)
	mockTemplates.AssertExpectations(t)
}

func TestUsersNotExist(t *testing.T) {
	users := NewUsers(map[string]string{"user_id": "username"})
	mockTemplates := new(mocks.Templates)

	recipients := []string{"not_a_user"}
	templateName := "template-name"
	data := make(map[string]string)

	shouter := NewShouter(users, mockTemplates)

	err := shouter.Send(recipients, templateName, data)

	assert.NotNil(t, err)
	mockTemplates.AssertExpectations(t)
}

func TestRenderTemplateError(t *testing.T) {
	users := NewUsers(map[string]string{"user_id": "username"})
	mockTemplates := new(mocks.Templates)

	recipients := []string{"user_id"}
	templateName := "template-name"
	data := make(map[string]string)

	mockTemplates.On("Render", templateName, data).Return("", fmt.Errorf("Expected error"))

	shouter := NewShouter(users, mockTemplates)

	err := shouter.Send(recipients, templateName, data)

	assert.NotNil(t, err)
	mockTemplates.AssertExpectations(t)
}
