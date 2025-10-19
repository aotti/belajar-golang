package main

func array_basic() [5]int8 {
	arr := [5]int8{1, 2, 3, 4}
	return arr
}

func slice_basic() []int8 {
	slice := []int8{9, 8, 7, 6, 5}
	return slice
}

func slice_with_map() []interface{} {
	arr := []interface{}{
		map[string]string{
			"name": "wawan",
		},
		map[string]string{
			"name": "dengkul",
		},
	}
	return arr
}

func map_basic() map[string]string {
	mp := map[string]string{
		"name": "wawan",
		"age":  "22",
	}
	return mp
}

func map_custom() map[string]interface{} {
	mp := map[string]interface{}{
		"name": "wawan",
		"age":  22,
	}
	return mp
}

func map_with_slice() map[string][]string {
	mp := map[string][]string{
		"name": {"wawan", "lele"},
		"age":  {"22", "81"},
	}
	return mp
}
