package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	DB, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
			log.Fatalf("failed to connect: %v", err)
	}
	defer DB.Close()

	if err := DB.Ping(); err != nil {
			log.Fatalf("failed to ping: %v", err)
	}

	log.Println("Successfully connected to PlanetScale!")
}
