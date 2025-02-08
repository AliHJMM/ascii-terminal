package functions

import (
	"fmt"
	"os"
	"strings"
)

func AlignText(text string, alignment string, arguments []string) string {
	numSpaces := len(arguments) - 1
	terminalWidth := GetConsoleWidth()

	if len(text) > terminalWidth {
		fmt.Println(Errors("fit"))
		os.Exit(0)
	}

	if alignment != "left" && alignment != "right" && alignment != "center" && alignment != "justify" {
		fmt.Println(Errors("align"))
		os.Exit(0)
	}

	var result strings.Builder

	switch alignment {
	case "left":
		text = ""
	case "right":
		text = strings.Repeat(" ", terminalWidth-len(text)-1)
	case "center":
		text = strings.Repeat(" ", (terminalWidth-len(text))/2)
	case "justify":
		text = strings.Repeat(" ", (terminalWidth-len(text))/numSpaces)
	}

	result.WriteString(text)

	return result.String()
}
