package db

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	_ "github.com/lib/pq"
	//"github.com/joho/godotenv"

)

var DB *sql.DB

func Init() {
	var err error

	// err = godotenv.Load()
    // if err != nil {
    //     log.Fatalf("Error loading .env file")
	// 	fmt.Println("error loading .env file")

    // }
	connStr := os.Getenv("DATABASE_URL") // Ensure DATABASE_URL is set in Vercel environment variables
	//fmt.Println(connStr);

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		fmt.Println("failed to connect to database")

	}

	// Set connection pool settings
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(0)

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
		fmt.Println("failed to ping database")

	}
}
