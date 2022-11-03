package utils

import (
	"os"
	"os/exec"
)

// CLEAR THE PREVIEWS TEXT
func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
