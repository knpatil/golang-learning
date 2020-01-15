package server

import (
	"fmt"
	"math"
)

func Hello() {
	fmt.Println("Hello, I am running ...")
}

func CircleInfo(radius float64) (float64, float64) {
	return math.Pi * radius * radius, 2 * math.Pi * radius
}

func Doubler1(x *int)  {
	*x = *x * 2 // *x *= 2
}

func doublerHelper(num *int) {
	*num *= 2
}

func Doubler(x *int) {
	doublerHelper(x)
}
