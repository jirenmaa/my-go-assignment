package books

import (
	"fmt"
	"strings"
)

// list every book in the collections
func ListBookService() {
	RenderBook(bookCollections.Books)
}

// list every borrowed book in the collections
func ListBorrowedBookService(borrower string) {
	result := Filter(bookCollections.Books, func(book Book) bool {
		return book.Status == "Borrowed" && strings.Contains(strings.ToLower(book.Borrower), strings.ToLower(borrower))
	})

	if len(result) == 0 {
		fmt.Println("Looks like you haven't borrow any book from this library.")
		// fmt.Println("(press any \"key\" to continue)")
		return
	}

	RenderBook(result)
}

// list every overdue book in the collection
func ListOverdueBookService() {
	result := Filter(bookCollections.Books, func(book Book) bool {
		return book.Status == "Overdue"
	})

	RenderBook(result)
}
