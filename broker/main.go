package main

import (
	"log"
	"net/http"
)

const PORT = "80"

type Config struct {
}

func main() {
	app := Config{}

	srv := &http.Server{
		Addr:    ":" + PORT,
		Handler: app.routes(),
	}

	log.Println("Listening at :", PORT)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
