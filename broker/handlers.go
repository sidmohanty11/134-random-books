package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

type RandomInt struct {
	Index int64 `json:"index"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(generateRandomInt())
	if err != nil {
		log.Println(err.Error())
		return
	}
	req, err := http.NewRequest("POST", "http://localhost:8082/random", bytes.NewBuffer(jsonData))

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

func generateRandomInt() RandomInt {
	return RandomInt{
		Index: int64(rand.Intn(134)),
	}
}
