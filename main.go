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

	// array, slice, map
	fmt.Println("array", array_basic())
	fmt.Println("slice", slice_basic())
	fmt.Println("map basic", map_basic())
	fmt.Println("map custom", map_custom())

	// function variadic
	fmt.Println("func variadic", functionVariadic(1, 2, 3, 4, 5))
	// function with type
	var numbers = []int8{1, 2, 3, 4, 5, 1, 1}
	fmt.Println("func with type", functionWithType(filterNumbers, numbers))
}
