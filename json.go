package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(
	w http.ResponseWriter, code int, payload interface{}) {

	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(code)
	w.Write(data)
	w.Header().Add("Content-type", "application/json")

}

func respondWithError(w http.ResponseWriter, code int, msgString string) {
	if code > 499 {
		type msg struct {
			Error string `json:"error"`
		}

		respondWithJSON(w, code, msg{
			Error: msgString,
		})
	}
}