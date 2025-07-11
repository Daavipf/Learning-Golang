package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	var output strings.Builder

	for i := 0; i < repeatCount; i++ {
		output.WriteString(character)
	}

	return output.String()
}
