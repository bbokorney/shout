package main

import "fmt"

// Shouter coordinates sending the notifiction(s)
type Shouter interface {
	Send(recipients []string, templateName string, data map[string]string) error
}

// NewShouter returns a new Shouter
func NewShouter(users Users, templates Templates, notifications Notifications) Shouter {
	return shouter{
		users:         users,
		templates:     templates,
		notifications: notifications,
	}
}

type shouter struct {
	users         Users
	templates     Templates
	notifications Notifications
}

// Send sends the notification(s)
func (s shouter) Send(recipients []string, templateName string, data map[string]string) error {
	// translate the user ids into Slack usernames
	var usernames []string
	for _, userID := range recipients {
		username, err := s.users.Lookup(userID)
		if err != nil {
			return fmt.Errorf("Error retrieving user: %s", userID)
		}
		usernames = append(usernames, username)
	}

	// render the message
	message, err := s.templates.Render(templateName, data)
	if err != nil {
		return fmt.Errorf("Error rendering message: %s", err)
	}

	fmt.Println("Rendered message: ", message)

	// send the message to the recipients
	for _, username := range usernames {
		if err := s.notifications.Send(username, message); err != nil {
			return fmt.Errorf("Error message %s to %s: %s", message, username, err)
		}
	}

	return nil
}
