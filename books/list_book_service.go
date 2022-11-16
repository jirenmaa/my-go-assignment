package books

// list every book in the collections
func ListBookService() {
	RenderBook(bookCollections.Books)
}

// list every borrowed book in the collections
func ListBorrowedBookService() {
	result := BookFilter(bookCollections.Books, func(book Book) bool {
		return book.Status == "Borrowed"
	})

	RenderBook(result)
}

// list every overdue book in the collection
func ListOverdueBookService() {
	result := BookFilter(bookCollections.Books, func(book Book) bool {
		return book.Status == "Overdue"
	})

	RenderBook(result)
}
