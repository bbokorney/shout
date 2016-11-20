package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsers(t *testing.T) {

	type testCase struct {
		userMapping      map[string]string
		userID           string
		expectedUsername string
		expectedErrorNil bool
	}

	testCases := []testCase{
		testCase{map[string]string{"user_id": "username"}, "user_id", "username", true},
		testCase{map[string]string{}, "user_id", "username", false},
	}

	for _, tc := range testCases {
		users := NewUsers(tc.userMapping)

		username, err := users.Lookup(tc.userID)

		if tc.expectedErrorNil {
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedUsername, username)
		} else {
			assert.NotNil(t, err)
		}
	}
}
