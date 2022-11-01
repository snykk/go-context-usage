package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

func queryWithTimeout(ctx context.Context, db *sql.DB) {
	queryCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	rows, err := db.QueryContext(queryCtx, "SELECT name FROM users")
	if err != nil {
		fmt.Println("Query failed:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Println("Failed to scan row:", err)
			return
		}
		fmt.Println("User:", name)
	}
}

func main() {
	// Setup database
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table and insert sample data
	db.Exec("CREATE TABLE users (name TEXT)")
	db.Exec("INSERT INTO users (name) VALUES ('Alice'), ('Bob'), ('Charlie')")

	ctx := context.Background()
	queryWithTimeout(ctx, db)
}
