package main

import "fmt"

type Book struct {
	Title string
	Author string
	Copies int
}

func printBook(book Book) {
	fmt.Printf("%v by %v - %v copies\n", book.Title,  book.Author, book.Copies)
}

func defineBook() {
		var book = Book{ 
		Title: "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Copies: 5,
	}
	printBook(book)
}

func main() {
	fmt.Println("Books in stock:")
	defineBook()
}
