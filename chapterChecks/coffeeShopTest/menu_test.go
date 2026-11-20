package menu_test

import (
	"menu"
	"slices"
	"testing"
)

func TestGetAllItems_ListInventory(t *testing.T) {
	t.Parallel()
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

	got := menu.GetAllItems()

	if !slices.Equal(want, got) {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGetItem_Found(t *testing.T) {
	t.Parallel()
	want :=	menu.Menu { 
			ID: "abc",
			Name: "Coffee",
			PriceCents: 5,
			InStock: true,
	}
	got, ok := menu.GetItem("abc")
	
	if !ok {
		t.Fatalf("Item not found")
	}

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}

func TestGetItem_NotFound(t *testing.T) {
	t.Parallel()
	_ , ok := menu.GetItem("123")
	if ok {
		t.Fatalf("want false for nonexistent ID, got true")
	}
}


func TestGetInStockItem_ReturnInStock(t *testing.T) {
	t.Parallel()
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

	got := menu.GetInStockItems()

	if !slices.Equal(want, got) {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}