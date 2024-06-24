package main

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"

	//"fyne.io/fyne/v2/layout"
	"breakdown/src/calculator"
	"io"

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

func initItemRow() fyne.CanvasObject {
	var row_height float32 = 15.0
	var row_width float32 = 70.0

	l, _ := url.Parse("https://buyee.jp")
	link := widget.NewHyperlink("hi", l)
	lbox := container.NewGridWrap(fyne.NewSize(200, row_height), link)

	person := widget.NewLabel("p")
	pbox := container.NewGridWrap(fyne.NewSize(150, row_height), person)

	s_item := fmt.Sprintf("|%d|", 0)
	yen := widget.NewLabel(s_item)
	//yen.Resize(fyne.NewSize(row_width, row_height))
	ybox := container.NewGridWrap(fyne.NewSize(row_width, row_height), yen)

	ship := widget.NewLabel(fmt.Sprintf("|%d|", 0))
	//ship.Resize(fyne.NewSize(row_width, row_height))
	//sbox := container.NewGridWrap(fyne.NewSize(row_width, row_height), ship)

	row := container.NewHBox(pbox, lbox, ybox, ship)

	return row
}

func initShipRow() fyne.CanvasObject {
	var row_height float32 = 20.0
	var row_width float32 = 70.0

	name := widget.NewLabel("nametext")
	nbox := container.NewGridWrap(fyne.NewSize(100, row_height), name)

	es := widget.NewLabel("text")
	ebox := container.NewGridWrap(fyne.NewSize(row_width, row_height), es)

	a := widget.NewLabel("text")
	abox := container.NewGridWrap(fyne.NewSize(row_width, row_height), a)

	p := widget.NewLabel("text")

	return container.NewHBox(nbox, ebox, abox, p)

}

func setShipRow(c calculator.Calculator, lii widget.ListItemID, con *fyne.Container) {
	//fmt.Println("hi")
	//fmt.Println(c.People[lii].Name)
	doubleContainerSetText(c.People[lii].Name, con.Objects[0].(*fyne.Container))
	doubleContainerSetText(fmt.Sprintf("%d", c.People[lii].Ship_b4), con.Objects[1].(*fyne.Container))
	doubleContainerSetText(fmt.Sprintf("%d", c.People[lii].Ship_total), con.Objects[2].(*fyne.Container))
	con.Objects[3].(*widget.Label).SetText(fmt.Sprintf("%.2f", c.People[lii].Proportion))
	//doubleContainerSetText(fmt.Sprintf("%.2f", c.People[lii].Proportion), con.Objects[3].(*fyne.Container))

}

func itemEdit(c calculator.Calculator, i widget.ListItemID, w fyne.Window) {
	desc := widget.NewEntry()
	desc.SetText(c.Items[i].Desc)

	link := widget.NewEntry()
	link.SetText(c.Items[i].Link)

	person := widget.NewSelect(
		c.GetPeople(),
		func(s string) {

		},
	)

	yen := widget.NewEntry()
	yen.SetText(fmt.Sprintf("%d", c.Items[i].Yen))

	ship := widget.NewEntry()
	ship.SetText(fmt.Sprintf("%d", c.Items[i].Shipping))

	//person.PlaceHolder = c.Items[i].Person.Name
	person.Selected = c.Items[i].Person.Name

	dialog.ShowForm("Edit Item", "Save", "Cancel",
		[]*widget.FormItem{
			widget.NewFormItem("Desc", desc),
			widget.NewFormItem("Link", link),
			widget.NewFormItem("Person", person),
			widget.NewFormItem("Price", yen),
			widget.NewFormItem("Shipping", ship),
		},
		func(b bool) {
			if b {
				//validation time....

				if person.Selected == "" || person.Selected == person.PlaceHolder {
					dialog.ShowError(
						errors.New("person must be selected"),
						w,
					)
					return
				}

				if len(desc.Text) == 0 || len(desc.Text) > 40 {
					dialog.ShowError(
						errors.New("desc invalid"),
						w,
					)
					return
				}

				_, err := url.Parse(link.Text)

				if err != nil {
					dialog.ShowError(err, w)
					return
				}

				n, e := strconv.Atoi(yen.Text)

				if e != nil {
					dialog.ShowError(e, nil)
					return
				} else if n < 0 {
					dialog.ShowError(errors.New("cannot be negative"), w)
					return
				}

				s, e1 := strconv.Atoi(ship.Text)

				if e1 != nil {
					dialog.ShowError(e, nil)
					return
				} else if s < 0 {
					dialog.ShowError(errors.New("cannot be negative"), w)
					return
				}

				c.Items[i].Desc = desc.Text
				c.Items[i].Person = c.People[c.GetPerson(person.Selected)]
				c.Items[i].Link = link.Text
				c.Items[i].Yen = n
				c.Items[i].Shipping = s

			}
		},
		w,
	)

}

