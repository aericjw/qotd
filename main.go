package main

import (
	"fmt"
	_ "embed"

	"github.com/wailsapp/wails"
)

type MyData struct {
	A string
	B float64
	C int64
}

// Format: { A: "", B: 0.0, C: 0 }
func basic(data map[string]interface{}) string {
	var result MyData
	fmt.Printf("data: %#v\n", data)

	err := mapstructure.Decode(data, &result)
	if err != nil {
		// TODO
	}
	fmt.Printf("result: %#V\n", result)
	return "Hello World!"
}

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "Quote of the Day",
	})
	app.Bind(hello)
	app.Run()
}
