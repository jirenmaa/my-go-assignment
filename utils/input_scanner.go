package utils

import (
	"bufio"
	"fmt"
	"os"
)

// getting input from user that used whitespaces
func GetScanner(content string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(content)

	if scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

func InputTemplate(content string) string {
	command := GetScanner(content)
	return command
}
