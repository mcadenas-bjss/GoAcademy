package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// StringPrompt asks for a string value using the label
func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

// StringPrompt asks for a string value using the label
func IntPrompt(label string) int {
	var s string

	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')

		if s == "\n" {
			fmt.Println("Defaulting to 0")
			return 0
		}

		if s != "" {
			break
		}
	}

	n, err := strconv.Atoi(strings.TrimSpace(s))

	if err != nil {
		fmt.Println("Must be a number.")
		return IntPrompt(label)
	}

	return n
}
