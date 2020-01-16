package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

// ----------- rectangle object
type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

/// circle
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// not method
func measure(g geometry) {
	fmt.Println("Area: ", g.area())
	fmt.Println(g, g.area(), g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	// The circle and rect struct types both implement
	// the geometry interface so we can use instances of
	// of these structs as arguments to measure()
	measure(r)
	measure(c)
}
