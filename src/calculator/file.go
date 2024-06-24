package calculator

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Calculator) SaveAsString() ([]byte, error) {
	// Convert the Calculator struct to JSON
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %v", err)
	}

	return data, nil
}

func LoadFromFile(data []byte) (Calculator, error) {
	var c Calculator

	// Convert the JSON data to a Calculator struct
	err := json.Unmarshal(data, &c)
	if err != nil {
		return c, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return c, nil
}

func (c *Calculator) ShowAsTablestring() string {
	var builder strings.Builder

	// Table for Items
	builder.WriteString("Person\tPrice\tLink\tDescription\n")
	for _, item := range c.Items {
		price := item.Yen
		builder.WriteString(fmt.Sprintf("%s\t%d\t%s\t%s\n", item.Person.Name, price, item.Link, item.Desc))
	}

	builder.WriteString("\n") // Add a newline to separate tables

	// Table for People
	builder.WriteString("Person\tItem Total\tShip Before\tShip Total\tProportion\n")
	for _, person := range c.People {
		builder.WriteString(fmt.Sprintf("%s\t%d\t%d\t%d\t%.2f\n", person.Name, person.Item_total, person.Ship_b4, person.Ship_total, person.Proportion))
	}

	return builder.String()
}
