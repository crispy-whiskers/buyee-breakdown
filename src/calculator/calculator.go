package calculator

import "fmt"

type Person struct {
	Name       string  //name
	Proportion float32 //percentage
	Iou        int
	Ship_b4    int
	Ship_total int
	Item_total int
}

type Item struct {
	Desc     string
	Link     string
	Person   Person
	Yen      int
	Shipping int
}

type Calculator struct {
	People         []Person
	Items          []Item
	total_shipping int
	batched        int
}

func removeIndexPerson(s []Person, index int) []Person {
	ret := make([]Person, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func RemoveIndexItem(s []Item, index int) []Item {
	ret := make([]Item, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func findPerson(s []Item, p Person) int {
	for i, e := range s {
		if e.Person == p {
			return i
		}
	}
	return -1
}

func (c *Calculator) PurgePerson(p Person) {
	for i := findPerson(c.Items, p); i != -1; i = findPerson(c.Items, p) {
		if i != -1 {
			RemoveIndexItem(c.Items, i)
		}
	}
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
		e.Proportion = float32(e.Ship_b4) / float32(c.total_shipping)
		e.Ship_total = int(float32(e.Proportion) * float32(c.batched)) //truncate
		e.Iou = e.Item_total + e.Ship_total
	}

}

func (c *Calculator) Add_person(name string) *Person {
	p := new(Person)
	p.Name = name
	c.People = append(c.People, *p)
	return p
}

func (c *Calculator) Remove_person(name string) bool {
	exists := false
	for i, e := range c.People {
		if e.Name == name {
			exists = true
			c.People = removeIndexPerson(c.People, i)

			c.PurgePerson(e)

			break
		}
	}
	return exists
}

func (c *Calculator) IsPerson(name string) bool {
	exists := false
	for _, e := range c.People {
		if e.Name == name {
			exists = true

			break
		}
	}
	return exists
}

func (c *Calculator) GetPerson(name string) Person {
	for _, e := range c.People {
		if e.Name == name {
			return e

		}
	}
	return Person{"DNE", 0, -1, -1, -1, -1}
}

func (c *Calculator) PrintSelf() {
	for _, e := range c.People {
		fmt.Println(e.Name)
	}
}

func (c *Calculator) AddItem(link string, desc string, p Person, yen int, shipping int) {
	i := new(Item)
	i.Link = link
	i.Person = p
	i.Yen = yen
	i.Shipping = shipping
	i.Desc = desc
	c.Items = append(c.Items, *i)
}
func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func (c *Calculator) GetPeople() []string {
	return Map(c.People, func(p Person) string {
		return p.Name
	})
}

func (c *Calculator) RemoveItem(desc string) {
	for i, e := range c.Items {
		if e.Desc == desc {
			c.Items = RemoveIndexItem(c.Items, i)
		}
	}
}
