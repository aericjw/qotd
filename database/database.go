package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Quote struct {
	QuoteID   int64
	AuthorID  int64
	Quote     string
	Last_Used string
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "db",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	quotes, err := getQuotes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Quotes: %v\n", quotes)
}

func getQuotes() ([]Quote, error) {
	var quotes []Quote

	rows, err := db.Query("SELECT * FROM quotes")
	if err != nil {
		return nil, fmt.Errorf("getQuotes: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var quote Quote
		if err := rows.Scan(&quote.QuoteID, &quote.AuthorID, &quote.Quote, &quote.Last_Used); err != nil {
			return nil, fmt.Errorf("getQuotes: %v", err)
		}
		quotes = append(quotes, quote)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getQuotes: %v", err)
	}

	return quotes, nil
}
