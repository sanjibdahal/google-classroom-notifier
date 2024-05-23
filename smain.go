package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello Fyne")

	myWindow.SetContent(container.NewVBox(
		widget.NewLabel("Welcome to Fyne!"),
		container.New(layout.NewCenterLayout(),
			widget.NewButton("Quit", func() {
				myApp.Quit()
			})),
	))

	myWindow.ShowAndRun()
}
