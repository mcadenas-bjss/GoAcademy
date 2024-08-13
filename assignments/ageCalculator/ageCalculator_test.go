package ageCalculator_test

import (
	"fmt"
	"testing"

	"github.com/mcadenas-bjss/GoAcademy/assignments/ageCalculator"
)

func TestAgeCalculator(t *testing.T) {
	expected := 33
	actual := ageCalculator.AgeCalculator("1991-03-06")

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func ExampleAgeCalculator() {
	age := ageCalculator.AgeCalculator("1991-03-06")
	fmt.Println(age)
	// Output: 33
}
