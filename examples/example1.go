package main

import (
	"fmt"
	"server"
)

func main() {
	server.Hello()

	area, circumference := server.CircleInfo(105);

	fmt.Println(server.CircleInfo(23))
	fmt.Printf("Area: %.2f\nCircumference: %.4f", area, circumference)
}


