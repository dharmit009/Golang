package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)

}

func main() {
	app := app.New()
	window := app.NewWindow("Clock")

	clock := widget.NewLabel("")
	updateTime(clock)

	// anonymous go routine ...
	// To update the time paralley in the background
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	window.SetContent(clock)

	window.ShowAndRun()

}
