package geom

import "testing"

func TestRectangle(t *testing.T) {

	rect := Rectangle{5, 10}

	if rect.Area() != 50 {
		t.Error("Wrong area of rectangle")
	}

	if rect.Perimeter() != 30 {
		t.Error("Wrong perimeter of rectangle")
	}

}
