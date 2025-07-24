package routes

import (
	"chat-web-app/config"
	"chat-web-app/models"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var msg models.Message
	json.NewDecoder(r.Body).Decode(&msg)

	msg.Timestamp = time.Now()

	collection := config.DB.Collection("messages")
	_, err := collection.InsertOne(r.Context(), msg)
	if err != nil {
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Message sent"})
}

func GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sender := r.URL.Query().Get("sender")
	receiver := r.URL.Query().Get("receiver")

	if sender == "" || receiver == "" {
		http.Error(w, "Missing sender or receiver", http.StatusBadRequest)
		return
	}

	collection := config.DB.Collection("messages")
	filter := bson.M{
		"$or": []bson.M{
			{"sender": sender, "receiver": receiver},
			{"sender": receiver, "receiver": sender},
		},
	}

	cursor, err := collection.Find(r.Context(), filter)
	if err != nil {
		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(r.Context())

	var messages []models.Message
	for cursor.Next(r.Context()) {
		var msg models.Message
		cursor.Decode(&msg)
		messages = append(messages, msg)
	}

	json.NewEncoder(w).Encode(messages)
}
