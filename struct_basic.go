package main

type Customer struct {
	name, address string
	age           int8
}

func struct_basic() Customer {
	var customer = Customer{
		name:    "wawan",
		address: "semarang",
		age:     22,
	}
	return customer
}

func struct_custom() []Customer {
	var customers = []Customer{
		{name: "test 1", address: "place 1", age: 1},
		{name: "test 2", address: "place 2", age: 2},
		{name: "test 3", address: "place 3", age: 3},
	}
	return customers[:]
}
