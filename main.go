package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting up...")

	parsedTemplates, err := ParseTemplates("templates")
	if err != nil {
		log.Fatal(err)
	}

	users := NewUsers(make(map[string]string))
	templates := NewTemplates(parsedTemplates)
	notifications := NewNotifications("http://fake-url")
	shouter := NewShouter(users, templates, notifications)
	shoutHandler := NewShoutHandler(shouter)

	http.HandleFunc("/shout", shoutHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
