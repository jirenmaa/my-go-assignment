package choices

import (
	"bufio"
	"fmt"
	"os"
)

// getting input from user that used whitespaces
func getScanner(content string) string {
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
