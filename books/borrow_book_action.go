package books

import (
	"errors"
	"os"
	"strconv"
)

func (collection *Books) BorrowBookAction(isbn string, durations string) (string, error) {
	var message string

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	duration, err := strconv.Atoi(durations)
	if err != nil {
		panic(err)
	}

	if duration < 1 || duration > 7 {
		message = "you cannot borrow a book longer than 7 or lower than 1 days!"
		return message, errors.New("")
	}

	for index, book := range collection.Books {
		// user can only borrow a book if only the book status is "Available"
		if book.Status == "Available" && book.ISBN == isbn {
			collection.Books[index].Status = "Borrowed"
			collection.Books[index].Borrower = hostname
			collection.Books[index].Deadline = duration
			break
		}
	}

	return message, nil
}
