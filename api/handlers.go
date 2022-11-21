package main

import (
	"encoding/json"
	"net/http"
)

type JSONPayload struct {
	Index     int64  `json:"index"`
	Title     string `json:"title"`
	Book      string `json:"book"`
	PlainText string `json:"plainText"`
}

func (app *Config) Random(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONPayload
	var jsonResponse jsonResponse

	_ = app.readJSON(w, r, &requestPayload)

	if requestPayload.Index != 0 {
		// check if we have it in our database
		book, err := app.Models.Book.Fetch(requestPayload.Index)
		if err != nil && err.Error() == "mongo: no documents in result" {
			resp, err := app.fetchFromScraper(w, r)
			if err != nil {
				app.errorJSON(w, err, http.StatusBadRequest)
				return
			}

			// store it for future reference
			result, err := app.Models.Book.Create(resp)
			if err != nil {
				app.errorJSON(w, err, http.StatusBadRequest)
				return
			}

			jsonResponse.Error = false
			jsonResponse.Message = "Successfully fetched a book for you! <From Summarizer>"
			jsonResponse.Data = result
			app.writeJSON(w, http.StatusAccepted, jsonResponse)
			return
		} else {
			app.errorJSON(w, err, http.StatusBadRequest)
		}

		jsonResponse.Error = false
		jsonResponse.Message = "Successfully fetched a book for you! <From API>"
		jsonResponse.Data = book
		app.writeJSON(w, http.StatusAccepted, jsonResponse)
	}
}

func (app *Config) fetchFromScraper(w http.ResponseWriter, r *http.Request) (JSONPayload, error) {
	req, err := http.NewRequest("GET", "http://localhost:8050/books/random", nil)

	if err != nil {
		return JSONPayload{}, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return JSONPayload{}, err
	}

	defer resp.Body.Close()

	type Response struct {
		Error   bool        `json:"error"`
		Message string      `json:"message"`
		Data    JSONPayload `json:"data"`
	}

	var jsonFromSummarizer Response

	err = json.NewDecoder(resp.Body).Decode(&jsonFromSummarizer)

	data := jsonFromSummarizer.Data
	return data, nil
}