func setRow(i calculator.Item, con *fyne.Container) {
	doubleContainerSetText(i.Person.Name, con.Objects[0].(*fyne.Container))
	con.Objects[1].(*fyne.Container).Objects[0].(*widget.Hyperlink).URL, _ = url.Parse(i.Link)
	con.Objects[1].(*fyne.Container).Objects[0].(*widget.Hyperlink).Text = i.Desc
	doubleContainerSetText(fmt.Sprintf("|%d|", i.Yen), con.Objects[2].(*fyne.Container))
	con.Objects[3].(*widget.Label).SetText(fmt.Sprintf("|%d|", i.Shipping))
	//doubleContainerSetText(fmt.Sprintf("|%d|", i.Shipping), con.Objects[3].(*fyne.Container))

}

func doubleContainerSetText(text string, o *fyne.Container) {
	o.Objects[0].(*widget.Label).SetText(text)
}

func main() {
	c := new(calculator.Calculator)

	a := app.New()

	w := a.NewWindow("Breakdown Generator")

	breakdown_table := widget.NewList(
		func() int { return len(c.Items) },
		func() fyne.CanvasObject { //create
			return initItemRow()
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) { //update
			hbox := co.(*fyne.Container)

			setRow(c.Items[lii], hbox)

		},
	)
	breakdown_table.OnSelected = func(id widget.ListItemID) {
		itemEdit(*c, id, w)
		c.Sum_shipping()
		breakdown_table.UnselectAll()
		//breakdown_table.Refresh()
		breakdown_table.RefreshItem(id)
	}

	var row_height float32 = 20.0
	var row_width float32 = 70.0

	per := widget.NewLabel("Person")
	pbox := container.NewGridWrap(fyne.NewSize(150, row_height), per)

	link := widget.NewLabel("Item")
	lbox := container.NewGridWrap(fyne.NewSize(200, row_height), link)

	yen := widget.NewLabel("Yen+Fees")
	ybox := container.NewGridWrap(fyne.NewSize(row_width, row_height), yen)

	ship := widget.NewLabel("Indiv. Ship")
	sbox := container.NewGridWrap(fyne.NewSize(row_width, row_height), ship)

	bdown_header := container.NewHBox(
		pbox,
		lbox,
		ybox,
		sbox,
	)
	add_item_button := widget.NewButtonWithIcon("Add", theme.ContentAddIcon(),
		func() {
			desc := widget.NewEntry()

			link := widget.NewEntry()

			person := widget.NewSelect(
				c.GetPeople(),
				func(s string) {

				},
			)

			yen := widget.NewEntry()

			ship := widget.NewEntry()

			person.PlaceHolder = "Select person"

			form := widget.NewForm(
				widget.NewFormItem("Description", desc),
				widget.NewFormItem("Link", link),
				widget.NewFormItem("Person", person),
				widget.NewFormItem("Price", yen),
				widget.NewFormItem("Shipping", ship),
			)

			d := dialog.NewForm(
				"Add Item",
				"Add",
				"Cancel",
				form.Items,
				func(b bool) {
					if b {
						//validation time....

						if person.Selected == "" || person.Selected == person.PlaceHolder {
							dialog.ShowError(
								errors.New("person must be selected"),
								w,
							)
							return
						}

						if len(desc.Text) == 0 || len(desc.Text) > 40 {
							dialog.ShowError(
								errors.New("desc invalid"),
								w,
							)
							return
						}

						_, err := url.Parse(link.Text)

						if err != nil {
							dialog.ShowError(err, w)
							return
						}
						if yen.Text == "" || ship.Text == "" {
							dialog.ShowError(errors.New("must fill all fields"), w)
							return
						}
						n, e := strconv.Atoi(yen.Text)

						if e != nil {
							dialog.ShowError(e, nil)
							return
						} else if n < 0 {
							dialog.ShowError(errors.New("cannot be negative"), w)
							return
						}

						s, e1 := strconv.Atoi(ship.Text)

						if e1 != nil {
							dialog.ShowError(e, nil)
							return
						} else if s < 0 {
							dialog.ShowError(errors.New("cannot be negative"), w)
							return
						}
						p_pointer := c.People[c.GetPerson(person.Selected)]
						c.AddItem(link.Text, desc.Text,
							p_pointer, n, s)
						breakdown_table.Refresh()
					}
				},
				w,
			)
			d.Show()
		},
	)
	rm_item_button := widget.NewButtonWithIcon("Remove", theme.ContentRemoveIcon(),
		func() {
			items := widget.NewSelect(
				calculator.Map(c.Items, func(i calculator.Item) string {
					return i.Desc
				}),
				func(s string) {},
			)

			items.PlaceHolder = "Select an item"

			form := widget.NewForm(
				widget.NewFormItem("Remove item", items),
			)

			d := dialog.NewForm(
				"Remove Item",
				"Remove",
				"Cancel",
				form.Items,
				func(b bool) {
					if b && items.Selected != "" && items.Selected != items.PlaceHolder {
						c.RemoveItem(items.Selected)
						breakdown_table.Refresh()
					}
				},
				w,
			)
			d.Show()
		},
	)
	bd_buttons := container.NewHBox(add_item_button, rm_item_button)

	bdown_header_box := container.NewVBox(
		bd_buttons,
		bdown_header,
	)
	//breakdown_table.Resize()
	breakdown_view := container.NewVBox(
		bdown_header_box, //top

		container.NewGridWrap(fyne.NewSize(600, 400), breakdown_table),
	)

	estimated := widget.NewLabel("Estimated shipping: 0")
	actual := widget.NewLabel("Actual shipping:")
	actual_entry := widget.NewEntry()
	if c.Batched > 0 {
		actual_entry.SetText(fmt.Sprintf("%d", c.Batched))
	}
	if c.Total_shipping > 0 {
		estimated.SetText(fmt.Sprintf("Estimated shipping: %d", c.Total_shipping))
	}
	ship_header := container.NewHBox(
		container.NewGridWrap(fyne.NewSize(70, 20), widget.NewLabel("Person")),
		container.NewGridWrap(fyne.NewSize(100, 20), widget.NewLabel("Est. Shipping")),
		container.NewGridWrap(fyne.NewSize(50, 20), widget.NewLabel("Actual")),
		container.NewGridWrap(fyne.NewSize(50, 20), widget.NewLabel("Proportion")),
	)

	ship_list := widget.NewList(
		func() int {
			return len(c.People)
		},
		func() fyne.CanvasObject {
			return initShipRow()
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			setShipRow(*c, lii, co.(*fyne.Container))
		},
	)

	ship_details := container.NewVBox(
		estimated,
		container.NewHBox(actual, actual_entry),
		ship_header,
		container.NewGridWrap(fyne.NewSize(600, 400), ship_list),
	)

	calc_shipping := widget.NewButtonWithIcon("Calculate", theme.ViewRefreshIcon(),
		func() {
			//fmt.Println("yeeet")
			n, e := strconv.Atoi(actual_entry.Text)
			if e == nil {
				c.Batched = n
			}

			c.Break_shipping_down()
			ship_details.Refresh()
		},
	)
	shipping_header := container.NewHBox(calc_shipping)

	shipping_view := container.NewBorder(
		shipping_header,
		nil,
		nil,
		nil,
		ship_details,
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
		fyne.NewSize(200, 30),
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

	rm_person_button := widget.NewButtonWithIcon("Remove", theme.ContentRemoveIcon(),
		func() {
			ne := widget.NewEntry()

			form := widget.NewForm(
				widget.NewFormItem("Remove person", ne),
			)

			d := dialog.NewForm(
				"Remove Person",
				"Remove",
				"Cancel",
				form.Items,
				func(b bool) {
					if len(ne.Text) > 3 && len(ne.Text) < 16 && c.IsPerson(ne.Text) {
						dialog.ShowConfirm("Confirm Removal",
							"Are you sure? Removing this person removes all their items, if any.",
							func(confirmed bool) {
								if confirmed {
									c.Remove_person(ne.Text)
									people_list.Refresh()
								}
							}, w)

					}
				},
				w,
			)
			d.Show()
		},
	)

	button_header := container.NewHBox(
		add_person_button,
		rm_person_button,
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

	save_dialog := dialog.NewFileSave(
		func(uc fyne.URIWriteCloser, err error) {
			dialog.ShowInformation("Notice", "Save the file with a \".json\" extension for ease of use.", w)

			if err != nil {
				dialog.ShowError(err, w)

			}
			if uc == nil {
				return //cancel save
			}
			data, e := c.SaveAsString()
			if e != nil {
				dialog.ShowError(err, w)
				return
			}

			if data == nil {
				dialog.ShowError(errors.New("unable to create JSON"), w)
				return
			}

			_, err = uc.Write(data)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			dialog.ShowInformation("Success", "File saved successfully.", w)
			uc.Close()
		},
		w,
	)
	read_dialog := dialog.NewFileOpen(
		func(uc fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if uc == nil {
				// User canceled the dialog
				return
			}

			// Read the content from the selected file
			data, err := io.ReadAll(uc)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			*c, err = calculator.LoadFromFile(data)
		},
		w,
	)
	sview := container.NewVBox(
		widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(),
			func() {
				save_dialog.Show()
			}),
	)

	oview := container.NewVBox(
		widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(),
			func() {
				read_dialog.Show()
			}),
	)

	s_entry := widget.NewMultiLineEntry()
	scroller := container.NewGridWrap(fyne.NewSize(600, 400), container.NewScroll(s_entry))
	entry_button := widget.NewButtonWithIcon("Spreadsheet Export", theme.FileTextIcon(),
		func() {
			s_entry.SetText(c.ShowAsTablestring())
			dialog.ShowCustom("Spreadsheet Export", "Ok", scroller, w)
		},
	)

	save_view := container.NewVBox(
		sview,
		oview,
		entry_button,
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
