package cli

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
)

// DisplayMenu shows the main menu with ASCII art
func DisplayMenu() {
	color.Green("Welcome to the Port Scanner")
	color.Cyan("1. 🚀 Start Scan")
	color.Cyan("2. ❌ Exit")
}

// GetUserChoice gets the user's choice from the menu
func GetUserChoice() string {
	reader := bufio.NewReader(os.Stdin)
	color.Cyan("Enter your choice: ")
	choice, _ := reader.ReadString('\n')
	return strings.TrimSpace(choice)
}
