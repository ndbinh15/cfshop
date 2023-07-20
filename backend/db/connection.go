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
	// Access the desired database and collection
	db := client.Database("myDatabase")  // Replace 'mydatabase' with your actual database name
	collection := db.Collection("users") // Replace 'users' with your actual collection name

	return collection
}
func GetProductCollection() *mongo.Collection {
	// Access the desired database and collection
	db := client.Database("myDatabase")     // Replace 'mydatabase' with your actual database name
	collection := db.Collection("products") // Replace 'users' with your actual collection name

	return collection
}
func GetCategoryCollection() *mongo.Collection {
	// Access the desired database and collection
	db := client.Database("myDatabase")       // Replace 'mydatabase' with your actual database name
	collection := db.Collection("categories") // Replace 'users' with your actual collection name

	return collection
}
