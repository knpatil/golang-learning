package main

import "fmt"

type Mountain struct {
	Name      string
	Elevation int
}

type Climb struct {
	Mountain
	Climber string
}

func (m Mountain) HowBig() string {
	if m.Elevation > 8800 {
		return "HUGE"
	}
	return "NORMAL"
}

func main() {

	c1 := Climb{}
	fmt.Println(c1)

	c1.Name = "K2"
	c1.Elevation = 8111 // i.e., missed the summit by 500 m
	c1.Climber = "Arjun Climber"

	fmt.Println(c1)

	c2 := Climb{
		Mountain{"K2", 8211},
		"Nakul Climber",
	}
	fmt.Println(c2)

	everest := Mountain{Elevation: 8848, Name: "Mt. Everest"}
	fmt.Println(everest.HowBig())
}
