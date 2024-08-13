package main

import (
	"flag"
	"fmt"

	"github.com/mcadenas-bjss/GoAcademy/utils"
)

// This uses the package workingWithStrings and the function ConcatenateStrings.
func main() {
	var first, middle, last string

	// Get the command line arguments
	flag.StringVar(&first, "first", "", "First name")
	flag.StringVar(&middle, "middle", "", "Middle name")
	flag.StringVar(&last, "last", "", "Last name")

	flag.Parse()

	if len(flag.Args()) < 1 {
		first = utils.StringPrompt("Enter your first name:")
		middle = utils.StringPrompt("Enter your middle name:")
		last = utils.StringPrompt("Enter your last name:")

	}

	// Print the result
	fmt.Printf("%s %s %s", first, middle, last)
}
