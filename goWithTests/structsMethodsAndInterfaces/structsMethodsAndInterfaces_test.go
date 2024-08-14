package structsmethodsandinterfaces

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func ExamplePerimeter() {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	fmt.Println(Perimeter(rectangle))
	// Output: 40
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{Width: 12.0, Height: 6.0}, 72.0},
		{Circle{Radius: 10.0}, 314.1592653589793},
		{Triangle{Base: 12.0, Height: 6.0}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}
}

func ExampleArea() {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	fmt.Println(rectangle.Area())
	// Output: 100
}
