package label

import "fmt"

type Book struct {
	Title string 
	Author string
	PriceCents float32
	Copies int
}

func ShelfLabel(book Book) string {
	end := "copies"
	if book.Copies == 1 {
		end = "copy"
	}
	
	return fmt.Sprintf("%v by %v ($%v, %v %v)", book.Title, book.Author, book.PriceCents, book.Copies, end)
}
