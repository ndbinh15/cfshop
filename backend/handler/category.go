package handler

import (
	"context"

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

// Find a category by ID
func FindCategoryByID(client *mongo.Client, categoryID primitive.ObjectID) (*Category, error) {
	collection := client.Database("coffee_shop").Collection("categories")
	var category Category
	err := collection.FindOne(context.TODO(), bson.M{"_id": categoryID}).Decode(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
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
