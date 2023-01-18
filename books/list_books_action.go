package books

import (
	"fmt"
	"strings"
)

func (collection *Books) ListBorrowedBookAction(borrower string) {
	var borrowedBooks []Book = Filter(collection.Books, func(book Book) bool {
		// check if the given name is the same as the book borrower
		// and skip it if the book status is "available"
		isTheBorrower := strings.EqualFold(book.Borrower, borrower)
		return !(book.Status == "Available") && isTheBorrower
	})

	if len(borrowedBooks) == 0 {
		fmt.Print("\nlooks like you haven't borrow any book from this library")
		return
	}

	collection.DisplayBookshelf(borrowedBooks)
}

func (collection *Books) ListOverdueBookAction() {
	var overdueBooks []Book = Filter(collection.Books, func(book Book) bool {
		return book.Status == "Overdue"
	})

	collection.DisplayBookshelf(overdueBooks)
}
