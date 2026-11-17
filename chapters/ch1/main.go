package main

import "fmt"

func hello() {
	var book = "Ready player One"
	var author = "some guy"
	printBook(book, author)
	book = "Book of things"
	author = "Author Guy"
	printBook(book, author)
}

func printBook(book, author string) {
	fmt.Println(book, "by", author)
}

func main() {
	hello()
}
