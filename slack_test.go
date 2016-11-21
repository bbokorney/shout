package main

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var slackTest bool

func init() {
	flag.BoolVar(&slackTest, "slack-test", false, "Run test with actual Slack API")
}

func TestSlackSend(t *testing.T) {
	if !slackTest {
		t.Skip()
	}

	slackURL := os.Getenv("SLACK_URL")
	assert.NotEmpty(t, slackURL)

	notifications := NewNotifications(slackURL)
	err := notifications.Send("bbokorney", "Hey there from test\nDoes this work?")
	assert.Nil(t, err)
}
