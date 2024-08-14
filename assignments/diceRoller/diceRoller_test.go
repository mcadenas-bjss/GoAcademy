package diceRoller_test

import (
	"reflect"
	"testing"

	"github.com/mcadenas-bjss/GoAcademy/assignments/diceRoller"
)

func TestRoll(t *testing.T) {
	actual := diceRoller.Roll(6)
	if actual < 1 || actual > 6 {
		t.Errorf("Expected a number between 1 and 6, but got %d", actual)
	}
}

func TestRoll2d6X50(t *testing.T) {
	actual := diceRoller.Roll2d6X50()
	expected := []int{1}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
