package main

import (
	"fmt"
	"math"
)

func main() {
	var name, age = "Ranveer Singh", 33
	fmt.Printf("%s\n", "Hello, world!")
	fmt.Printf("%s is %d years old\n", name, age)
	fmt.Print("Enter an integer: ")

	var num int
	fmt.Scan(&num) // read from keyboard–won't work in the playground
	fmt.Printf("The square root of %d is %f\n", num, math.Sqrt(float64(num)))

	var f float64 = 38730204.3832
	fmt.Printf("%v\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%.2f\n", f)
	var i int = 13
	fmt.Printf("|%5d|\n", i)
	fmt.Printf("|%-5d|\n", i)
	var b bool = false
	fmt.Printf("%t\n", b)
	fmt.Printf("%T %T %T\n", f, i, b)

}
