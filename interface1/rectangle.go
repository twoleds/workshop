package main

import "fmt"

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return (r.width + r.height) * 2
}

func (r Rectangle) String() string {
	return fmt.Sprintf(
		"RECTANGLE [width=%.3f, height=%.3f]",
		r.width,
		r.height,
	)
}
