package main

import "fmt"
import "books"

func main() {
	book:= books.Book {
		Title: "Engineering in PLain Sight",
		Author: "Grady Hillhouse",
		Copies: 2,
	}
	fmt.Println(books.BookToString(book))
}
