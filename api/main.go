package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const webPort = "8082"

type Config struct {
	Models Models
}

func main() {
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Fatal(err)
	}
	client = mongoClient

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: New(client),
	}

	log.Println("Starting server on port", webPort)

	srv := &http.Server{
		Addr:    ":" + webPort,
		Handler: app.routes(),
	}

	log.Fatal(srv.ListenAndServe())
}

func connectToMongo() (*mongo.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error Connecting:", err.Error())
		return nil, err
	}
	return c, nil
}
