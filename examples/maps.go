package main

import "fmt"

func main() {
	continents := map[string]string{
		"United States": "North America",
		"India":         "Asia",
		"Namibia":       "Africa",
		"Germany":       "Europe",
		"Mexico":        "North America",
		"Ireland":       "Europe",
		"Iceland":       "Europe",
	}
	c := "Iceland"
	fmt.Println(c, "is in", continents[c])
	fmt.Printf("%#v\n", continents)
	c = "Freedonia"
	fmt.Println(c, "is in", continents[c])

	fmt.Println("-------------------------------------------")

	// To create an empty map, use make
	m := make(map[string]int)
	// Set key/value pairs
	m["key1"] = 14
	m["key2"] = -2
	m["key3"] = 0
	fmt.Println(m)
	fmt.Println("len:", len(m))
	delete(m, "key2")
	fmt.Println(m)

	// Optional second return value indicates if key was in map.
	// Used to disambiguate between missing keys and zero values.
	_, present := m["key2"]
	fmt.Println("key2 present?", present)
	fmt.Println("-------------------------------------------")
}
