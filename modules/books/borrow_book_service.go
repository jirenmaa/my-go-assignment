package books

import "os"

func BorrowBookService(isbn string) (string, bool) {
	var isBookBorrowed bool = true
	var messageErr string = "It appears that the book you are attempting to borrow has already been borrowed.\nTry borrowing another book."

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	for index, book := range globalBook.Books {
		// user can only borrow a book if only the book status is Available
		if book.Status == "Available" && book.ISBN == isbn {
			isBookBorrowed = false
			globalBook.Books[index].Status = "Borrowed"
			globalBook.Books[index].Borrower = hostname
			break
		}
	}

	return messageErr, isBookBorrowed
}
