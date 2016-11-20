package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting up...")

	users := NewUsers(make(map[string]string))
	templates := NewTemplates()
	shouter := NewShouter(users, templates)
	shoutHandler := NewShoutHandler(shouter)

	http.HandleFunc("/shout", shoutHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
