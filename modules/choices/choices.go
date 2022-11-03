package choices

import (
	"fmt"

	"github.com/jirenmaa/gundar-go-library/modules/utils"

	"github.com/jirenmaa/gundar-go-library/modules/books"
	"github.com/jirenmaa/gundar-go-library/modules/commands"
)

// [SEARCH BOOK] user can search book here
func ChoiceOne() {
	books.ListBookService()
	fmt.Print("\n")
	selected_book := commands.SearchForBook()

	book, msg, isFound := books.SearchBookService(selected_book)

	if !isFound {
		fmt.Printf("\n%s", msg)
	} else {
		books.RenderBook(book)
	}
	fmt.Scanln()
	utils.ClearScreen()
}

// [BORROW BOOK] user are borrow a book here
func ChoiceTwo() {
	books.ListBookService()
	fmt.Print("\n")
	selected_book := commands.ChooseBookToBorrow()

	msg, err := books.BorrowBookService(selected_book)
	if err {
		fmt.Printf("\n%s", msg)
		fmt.Scanln()
	}
	utils.ClearScreen()
}

// [RETURN BORROWED BOOK] use can return the borrowed book here
func ChoiceThree() {
	books.ListBookService()
	fmt.Print("\n")
	selected_book := commands.ChooseBookToReturn()

	msg, err := books.ReturnBookService(selected_book)
	if !err {
		fmt.Printf("\n%s", msg)
		fmt.Scanln()
	}
	utils.ClearScreen()
}

// [BORROWED BOOKS] user can see all borrowed book here
func ChoiceFour() {
	books.ListBorrowedBookService()
	fmt.Print("\n")
	fmt.Scanln()
	utils.ClearScreen()
}

// [OVERDUE BOOKS] user can see all overdue borrowed book here
func ChoiceFive() {
	books.ListOverdueBookService()
	fmt.Print("\n")
	fmt.Scanln()
	utils.ClearScreen()
}
