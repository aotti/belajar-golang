package helper

import (
	"fmt"
	"regexp"
)

func Filter_String(str string) {
	var pattern = `foo.*`
	var matched1, err1 = regexp.MatchString(pattern, str)
	var matched2 = regexp.MustCompile(pattern).Find([]byte(str))

	fmt.Println("filter string", matched1, err1)
	fmt.Printf("filter string %q", matched2)
}
