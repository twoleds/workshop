package main

import (
	"fmt"
)

type Geometry interface {
	Area() float64
	Perimeter() float64
}

func printGeometry(g Geometry) {
	fmt.Println(g)
	fmt.Printf("\tArea: %.3f\n", g.Area())
	fmt.Printf("\tPerimeter: %.3f\n", g.Perimeter())
}

func main() {

	printGeometry(Rectangle{10, 5})
	printGeometry(Circle{5})

}
