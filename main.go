package main

import (
	"embed"
	"log"
	"math/rand"
	"quote-of-the-day/database"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

func getRandomNumber(length int) int {
	// Random number generated between given
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(length)
	return result
}

func convertStringToDate(date string) time.Time {
	layout := "2006-01-02 15:04:05"
	result, err := time.Parse(layout, date)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func getQuoteToDisplay(quotes []database.Quote) database.Quote {
	// This function finds the Quote that has been unused for the longest amount of time
	// If multiple quotes have been unused for the longest amount of time, pick one randomly

	currentTime := time.Now()
	var longestUnusedInterval time.Duration
	var longestUnusedQuotes []database.Quote

	// Interval is the difference between now and the last time the quote was used

	for i := 0; i < len(quotes); i++ {
		interval := currentTime.Sub(convertStringToDate(quotes[i].Last_Used))
		if interval > longestUnusedInterval {
			longestUnusedInterval = interval
			longestUnusedQuotes = []database.Quote{quotes[i]}
		} else if interval == longestUnusedInterval {
			longestUnusedQuotes = append(longestUnusedQuotes, quotes[i])
		}
	}

	numUnusedQuotes := len(longestUnusedQuotes)
	if numUnusedQuotes > 1 {
		randIndex := getRandomNumber(numUnusedQuotes)
		return longestUnusedQuotes[randIndex]
	}

	return longestUnusedQuotes[0]
}

func DisplayQuote() database.Quote {
	db := database.Connect()
	quotes, err := database.GetQuotes(db)
	if err != nil {
		log.Fatal(err)
	}
	return getQuoteToDisplay(quotes)
}

func main() {
	var assets embed.FS

	app := NewApp()

	err := wails.Run(&options.App{
		Title:            "Quote of the Day",
		Width:            800,
		Height:           480,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
