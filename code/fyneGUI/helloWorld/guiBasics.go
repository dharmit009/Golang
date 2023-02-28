package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {

	app := app.New()
	window := app.NewWindow("# GUI Basics #")
	label1 := widget.NewLabel("Hello World")
	window.SetContent(label1)
	window.ShowAndRun()
}
