package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

type Models struct {
	Book Book
}

type Book struct {
	ID        string    `json:"id" bson:"_id"`
	Title     string    `json:"title" bson:"title"`
	PlainText string    `json:"plainText" bson:"plainText"`
	Book      string    `json:"book" bson:"book"`
	Index     int64     `json:"index" bson:"index"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func New(mongo *mongo.Client) Models {
	client = mongo
	return Models{
		Book: Book{},
	}
}

func (b *Book) Create(book JSONPayload) (*mongo.InsertOneResult, error) {
	collection := client.Database("books").Collection("books")
	res, err := collection.InsertOne(context.TODO(), Book{
		ID:        primitive.NewObjectID().Hex(),
		Title:     book.Title,
		PlainText: book.PlainText,
		Book:      book.Book,
		Index:     book.Index,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error Inserting:", err.Error())
		return nil, err
	}
	return res, nil
}

func (b *Book) Fetch(index int64) (*Book, error) {
	collection := client.Database("books").Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var book Book
	err := collection.FindOne(ctx, bson.M{"index": index}).Decode(&book)
	if err != nil {
		log.Println("Error Finding:", err.Error())
		return nil, err
	}
	return &book, nil
}
