package books

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

// create a table view of the book collection
func RenderBook(books []Book) {
	var tableHeader []string = []string{"isbn", "title", "author", "pages", "status", "deadline", "borrower"}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(tableHeader)

	for _, v := range books {
		var deadline string = strconv.Itoa(v.Deadline)
		if v.Deadline < 1 && (v.Status == "Available") {
			deadline = "------"
		} else {
			deadline += " days"
		}

		table.Append([]string{v.ISBN, v.Title, v.Author, strconv.Itoa(v.Pages), v.Status, deadline, v.Borrower})
	}
	table.Render()
}

// filter selection based on condition
func Filter(books []Book, condition func(Book) bool) []Book {
	result := []Book{}

	for i := range books {
		if condition(books[i]) {
			result = append(result, books[i])
		}
	}
	return result
}

// generate random number between 0 to max
func RandomNumberGenerator(max int) int64 {
	// When the same seed is used, rand.New will always provide the same result,
	// thus we need use time packages to build a unique ID so that the seed is never the same.
	source := rand.NewSource(time.Now().UnixNano()) // seed
	random := rand.New(source)                      // generator

	return int64(random.Intn(max))
}
