package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting up...")

	users := NewUsers(make(map[string]string))
	templates := NewTemplates(make(map[string]*template.Template))
	notifications := NewNotifications()
	shouter := NewShouter(users, templates, notifications)
	shoutHandler := NewShoutHandler(shouter)

	http.HandleFunc("/shout", shoutHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
