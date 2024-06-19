package calculator

type Person struct {
	string             //name
	proportion float32 //percentage
	iou        float32
}

type Item struct {
	link     []string
	person   []Person
	yen      float32
	shipping float32
}

type Calculator struct {
	people         []Person
	items          []Item
	total_shipping float32
}
