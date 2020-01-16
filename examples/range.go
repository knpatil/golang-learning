package main

import "fmt"

func main() {
	nums := []int{5, 13, 27, 19, -1}
	sum := 0
	for index, num := range nums {
		sum += num
		fmt.Println("Added element", index)
	}
	fmt.Println("sum of", nums, "is", sum)

	fmt.Println("------------")

	nums = []int{5, 13, 27, 19, -1}
	sum = 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum of", nums, "is", sum)

	fmt.Println("------------")
	// range over maps

	continents := map[string]string{
		"United States": "North America",
		"India":         "Asia",
		"Namibia":       "Africa",
		"Germany":       "Europe",
		"Mexico":        "North America",
		"Ireland":       "Europe",
		"Iceland":       "Europe",
	}
	for country, continent := range continents {
		fmt.Println(country, "is in", continent)
	}

	fmt.Println("------------")
	// just keys
	for country := range continents {
		fmt.Println(country)
	}
	fmt.Println("------------")

	// range on a string
	for index, char := range "Golang" {
		// print each letter as a rune (int32) and character
		fmt.Printf("%d %d %c\n", index, char, char)
	}
}
