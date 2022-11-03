package commands

import (
	"bufio"
	"fmt"
	"os"
)

func getScanner(content string) string {
	// getting input from user that used whitespaces

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(content)

	if scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

func InputTemplate(content string) string {
	command := getScanner(content)
	return command
}
