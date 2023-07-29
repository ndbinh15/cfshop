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
	Category     string             `json:"category"`
}

// Insert a new product
func InsertProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method InsertProduct not allowed", http.StatusMethodNotAllowed)
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"success": true,
		"message": "Product created successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method GetAllProducts not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		log.Println("$GET all product success")
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

func GetProductCount(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method GetProductCount not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		log.Println("$GET product count success")
	}
	// Get the MongoDB collection for products
	productCollection := db.GetProductCollection()

	// Query the database for the count of all products
	count, err := productCollection.CountDocuments(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a JSON response containing the count
	response := struct {
		Count int64 `json:"count"`
	}{
		Count: count,
	}

	// Encode the response as JSON and send it back
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
