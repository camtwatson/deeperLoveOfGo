package books

import "fmt"

type Book struct {
	Title string
	Author string
	Copies int
	ID string
}

var catalog = []Book{
	{
		Title: "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
		ID: "3aqfr",
	},
	{
		Title: "White Heat",
		Author: "Dominic Sandbrook",
		Copies: 2,
		ID: "5ytr7",
	},
}

func BookToString(book Book) string {
	result := fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
	return result
}

func GetAllBooks() []Book {
	return catalog
}

func GetBook(id string) (Book, bool) {
	for _, book := range catalog {
		if book.ID == id {
			return book, true
		}
	}
	return Book{}, false
}