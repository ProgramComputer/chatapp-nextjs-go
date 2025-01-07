package api

import (
	"encoding/json"
	"net/http"
		"fmt"
	"backend/db"
	"backend/models"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {


	w.Header().Set("Access-Control-Allow-Origin", "*")

	db.Init()//for serverless
		
	rows, err := db.DB.Query("SELECT username, content, timestamp FROM messages ORDER BY timestamp DESC LIMIT 50")
	if err != nil {
		http.Error(w, "Failed to retrieve messages", http.StatusInternalServerError)
		fmt.Println("failed to retrieve messages");
		return
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.Username, &msg.Content, &msg.Timestamp); err != nil {
			http.Error(w, "Failed to scan message", http.StatusInternalServerError)
			fmt.Println("failed to scan message")

			return
		}
		messages = append(messages, msg)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
