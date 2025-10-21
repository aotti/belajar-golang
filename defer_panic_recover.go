package main

import "fmt"

/*
@funcName name of function
*/
func functionDefer(funcName string) {
	// recover == catch error,
	// must place in defer function
	var message = recover()
	if message != nil {
		fmt.Println("error message:", message)
	}
	fmt.Println(funcName, "over")
}

// function to reproduce panic / error
func functionPanic(status bool) string {
	defer functionDefer("func panic")
	if status {
		panic("function interrupted")
	}
	return "func panic success"
}
