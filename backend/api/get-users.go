package api

import (
	"encoding/json"
	"net/http"
	"fmt"
	"backend/db"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	db.Init()

	rows, err := db.DB.Query("SELECT username FROM users")
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		fmt.Println("failed to retrieve users")
		return
	}
	defer rows.Close()

	var users []string
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			http.Error(w, "Failed to scan user", http.StatusInternalServerError)
			fmt.Println("failed to scan user")

			return
		}
		users = append(users, username)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
