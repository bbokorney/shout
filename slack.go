package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Notifications handles communicating with Slack
type Notifications interface {
	Send(recipient, message string) error
}

// NewNotifications creates a new Notifications
func NewNotifications(slackURL string) Notifications {
	return notifications{
		slackURL: slackURL,
	}
}

type notifications struct {
	slackURL string
}

func (n notifications) Send(recipient, message string) error {
	msgBody := slackMessageBody{
		Channel: recipient,
		Text:    message,
	}

	marshalledBody, err := json.Marshal(msgBody)
	if err != nil {
		return fmt.Errorf("Error marshalling Slack request body: %s", err)
	}

	resp, err := http.Post(n.slackURL, "application/json", bytes.NewReader(marshalledBody))
	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("Error sending Slack request, unexpected response code %d: %s", resp.StatusCode, err)
		respBody, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			errMsg = fmt.Sprintf("%s: %s", errMsg, string(respBody))
		}
		return errors.New(errMsg)
	}

	return nil
}

type slackMessageBody struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}
