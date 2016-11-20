package main

import "net/http"

func shout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
