// Package geom implements geometry objects
package geom

// The interface provides basic methods for geometry objects
type Geometry interface {
	// It calculates and returns area of a geometry object
	Area() float64

	// It calculates and returns perimeter of a geometry object
	Perimeter() float64
}
