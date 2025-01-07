package api

import (
	"encoding/json"
	"net/http"
	"fmt"
	"backend/db"
	"backend/models"
)

func InsertUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db.Init()
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		fmt.Println("failed request")
		return
	}

	var msg models.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Failed to parse message", http.StatusBadRequest)
		fmt.Println("failed parse")

		return
	}

	_, err := db.DB.Exec(`INSERT INTO users (username) VALUES ('`+ msg.Username+ `') ON CONFLICT DO NOTHING`)
	if err != nil {
		
		fmt.Println("Failed to insert user:", err)


		return
	}

	w.WriteHeader(http.StatusOK)
}
