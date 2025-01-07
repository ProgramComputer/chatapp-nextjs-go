package main

import (
	"log"
	"fmt"
	"net/http"
	"backend/db"
	"backend/api"
)
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "ROOT")
}
//Local entry
func main() {
	// Initialize the database connection
	db.Init()

	// Register HTTP handlers
	http.HandleFunc("/", HomeHandler)

	http.HandleFunc("/api/send-message", api.SendMessage)
	http.HandleFunc("/api/get-messages", api.GetMessages)
	http.HandleFunc("/api/get-users", api.GetUsers)
	http.HandleFunc("/api/insert-user", api.InsertUser)
	http.HandleFunc("/api/delete-user", api.DeleteUser)



	// Start the server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
