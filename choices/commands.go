package choices1

import (
	"fmt"
	"strconv"
)

func DisplayBookShelfMenu() int {
	var command string

	fmt.Println("=================================")
	fmt.Println("List of available commands :")
	fmt.Println()
	fmt.Println(" [1] 🔍 Search a Book by Name")
	fmt.Println(" [2] 📚 Borrow a Book")
	fmt.Println(" [3] 📘 Return Borrowed Book")
	fmt.Println(" [4] 🎒 Add New Book")
	fmt.Println(" [5] 📙 Show All Borrowed Book")
	fmt.Println(" [6] 📕 Show All Overdue Book")
	fmt.Println()
	fmt.Println("=================================")

	fmt.Print("\n$ ")
	fmt.Scanln(&command)

	value, err := strconv.Atoi(command)
	if err != nil {
		panic(err)
	}

	return value
}
