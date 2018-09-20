package geom

import (
	"math"
	"testing"
)

func TestCircle(t *testing.T) {

	circle := Circle{5}

	if circle.Area() != (math.Pi * 25) {
		t.Error("Wrong area of circle")
	}

	if circle.Perimeter() != (10 * math.Pi) {
		t.Error("Wrong perimeter of circle")
	}

}
