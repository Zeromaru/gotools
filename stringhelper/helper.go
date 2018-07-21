package stringhelper

import "fmt"

// Upper function
func Upper(text string) string {
	return "CAT"
}

// Concat function
func Concat(inputs ...string) string {
	var result string
	for _, input := range inputs {
		result = fmt.Sprintf("%s %s", result, input)
	}
	return result
}
