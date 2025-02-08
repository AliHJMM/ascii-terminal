package functions

import (
	"strings"
)

func FormatAlphabetString(input string) string {
	formattedString := strings.ReplaceAll(input, "\\t", "    ")
	return formattedString
}
