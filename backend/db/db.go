package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to DB: %v", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatalf("DB connection test failed: %v", err)
    }

    fmt.Println("Connected to the database!")
    return db
}