package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func UNUSED(x ...interface{}) {}

func main() {
	a := app.New()

	w := a.NewWindow("Breakdown Generator")

	hello := widget.NewLabel("Batch Order Breakdown")
	breakdown_view := container.New(
		layout.NewGridLayout(4), //use new scroll somewhere here
		hello,
	)

	shipping_view := container.NewCenter(
		hello,
	)

	people_view := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		hello,
	)

	save_view := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		hello,
	)

	tabs := container.NewAppTabs(
		container.NewTabItem("Items", breakdown_view),
		container.NewTabItem("Shipping Calc", shipping_view),
		container.NewTabItem("People", people_view),
		container.NewTabItem("File", save_view),
	)

	UNUSED(tabs)

	w.SetContent(container.NewVBox(
		tabs,
	))

	w.ShowAndRun()
}
