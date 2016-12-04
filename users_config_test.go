package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDataDir = "testdata"

func TestUserConfig(t *testing.T) {
	type testCase struct {
		filename         string
		exepctedMapping  map[string]string
		expectedErrorNil bool
	}

	testCases := []testCase{
		testCase{"does-not-exist.yml", nil, false},
		testCase{"single-user.yml", map[string]string{
			"user_id": "@username"}, true},
		testCase{"multiple-users.yml", map[string]string{
			"alice@example.com": "@alice",
			"Robert Johnson":    "@robby"}, true},
		testCase{"not-yaml.txt", map[string]string{"who": "cares"}, false},
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
