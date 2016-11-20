package main

// Shouter coordinates sending the notifiction(s)
type Shouter interface {
	Send(recipients []string, templateName string, data map[string]string) error
}

// NewShouter returns a new Shouter
func NewShouter() Shouter {
	return shouter{}
}

type shouter struct {
}

// Send sends the notification(s)
func (s shouter) Send(recipients []string, templateName string, data map[string]string) error {

	return nil
}
