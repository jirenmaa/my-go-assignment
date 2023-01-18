package books

import (
	"errors"
	"os"
)

func (collection *Books) ReturnBookAction(isbn string) (string, error) {
	var message string = "Are you certain the ISBN you input are correct?"
	var err error = nil

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	for index, book := range collection.Books {
		// check if the the user is borrowing this book
		if book.Borrower == hostname && book.ISBN == isbn {
			collection.Books[index].Borrower = "-----"
			collection.Books[index].Deadline = 0
			collection.Books[index].Status = "Available"
			break
		}

		if book.Borrower != hostname && book.ISBN == isbn {
			message = "You currently are not borrowing this book"
			err = errors.New("")
		}
	}

	return message, err
}
