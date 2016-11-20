package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShout(t *testing.T) {
	type testCase struct {
		recipients       []string
		templateName     string
		data             map[string]string
		users            Users
		expectedErrorNil bool
	}

	testCases := []testCase{
		testCase{
			recipients:       []string{"user_id"},
			templateName:     "template-name",
			data:             nil,
			users:            NewUsers(map[string]string{"user_id": "username"}),
			expectedErrorNil: false,
		},
		testCase{
			recipients:       []string{"does_not_exist"},
			templateName:     "template-name",
			data:             nil,
			users:            NewUsers(map[string]string{"user_id": "username"}),
			expectedErrorNil: true,
		},
	}

	for _, tc := range testCases {
		shouter := NewShouter(tc.users)
		err := shouter.Send(tc.recipients, tc.templateName, tc.data)

		assert.Equal(t, tc.expectedErrorNil, err != nil)
	}
}
