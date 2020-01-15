package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Enter a number : ")
	var num int
	fmt.Scan(&num)
	for num != 0 {
		fmt.Println(fizzBuzz(num))
		fmt.Print("Enter a number : ")
		fmt.Scan(&num)
	}
	fmt.Println("Finished.")
}

func fizzBuzz(num int) string {
	switch {
	case num%3 == 0 && num%5 == 0:
		return "fizzbuzz"
	case num%3 == 0:
		return "fizz"
	case num%5 == 0:
		return "buzz"
	default:
		return strconv.Itoa(num)
	}
}
