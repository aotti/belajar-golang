package main

func type_checker(value interface{}) interface{} {
	switch value.(type) {
	case string:
		return "tipe value adalah string"
	case int:
		return "tipe value adalah integer"
	default:
		return "tipe value tidak tau"
	}
}
