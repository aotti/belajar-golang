package main

type Filter func([]int8) []int8

func filterNumbers(numbers []int8) []int8 {
	var filteredNumbers = make([]int8, len(numbers))
	// filter number 1
	for i, v := range numbers {
		if v != 1 {
			filteredNumbers[i] = v
		}
	}
	return filteredNumbers[:]
}

func functionWithType(filterNumbers Filter, numbers []int8) int8 {
	// filter numbers before sum
	var filteredNumbers = filterNumbers(numbers)

	// sum numbers
	var sum int8 = 0
	for _, number := range filteredNumbers {
		sum += number
	}
	return sum
}
