package menu

import (
	"maps"
	"slices"
)

type Menu struct {
	ID string
	Name string
	PriceCents int
	InStock bool
}
func GetMenu() map[string]Menu {
	 return map[string]Menu {
		"abc": {
			ID: "abc",
			Name: "Coffee",
			PriceCents: 5,
			InStock: true,
		},
		"qwe": {
			ID: "qwe",
			Name: "Water",
			PriceCents: 1,
			InStock: true,
		},
		"xyz": {
			ID: "xyz",
			Name: "Muffin",
			PriceCents: 3,
			InStock: false,
		},

	}
}


func GetAllItems(item map[string]Menu) []Menu {
	return slices.Collect(maps.Values(item))
}

func GetItem(item map[string]Menu, id string) (Menu, bool) {
	order, ok := item[id]
	return order, ok
}

func GetInStockItems(item map[string]Menu) []Menu {
	var instock []Menu
	for _, menu := range item {
		if menu.InStock {
			instock = append(instock, menu)
		}
	}
	return instock
}

func AddItem(item map[string]Menu, newItem Menu) {
	item[newItem.ID] = newItem
}