package main

import "fmt"

// Users stores the user ID -> Slack username
type Users interface {
	Lookup(id string) (string, error)
}

// NewUsers returns a new Users
func NewUsers(mapping map[string]string) Users {
	return users{
		mapping: mapping,
	}
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
