package handler

import (
	"cfshop/backend/db"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	UserName    string             `bson:"username"`
	Email       string             `bson:"email"`
	PhoneNumber string             `bson:"phonenumber"`
	Address     string             `bson:"address"`
	Password    string             `bson:"password"`
}

func CreateUser(user User) (*primitive.ObjectID, error) {
	collection := db.GetUserCollection()

	// Hash the password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		log.Println("Failed to hash password:", err)
		return nil, err
	}

	// Set the hashed password
	user.Password = hashedPassword

	// Create the user entry
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("Failed to create user:", err)
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID)
	user.ID = id

	return &id, nil
}

func GetUserByID(id string) (*User, error) {
	collection := db.GetUserCollection()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid object ID:", err)
		return nil, err
	}

	filter := bson.M{"_id": objID}
	var user User
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Println("Failed to find user:", err)
		return nil, err
	}

	return &user, nil
}

func UpdateUser(user User) error {
	collection := db.GetUserCollection()

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{
		"name":        user.Name,
		"email":       user.Email,
		"phonenumber": user.PhoneNumber,
		"address":     user.Address,
	}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Failed to update user:", err)
		return err
	}

	return nil
}

func DeleteUser(id string) error {
	collection := db.GetUserCollection()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid object ID:", err)
		return err
	}

	filter := bson.M{"_id": objID}

	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("Failed to delete user:", err)
		return err
	}

	return nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
	// return password, nil
}

func AuthenticateUser(username, password string) error {
	collection := db.GetUserCollection()

	// Find the user with the provided username
	filter := bson.M{"username": username}
	var user User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Println("User not found:", err)
		return errors.New("Invalid username or password")
	}
	// Verify the password
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	// if err != nil {
	// 	log.Println("Invalid password", err)
	// 	return errors.New("Invalid username or password")
	// }

	if user.Password != password {
		log.Println("invalid pass")
		return errors.New("Invalid")
	}

	log.Println("User authenticated successfully")
	return nil
}
