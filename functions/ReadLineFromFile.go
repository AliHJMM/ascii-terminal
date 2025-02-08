package functions

import (
	"bufio"
	"fmt"
	"os"
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
		if currentLine == lineIndex {
			result = scanner.Text()
			break
		}
		currentLine++
	}
	return result
}
