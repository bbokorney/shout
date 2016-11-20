package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDataDir = "testdata"

type userConfigTestCase struct {
	filename         string
	exepctedMapping  map[string]string
	expectedErrorNil bool
}

func TestUserConfig(t *testing.T) {
	testCases := []userConfigTestCase{
		userConfigTestCase{"does-not-exist.yml", nil, false},
		userConfigTestCase{"single-user.yml", map[string]string{
			"user_id": "username"}, true},
		userConfigTestCase{"multiple-users.yml", map[string]string{
			"alice@example.com": "alice",
			"Robert Johnson":    "robby"}, true},
		userConfigTestCase{"not-yaml.txt", map[string]string{"who": "cares"}, false},
	}

	for _, tc := range testCases {
		mapping, err := ReadUsersFile(filepath.Join(testDataDir, tc.filename))

		if tc.expectedErrorNil {
			assert.Nil(t, err)
			assert.Equal(t, tc.exepctedMapping, mapping)
		} else {
			assert.NotNil(t, err)
		}
	}
}
