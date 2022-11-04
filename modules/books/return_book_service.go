package books

import "os"

func ReturnBookService(isbn string) (string, bool) {
	var isReturned bool = false
	var messageErr string = "Something went terribly wrong. Are you certain that the ISBN is correct?"

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	for index, book := range globalBook.Books {
		// check if the borrower are actually the user and not the other
		if book.Borrower == hostname && book.ISBN == isbn {
			isReturned = true
			globalBook.Books[index].Borrower = "------"
			globalBook.Books[index].Status = "Available"
			break
		}
	}

	return messageErr, isReturned
}
