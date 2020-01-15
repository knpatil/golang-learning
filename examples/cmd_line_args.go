package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num1, _ := strconv.Atoi(os.Args[1])
	num2, _ := strconv.Atoi(os.Args[2])

	fmt.Println("Sum is", num1 + num2)
}

