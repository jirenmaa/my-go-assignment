package books

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jirenmaa/gunadarma/gundar-go-library/utils"
)

func GettingInput() (string, string, string, int) {
	title := utils.InputTemplate("Type the book title \t\t$ ")
	author := utils.InputTemplate("Enter the book's author\t\t$ ")
	publisher := utils.InputTemplate("Enter the book's publisher\t$ ")
	pages := utils.InputTemplate("Enter the number of book pages\t$ ")

	page, err := strconv.Atoi(pages)
	if err != nil {
		panic(err)
	}

	return title, author, publisher, page
}

func CreateBookService() (string, bool) {
	var title, author, publisher string
	var pages int

	title, author, publisher, pages = GettingInput()
	isbn := fmt.Sprintf("%s%s", "9", strconv.Itoa(int(RandomNumberGenerator(999999999999))))
	timestamp := time.Now().Format(time.RFC3339)

	bookCollections.Books = append(bookCollections.Books, Book{
		ISBN:        isbn,
		Title:       title,
		Subtitle:    "------",
		Author:      author,
		Published:   timestamp,
		Publisher:   publisher,
		Pages:       pages,
		Description: "------",
		Status:      "Available",
		Borrower:    "------",
	})

	return "", true
}
