package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	//"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func UNUSED(x ...interface{}) {}

func main() {
	a := app.New()

	w := a.NewWindow("Breakdown Generator")

	hello := widget.NewLabel("Batch Order Breakdown")

	hello_row := container.NewHBox(hello)
	breakdown_table := widget.NewList(
		func() int { return 1 },
		func() fyne.CanvasObject { //create
			return widget.NewLabel("hi")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) { //update
			co.(*widget.Label).SetText("hi")
		},
	)
	//call append()
	//then call refresh

	breakdown_grid := container.NewAdaptiveGrid(
		4,
		hello_row,
		hello_row,
		hello_row,
		hello_row,
		hello_row,
	)
	breakdown_view := container.NewVScroll(breakdown_table)
	s := new(fyne.Size)
	s.Height = 500
	s.Width = 500
	breakdown_view.SetMinSize(*s)

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

	UNUSED(breakdown_grid)

	w.SetContent(container.NewVBox(
		tabs,
	))

	w.ShowAndRun()
}
