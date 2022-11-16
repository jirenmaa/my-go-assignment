package main

import (
	"fmt"

	"github.com/jirenmaa/gunadarma/gundar-go-library/utils"

	"github.com/jirenmaa/gunadarma/gundar-go-library/books"
	"github.com/jirenmaa/gunadarma/gundar-go-library/choices"
)

func main() {
	books.InitBookService()

	for {
		books.ListBookService()
		fmt.Print("\n")

		selected_command := choices.ShowAvailableCommands()
		utils.ClearScreen()

		switch selected_command {
		case 1: choices.SearchForBook() // search book
		case 2: choices.ListBookCollection() // borrow a book
		case 3: choices.ReturnBorrowedBook() // return a book
		case 4: choices.ListBorrowedBookCollection() // list all borrowed book
		case 5: choices.ListOverdueBookCollection() // list all overdue book
		}
	}
}
