package main

import "fmt"

func main() {
	s := f1()
	f2(s)
	slicing_slices()
}

func slicing_slices() {
	slice := make([]string, 0)
	slice = append(slice, "apple", "lemon", "mango", "banana", "cherry", "lime")
	fmt.Println(slice)
	// Slices support a "slice" operator with the syntax slice[low:high]
	slice1 := slice[2:5]
	fmt.Println("slice1[2:5]", slice1)
	// This slices up to (but excluding) `s[5]`.
	slice2 := slice[:5]
	fmt.Println("slice2[:5]", slice2)
	// And this slices up from (and including) `s[2]`.
	slice3 := slice[2:]
	fmt.Println("slice3[2:]", slice3)
	// We can declare and initialize a variable for slice
	// in a single line as well.
	t := []string{"raspberry", "orange", "guava"}
	fmt.Println("declared slice:", t)
}

func f1() []string {
	s := make([]string, 3)
	fmt.Println("initial slice:", s)

	s[0] = "zero" // set just like an array
	s[1] = "one"
	s[2] = "two"

	fmt.Println("now:", s)

	fmt.Println("get:", s[2]) // get like an array

	fmt.Println("len of slice:", len(s))

	return s
}

func f2(s []string) {
	s = append(s, "three")
	s = append(s, "four", "five")
	fmt.Println("after append:", s)

	copy_of_s := make([]string, len(s))
	copy(copy_of_s, s)
	fmt.Println("copy:", copy_of_s)
}
