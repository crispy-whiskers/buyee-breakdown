package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	//"fyne.io/fyne/v2/layout"
	"breakdown/src/calculator"

	"fyne.io/fyne/v2/widget"
)

func UNUSED(x ...interface{}) {}

func makePersonRow(p calculator.Person) fyne.CanvasObject {
	name := widget.NewLabel(p.Name)

	s_item := fmt.Sprintf("|%.2f|", p.Item_total)
	itemtotal := widget.NewLabel(s_item)

	iou := widget.NewLabel(fmt.Sprintf("|%.2f|", p.Iou))
	b4 := widget.NewLabel(fmt.Sprintf("|%.2f|", p.Ship_b4))
	batch := widget.NewLabel(fmt.Sprintf("|%.2f|", p.Ship_total))

	row := container.NewHBox(name, itemtotal, iou, b4, batch)

	return row
}

func main() {
	c := new(calculator.Calculator)
	c.Add_person("catto")
	c.Add_person("reverie")

	a := app.New()

	w := a.NewWindow("Breakdown Generator")

	hello := widget.NewLabel("Batch Order Breakdown")

	breakdown_table := widget.NewList(
		func() int { return len(c.Items) },
		func() fyne.CanvasObject { //create
			return widget.NewLabel("hi")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) { //update
			co.(*widget.Label).SetText("hi")
		},
	)

	//call append() on table
	//then call refresh

	breakdown_box := container.NewVScroll(breakdown_table)
	s := new(fyne.Size)
	s.Height = 500
	s.Width = 500
	breakdown_box.SetMinSize(*s)

	breakdown_view := container.NewBorder(
		nil, //top
		nil, //bottom
		nil,
		nil,
		breakdown_box,
	)

	shipping_view := container.NewCenter(
		hello,
	)

	people_list := widget.NewList(
		func() int { return len(c.People) },
		func() fyne.CanvasObject { //create
			return widget.NewLabel("hi")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) { //update
			co.(*widget.Label).SetText(c.People[lii].Name)
		},
	)

	people_view := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		people_list,
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

	w.SetContent(container.NewVBox(
		tabs,
	))

	w.ShowAndRun()
}
