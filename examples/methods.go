package main

import (
	"fmt"
	"math"
)

/**
Go is not object oriented...or is it?
Go does not have classes, but we can define methods on struct types
types and methods allow for an OO style of progamming
*/

type triangle struct {
	side1, side2, side3 int
}

// compute area of triangle using Heron's formula
func (t triangle) area() int {
	s := (t.side1 + t.side2 + t.side3) / 2
	return int(math.Sqrt(float64(s * (s - t.side1) * (s - t.side2) * (s - t.side3))))
}

// compute perimeter of triangle
func (t triangle) perimeter() int {
	return t.side1 + t.side2 + t.side3
}

/**
we generally want all of our methods to either accept a pointer or a type,
so the consumer of our code is not confused (when to pass a pointer vs. when to pass a value)
*/

// scale the sides of a triangle
func (t *triangle) scale(factor int) {
	t.side1 *= factor
	t.side2 *= factor
	t.side3 *= factor
}

func main() {

	t := triangle{3, 4, 5}
	fmt.Println("area: ", t.area(), "\nperim:", t.perimeter())

	t.scale(5)
	fmt.Println("sides: ", t.side1, t.side2, t.side3)
}
