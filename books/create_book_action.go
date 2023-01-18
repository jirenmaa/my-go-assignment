package books

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	util "github.com/jirenmaa/gunadarma/gundar-go-library/utils"
)

func getInputFromUser() (string, string, string, int) {
	fmt.Println("\n(enter \"q\" to quit)")

	titles := util.InputTemplate("Enter book titles \t\t$ ") // q
	author := util.InputTemplate("Enter book author \t\t$ ") // q
	patron := util.InputTemplate("Enter book patron \t\t$ ") // q
	pages := util.InputTemplate("Enter book pages  \t\t$ ")  // 0

	page, err := strconv.Atoi(pages)
	if err != nil {
		panic(err)
	}

	return titles, author, patron, page
}

func (collection *Books) CreateBookAction() (string, error) {
	var message string

	title, author, publisher, pages := getInputFromUser()

	if title+author+publisher == "qqq" {
		return "q", errors.New("")
	}

	// this is for published book when book is created
	var timestamp string = time.Now().Format(time.RFC3339)
	// books title cannot be same as the that already existed
	var duplicate []Book = Filter(collection.Books, func(book Book) bool {
		return strings.EqualFold(book.Title, title)
	})

	// random number that has length of 12
	_r := strconv.Itoa(int(RandomNumberGenerator(11111111111, 999999999999)))
	var isbn string = fmt.Sprintf("%s%s", "9", _r)

	if len(duplicate) == 0 {
		collection.Books = append(collection.Books, Book{
			ISBN:      isbn,
			Title:     title,
			Author:    author,
			Status:    "Available",
			Published: timestamp,
			Publisher: publisher,
			Pages:     pages,
			Borrower:  "------",
			Subtitle:  "------",
		})

		message = "New Book was Added"
		return message, nil
	}

	message = fmt.Sprintf("\nBook with title \"%s\" already exists", title)
	return message, errors.New("")
}
