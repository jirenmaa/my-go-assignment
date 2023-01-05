package books

import (
	"encoding/json"
	"io/ioutil"
	"os"

	// "time"

	"github.com/Pallinder/sillyname-go"
)

var bookCollections Books

func applyStatus() {
	for index := 0; index < len(bookCollections.Books); index++ {
		var borrower string = sillyname.GenerateStupidName()
		var randomize int64 = RandomNumberGenerator(3)
		var status string = []string{"Available", "Borrowed", "Overdue"}[randomize]

		// determine if the books are now BORROWED or OVERDUE.
		// If so, set the book borrower to the specified borrower name; else, use default.
		var alreadyBorrowed bool = status == "Borrowed" || status == "Overdue"
		if alreadyBorrowed {
			bookCollections.Books[index].Borrower = borrower

			var deadline int64 = RandomNumberGenerator(5) + 2
			bookCollections.Books[index].Deadline = int(deadline)
		} else {
			bookCollections.Books[index].Borrower = "------"
		}

		bookCollections.Books[index].Status = status
	}
}

// NOTE: This method should be called first in this package
func InitBookServiceJSON() {
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
	applyStatus()
}

// NOTE: This method should be called first in this package (without using json file)
func InitBookServiceArrs() {
	var collections []Book = []Book{
		{
			ISBN:        "9781593279509",
			Title:       "Eloquent JavaScript, Third Edition",
			Subtitle:    "A Modern Introduction to Programming",
			Author:      "Marijn Haverbeke",
			Published:   "2018-12-04T00:00:00.000Z",
			Publisher:   "No Starch Press",
			Pages:       472,
			Description: "JavaScript lies at the heart of almost every modern web application, from social apps like Twitter to browser-based game frameworks like Phaser and Babylon. Though simple for beginners to pick up and play with, JavaScript is a flexible, complex language that you can use to build full-scale applications."},
		{
			ISBN:        "9781491943533",
			Title:       "Practical Modern JavaScript",
			Subtitle:    "Dive into ES6 and the Future of JavaScript",
			Author:      "Nicolás Bevacqua",
			Published:   "2017-07-16T00:00:00.000Z",
			Publisher:   "O'Reilly Media",
			Pages:       334,
			Description: "To get the most out of modern JavaScript, you need learn the latest features of its parent specification, ECMAScript 6 (ES6). This book provides a highly practical look at ES6, without getting lost in the specification or its implementation details."},
		{
			ISBN:        "9781593277574",
			Title:       "Understanding ECMAScript 6",
			Subtitle:    "The Definitive Guide for JavaScript Developers",
			Author:      "Nicholas C. Zakas",
			Published:   "2016-09-03T00:00:00.000Z",
			Publisher:   "No Starch Press",
			Pages:       352,
			Description: "ECMAScript 6 represents the biggest update to the core of JavaScript in the history of the language. In Understanding ECMAScript 6, expert developer Nicholas C. Zakas provides a complete guide to the object types, syntax, and other exciting changes that ECMAScript 6 brings to JavaScript."},
		{
			ISBN:        "9781449365035",
			Title:       "Speaking JavaScript",
			Subtitle:    "An In-Depth Guide for Programmers",
			Author:      "Axel Rauschmayer",
			Published:   "2014-04-08T00:00:00.000Z",
			Publisher:   "O'Reilly Media",
			Pages:       460,
			Description: "Like it or not, JavaScript is everywhere these days -from browser to server to mobile- and now you, too, need to learn the language or dive deeper than you have. This concise book guides you into and through JavaScript, written by a veteran programmer who once found himself in the same position."},
		{
			ISBN:        "9781449331818",
			Title:       "Learning JavaScript Design Patterns",
			Subtitle:    "A JavaScript and jQuery Developer's Guide",
			Author:      "Addy Osmani",
			Published:   "2012-08-30T00:00:00.000Z",
			Publisher:   "O'Reilly Media",
			Pages:       254,
			Description: "With Learning JavaScript Design Patterns, you'll learn how to write beautiful, structured, and maintainable JavaScript by applying classical and modern design patterns to the language. If you want to keep your code efficient, more manageable, and up-to-date with the latest best practices, this book is for you."},
		{
			ISBN:        "9798602477429",
			Title:       "You Don't Know JS Yet",
			Subtitle:    "Get Started",
			Author:      "Kyle Simpson",
			Published:   "2020-01-28T00:00:00.000Z",
			Publisher:   "Independently published",
			Pages:       143,
			Description: "The worldwide best selling You Don't Know JS book series is back for a 2nd edition: You Don't Know JS Yet. All 6 books are brand new, rewritten to cover all sides of JS for 2020 and beyond."},
		{
			ISBN:        "9781484200766",
			Title:       "Pro Git",
			Subtitle:    "Everything you neeed to know about Git",
			Author:      "Scott Chacon and Ben Straub",
			Published:   "2014-11-18T00:00:00.000Z",
			Publisher:   "Apress; 2nd edition",
			Pages:       458,
			Description: "Pro Git (Second Edition) is your fully-updated guide to Git and its usage in the modern world. Git has come a long way since it was first developed by Linus Torvalds for Linux kernel development. It has taken the open source world by storm since its inception in 2005, and this book teaches you how to use it like a pro."},
		{
			ISBN:        "9781484242216",
			Title:       "Rethinking Productivity in Software Engineering",
			Subtitle:    "",
			Author:      "Caitlin Sadowski, Thomas Zimmermann",
			Published:   "2019-05-11T00:00:00.000Z",
			Publisher:   "Apress",
			Pages:       310,
			Description: "Get the most out of this foundational reference and improve the productivity of your software teams. This open access book collects the wisdom of the 2017 \"Dagstuhl\" seminar on productivity in software engineering, a meeting of community leaders, who came together with the goal of rethinking traditional definitions and measures of productivity."},
	}

	// update book collections
	bookCollections = Books{Books: collections}

	// apply both the status and borrower for the book
	applyStatus()
}
