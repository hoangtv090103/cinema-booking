package dbpg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func Connect() (*sql.DB, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	// Ping the database to verify the connection immediately.
	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("unable to verify database connection: %v", err)
	}

	// Optional: Set connection pool settings
	conn.SetMaxOpenConns(25)   // Max open connections
	conn.SetMaxIdleConns(25)   // Max idle connections
	conn.SetConnMaxLifetime(0) // No connection lifetime limit

	return conn, nil
}
