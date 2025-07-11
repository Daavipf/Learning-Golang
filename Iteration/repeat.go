package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	var output strings.Builder

	for i := 0; i < repeatCount; i++ {
		output.WriteString(character)
	}

	return output.String()
}
