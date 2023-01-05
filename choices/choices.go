package choices

import (
	"fmt"

	"github.com/jirenmaa/gunadarma/gundar-go-library/utils"

	"github.com/jirenmaa/gunadarma/gundar-go-library/books"
)

func SearchForBook() {
	books.ListBookService()
	fmt.Print("\n")
	selected_book := utils.InputTemplate("Input book name $ ")

	book, msg, isFound := books.SearchBookService(selected_book)

	if !isFound {
		fmt.Printf("\n%s", msg)
	} else {
		books.RenderBook(book)
	}
	fmt.Scanln()
	utils.ClearScreen()
}

func BorrowBook() {
	for {
		books.ListBookService()
		fmt.Print("\n")
		fmt.Println("(type \"q\" to quit)")

		selected_book := utils.InputTemplate("Input book ISBN to borrow $ ")
		if selected_book == "q" {
			break
		}

		duration := utils.InputTemplate("Duration (number in days) $ ")
		msg, err := books.BorrowBookService(selected_book, duration)

		if err {
			fmt.Printf("\n%s", msg)
			fmt.Scanln()
		}

		utils.ClearScreen()
	}
	utils.ClearScreen()
}

func ReturnBorrowedBook() {
	books.ListBookService()
	fmt.Print("\n")
	selected_book := utils.InputTemplate("Input book ISBN to return $ ")

	msg, err := books.ReturnBookService(selected_book)
	if !err {
		fmt.Printf("\n%s", msg)
		fmt.Scanln()
	}
	utils.ClearScreen()
}

func AddBookToCollection() {
	msg, err := books.CreateBookService()
	if !err {
		fmt.Printf("\n%s", msg)
		fmt.Scanln()
	}
	utils.ClearScreen()
}

func ListBorrowedBookCollection() {
	books.ListBookService()
	fmt.Print("\n")
	borrower := utils.InputTemplate("Input borrower name $ ")

	books.ListBorrowedBookService(borrower)
	fmt.Println("\n(press any \"key\" to continue)")
	fmt.Scanln()
	utils.ClearScreen()
}

func ListOverdueBookCollection() {
	books.ListOverdueBookService()
	fmt.Print("\n")
	fmt.Scanln()
	utils.ClearScreen()
}
