package main

import (
	"books"
	"fmt"
	"os"
)

func main() {
	catalog := books.GetCatalog()
	if len(os.Args) != 2 {
		fmt.Println("Usage: find <BOOK ID>")
		return
	}
	
	id := os.Args[1]
	book, ok := books.GetBook(catalog, id)
	if !ok {
		fmt.Println("Sorry, I couldn't find that book in the catalog.")
		return
	}
	fmt.Println(books.BookToString(book))
}