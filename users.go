package main

import "fmt"

// Users stores the user ID -> Slack username
type Users interface {
	Lookup(id string) (string, error)
}

type users struct {
	mapping map[string]string
}

func (u users) Lookup(id string) (string, error) {
	username, exists := u.mapping[id]
	if !exists {
		return "", fmt.Errorf("User with ID %s not found", id)
	}
	return username, nil
}
