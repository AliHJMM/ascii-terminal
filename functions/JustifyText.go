package functions

import (
	"fmt"
	"strings"
)

func JustifyText(words []string, filename string, textColor string, highlightedLetters string, alignment string, outputFile string) string {
	finalOutput := ""
	tempResult := ""
	currentResult := ""
	intermediateResult := ""
	space := ""
	var formattedLines []string
	var temporaryLine string

	for _, line := range words {
		wordsInLine := strings.Split(line, " ")
		wordsInLine = FilterEmptyArgs(wordsInLine)
		for i := 0; i < 8; i++ {

			for _, word := range wordsInLine {
				currentResult, intermediateResult = FormatAndColorWord(i, word, filename, textColor, highlightedLetters, alignment, currentResult, intermediateResult)
				if len(wordsInLine) == 1 {
					space = AlignText(intermediateResult, "left", wordsInLine)
					currentResult = space + currentResult
					fmt.Println(currentResult)
				}
				if i == 0 {
					temporaryLine = temporaryLine + intermediateResult
				}
				formattedLines = append(formattedLines, currentResult)
				currentResult = ""
				intermediateResult = ""
			}
			if outputFile != "" {
				_, tempResult = FormatAndColorWord(i, line, filename, textColor, highlightedLetters, alignment, currentResult, intermediateResult)
			}
			finalOutput = finalOutput + "\n" + tempResult
			if len(wordsInLine) > 1 {
				space = AlignText(temporaryLine, alignment, wordsInLine)
				resultWithSpace := strings.Join(formattedLines, space)
				fmt.Println(resultWithSpace)
			}
			formattedLines = nil
		}
		temporaryLine = ""
	}

	return finalOutput
}
