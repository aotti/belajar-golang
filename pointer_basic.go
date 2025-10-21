package main

import "fmt"

type Address struct {
	city, province string
}

func change_data(address *Address) {
	address.city = "Tangerang"
}

func (address *Address) change_data_struct() {
	address.city = "Denpasar"
	address.province = "Bali"
}

func pointer_basic() string {
	var address1 = Address{city: "Jakarta", province: "DKI Jakarta"}
	// set value by copying address1 data
	var address2 = address1
	address2.city = "Bogor"
	address2.province = "Jawa Barat"
	// set value by referencing to var address1, not copying the data
	var address3 = &address1
	// set value by referencing to the data itself, not variable
	*address3 = Address{city: "Paris", province: "France"}
	// create empty pointer, not referencing to any data
	var address4 = new(Address)
	// change data from function
	change_data(&address2)
	// change data with struct function
	address4.change_data_struct()

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)
	return ""
}
