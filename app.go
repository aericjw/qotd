package main

import (
	"context"
	"fmt"

	"quote-of-the-day/database"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetQuote() database.Quote {
	db := database.Connect()
	quotes, err := database.GetQuotes(db)
	if err == nil {
		return quotes[0]
	}
	return database.Quote{QuoteID: 1, AuthorID: 1, Quote: err.Error(), Last_Used: "Error"}
}