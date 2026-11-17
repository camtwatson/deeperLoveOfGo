package label_test

import "testing"
import "label"

func TestShelfLabel_OutputString(t *testing.T) {
	book := label.Book {
		Title: "The Hobbit",
		Author: "J.R.R. Tolkien",
		PriceCents: 14.99,
		Copies: 3,
	}

	want := "The Hobbit by J.R.R. Tolkien ($14.99, 3 copies)"
	got := label.ShelfLabel(book)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}

func TestShelfLabel_OneCopy(t *testing.T) {
		book := label.Book {
		Title: "The Hobbit",
		Author: "J.R.R. Tolkien",
		PriceCents: 14.99,
		Copies: 1,
	}

	want := "The Hobbit by J.R.R. Tolkien ($14.99, 1 copy)"
	got := label.ShelfLabel(book)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}
