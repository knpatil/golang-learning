package main

import "fmt"

/**

typed collections of fields
useful for grouping data together (like structs in other languages)

you can name the fields when initializing a struct
you can leave out fields, which will be set to zero
& operator yields a pointer to the struct
use a dot to access the fields

*/

type mountain struct {
	name      string
	elevation int
}

/**
Anonymous struct
why an anonymous struct?
e.g., grouped globals
*/

var config struct {
	APIKey string
	IPAddr string
}

func main() {

	k2 := mountain{"K2", 8611} // create a new struct
	fmt.Println(k2)
	fmt.Printf("%#v\n", k2)

	everest := mountain{elevation: 8848, name: "Mt. Everest"}
	fmt.Printf("%#v\n", everest)
	fmt.Println(mountain{name: "Flat"})
	fmt.Println(&mountain{name: "Long's Peak", elevation: 4346})

	nd := mountain{name: "Nanda Devi", elevation: 7815}
	fmt.Println(nd.name)

	sp := &nd
	fmt.Println(sp.elevation)

	sp.elevation++
	fmt.Println(sp.elevation)

	config.APIKey = "BADC0C0A"
	config.IPAddr = "1.1.1.1"

	fmt.Println("Config: ", config)

}
