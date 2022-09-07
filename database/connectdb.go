package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Quote struct {
	QuoteID   int64
	Author    string
	Quote     string
	Last_Used string
}

func Connect() *sql.DB {
	var db *sql.DB

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
	return db
}

func GetQuotes(db *sql.DB) ([]Quote, error) {
	var quotes []Quote

	rows, err := db.Query("SELECT * FROM quotes")
	if err != nil {
		return nil, fmt.Errorf("getQuotes: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var quote Quote
		if err := rows.Scan(&quote.QuoteID, &quote.Author, &quote.Quote, &quote.Last_Used); err != nil {
			return nil, fmt.Errorf("getQuotes: %v", err)
		}
		quotes = append(quotes, quote)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getQuotes: %v", err)
	}

	return quotes, nil
}

func AddQuote(db *sql.DB, quote Quote) (int64, error) {
	result, err := db.Exec("INSERT INTO quotes (author, quote, last_used) VALUES (?,?,NOW())", &quote.Author, &quote.Quote)
	if err != nil {
		return 0, fmt.Errorf("addQuote: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addQuote: %v", err)
	}
	return id, nil
}
