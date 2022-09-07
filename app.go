package main

import (
	"context"
	"database/sql"
	"quote-of-the-day/database"
)

// App struct
type App struct {
	ctx context.Context
	db  *sql.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.db = database.Connect()
}

// Greet returns a greeting for the given name
func (a *App) DisplayQuote() database.Quote {
	quotes, err := database.GetQuotes(a.db)

	if err == nil {
		return quotes[0]
	}

	return database.Quote{QuoteID: -1, Author: "Error", Quote: err.Error(), Last_Used: "Error"}
}
