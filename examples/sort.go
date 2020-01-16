package main

import (
	"fmt"
	"sort"
)

func main() {

	// sorting strings
	strs := []string{"pear", "apple", "lemon"}
	sort.Strings(strs)
	fmt.Println("sorted:", strs)

	// Sorting integers
	ints := []int{7, 2, -4}
	sort.Ints(ints)
	fmt.Println("sorted:", ints)

	// Check if a slice is already sorted
	s := sort.StringsAreSorted(strs)
	fmt.Println("Are strings sorted?", s)

}
