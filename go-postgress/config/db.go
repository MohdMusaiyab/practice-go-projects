package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Build DSN from env variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)

	// Open connection
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Error opening DB: %v\n", err)
	}

	// Verify connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to DB: %v\n", err)
	}

	fmt.Println("✅ Connected to PostgreSQL successfully!")
}
