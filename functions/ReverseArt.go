package functions

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Store the reversed output
var ReversedResult string

func ReverseArt() {
	if len(os.Args) != 2 || len(os.Args[1]) < 9 || os.Args[1][0:10] != "--reverse=" {
		fmt.Println(Errors("reverseUsage"))
		os.Exit(0)
	}

	artFile := os.Args[1][10:]
	file, err := os.Open(artFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	if ProcessFile(file) {
		ReversedResult = ""

		artContent, err := os.ReadFile(artFile)
		if err != nil {
			fmt.Println(Errors("textFileError"))
			os.Exit(0)
		}
		artData := string(artContent)
		asciiArt := strings.Split(artData, "\n")

		// Detect banner file
		BannerFileName := "standard.txt"
		fContent, err := os.ReadFile(BannerFileName)
		if err != nil {
			fmt.Println(Errors("fileContentError"))
			os.Exit(0)
		}
		BannerLines := strings.Split(string(fContent), "\n")

		// Process ASCII Art
		c := 0
		if len(asciiArt) > 9 {
			for i := 0; i < len(asciiArt)-1; {
				if len(asciiArt[i]) > 0 {
					c = i + 8
					if len(BannerLines)-c < 8 && len(BannerLines)-c != 0 {
						os.Exit(0)
					}
					reverseArt(BannerLines, asciiArt[i:c], 0, 0, 1)
					i = i + 8
				} else {
					i++
				}
			}
		} else {
			reverseArt(BannerLines, asciiArt, 0, 0, 1)
		}

		// Print only the reversed text
		fmt.Println(ReversedResult)
	}

	os.Exit(0)
}

var idx = 0

func reverseArt(BannerLines []string, asciiArt []string, SymbolFound int, ArtLine int, BannerLine int) {
	if SymbolFound == len(asciiArt[ArtLine]) {
		return
	}

	bannerWidth := len(BannerLines[BannerLine])

	if SymbolFound+bannerWidth <= len(asciiArt[ArtLine]) {
		if ArtLine < 7 {
			if BannerLines[BannerLine+ArtLine] == asciiArt[ArtLine][SymbolFound:SymbolFound+bannerWidth] {
				reverseArt(BannerLines, asciiArt, SymbolFound, ArtLine+1, BannerLine)
			} else {
				idx = BannerLine + 9
				if len(BannerLines)-1-idx < 9 && len(BannerLines)-1-idx != 0 {
					os.Exit(0)
				}
				reverseArt(BannerLines, asciiArt, SymbolFound, 0, idx)
			}
		} else {
			r := ((BannerLine - 1) / 9) + 32
			ReversedResult += string(r)
			reverseArt(BannerLines, asciiArt, SymbolFound+bannerWidth, 0, 1)
		}
	} else {
		idx = BannerLine + 9
		if len(BannerLines)-1-idx < 9 && len(BannerLines)-1-idx != 0 {
			os.Exit(0)
		}
		reverseArt(BannerLines, asciiArt, SymbolFound, 0, idx)
	}
}

func ProcessFile(file io.Reader) bool {
	scanner := bufio.NewScanner(file)
	lineCount, letterWidth := 0, 0
	counter := 0

	for scanner.Scan() {
		counter++
		line := scanner.Text()
		if lineCount == 0 {
			letterWidth = len(line)
		}
		if line == "" && lineCount == 0 {
			continue
		} else if line == "" && lineCount != 0 {
			fmt.Printf("Error: Empty line at line %d\n", counter)
		}
		if len(line) != letterWidth {
			fmt.Printf("Error: Invalid line length at line %d\n", counter)
			os.Exit(0)
		}

		lineCount++
		if lineCount == 8 {
			lineCount = 0
		}
	}
	if lineCount != 0 {
		fmt.Printf("Error: Incomplete input at line %d\n", counter)
		os.Exit(0)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return true
}
