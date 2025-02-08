package functions

import (
	"fmt"
	"os"
)

func SaveContentToFile(data string, fileName string) error {
	outputFile, fileErr := os.Create(fileName)
	if fileErr != nil {
		return fileErr
	}
	defer outputFile.Close()

	data = data + "\n\n"

	_, writeErr := fmt.Fprint(outputFile, data)
	if writeErr != nil {
		return writeErr
	}

	return nil
}
