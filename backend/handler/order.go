package handler

import (
	"cfshop/backend/db"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Order struct {
	ID        primitive.ObjectID `json:"_id,omitempty"`
	UserID    string             `json:"userid"`
	CartID    string             `json:"cartid"`
	Items     []OrderItem        `json:"items"`
	TotalBill string             `json:"totalBill"`
	Status    string             `json:"status"`
}

type OrderItem struct {
	ProductID   primitive.ObjectID `json:"productId"`
	ProductName string             `json:"productName"`
	Quantity    int                `json:"quantity"`
	Price       string             `json:"price"`
	Image       string             `json:"image"`
}

type OrderData struct {
	Product []struct {
		Product  ProductData `json:"product"`
		Quantity int         `json:"quantity"`
		Price    string      `json:"price"`
	} `json:"product"`
	UserID     string `json:"userid"`
	CartID     string `json:"cartid"`
	TotalPrice string `json:"totalPrice"`
	TotalBill  string `json:"totalBill"`
}

type ProductData struct {
	ID          primitive.ObjectID `json:"_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Quantity    int                `json:"quantity"`
	Image       string             `json:"image"`
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method CreateOrder not allowed", http.StatusMethodNotAllowed)
		return
	}

	var orderData OrderData

	err := json.NewDecoder(r.Body).Decode(&orderData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := Order{
		ID:        primitive.NewObjectID(),
		UserID:    orderData.UserID,
		CartID:    orderData.CartID,
		TotalBill: orderData.TotalBill,
		Status:    "In-Progress",
		Items:     make([]OrderItem, 0),
	}

	for _, prod := range orderData.Product {
		order.Items = append(order.Items, OrderItem{
			ProductID:   prod.Product.ID,
			ProductName: prod.Product.Name,
			Quantity:    prod.Quantity,
			Price:       prod.Price,
			Image:       prod.Product.Image,
		})
	}

	orderCollection := db.GetOrderCollection()
	_, err = orderCollection.InsertOne(r.Context(), order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = DeleteCartByCartID(order.CartID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"success": true,
		"message": "Order created successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func GetOrdersByUserID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method GetOrdersByUserID not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Missing userID parameter", http.StatusBadRequest)
		return
	}
	filter := bson.M{"userid": userID}

	orderCollection := db.GetOrderCollection()

	findOptions := options.Find()
	findOptions.SetSort(bson.M{"_id": -1})

	cur, err := orderCollection.Find(r.Context(), filter, findOptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(r.Context())

	var orders []Order
	for cur.Next(r.Context()) {
		var order Order
		err := cur.Decode(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)
	}

	if err := cur.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func GetOrderCount(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method GetOrdersByUserID not allowed", http.StatusMethodNotAllowed)
		return
	}
	orderCollection := db.GetOrderCollection()

	count, err := orderCollection.CountDocuments(r.Context(), bson.M{})
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

func GetInProgressOrderCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method GetInProgressOrderCount not allowed", http.StatusMethodNotAllowed)
		return
	}

	orderCollection := db.GetOrderCollection()

	filter := bson.M{"status": "In-Progress"}

	count, err := orderCollection.CountDocuments(r.Context(), filter)
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

func GetCompletedOrderCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method GetCompletedOrderCount not allowed", http.StatusMethodNotAllowed)
		return
	}

	orderCollection := db.GetOrderCollection()

	filter := bson.M{"status": "Completed"}

	count, err := orderCollection.CountDocuments(r.Context(), filter)
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

func GetAllInProgressOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method GetAllInProgressOrders not allowed", http.StatusMethodNotAllowed)
		return
	}

	filter := bson.M{"status": "In-Progress"}

	orderCollection := db.GetOrderCollection()

	findOptions := options.Find()
	findOptions.SetSort(bson.M{"_id": -1})

	cur, err := orderCollection.Find(r.Context(), filter, findOptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(r.Context())

	var orders []Order
	for cur.Next(r.Context()) {
		var order Order
		err := cur.Decode(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)
	}

	if err := cur.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func GetAllCompletedOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method GetAllInProgressOrders not allowed", http.StatusMethodNotAllowed)
		return
	}

	filter := bson.M{"status": "Completed"}

	orderCollection := db.GetOrderCollection()

	findOptions := options.Find()
	findOptions.SetSort(bson.M{"_id": -1})

	cur, err := orderCollection.Find(r.Context(), filter, findOptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(r.Context())

	var orders []Order
	for cur.Next(r.Context()) {
		var order Order
		err := cur.Decode(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)
	}

	if err := cur.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func MarkOrderAsDoneByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method MarkOrderAsDoneByID not allowed", http.StatusMethodNotAllowed)
		return
	}

	orderID := r.URL.Query().Get("id")
	if orderID == "" {
		http.Error(w, "Missing orderID parameter", http.StatusBadRequest)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	orderCollection := db.GetOrderCollection()

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"status": "Completed"}}

	_, err = orderCollection.UpdateOne(r.Context(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"success": true,
		"message": "Order marked as completed successfully",
	}
	json.NewEncoder(w).Encode(response)
}
