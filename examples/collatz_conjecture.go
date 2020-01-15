package main

import "fmt"

func main() {
	fmt.Print("Enter a number : ")
	var num int
	fmt.Scan(&num)
	for num != 0 {
		collatzConjecture(num)
		fmt.Print("\nEnter a number : ")
		fmt.Scan(&num)
	}
	fmt.Println("Finished.")
}

func collatzConjecture(num int) {
	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			num *= 3
			num += 1
		}
		fmt.Print(num, " ")
	}
}
