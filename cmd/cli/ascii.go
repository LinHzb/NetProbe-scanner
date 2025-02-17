package cli

import (
	"io/ioutil"
	"log"

	"github.com/fatih/color"
)

// DisplayASCIIArt reads and displays ASCII art from a file
func DisplayASCIIArt(filepath string) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Failed to read ASCII art file: %v", err)
	}
	color.Magenta(string(data))
}
