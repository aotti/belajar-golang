package main

import (
	"fmt"
	"strings"
)

func main() {
	// variables
	nameBasic, ageBasic, hobbyBasic := variableBasic()
	fmt.Println("variable basic", nameBasic, ageBasic, hobbyBasic)
	nameBasic = "terlele"
	fmt.Println("variable basic", nameBasic, ageBasic, hobbyBasic)

	nameConst, ageConst := variableConst()
	fmt.Println("variable const", strings.Split(nameConst, ""), ageConst)
	nameConst = "terkadang"
	fmt.Println("variable const", nameConst, ageConst)
	fmt.Println()

	// array, slice, map
	fmt.Println("array basic", array_basic())
	fmt.Println("slice basic", slice_basic())
	fmt.Println("slice with map", slice_with_map())
	fmt.Println("slice with map", slice_with_map()[0])
	fmt.Println("map basic", map_basic())
	fmt.Println("map custom", map_custom())
	fmt.Println("map with slice", map_with_slice())
	fmt.Println("map with slice", map_with_slice()["name"])
	fmt.Println()

	// function variadic
	fmt.Println("func variadic", functionVariadic(1, 2, 3, 4, 5))
	// function with type
	var numbers = []int8{1, 2, 3, 4, 5, 1, 1}
	fmt.Println("func with type", functionWithType(filterNumbers, numbers))
	// function anonym
	var functionAnon = func() string {
		return "lelemao"
	}
	fmt.Println("func anonymous", functionAnon())
}
