package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() {
	// MongoDB Atlas connection string
	connectionString := "mongodb+srv://cfshop.qtejzjh.mongodb.net/?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority&tlsCertificateKeyFile=../X509-cert-4058952649266884251.pem"

	// Create a MongoDB client
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB Atlas successfully.")
}

func GetUserCollection() *mongo.Collection {
	db := client.Database("myDatabase2")
	collection := db.Collection("users")

	return collection
}
func GetProductCollection() *mongo.Collection {
	db := client.Database("myDatabase2")
	collection := db.Collection("products")

	return collection
}
func GetCategoryCollection() *mongo.Collection {
	db := client.Database("myDatabase2")
	collection := db.Collection("categories")

	return collection
}
func GetCartCollection() *mongo.Collection {
	db := client.Database("myDatabase2")
	collection := db.Collection("cart")

	return collection
}
func GetOrderCollection() *mongo.Collection {
	db := client.Database("myDatabase2")
	collection := db.Collection("order")

	return collection
}
