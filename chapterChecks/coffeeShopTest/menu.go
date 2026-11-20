package menu

import (
)

type Menu struct {
	ID string
	Name string
	PriceCents int
	InStock bool
}

var items = []Menu {
	{
		ID: "abc",
		Name: "Coffee",
		PriceCents: 5,
		InStock: true,
	},
	{
		ID: "qwe",
		Name: "Water",
		PriceCents: 1,
		InStock: true,
	},
	{
		ID: "xyz",
		Name: "Muffin",
		PriceCents: 3,
		InStock: false,
	},

}


func GetAllItems() []Menu {
	return items
}

func GetItem(id string) (Menu, bool) {
	for _, menu := range items {
		if menu.ID == id {
			return menu, true
		}
	}
	return Menu{}, false
}

func GetInStockItems() []Menu {
	var instock []Menu
	for _, menu := range items {
		if menu.InStock {
			instock = append(instock, menu)
		}
	}
	return instock
}