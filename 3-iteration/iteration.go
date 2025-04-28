package iteration

import "strings"

// Repeat takes a character and an integer (repeatCount) and return the repeated character repeatCount times
func Repeat(character string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
