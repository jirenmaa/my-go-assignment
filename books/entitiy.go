package books

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/Pallinder/sillyname-go"
)

type Books struct {
	Books []Book `json:"books"`
}

type Book struct {
	ISBN      string `json:"isbn"`
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	Author    string `json:"author"`
	Published string `json:"published"`
	Publisher string `json:"publisher"`
	Pages     int    `json:"pages"`
	Status    string `json:"status"`
	Deadline  int    `json:"deadline"`
	Borrower  string `json:"borrower"`
}

func (collection *Books) Configure(using string) error {
	if using == "json" {
		collection.InitBookFromJSON()
	} else {
		collection.InitBookFromArray()
	}

	collection.UpdateBookStatus()

	return nil
}

func (collection *Books) UpdateBookStatus() error {
	collectionSize := len(collection.Books)

	for index := 0; index < collectionSize; index++ {
		var randomized int64 = RandomNumberGenerator(0, 3)
		var status []string = []string{"Available", "Borrowed", "Overdue"}

		var borrowerName string = sillyname.GenerateStupidName()
		var newStatus string = status[randomized]
		var deadline int64 = RandomNumberGenerator(1, 7)

		collection.Books[index].Status = newStatus
		collection.Books[index].Borrower = borrowerName
		collection.Books[index].Deadline = int(deadline)

		// available book should not have any borrower
		if newStatus == "Available" {
			collection.Books[index].Borrower = "------"
			collection.Books[index].Deadline = 0
		}
	}

	return nil
}

func (collection *Books) InitBookFromJSON() error {
	file, err := os.Open("./books/books.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	bytes, err_read := ioutil.ReadAll(file)
	if err_read != nil {
		return errors.New("")
	}

	err_marshal := json.Unmarshal(bytes, &collection)
	if err_marshal != nil {
		return errors.New("")
	}

	return nil
}

func (collection *Books) InitBookFromArray() error {
	var local []Book = []Book{
		{
			ISBN:      "9781593279509",
			Title:     "Eloquent JavaScript, Third Edition",
			Subtitle:  "A Modern Introduction to Programming",
			Author:    "Marijn Haverbeke",
			Published: "2018-12-04T00:00:00.000Z",
			Publisher: "No Starch Press",
			Pages:     472},
		{
			ISBN:      "9781491943533",
			Title:     "Practical Modern JavaScript",
			Subtitle:  "Dive into ES6 and the Future of JavaScript",
			Author:    "Nicolás Bevacqua",
			Published: "2017-07-16T00:00:00.000Z",
			Publisher: "O'Reilly Media",
			Pages:     334},
		{
			ISBN:      "9781593277574",
			Title:     "Understanding ECMAScript 6",
			Subtitle:  "The Definitive Guide for JavaScript Developers",
			Author:    "Nicholas C. Zakas",
			Published: "2016-09-03T00:00:00.000Z",
			Publisher: "No Starch Press",
			Pages:     352},
		{
			ISBN:      "9781449365035",
			Title:     "Speaking JavaScript",
			Subtitle:  "An In-Depth Guide for Programmers",
			Author:    "Axel Rauschmayer",
			Published: "2014-04-08T00:00:00.000Z",
			Publisher: "O'Reilly Media",
			Pages:     460},
		{
			ISBN:      "9781449331818",
			Title:     "Learning JavaScript Design Patterns",
			Subtitle:  "A JavaScript and jQuery Developer's Guide",
			Author:    "Addy Osmani",
			Published: "2012-08-30T00:00:00.000Z",
			Publisher: "O'Reilly Media",
			Pages:     254},
		{
			ISBN:      "9798602477429",
			Title:     "You Don't Know JS Yet",
			Subtitle:  "Get Started",
			Author:    "Kyle Simpson",
			Published: "2020-01-28T00:00:00.000Z",
			Publisher: "Independently published",
			Pages:     143},
		{
			ISBN:      "9781484200766",
			Title:     "Pro Git",
			Subtitle:  "Everything you neeed to know about Git",
			Author:    "Scott Chacon and Ben Straub",
			Published: "2014-11-18T00:00:00.000Z",
			Publisher: "Apress; 2nd edition",
			Pages:     458},
		{
			ISBN:      "9781484242216",
			Title:     "Rethinking Productivity in Software Engineering",
			Subtitle:  "",
			Author:    "Caitlin Sadowski, Thomas Zimmermann",
			Published: "2019-05-11T00:00:00.000Z",
			Publisher: "Apress",
			Pages:     310},
	}

	collection = &Books{Books: local}

	return nil
}
