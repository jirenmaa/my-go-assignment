package books

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// RENDER GIVEN BOOK DATA IN TABLE
func RenderBook(books []Book) {
	var tableHeader []string = []string{"isbn", "title", "author", "pages", "status", "borrower"}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(tableHeader)

	for _, v := range books {
		table.Append([]string{v.ISBN, v.Title, v.Author, strconv.Itoa(v.Pages), v.Status, v.Borrower})
	}
	table.Render()
}

// FILTER GIVEN BOOK DATA WITH GIVEN CONDITION
func BookFilter(books []Book, condition func(Book) bool) []Book {
	result := []Book{}

	for i := range books {
		if condition(books[i]) {
			result = append(result, books[i])
		}
	}
	return result
}
