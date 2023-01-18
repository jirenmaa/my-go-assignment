package choices1

import (
	"fmt"

	books "github.com/jirenmaa/gunadarma/gundar-go-library/books"
	util "github.com/jirenmaa/gunadarma/gundar-go-library/utils"
)

func ChoiceSearchBook(collections *books.Books) {
	for {
		collections.DisplayBookshelf(collections.Books)

		fmt.Println("\n(type \"q\" to quit)")

		var prompt string = util.InputTemplate("Enter Book Name $ ")
		if prompt == "q" { // quit
			break
		}

		books, msg, found := collections.SearchBookAction(prompt)

		if !found {
			fmt.Printf("\n%s", msg)
		} else {
			fmt.Println()
			// display all founded book that has similar to the prompt
			collections.DisplayBookshelf(books)
			fmt.Print("\n(press any \"key\" to continue)")
		}

		fmt.Scanln()
		util.ClearScreen()
	}

	util.ClearScreen()
}

func ChoiceBorrowBook(collections *books.Books) {
	for {
		collections.DisplayBookshelf(collections.Books)

		fmt.Println("\n(type \"q\" to quit)")

		var prompt string = util.InputTemplate("Enter ISBN $ ")
		if prompt == "q" { // quit
			break
		}

		duration := util.InputTemplate("Duration (number in days) $ ")
		msg, err := collections.BorrowBookAction(prompt, duration)
		if err != nil {
			fmt.Printf("\n%s", msg)
			fmt.Scanln()
		}

		util.ClearScreen()
	}

	util.ClearScreen()
}

func ChoiceReturnBook(collections *books.Books) {
	for {
		collections.DisplayBookshelf(collections.Books)

		fmt.Println("\n(type \"q\" to quit)")

		var prompt string = util.InputTemplate("Enter ISBN $ ")
		if prompt == "q" { // quit
			break
		}

		msg, err := collections.ReturnBookAction(prompt)
		if err != nil {
			fmt.Printf("\n%s", msg)
			fmt.Scanln()
		}

		util.ClearScreen()
	}

	util.ClearScreen()
}

func ChoiceAddNewBook(collections *books.Books) {
	for {
		msg, err := collections.CreateBookAction()
		if err != nil && msg == "q" {
			break
		}

		fmt.Println(msg)
		fmt.Scanln()
		util.ClearScreen()
	}

	util.ClearScreen()
}

func ChoiceDisplayBorrowedBooks(collections *books.Books) {
	for {
		collections.DisplayBookshelf(collections.Books)

		fmt.Println("\n(type \"q\" to quit)")

		var prompt string = util.InputTemplate("Enter Borrower Name $ ")
		if prompt == "q" { // quit
			break
		}

		collections.ListBorrowedBookAction(prompt)
		fmt.Print("\n(press any \"key\" to continue)")
		fmt.Scanln()
		util.ClearScreen()
	}

	util.ClearScreen()
}

func ChoiceDisplayOverduesBooks(collections *books.Books) {
	collections.ListOverdueBookAction()

	fmt.Print("\n(press any \"key\" to continue)")
	fmt.Scanln()
	util.ClearScreen()
}
