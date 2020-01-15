package main

import "fmt"

func main() {
	fmt.Println("Data Structures in Go Demo: ")
	fmt.Println("-------------------------------")

	oneDArrays()
	fmt.Println("-------------------------------")

	initializeArrays()
	fmt.Println("-------------------------------")

	twoD()
	fmt.Println("-------------------------------")

}

func initializeArrays() {
	// ellipsis (...) can be used in place of size
	// if array is being intialized
	nums := [...]float32{1.4, 2.67, -3.56}
	fmt.Println("len(nums) =", len(nums))
	chars := [...]rune{'G', 'o', 'l', 'a', 'n', 'g', '€'}
	fmt.Printf("%q\n", chars) // %q == quoted
	fmt.Println(chars)        // values are actually int32
}

func oneDArrays() {
	var a [10]int // array of 10 integers, initialized to 0 by default
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		a[i] = i + 101
	}
	fmt.Println(a)
	fmt.Println("a[3] =", a[3]) // array indexing
	// Arrays can be initialized when they are declared
	cities := [5]string{"Hyderabad", "Delhi", "Pune", "Kochi", "Chennai"}
	fmt.Println("Cities:", cities)
}
func twoD() {
	var twoD [3][5]int // 2-dimensional array
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, twoD[i])
	}
	fmt.Println("\nAll at once:", twoD)
}
