package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func fileMenu() *fyne.Menu {
	item1 := fyne.NewMenuItem("A", func() {})
	item2 := fyne.NewMenuItem("B", func() {})
	return fyne.NewMenu("File", item1, item2)
}

func helpMenu(a fyne.App) *fyne.Menu {
	item1 := fyne.NewMenuItem("C", func() {})
	item2 := fyne.NewMenuItem("D", func() { showWindow(a) })
	return fyne.NewMenu("Help", item1, item2)
}

func showWindow(a fyne.App) {
	w := a.NewWindow("quick test")
	w.SetContent(widget.NewLabel("quick test"))
	w.Resize(fyne.NewSize(200, 200))
	w.Show()
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	content := container.NewVBox(
		widget.NewLabel("The top row of the VBox"),
		container.NewHBox(
			widget.NewLabel("Label 1"),
			widget.NewLabel("Label 2"),
		),
	)

	content.Add(widget.NewButton("Add more items", func() {
		content.Add(widget.NewLabel("Added"))
	}))

	mainMenu := fyne.NewMainMenu(fileMenu(), helpMenu(myApp))
	myWindow.SetMainMenu(mainMenu)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
