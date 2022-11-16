package choices

import (
	"fmt"

	"github.com/jirenmaa/gunadarma/gundar-go-library/utils"

	"github.com/jirenmaa/gunadarma/gundar-go-library/books"
)

func SearchForBook() {
	books.ListBookService()
	fmt.Print("\n")
	selected_book := InputTemplate("Input book name $ ")

	book, msg, isFound := books.SearchBookService(selected_book)

	if !isFound {
		fmt.Printf("\n%s", msg)
	} else {
		books.RenderBook(book)
	}
	fmt.Scanln()
	utils.ClearScreen()
}

func ListBookCollection() {
	books.ListBookService()
	fmt.Print("\n")
	selected_book := InputTemplate("Input book ISBN to borrow $ ")

	msg, err := books.BorrowBookService(selected_book)
	if err {
		fmt.Printf("\n%s", msg)
		fmt.Scanln()
	}
	utils.ClearScreen()
}

func ListBorrowedBookCollection() {
	books.ListBorrowedBookService()
	fmt.Print("\n")
	fmt.Scanln()
	utils.ClearScreen()
}

func ListOverdueBookCollection() {
	books.ListOverdueBookService()
	fmt.Print("\n")
	fmt.Scanln()
	utils.ClearScreen()
}

func ReturnBorrowedBook() {
	books.ListBookService()
	fmt.Print("\n")
	selected_book := InputTemplate("Input book ISBN to return $ ")

	msg, err := books.ReturnBookService(selected_book)
	if !err {
		fmt.Printf("\n%s", msg)
		fmt.Scanln()
	}
	utils.ClearScreen()
}
