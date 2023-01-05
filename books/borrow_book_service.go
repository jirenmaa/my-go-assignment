package books

import (
	"os"
	"strconv"
)

func BorrowBookService(isbn string, duration string) (string, bool) {
	var isBookBorrowed bool = true
	var messageErr string = "It appears that the book you are attempting to borrow has already been borrowed.\nTry borrowing another book."

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	durations, err := strconv.Atoi(duration)
	if err != nil {
		panic(err)
	}

	if durations > 7 {
		messageErr = "You cannot borrow a book longer than 7 days!"
	}

	for index, book := range bookCollections.Books {
		// user can only borrow a book if only the book status is Available
		if book.Status == "Available" && book.ISBN == isbn {
			isBookBorrowed = false
			bookCollections.Books[index].Status = "Borrowed"
			bookCollections.Books[index].Borrower = hostname
			bookCollections.Books[index].Deadline = durations
			break
		}
	}

	return messageErr, isBookBorrowed
}
