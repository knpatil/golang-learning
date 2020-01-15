package main

import "fmt"

func main() {
	fmt.Print("Enter a number : ")
	var num float64
	fmt.Scan(&num)
	for num != 0 {
		fmt.Println("Square root is : ", sqrt(num))
		fmt.Print("Enter a number : ")
		fmt.Scan(&num)
	}
	fmt.Println("Finished.")
}

func sqrt(num float64) float64 {
	guess := 1.0
	for i := 1; i <= 25; i++ {
		fmt.Println("guess: ", guess)
		guess = guess - (guess*guess-num)/(2*guess)
	}
	return guess
}
