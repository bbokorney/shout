package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// NewShoutHandler returns a new shout HTTP handler
func NewShoutHandler(shouter Shouter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req shoutRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Bad request format: %s", err)))
			return
		}

		log.Printf("Request shout received: %+v", req)

		err = validateRequest(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Bad request format: %s", err)))
			return
		}

		err = shouter.Send(req.Recipients, req.Template, req.Data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}

type shoutRequest struct {
	Recipients []string          `json:"recipients"`
	Template   string            `json:"template"`
	Data       map[string]string `json:"data"`
}

func validateRequest(req shoutRequest) error {
	if len(req.Recipients) == 0 {
		return fmt.Errorf("Must have at least one recipient")
	}

	for _, r := range req.Recipients {
		if len(r) == 0 {
			return fmt.Errorf("Recipients must not be empty")
		}
	}

	if len(req.Template) == 0 {
		return fmt.Errorf("Template name must not be empty")
	}

	return nil
}
