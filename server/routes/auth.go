package routes

import (
	"chat-web-app/config"
	"chat-web-app/models"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	collection := config.DB.Collection("users")

	// Check if user already exists
	var existing models.User
	err := collection.FindOne(r.Context(), bson.M{"email": user.Email}).Decode(&existing)
	if err == nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "User already exists"})
		return
	}

	_, err = collection.InsertOne(r.Context(), user)
	if err != nil {
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Registration successful"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds models.User
	json.NewDecoder(r.Body).Decode(&creds)

	collection := config.DB.Collection("users")

	var existing models.User
	err := collection.FindOne(r.Context(), bson.M{
		"email":    creds.Email,
		"password": creds.Password,
	}).Decode(&existing)

	if err == mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid credentials"})
		return
	} else if err != nil {
		http.Error(w, "Login failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful", "name": existing.Name})
}
