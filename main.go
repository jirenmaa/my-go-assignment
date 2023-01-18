package main

import (
	"fmt"

	books "github.com/jirenmaa/gunadarma/gundar-go-library/books"
	choice "github.com/jirenmaa/gunadarma/gundar-go-library/choices"
	util "github.com/jirenmaa/gunadarma/gundar-go-library/utils"
)

func main() {
	var collections books.Books

	collections.Configure("json")

	for {
		collections.DisplayBookshelf(collections.Books)
		fmt.Println()

		var prompt int = choice.DisplayBookShelfMenu()
		util.ClearScreen()

		switch prompt {
		case 1:
			choice.ChoiceSearchBook(&collections)
		case 2:
			choice.ChoiceBorrowBook(&collections)
		case 3:
			choice.ChoiceReturnBook(&collections)
		case 4:
			choice.ChoiceAddNewBook(&collections)
		case 5:
			choice.ChoiceDisplayBorrowedBooks(&collections)
		case 6:
			choice.ChoiceDisplayOverduesBooks(&collections)
		}
	}
}
