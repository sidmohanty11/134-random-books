package main

import (
	"log"
	"net/http"
)

const PORT = "8081"
const HOST = "http://localhost"

type Config struct {
}

func main() {
	app := Config{}

	srv := &http.Server{
		Addr:    ":" + PORT,
		Handler: app.routes(),
	}

	log.Println("Listening at " + HOST + ":" + PORT)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
