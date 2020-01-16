package main

import "fmt"

func product(nums ...int) {
	prod := 1
	for _, num := range nums {
		prod *= num
	}
	fmt.Println(nums, prod)
}

func main() {
	// Variadic functions can be called in the usual way with individual arguments
	product(3, 4)
	product(2, 3, 4, 5, 6)
	product()
}
