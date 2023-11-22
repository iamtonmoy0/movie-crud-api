package controllers

import (
	"context"
	"fmt"
	"log"
	"os"

	model "github.com/iamtonmoy0/movie-crud-api/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const connectionString = os.Getenv("DB")
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// connect with mongoDB

func init() {
	//client option
	clientOption := options.Client().ApplyURI(os.Getenv("DB"))

	//connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie with ID: ", inserted.InsertedID)
}
