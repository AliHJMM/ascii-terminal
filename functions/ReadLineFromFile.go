package functions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadLineFromFile(lineIndex int, filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(Errors("banner"))
		os.Exit(0)
	}
	defer file.Close()

	var result string
	scanner := bufio.NewScanner(file)
	currentLine := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Remove any trailing carriage returns
		line = strings.TrimRight(line, "\r")
		if currentLine == lineIndex {
			result = line
			break
		}
		currentLine++
	}
	
	
	return result
}
