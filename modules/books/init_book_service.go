package books

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/Pallinder/sillyname-go"
)

var globalBook Books

// GENERATE RANDOM NUMBER BETWEEN 0 AND 2
func generateRandomStatus(max int) int64 {
	// because the rand.New will always generate same result when
	// the same seed is given, then we should create a uniqe ID
	// using time package to make the seed is not always the same
	source := rand.NewSource(time.Now().UnixNano()) // seed
	random := rand.New(source)                      // generator

	return int64(random.Intn(max))
}

func applyRandomStatus(index int) {
	var borrower string = sillyname.GenerateStupidName()
	var randomize int64 = generateRandomStatus(3)
	var status string = []string{"Available", "Borrowed", "Overdue"}[randomize]

	var alreadyBorrowed bool = status == "Borrowed" || status == "Overdue"
	// check if the book status are already BORROWED or OVERDUE if yes, than
	// set the book borrower with given borrower name, else use default
	if alreadyBorrowed {
		globalBook.Books[index].Borrower = borrower
	} else {
		globalBook.Books[index].Borrower = "------"
	}

	globalBook.Books[index].Status = status
}

// NOTE: THIS FUNCTION WILL BE CALLED ONLY ONCE AND
// THIS FUNCTION SHOULD BE CALLED FIRST BEFORE ALL THE FUNCTION IN THIS PACKAGE
func InitBookService() {
	jsonBookDatas, err := os.Open("./data/books.json")
	if err != nil { panic(err) }

	// close json file
	defer jsonBookDatas.Close()

	// read the jsonBookDatas then return as bytes
	byteValues, err := ioutil.ReadAll(jsonBookDatas)
	if err != nil { panic(err) }
	// parse the JSON encoded data and stores the result
	json.Unmarshal(byteValues, &globalBook)

	// apply each the status and borrower for the book
	for index := 0; index < len(globalBook.Books); index++ {
		applyRandomStatus(index)
	}
}
