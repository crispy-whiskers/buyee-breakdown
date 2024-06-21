package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"

	//"fyne.io/fyne/v2/layout"
	"breakdown/src/calculator"

	"fyne.io/fyne/v2/widget"
)

func UNUSED(x ...interface{}) {}

func initPersonRow() fyne.CanvasObject {
	var row_height float32 = 20.0
	var row_width float32 = 400.0
	name := widget.NewLabel("nametext")
	nbox := container.NewGridWrap(fyne.NewSize(200, 20), name)

	s_item := fmt.Sprintf("|%d|", 0)
	itemtotal := widget.NewLabel(s_item)
	itemtotal.Resize(fyne.NewSize(row_width, row_height))

	iou := widget.NewLabel(fmt.Sprintf("|%d|", 1))
	b4 := widget.NewLabel(fmt.Sprintf("|%d|", 0))
	batch := widget.NewLabel(fmt.Sprintf("|%d|", 0))

	iou.Resize(fyne.NewSize(row_width, row_height))
	b4.Resize(fyne.NewSize(row_width, row_height))
	batch.Resize(fyne.NewSize(row_width, row_height))

	row := container.NewHBox(nbox, itemtotal, iou, b4, batch)

	return row
}

func doubleContainerSetText(text string, o *fyne.Container) {
	o.Objects[0].(*widget.Label).SetText(text)
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
			return hello
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) { //update

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
			return initPersonRow()
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) { //update
			hbox := co.(*fyne.Container)
			//hbox.Objects[0].(*fyne.Container).Objects[0].(*widget.Label).SetText(c.People[lii].Name)

			doubleContainerSetText(c.People[lii].Name, hbox.Objects[0].(*fyne.Container))

			s_item := fmt.Sprintf("|%d|", c.People[lii].Item_total)
			hbox.Objects[1].(*widget.Label).SetText(s_item)

			hbox.Objects[2].(*widget.Label).SetText(fmt.Sprintf("|%d|", c.People[lii].Iou))
			hbox.Objects[3].(*widget.Label).SetText(fmt.Sprintf("|%d|", c.People[lii].Ship_b4))
			hbox.Objects[4].(*widget.Label).SetText(fmt.Sprintf("|%d|", c.People[lii].Ship_total))

		},
	)

	name_header_space := container.NewGridWrap(
		fyne.NewSize(200, 20),
		widget.NewLabel("Person"),
	)
	people_row_header := container.NewHBox(
		name_header_space,
		widget.NewLabel("Item $"),
		widget.NewLabel("Owed"),
		widget.NewLabel("Single"),
		widget.NewLabel("Split"),
	)

	add_person_button := widget.NewButtonWithIcon("Add", theme.ContentAddIcon(),
		func() {
			ne := widget.NewEntry()

			form := widget.NewForm(
				widget.NewFormItem("New person", ne),
			)

			d := dialog.NewForm(
				"Add Person",
				"Add",
				"Cancel",
				form.Items,
				func(b bool) {
					if len(ne.Text) > 3 && len(ne.Text) < 16 {

						c.Add_person(ne.Text)
						people_list.Refresh()
					}
				},
				w,
			)
			d.Show()
		},
	)

	button_header := container.NewHBox(
		add_person_button,
	)
	people_header := container.NewVBox(
		button_header,
		people_row_header,
	)

	people_view := container.NewBorder(
		people_header,
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
