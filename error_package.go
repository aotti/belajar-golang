package main

import "errors"

func error_test(number int8) (int8, error) {
	if number == 0 {
		return 0, errors.New("angka tidak boleh 0")
	} else {
		return number, nil
	}
}
