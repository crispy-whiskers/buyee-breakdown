package calculator

type Person struct {
	string             //name
	proportion float32 //percentage
	iou        float32
	ship_b4    float32
	ship_total float32
	item_total float32
}

type Item struct {
	link     string
	person   Person
	yen      float32
	shipping float32
}

type Calculator struct {
	people         []Person
	items          []Item
	total_shipping float32
	batched        float32
}

func (c *Calculator) Sum_shipping() {

	for _, e := range c.items {
		c.total_shipping += e.shipping
		e.person.ship_b4 += e.shipping
		e.person.item_total += e.yen
	}
}

func (c *Calculator) Break_shipping_down() {
	c.Sum_shipping()

	if c.total_shipping == 0 || c.batched == 0 {
		return
	}

}
