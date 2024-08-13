package listStuff_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mcadenas-bjss/GoAcademy/assignments/listStuff"
)

func TestArray1To10(t *testing.T) {
	expected := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	actual := listStuff.Array1To10()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestSortIntArray(t *testing.T) {
	input := listStuff.Array1To10()

	t.Run("Sorts in desc order", func(t *testing.T) {
		actual := listStuff.SortIntArray(input, -1)
		expected := [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

		compareSlices(t, expected[:], actual[:])
	})

	t.Run("Sorts in asc order", func(t *testing.T) {
		actual := listStuff.SortIntArray(input, 1)
		expected := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		compareSlices(t, expected[:], actual[:])
	})
}

func ExampleSortIntArray() {
	input := listStuff.Array1To10()
	sorted := listStuff.SortIntArray(input, -1)
	fmt.Println(sorted)
	// Output: [10 9 8 7 6 5 4 3 2 1]
}

func TestSortEvenAscOddDesc(t *testing.T) {
	input := listStuff.Array1To10()

	even, odd := listStuff.SortEvenAscOddDesc(input[:])
	expectedEven := [5]int{2, 4, 6, 8, 10}
	expectedOdd := [5]int{9, 7, 5, 3, 1}

	compareSlices(t, expectedEven[:], even)
	compareSlices(t, expectedOdd[:], odd)
}

func ExampleSortEvenAscOddDesc() {
	input := listStuff.Array1To10()
	even, odd := listStuff.SortEvenAscOddDesc(input[:])
	fmt.Println(even)
	fmt.Println(odd)
	// Output:
	// [2 4 6 8 10]
	// [9 7 5 3 1]
}

func compareSlices(t testing.TB, expected, actual []int) {
	t.Helper()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
