package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var integrationTest bool

func init() {
	flag.BoolVar(&integrationTest, "integration-test", false, "Run integration test with actual Slack API")
}

func TestIntegration(t *testing.T) {
	if !integrationTest {
		t.Skip()
	}

	slackURL := os.Getenv("SLACK_URL")
	assert.NotEmpty(t, slackURL)

	notifications := NewNotifications(slackURL)

	userMap, err := ReadUsersFile(filepath.Join(testDataDir, "real-user.yml"))
	assert.Nil(t, err)
	users := NewUsers(userMap)

	parsedTemplates, err := ParseTemplates(testTemplatesDir)
	assert.Nil(t, err)
	templates := NewTemplates(parsedTemplates)

	shouter := NewShouter(users, templates, notifications)
	shoutHandler := NewShoutHandler(shouter)

	ts := httptest.NewServer(shoutHandler)
	defer ts.Close()

	requestBody := `
{
  "recipients":["baker"],
  "template":"years-old",
  "data":{
    "name":"Rufus",
    "years":"eleventeen"
  }
}`
	res, err := http.Post(fmt.Sprintf("%s/%s", ts.URL, "shout"),
		"application/json", strings.NewReader(requestBody))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusAccepted, res.StatusCode)
}
