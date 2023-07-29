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

type Category struct {
	ID   primitive.ObjectID `json:"_id,omitempty"`
	Name string             `json:"name"`
}

// Insert a new category
func InsertCategory(client *mongo.Client, category Category) error {
	collection := client.Database("coffee_shop").Collection("categories")
	_, err := collection.InsertOne(context.TODO(), category)
	return err
}

func GetAllCategory(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method GetAllCategory not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		log.Println("$GET all category success")
	}
	// Get the MongoDB collection for products
	categoryCollection := db.GetCategoryCollection()

	// Fetch all products from the collection
	cursor, err := categoryCollection.Find(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(r.Context())

	// Collect products in a slice
	var categories []Category
	for cursor.Next(r.Context()) {
		var category Category
		err := cursor.Decode(&category)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}
	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode categories as JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(categories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Update a category
func UpdateCategory(client *mongo.Client, categoryID primitive.ObjectID, updatedCategory Category) error {
	collection := client.Database("coffee_shop").Collection("categories")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": categoryID}, bson.M{"$set": updatedCategory})
	return err
}

// Delete a category
func DeleteCategory(client *mongo.Client, categoryID primitive.ObjectID) error {
	collection := client.Database("coffee_shop").Collection("categories")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": categoryID})
	return err
}
