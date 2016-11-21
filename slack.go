package main

// Notifications handles communicating with Slack
type Notifications interface {
	Send(recipient, message string) error
}

// NewNotifications creates a new Notifications
func NewNotifications() Notifications {
	return notifications{}
}

type notifications struct {
}

func (n notifications) Send(recipient, message string) error {
	return nil
}
