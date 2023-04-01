package main

import "fmt"

// purpose of this to learn - go
// completed - struct, slice, map ,interface, pointer in go
// to run this use below command
// go run main.go deck.go structdemo.go interface.go

func main() {
	// Go is a static type lang, cant change variable type later
	cards, _ := newDeckFromFile("mycards.txt")
	cards.shuffle()
	// same package
	cards.print()

	// structs
	e := employee{name: "Devendra", company: "CleverTap"}
	fmt.Println(e)
	ep := &e
	ep.changeCompany("-- ")
	fmt.Println(e)

	e.changeCompany("google")
	fmt.Println(e)

	// map
	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
		"black": "#eeeeee",
	}

	// Iterate over map
	for k, v := range colors {
		fmt.Println(k, v)
	}

	//fmt.Println(colors)
	// delete from map
	// delete(colors, "red")

	fmt.Println(colors)

	//  ways to create map
	//  var emptyMap map[string]string
	//  fmt.Println(emptyMap)

	//  emptyMap2 := make(map[string]string)
	//  fmt.Println(emptyMap2)

	//interfaces

	gp := goodprinter{}
	print(gp)

	bp := badprinter{}
    print(bp)
}
