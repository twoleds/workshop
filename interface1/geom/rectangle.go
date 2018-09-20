package geom

import "fmt"

// Rectangle implements `Geometry` interface.
type Rectangle struct {
	Width, Height float64
}

// It calculates and returns area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func (r Rectangle) String() string {
	return fmt.Sprintf(
		"RECTANGLE [Width=%.3f, Height=%.3f]",
		r.Width,
		r.Height,
	)
}
