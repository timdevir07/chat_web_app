// routes/login.go
package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"chat-web-app/config"
	"chat-web-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var creds models.User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// MongoDB: Find user by email
	var user models.User
	err := config.UserCollection.FindOne(context.TODO(), bson.M{"email": creds.Email}).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "User not found"})
		return
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "Incorrect password"})
		return
	}

	// Success
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
		"name":    user.Name,
	})
}
