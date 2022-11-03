package books

import "strings"

func isBookAlreadyFound(array []Book, content string) bool {
	for _, item := range array {
		return content == item.Title
	}

	return false
}

// LIST ALL BOOK FROM BOOK COLLECTION WITH GIVEN CONDITION
func SearchBookService(book_name string) ([]Book, string, bool) {
	// a bucket for the books that matches with condition
	var searchedBook []Book
	// split the searched book name, and search it by every word instead the whole word
	var splitContent []string = strings.Split(book_name, " ")

	var isBookFound bool = false
	var messageErr string = "It looks that the book you are looking for is not in our collection.\nTry looking for another book."

		for _, word := range splitContent {
			// loop all the book from the collection
			for _, book := range globalBook.Books {
				// check if the "searched string" match with the "book title"
				if strings.Contains(strings.ToLower(book.Title), strings.ToLower(word)) && !isBookAlreadyFound(searchedBook, book.Title) {
					isBookFound = true
					searchedBook = append(searchedBook, book)
				}
			}
	}

	return searchedBook, messageErr, isBookFound
}
