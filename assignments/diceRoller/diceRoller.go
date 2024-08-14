package diceRoller

import (
	"fmt"
	"math/rand/v2"
	"sort"
)

func Roll(d int) int {
	n := rand.IntN(d)
	if n == 0 {
		n = 1
	}

	return n
}

func Roll2d6X50() []int {
	rolls := make([]int, 0)

	for i := 0; i < 50; i++ {
		roll := (Roll(6) + Roll(6))
		rolls = append(rolls, roll)
	}

	sort.Slice(rolls, func(i, j int) bool {
		return rolls[i]-rolls[j] < 0
	})

	fmt.Println(rolls)

	return rolls
}
