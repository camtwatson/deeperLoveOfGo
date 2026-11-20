package main

import (
	"fmt"
	"label"
)

func main() {
	book := label.Book {
		Title: "The Sun Also Rises",
		Author: "Ernest Hemmingway",
		PriceCents: 15.50,
		Copies: 5,
	}
	
	fmt.Println(label.ShelfLabel(book))
}