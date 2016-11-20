package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bbokorney/shout/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHttpBadFormat(t *testing.T) {
	mockShouter := new(mocks.Shouter)
	requestBody := `{"key":"value"}`
	ts := httptest.NewServer(NewShoutHandler(mockShouter))
	defer ts.Close()

	res, err := http.Post(fmt.Sprintf("%s/%s", ts.URL, "shout"),
		"application/json", strings.NewReader(requestBody))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	mockShouter.AssertExpectations(t)
}

func TestHttpInvalidJson(t *testing.T) {
	mockShouter := new(mocks.Shouter)
	requestBody := `Not valid json`
	ts := httptest.NewServer(NewShoutHandler(mockShouter))
	defer ts.Close()

	res, err := http.Post(fmt.Sprintf("%s/%s", ts.URL, "shout"),
		"application/json", strings.NewReader(requestBody))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	mockShouter.AssertExpectations(t)
}

func TestHttpInvalidMethod(t *testing.T) {
	mockShouter := new(mocks.Shouter)
	ts := httptest.NewServer(NewShoutHandler(mockShouter))
	defer ts.Close()

	res, err := http.Get(fmt.Sprintf("%s/%s", ts.URL, "shout"))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
	mockShouter.AssertExpectations(t)
}

func TestHttpNoShoutError(t *testing.T) {
	mockShouter := new(mocks.Shouter)
	mockShouter.On("Send", []string{"user_id"},
		"template-name", map[string]string{"key": "value"}).
		Return(nil)

	requestBody := `{"recipients": ["user_id"],
	"template":"template-name",
	"data":{
		"key":"value"
	}
}`
	ts := httptest.NewServer(NewShoutHandler(mockShouter))
	defer ts.Close()

	res, err := http.Post(fmt.Sprintf("%s/%s", ts.URL, "shout"),
		"application/json", strings.NewReader(requestBody))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusAccepted, res.StatusCode)
	mockShouter.AssertExpectations(t)
}
func TestHttpShoutError(t *testing.T) {
	mockShouter := new(mocks.Shouter)
	mockShouter.On("Send", []string{"user_id"},
		"template-name", map[string]string{"key": "value"}).
		Return(fmt.Errorf("Shouter.Send error"))

	requestBody := `{"recipients": ["user_id"],
	"template":"template-name",
	"data":{
		"key":"value"
	}
}`
	ts := httptest.NewServer(NewShoutHandler(mockShouter))
	defer ts.Close()

	res, err := http.Post(fmt.Sprintf("%s/%s", ts.URL, "shout"),
		"application/json", strings.NewReader(requestBody))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	mockShouter.AssertExpectations(t)
}

func TestValidate(t *testing.T) {
	type testCase struct {
		request          shoutRequest
		expectedErrorNil bool
	}

	testCases := []testCase{
		testCase{
			request:          shoutRequest{[]string{"user_id"}, "template-name", nil},
			expectedErrorNil: true,
		},
		testCase{
			request:          shoutRequest{[]string{}, "template-name", nil},
			expectedErrorNil: false,
		},
		testCase{
			request:          shoutRequest{[]string{"user_id", ""}, "template-name", nil},
			expectedErrorNil: false,
		},
		testCase{
			request:          shoutRequest{[]string{"user_id"}, "", nil},
			expectedErrorNil: false,
		},
	}

	for _, tc := range testCases {
		err := validateRequest(tc.request)
		assert.Equal(t, tc.expectedErrorNil, err == nil)
	}
}
