package handler

import (
	"cfshop/backend/db"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID           primitive.ObjectID `json:"_id,omitempty"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	Price        float64            `json:"price"`
	Quantity     int64              `json:"quantity"`
	Ingredients  []string           `json:"ingredients"`
	Availability bool               `json:"availability"`
	Category     string             `json:"category"`
	Image        string             `json:"image"`
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

	if product.Quantity == 0 {
		product.Quantity = 0
	}

	if product.Image == "" {
		product.Image = "default.jpg"
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

	// // Return a success response
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// response := map[string]interface{}{
	// 	"success": true,
	// 	"message": "Product created successfully",
	// }
	// json.NewEncoder(w).Encode(response)
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

	// // Return a success response
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// response := map[string]interface{}{
	// 	"success": true,
	// 	"message": "Product created successfully",
	// }
	// json.NewEncoder(w).Encode(response)
}

func UpdateProductQuantity(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method UpdateProductQuantity not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		log.Println("$POST UpdateProductQuantity success")
	}

	// Parse the request body to get the JSON data
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println("error decode")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the MongoDB collection for products
	productCollection := db.GetProductCollection()

	// Check if the product exists based on its name
	filter := bson.M{"name": product.Name}
	log.Println(filter)
	existingProduct := productCollection.FindOne(r.Context(), filter)

	if existingProduct.Err() == nil {
		update := bson.M{"$set": bson.M{"quantity": product.Quantity}}
		_, err := productCollection.UpdateOne(r.Context(), filter, update)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		log.Println("error found")
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// w.WriteHeader(http.StatusOK)
	// Return a success response
	w.Header().Set("Content-Type", "application/json")
	log.Println("error 1")
	w.WriteHeader(http.StatusOK)
	log.Println("error 2")
	response := map[string]interface{}{
		"success": true,
		"message": "Add product's quantity successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method DeleteProduct not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		log.Println("$POST DeleteProduct success")
	}

	var request struct {
		Name string `json:"name"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productCollection := db.GetProductCollection()

	filter := bson.M{"name": request.Name}
	log.Println(request.Name)
	log.Println(filter)
	log.Println(productCollection.DeleteOne(r.Context(), filter))
	_, err = productCollection.DeleteOne(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"success": true,
		"message": "Product deleted successfully",
	}
	json.NewEncoder(w).Encode(response)
}
