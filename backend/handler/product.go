package handler

import (
	"cfshop/backend/db"
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

	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the MongoDB collection for products
	productCollection := db.GetProductCollection()

	// Check if the name of the product is unique
	// filter := bson.M{"name": product.Name}
	// log.Println(filter)
	// count, err := productCollection.CountDocuments(context.TODO(), filter)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// if count > 0 {
	// 	http.Error(w, "Product name already exists", http.StatusBadRequest)
	// 	return
	// }

	productID := primitive.NewObjectID()
	product.ID = productID

	if product.Quantity == 0 {
		product.Quantity = 0
	}

	if product.Image == "" {
		product.Image = "default.jpg"
	}

	_, err = productCollection.InsertOne(r.Context(), product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"success": true,
		"message": "Product created successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method GetProductById not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		log.Println("$GET GetProductById success")
	}

	id := r.URL.Query().Get("id")
	productID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	productCollection := db.GetProductCollection()

	filter := bson.M{"id": productID}
	log.Println(filter)
	var product Product
	err = productCollection.FindOne(r.Context(), filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method GetAllProducts not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		log.Println("$GET all product success")
	}
	productCollection := db.GetProductCollection()

	cursor, err := productCollection.Find(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(r.Context())

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
	productCollection := db.GetProductCollection()

	count, err := productCollection.CountDocuments(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Count int64 `json:"count"`
	}{
		Count: count,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateProductQuantity(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		http.Error(w, "Method UpdateProductQuantity not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		log.Println("$PUT UpdateProductQuantity success")
	}

	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println("error decode")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productCollection := db.GetProductCollection()

	filter := bson.M{"name": product.Name}
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

func UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Method UpdateProductByID not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		log.Println("$PUT UpdateProductByID success")
	}

	var updatedProduct Product
	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productCollection := db.GetProductCollection()

	productID, err := primitive.ObjectIDFromHex(updatedProduct.ID.Hex())
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	filter := bson.M{"id": productID}
	existingProduct := productCollection.FindOne(r.Context(), filter)

	if existingProduct.Err() == nil {
		update := bson.M{"$set": updatedProduct}
		_, err := productCollection.UpdateOne(r.Context(), filter, update)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		log.Println("Product not found")
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"success": true,
		"message": "Product updated successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method DeleteProduct not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		log.Println("$DELETE DeleteProduct success")
	}

	idParam := r.URL.Query().Get("id")
	productID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	productCollection := db.GetProductCollection()

	filter := bson.M{"id": productID}
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
