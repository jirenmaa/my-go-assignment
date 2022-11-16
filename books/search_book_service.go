package books

import "strings"

func isBookAlreadyFound(array []Book, content string) bool {
	for _, item := range array {
		return content == item.Title
	}

	return false
}

func SearchBookService(book_name string) ([]Book, string, bool) {
	var searchedBook []Book
	// instead of using the entire book title, split the keyword search into its individual words.
	var splitContent []string = strings.Split(book_name, " ")

	var isBookFound bool = false
	var messageErr string = "It looks that the book you are looking for is not in our collection.\nTry looking for another book."

	for _, word := range splitContent {
		// loop over every book in the collection
		for _, book := range bookCollections.Books {
			// verify that the "searched string" and the "book title" match.
			if strings.Contains(strings.ToLower(book.Title), strings.ToLower(word)) && !isBookAlreadyFound(searchedBook, book.Title) {
				isBookFound = true
				searchedBook = append(searchedBook, book)
			}
		}
	}

	return searchedBook, messageErr, isBookFound
}
