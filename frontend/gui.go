package frontend

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateGUI(quoteDisplay string, authorDisplay string) {
	myApp := app.New()
	w := myApp.NewWindow("Quote of the Day")

	raster := canvas.NewRasterWithPixels(func(x, y, w, h int) color.Color { return color.White })

	quote := canvas.NewText(quoteDisplay, color.Black)
	author := canvas.NewText(authorDisplay, color.Black)
	settingsIcon, _ := fyne.LoadResourceFromPath("frontend/resources/icon.png")
	settings := widget.NewButtonWithIcon("", settingsIcon, func() {
		log.Println("tapped")
	})

	quoteContainer := container.New(layout.NewHBoxLayout(), quote)
	authorContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), author)
	settingsContainer := container.New(layout.NewCenterLayout(), settings)
	settingsGrid := container.New(layout.NewGridLayout(2), layout.NewSpacer(), layout.NewSpacer(), settingsContainer, layout.NewSpacer())
	authorAndSettingsContainer := container.New(layout.NewGridLayout(2), settingsGrid, authorContainer)

	content := container.New(layout.NewGridLayout(1), quoteContainer, authorAndSettingsContainer)
	render := container.New(layout.NewMaxLayout(), raster, content)

	w.SetContent(render)
	w.Resize(fyne.NewSize(600, 350))
	w.ShowAndRun()
}
