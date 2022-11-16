package books

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// create a table view of the book collection
func RenderBook(books []Book) {
	var tableHeader []string = []string{"isbn", "title", "author", "pages", "status", "borrower"}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(tableHeader)

	for _, v := range books {
		table.Append([]string{v.ISBN, v.Title, v.Author, strconv.Itoa(v.Pages), v.Status, v.Borrower})
	}
	table.Render()
}

// filter the provided book selection based on condition
func BookFilter(books []Book, condition func(Book) bool) []Book {
	result := []Book{}

	for i := range books {
		if condition(books[i]) {
			result = append(result, books[i])
		}
	}
	return result
}
