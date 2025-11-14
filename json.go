package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code >= 500 {
		log.Printf("Internal server error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	responseWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal payload: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
	// Alternatively, you can use the json.Encoder
	// json.NewEncoder(w).Encode(payload)
}
