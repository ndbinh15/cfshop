package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"cfshop/backend/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CartRequest struct {
	ID     primitive.ObjectID `json:"_id,omitempty"`
	UserID string             `json:"userid"`
	Items  []CartItem         `json:"items"`
}

type CartItem struct {
	ProductID primitive.ObjectID `json:"productId"`
	Quantity  int                `json:"quantity"`
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method AddToCart not allowed", http.StatusMethodNotAllowed)
		return
	}

	var cartReq CartRequest
	err := json.NewDecoder(r.Body).Decode(&cartReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := primitive.ObjectIDFromHex(cartReq.UserID)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	cartCollection := db.GetCartCollection()

	filter := bson.M{"userid": userID.Hex()}
	var existingCart CartRequest
	err = cartCollection.FindOne(r.Context(), filter).Decode(&existingCart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("User's cart not found. Creating a new cart...")
		} else {
			fmt.Println("Error while querying user's cart:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	fmt.Println(existingCart.Items, cartReq.Items)

	if err == mongo.ErrNoDocuments {
		var newCart CartRequest
		newCart.UserID = userID.Hex()
		newCart.Items = cartReq.Items
		newCartID := primitive.NewObjectID()
		newCart.ID = newCartID

		_, err = cartCollection.InsertOne(r.Context(), newCart)
		if err != nil {
			fmt.Println("Error creating a new cart:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("created a new cart")
	} else {
		updatedItems := updateCartItems(existingCart.Items, cartReq.Items)
		fmt.Println("update items: ", updatedItems)
		update := bson.M{"$set": bson.M{"items": updatedItems}}
		_, err = cartCollection.UpdateOne(r.Context(), filter, update, options.Update())

		if err != nil && err != mongo.ErrNoDocuments {
			fmt.Println("Error updating user's cart:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("updated exist cart")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"success": true,
		"message": "Items added to cart successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func updateCartItems(existingItems []CartItem, newItems []CartItem) []CartItem {
	updatedItems := append([]CartItem{}, existingItems...)

	for _, newItem := range newItems {
		found := false

		for i, existingItem := range updatedItems {
			if existingItem.ProductID == newItem.ProductID {
				updatedItems[i].Quantity += newItem.Quantity
				found = true
				break
			}
		}

		if !found {
			updatedItems = append(updatedItems, newItem)
		}
	}

	return updatedItems
}

func GetCartByUserID(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method AddToCart not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	cartCollection := db.GetCartCollection()

	filter := bson.M{"userid": userID.Hex()}
	log.Println(filter)
	var cart CartRequest
	err = cartCollection.FindOne(r.Context(), filter).Decode(&cart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Cart not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetCartByCartID(cartID string) (CartRequest, error) {
	cartCollection := db.GetCartCollection()

	filter := bson.M{"_id": cartID}

	var cart CartRequest

	err := cartCollection.FindOne(context.Background(), filter).Decode(&cart)
	if err != nil {
		return CartRequest{}, err
	}

	return cart, nil
}

func DeleteCartByCartID(cartID string) error {
	cartCollection := db.GetCartCollection()

	objID, err := primitive.ObjectIDFromHex(cartID)
	if err != nil {
		return err
	}

	filter := bson.M{"id": objID}
	_, err = cartCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
