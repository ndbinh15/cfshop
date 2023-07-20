package handler

import (
	"cfshop/backend/db"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	ID           primitive.ObjectID `json:"_id,omitempty"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	Price        float64            `json:"price"`
	Ingredients  []string           `json:"ingredients"`
	Availability bool               `json:"availability"`
	Category     Category           `json:"category"`
}

// Insert a new product
func InsertProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		log.Println("$POST add product success")
	}

	// Parse the request body to get the product data
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the MongoDB collection for products
	productCollection := db.GetProductCollection()

	// Insert the product into the products collection
	_, err = productCollection.InsertOne(r.Context(), product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusCreated)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		log.Println("$GET product success")
	}
	// Get the MongoDB collection for products
	productCollection := db.GetProductCollection()

	// Fetch all products from the collection
	cursor, err := productCollection.Find(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(r.Context())

	// Collect products in a slice
	var products []Product
	for cursor.Next(r.Context()) {
		var product Product
		err := cursor.Decode(&product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}
	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode products as JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Find a product by ID
func FindProductByID(client *mongo.Client, productID primitive.ObjectID) (*Product, error) {
	collection := client.Database("myDatabase").Collection("products")
	var product Product
	err := collection.FindOne(context.TODO(), bson.M{"_id": productID}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Update a product
func UpdateProduct(client *mongo.Client, productID primitive.ObjectID, updatedProduct Product) error {
	collection := client.Database("myDatabase").Collection("products")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": productID}, bson.M{"$set": updatedProduct})
	return err
}

// Delete a product
func DeleteProduct(client *mongo.Client, productID primitive.ObjectID) error {
	collection := client.Database("myDatabase").Collection("products")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": productID})
	return err
}
