package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting up...")

	http.HandleFunc("/shout", shout)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
