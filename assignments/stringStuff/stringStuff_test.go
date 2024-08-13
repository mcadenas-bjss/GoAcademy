package stringStuff_test

import (
	"fmt"
	"testing"

	"github.com/mcadenas-bjss/GoAcademy/assignments/stringStuff"
)

func TestConcatenateStrings(t *testing.T) {
	t.Run("returns one string from all arguments", func(t *testing.T) {
		actual := stringStuff.ConcatenateStrings("fizz", "buzz")
		expected := "fizzbuzz"

		expectStringEquals(t, actual, expected)
	})

	t.Run("returns empty string when no arguments", func(t *testing.T) {
		actual := stringStuff.ConcatenateStrings()
		expected := ""

		expectStringEquals(t, actual, expected)
	})
}

func expectStringEquals(t testing.TB, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func ExampleConcatenateStrings() {
	result := stringStuff.ConcatenateStrings("fizz", "buzz")
	fmt.Println(result)
	// Output: fizzbuzz
}
