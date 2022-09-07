package main

import (
	"strconv"

	"quote-of-the-day/database"
	"quote-of-the-day/frontend"
)

func main() {
	db := database.Connect()
	quotes, _ := database.GetQuotes(db)

	frontend.CreateGUI(quotes[0].Quote, strconv.Itoa(int(quotes[0].AuthorID)))

}
