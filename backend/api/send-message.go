package api

import (
	"encoding/json"
	"net/http"
	"fmt"
	"backend/db"
	"backend/models"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {

	db.Init()
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

	_, err := db.DB.Exec("INSERT INTO messages (username, content) VALUES ($1, $2)", msg.Username, msg.Content)
	if err != nil {
		fmt.Println("failed insert",err)

		return
	}
	

	w.WriteHeader(http.StatusOK)
}
