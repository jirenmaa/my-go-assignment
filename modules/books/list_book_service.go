package books

// LIST ALL BOOK FROM BOOK COLLECTION
func ListBookService() {
	var bookList []Book = globalBook.Books
	RenderBook(bookList)
}

// LIST ALL BORROWED BOOK FROM BOOK COLLECTION
func ListBorrowedBookService() {
	result := BookFilter(globalBook.Books, func(book Book) bool {
		return book.Status == "Borrowed"
	})

	RenderBook(result)
}

// LIST ALL OVERDUE BOOK FROM BOOK COLLECTION
func ListOverdueBookService() {
	result := BookFilter(globalBook.Books, func(book Book) bool {
		return book.Status == "Overdue"
	})

	RenderBook(result)
}
