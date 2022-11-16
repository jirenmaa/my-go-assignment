package books

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/Pallinder/sillyname-go"
)

var bookCollections Books

// generate random number between 0 to 2
func generateRandomStatus(max int) int64 {
	// When the same seed is used, rand.New will always provide the same result,
	// thus we need use time packages to build a unique ID so that the seed is never the same.
	source := rand.NewSource(time.Now().UnixNano()) // seed
	random := rand.New(source)                      // generator

	return int64(random.Intn(max))
}

func applyRandomStatus(index int) {
	var borrower string = sillyname.GenerateStupidName()
	var randomize int64 = generateRandomStatus(3)
	var status string = []string{"Available", "Borrowed", "Overdue"}[randomize]

	var alreadyBorrowed bool = status == "Borrowed" || status == "Overdue"
	// determine if the books are now BORROWED or OVERDUE.
	// If so, set the book borrower to the specified borrower name; else, use default.
	if alreadyBorrowed {
		bookCollections.Books[index].Borrower = borrower
	} else {
		bookCollections.Books[index].Borrower = "------"
	}

	bookCollections.Books[index].Status = status
}

// NOTE: This method should be called first in this package
func InitBookService() {
	jsonBookDatas, err := os.Open("./books/books.json")
	if err != nil {
		panic(err)
	}

	defer jsonBookDatas.Close()

	// then store as bytes after reading the jsonBookDatas
	byteValues, err := ioutil.ReadAll(jsonBookDatas)
	if err != nil {
		panic(err)
	}
	// analyze the JSON-encoded data and save the result
	json.Unmarshal(byteValues, &bookCollections)

	// apply both the status and borrower for the book
	for index := 0; index < len(bookCollections.Books); index++ {
		applyRandomStatus(index)
	}
}
