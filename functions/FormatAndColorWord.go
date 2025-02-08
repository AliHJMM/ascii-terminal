package functions

import (
	"fmt"
	"os"
	"strings"
)

func FormatAndColorWord(lineIndex int, targetWord string, fileName string, colorCode string, highlightLetters string, alignment string, formattedResult string, tempResult string) (string, string) {
	count := 0
	for index, letter := range targetWord {
		currentLine := ReadLineFromFile(1+(int(letter)-32)*9+lineIndex, fileName)
		tempResult += currentLine
		
		if strings.ContainsRune(highlightLetters, letter) {
			if index < (1 + len(targetWord) - len(highlightLetters)) {
				if highlightLetters == targetWord[index:index+len(highlightLetters)] {
					count = len(highlightLetters)
				}
			}
			if count != 0 {
				colorEscape := Color(colorCode)
				if colorEscape != "" {
					currentLine = colorEscape + currentLine + Color("clearFormat")
				} else {
					fmt.Println(Errors("availableColor"))
					os.Exit(0)
				}
				count--
			}
		} else if highlightLetters == "   " {
			colorEscape := Color(colorCode)
			if colorEscape != "" {
				currentLine = colorEscape + currentLine + Color("clearFormat")
			} else {
				fmt.Println(Errors("availableColor"))
				os.Exit(0)
			}
		}
		formattedResult += currentLine
	}
	return formattedResult, tempResult
}
