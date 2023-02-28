package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Handling Multiple Windows")
	w.SetMaster()

	button := widget.NewButton("Open New", func() {
		w2 := a.NewWindow("Second window")
		w2.SetContent(widget.NewLabel("More Content"))
		w2.Resize(fyne.NewSize(500, 500))
		w2.Show() //shows secondary window
	})

	w.SetContent(button)
	w.Padded()

	w.Show() //shows the window
	a.Run()  // runs the application

}
