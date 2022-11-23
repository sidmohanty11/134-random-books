package main

import (
	"encoding/json"
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", "http://localhost:8050/books/random", nil)

	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	defer resp.Body.Close()

	var jsonFromSummarizer jsonResponse

	err = json.NewDecoder(resp.Body).Decode(&jsonFromSummarizer)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	jsonFromSummarizer.Message = "Successfully fetched a book for you!<From Broker>"
	app.writeJSON(w, http.StatusAccepted, jsonFromSummarizer)
}
