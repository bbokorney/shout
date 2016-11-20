package main

import "fmt"

// Shouter coordinates sending the notifiction(s)
type Shouter interface {
	Send(recipients []string, templateName string, data map[string]string) error
}

// NewShouter returns a new Shouter
func NewShouter(users Users) Shouter {
	return shouter{
		users: users,
	}
}

type shouter struct {
	users Users
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

	return nil
}
