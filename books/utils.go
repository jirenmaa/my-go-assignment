package books

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

func (collection *Books) DisplayBookshelf(books []Book) {
	var header []string = []string{
		"isbn", "title", "author", "pages", "status", "deadline", "borrower",
	}

	// configure tablewriter
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)

	for _, book := range books {
		// change the type from int to string, to be concate with string later
		var deadline string = strconv.Itoa(book.Deadline)

		if book.Deadline <= 1 && (book.Status == "Available") {
			deadline = "------"
		} else {
			deadline += " days"
		}

		var body []string = []string{
			book.ISBN,
			book.Title,
			book.Author,
			strconv.Itoa(book.Pages),
			book.Status,
			deadline,
			book.Borrower,
		}

		table.Append(body)
	}

	table.Render()
}

func RandomNumberGenerator(min int, max int) int64 {
	// every time rand.new called, if the parameter does not provide a unique source
	// the result will always the same, so the params need to be unique
	source := rand.NewSource(time.Now().UnixNano())

	// generate random from seed
	random := rand.New(source)

	return int64(random.Intn(max-min) + min)
}

func Filter(books []Book, condition func(Book) bool) []Book {
	result := []Book{}

	for index := range books {
		if condition(books[index]) {
			result = append(result, books[index])
		}
	}

	return result
}
