package books

import "os"

func ReturnBookService(isbn string) (string, bool) {
	var isReturned bool = false
	var messageErr string = "Something went terribly wrong. Are you certain that the ISBN is correct?"

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	for index, book := range bookCollections.Books {
		// verify that the borrower is the actual user
		if book.Borrower == hostname && book.ISBN == isbn {
			isReturned = true
			bookCollections.Books[index].Borrower = "------"
			bookCollections.Books[index].Status = "Available"
			break
		}
	}

	return messageErr, isReturned
}
