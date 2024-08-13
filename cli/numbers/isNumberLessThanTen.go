package main

import (
	"flag"

	"github.com/mcadenas-bjss/GoAcademy/assignments/numberStuff"
	"github.com/mcadenas-bjss/GoAcademy/utils"
)

func main() {
	var n int

	flag.IntVar(&n, "number", 0, "a number")
	flag.Parse()

	if flag.NArg() == 0 {
		n = utils.IntPrompt("Enter a number")
	}

	if numberStuff.IsIntBetweenOneAndTen(n) {
		println("Yes")
	} else {
		println("No")
	}
}
