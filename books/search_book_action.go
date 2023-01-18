package books

import (
	"strings"
)

func foundBook(books []Book, word string) bool {
	for _, book := range books {
		return word == book.Title
	}
	return false
}

func (collection *Books) SearchBookAction(bookName string) ([]Book, string, bool) {
	var bucketBooks []Book
	// instead of using the entire book title, split the keyword search into its individual words.
	var words []string = strings.Split(bookName, " ")

	var alreadyFound bool = false
	var message string = "It looks that the book you are looking for is not in our collection.\nTry looking for another book."

	for _, word := range words {
		for _, book := range collection.Books {
			var match bool = strings.Contains(strings.ToLower(book.Title), strings.ToLower(word))

			// check if the "searched word" and the book title is match
			if match && !foundBook(bucketBooks, book.Title) {
				alreadyFound, message = true, ""
				bucketBooks = append(bucketBooks, book)
			}
		}
	}

	return bucketBooks, message, alreadyFound
}
