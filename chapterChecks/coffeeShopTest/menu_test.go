package menu_test

import (
	"cmp"
	"menu"
	"slices"
	"testing"
)

func TestingMenu() map[string]menu.Menu {
	 return map[string]menu.Menu {
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


func TestGetAllItems_ListInventory(t *testing.T) {
	t.Parallel()
	items := TestingMenu()
	want :=	[]menu.Menu { 
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

	got := menu.GetAllItems(items)
	slices.SortFunc(got, func(a, b menu.Menu) int {
		return cmp.Compare(a.ID, b.ID)
	})

	if !slices.Equal(want, got) {
		t.Fatalf("want: %#v, got %#v", want, got)
	}
}

func TestGetItem_Found(t *testing.T) {
	t.Parallel()
	items := TestingMenu()
	want :=	menu.Menu { 
			ID: "abc",
			Name: "Coffee",
			PriceCents: 5,
			InStock: true,
	}
	got, ok := menu.GetItem(items, "abc")
	
	if !ok {
		t.Fatalf("Item not found")
	}

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}

func TestGetItem_NotFound(t *testing.T) {
	t.Parallel()
	items := TestingMenu()
	_ , ok := menu.GetItem(items, "123")
	if ok {
		t.Fatalf("want false for nonexistent ID, got true")
	}
}


func TestGetInStockItem_ReturnInStock(t *testing.T) {
	t.Parallel()
	items := TestingMenu()
	want := []menu.Menu {
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
	}

	got := menu.GetInStockItems(items)
	slices.SortFunc(got, func(a, b menu.Menu) int {
		return cmp.Compare(a.ID, b.ID)
	})

	if !slices.Equal(want, got) {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}

func TestAddItem(t *testing.T) {
	t.Parallel()
	items := TestingMenu()
	_, ok := menu.GetItem(items, "asd")
	if ok {
		t.Fatalf("Item already exist")
	}

	menu.AddItem(items,menu.Menu {
			ID: "asd",
			Name: "Tea",
			PriceCents: 2,
			InStock: true,
		})
	
	_, ok = menu.GetItem(items, "asd") 
	if !ok {
		t.Fatalf("Item not added")
	}
}