package workingWithInts_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/mcadenas-bjss/GoAcademy/assignments/workingWithInts"
)

func TestIsIntBetweenOneAndTen(t *testing.T) {
	for _, n := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11} {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			actual := workingWithInts.IsIntBetweenOneAndTen(n)
			expected := (n >= 1 && n <= 10)

			if actual != expected {
				t.Errorf("Expected %t, got %t", expected, actual)
			}
		})
	}
}

func ExampleIsIntBetweenOneAndTen() {
	output := workingWithInts.IsIntBetweenOneAndTen(5)
	fmt.Println(output)
	// Output: true
}
