package calculator

type Person struct {
	Name       string  //name
	Proportion float32 //percentage
	Iou        float32
	Ship_b4    float32
	Ship_total float32
	Item_total float32
}

type Item struct {
	Link     string
	Person   Person
	Yen      float32
	Shipping float32
}

type Calculator struct {
	People         []Person
	Items          []Item
	total_shipping float32
	batched        float32
}

func RemoveIndexPerson(s []Person, index int) []Person {
	ret := make([]Person, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func RemoveIndexItem(s []Item, index int) []Item {
	ret := make([]Item, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func (c *Calculator) Sum_shipping() {

	for _, e := range c.Items {
		c.total_shipping += e.Shipping
		e.Person.Ship_b4 += e.Shipping
		e.Person.Item_total += e.Yen
	}
}

func (c *Calculator) Break_shipping_down() {
	c.Sum_shipping()

	if c.total_shipping == 0 || c.batched == 0 {
		return
	}

	for _, e := range c.People {
		e.Proportion = e.Ship_b4 / c.total_shipping
		e.Ship_total = e.Proportion * c.batched
		e.Iou = e.Item_total + e.Ship_total
	}

}

func (c *Calculator) Add_person(name string) {
	p := new(Person)
	p.Name = name
	c.People = append(c.People, *p)
}

func (c *Calculator) Remove_person(name string) bool {
	exists := false
	for i, e := range c.People {
		if e.Name == name {
			exists = true
			RemoveIndexPerson(c.People, i)
			break
		}
	}
	return exists
}
