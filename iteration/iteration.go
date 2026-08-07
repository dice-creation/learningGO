package iteration

import "strings"

// func Repeat(character string) string {
// 	var repeated string
// 	for i := 0; i < 5; i++ {
// 		repeated += character
// 	}

// 	return repeated
// }

// significantly faster 128.7 ns/op --> 24.39 ns/op

func Repeat(character string, repeatCount int) string {
	if(repeatCount == 0){
		return ""
	}
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}