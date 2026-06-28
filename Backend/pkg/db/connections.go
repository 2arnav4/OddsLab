package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

// InitDB initializes the PostgreSQL database connection pool.
func InitDB() error {
	// Retrieve the connection string from environment variables
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	var err error
	// Establish the connection pool
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Verify the connection is active
	if err = db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Configure the connection pool settings
	db.SetMaxOpenConns(25)                 // Maximum number of open connections to the database
	db.SetMaxIdleConns(25)                 // Maximum number of idle connections in the pool
	db.SetConnMaxLifetime(5 * time.Minute) // Maximum amount of time a connection may be reused (in seconds)

	log.Println("Database connection pool initialized successfully")
	return nil

}
func GetDB() *sql.DB {
	return db
}
