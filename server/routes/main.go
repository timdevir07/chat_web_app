package main

import (
	"log"
	"net/http"

	"chat-web-app/config"
	"chat-web-app/routes"
)

func main() {
	// Step 1: Connect to MongoDB
	config.ConnectDB()

	// Step 2: Routes
	// Auth Routes
	http.HandleFunc("/api/register", routes.RegisterHandler)
	http.HandleFunc("/api/login", routes.LoginHandler)

	// Chat Routes
	http.HandleFunc("/api/send", routes.SendMessageHandler)
	http.HandleFunc("/api/messages", routes.GetMessagesHandler)

	// Step 3: Start server
	log.Println("✅ Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
