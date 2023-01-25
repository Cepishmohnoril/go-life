package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	createWindow()
}

func createWindow() {
	myApp := app.New()
	mainWindow := myApp.NewWindow("Canvas")
	lifeCanvas := mainWindow.Canvas()
	mainContainer := container.New(layout.NewGridLayout(10))

	for i := 0; i < 100; i++ {
		circle := canvas.NewCircle(color.White)
		circle.Resize(fyne.NewSize(3, 3))
		mainContainer.Add(circle)
	}

	lifeCanvas.SetContent(mainContainer)

	mainWindow.Resize(fyne.NewSize(300, 300))
	mainWindow.ShowAndRun()
}
