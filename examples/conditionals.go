package main

import (
	"fmt"
	"time"
)
import "math/rand"

func someNumber() int32 {
	_, _, sec := time.Now().Clock() // get number of seconds after the hour
	rand.Seed(int64(sec))           // set random seed

	return rand.Int31n(30) - 15
}

func main() {
	if num := someNumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
