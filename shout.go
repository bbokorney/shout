package main

import "fmt"

// Shouter coordinates sending the notifiction(s)
type Shouter interface {
	Send(recipients []string, templateName string, data map[string]string) error
}

// NewShouter returns a new Shouter
func NewShouter(users Users, templates Templates) Shouter {
	return shouter{
		users:     users,
		templates: templates,
	}
}

type shouter struct {
	users     Users
	templates Templates
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

	return nil
}
