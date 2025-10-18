package main

func functionVariadic(numbers ...int8) int8 {
	var sum int8 = 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}
