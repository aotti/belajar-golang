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
	fmt.Println()

	// defer panic recover
	fmt.Println("1 defer, panic, recover", functionPanic(true))
	fmt.Println("2 defer, panic, recover", functionPanic(false))
	// error package 1
	var data, err = error_test(3)
	if err == nil {
		fmt.Println("error package 1", data)
	}
	fmt.Println("error package 1", err)
	// error package 2
	data, err = error_test(0)
	if err == nil {
		fmt.Println("error package 2", data)
	}
	fmt.Println("error package 2", err)
	fmt.Println()

	// struct
	fmt.Println("struct basic", struct_basic())
	fmt.Println("struct custom", struct_custom())
	fmt.Println("struct custom", struct_custom()[0].name, struct_custom()[1].address)
	var player = Player{
		username:   "awiwi",
		password:   "123",
		game_count: 5,
	}
	fmt.Println("struct method", player.struct_method())

	// type assertion
	fmt.Println("type assertion", type_checker(1))
	fmt.Println("type assertion", type_checker('9'))
	fmt.Println("type assertion", type_checker("test"))
	fmt.Println()

	// pointer
	fmt.Println("pointer basic", pointer_basic())
}
