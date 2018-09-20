package main

import (
	"fmt"
	"workshop/interface1/geom"
)

func printGeometry(g geom.Geometry) {
	fmt.Println(g)
	fmt.Printf("\tArea: %.3f\n", g.Area())
	fmt.Printf("\tPerimeter: %.3f\n", g.Perimeter())
}

func main() {

	printGeometry(geom.Rectangle{10, 5})
	printGeometry(geom.Circle{5})

}
