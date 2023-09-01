package handler

import (
	"cfshop/backend/db"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	UserName    string             `bson:"username"`
	Role        string             `bson:"role"`
	Email       string             `bson:"email"`
	PhoneNumber string             `bson:"phonenumber"`
	Address     string             `bson:"address"`
	Password    string             `bson:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method CreateUser not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		log.Println("$POST CreateUser success")
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userCollection := db.GetUserCollection()
	// Check if the username is unique
	filter := bson.M{"username": user.UserName}
	count, err := userCollection.CountDocuments(context.TODO(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if count > 0 {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	userID := primitive.NewObjectID()
	user.ID = userID
	user.Role = "user"

	_, err = userCollection.InsertOne(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"success": true,
		"message": "User Registered successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method GetAllUsers not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		log.Println("$GET GetAllUsers success")
	}

	userCollection := db.GetUserCollection()

	cursor, err := userCollection.Find(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(r.Context())

	// Collect users in a slice
	var users []User
	for cursor.Next(r.Context()) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUserRole(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method GetUserRole not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		log.Println("$GET GetUserRole success")
	}

	userId := r.URL.Query().Get("id")

	userCollection := db.GetUserCollection()
	user := User{}
	err := userCollection.FindOne(r.Context(), bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	role := user.Role

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUserCount(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method GetOrdersByUserID not allowed", http.StatusMethodNotAllowed)
		return
	}

	userCollection := db.GetUserCollection()

	count, err := userCollection.CountDocuments(r.Context(), bson.M{})
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

func GetUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method GetUserById not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		log.Println("$GET GetUserById success")
	}

	id := r.URL.Query().Get("id")
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userCollection := db.GetUserCollection()

	filter := bson.M{"_id": userID}
	log.Println(filter)
	var user User
	err = userCollection.FindOne(r.Context(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Method UpdateUserByID not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		log.Println("$PUT UpdateUserByID success")
	}

	var updatedUser User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userCollection := db.GetUserCollection()

	userID, err := primitive.ObjectIDFromHex(updatedUser.ID.Hex())
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": userID}
	existingUser := userCollection.FindOne(r.Context(), filter)

	if existingUser.Err() == nil {
		update := bson.M{"$set": updatedUser}
		_, err := userCollection.UpdateOne(r.Context(), filter, update)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		log.Println("User not found")
		http.Error(w, "User not found", http.StatusNotFound)
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

// func hashPassword(password string) (string, error) {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(hashedPassword), nil
// 	// return password, nil
// }

func AuthenticateUser(username, password string) (string, string, error) {
	collection := db.GetUserCollection()

	// Find the user with the provided username
	filter := bson.M{"username": username}
	var user User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Println("User not found:", err)
		return "", "", errors.New("Invalid username or password")
	}
	// Verify the password
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	// if err != nil {
	// 	log.Println("Invalid password", err)
	// 	return errors.New("Invalid username or password")
	// }

	if user.Password != password {
		log.Println("invalid pass")
		return "", "", errors.New("Invalid")
	}

	log.Println("User authenticated successfully")
	return user.ID.Hex(), user.Role, nil
}
