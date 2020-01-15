package main

import (
	"fmt"
	"server"
)

func main() {
	num := 10

	server.Doubler1(&num)
	fmt.Println("Doubled value : ", num)

	server.Doubler(&num)
	fmt.Println("Using double helper : ", num)
}
