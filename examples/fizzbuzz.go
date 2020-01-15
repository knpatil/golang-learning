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
		fmt.Println("fizzBuzz1 response : ", fizzBuzz(num))
		fmt.Println("fizzBuzz2 response : ", fizzBuzz2(num))
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

func fizzBuzz2(num int) string {
	result := ""
	if num%3 == 0 {
		result += "fizz"
	}
	if num%5 == 0 {
		result += "buzz"
	}
	if result == "" {
		result = strconv.Itoa(num)
	}
	return result
}
