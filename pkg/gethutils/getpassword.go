package gethutils

import (
	"fmt"
	"syscall"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

// GetPassword is a utility function for retrieving a key password via the terminal
func GetPassword() string {
	color.Red("\nPlease enter your key password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)
	fmt.Println()
	return password
}
