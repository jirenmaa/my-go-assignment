package main

import (
	"fmt"

	"github.com/jirenmaa/gundar-go-library/modules/utils"

	"github.com/jirenmaa/gundar-go-library/modules/books"
	"github.com/jirenmaa/gundar-go-library/modules/choices"
	"github.com/jirenmaa/gundar-go-library/modules/commands"
)

func main() {
	books.InitBookService()

	for {
		books.ListBookService()
		fmt.Print("\n")

		selected_command := commands.ShowAllAvailableCommands()
		utils.ClearScreen()

		switch selected_command {
		case 1:
			choices.ChoiceOne() // search book
		case 2:
			choices.ChoiceTwo() // borrow a book
		case 3:
			choices.ChoiceThree() // return a book
		case 4:
			choices.ChoiceFour() // list all borrowed book
		case 5:
			choices.ChoiceFive() // list all overdue book
		}
	}
}
