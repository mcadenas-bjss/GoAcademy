package main

import (
	"flag"
	"fmt"
)

// This uses the package workingWithStrings and the function ConcatenateStrings.
func main() {
	var first, middle, last string

	// Get the command line arguments
	flag.StringVar(&first, "first", "", "First name")
	flag.StringVar(&middle, "middle", "", "Middle name")
	flag.StringVar(&last, "last", "", "Last name")

	flag.Parse()

	// Print the result
	fmt.Printf("%s %s %s", first, middle, last)
}
