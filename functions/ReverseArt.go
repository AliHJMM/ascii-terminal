package functions

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

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
		fmt.Println("Art file:", artFile)

		artContent, err := os.ReadFile(artFile)
		if err != nil {
			fmt.Println(Errors("textFileError"))
			os.Exit(0)
		}
		artData := string(artContent)
		artDataTest := strings.TrimSpace(artData)
		asciiArt := strings.Split(artData, "\n")

		BannerFileName := "standard.txt"
		for _, symbol := range artDataTest {
			if strings.Contains(string(symbol), "o") {
				BannerFileName = "thinkertoy.txt"
				break
			} else if string(symbol) == "V" || string(symbol) == ")" || string(symbol) == "(" || string(symbol) == "/" ||
				string(symbol) == "\\" || string(symbol) == "<" || string(symbol) == ">" || string(symbol) == "'" {
				BannerFileName = "standard.txt"
				break
			} else {
				BannerFileName = "shadow.txt"
			}
		}

		fmt.Println("Using banner file:", BannerFileName)
		fContent, err := os.ReadFile(BannerFileName)
		if err != nil {
			fmt.Println(Errors("fileContentError"))
			os.Exit(0)
		}
		fontData := string(fContent)
		BannerLines := strings.Split(fontData, "\n")
		c := 0

		if len(asciiArt) > 9 {
			for i := 0; i < len(asciiArt)-1; {
				if len(asciiArt[i]) > 0 {
					c = i + 8
					if len(BannerLines)-c < 8 && len(BannerLines)-c != 0 {
						fmt.Println("Error: Invalid ascii art alignment.")
						os.Exit(0)
					}
					reverseArt(BannerLines, asciiArt[i:c], 0, 0, 1)
					fmt.Println()
					i = i + 8
				} else {
					fmt.Println()
					i = i + 1
				}
			}
		} else {
			reverseArt(BannerLines, asciiArt, 0, 0, 1)
		}
		fmt.Println("Reverse process complete.")
	}

	os.Exit(0)
}

var idx = 0

func reverseArt(BannerLines []string, asciiArt []string, SymbolFound int, ArtLine int, BannerLine int) {
	if SymbolFound != len(asciiArt[ArtLine]) {

		bannerWidth := len(BannerLines[BannerLine])

		if SymbolFound+bannerWidth <= len(asciiArt[ArtLine]) {
			if ArtLine < 7 {
				if BannerLines[BannerLine+ArtLine] == asciiArt[ArtLine][SymbolFound:SymbolFound+bannerWidth] { 
					reverseArt(BannerLines, asciiArt, SymbolFound, ArtLine+1, BannerLine) 
				} else {
					idx = BannerLine + 9
					if len(BannerLines)-1-idx < 9 && len(BannerLines)-1-idx != 0 {
						fmt.Println("Reverse process complete.")
						os.Exit(0)
					}
					reverseArt(BannerLines, asciiArt, SymbolFound, 0, idx) 
				}
			} else { 
				r := ((BannerLine - 1) / 9) + 32 
				fmt.Printf("Character: %c\n", r)
				reverseArt(BannerLines, asciiArt, SymbolFound+bannerWidth, 0, 1) 
			}
		} else {
			idx = BannerLine + 9
			if len(BannerLines)-1-idx < 9 && len(BannerLines)-1-idx != 0 {
				fmt.Println("Error: Invalid ascii art or unsupported banner format.")
				os.Exit(0)
			}
			reverseArt(BannerLines, asciiArt, SymbolFound, 0, idx) 
		}

	}
}

func ProcessFile(file io.Reader) bool {
	scanner := bufio.NewScanner(file)
	lineCount, wordNum, letterWidth := 0, 0, 0
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
			// Check for invalid line length
			fmt.Printf("Error: Invalid line length at line %d\n", counter)
			os.Exit(0)
		}

		lineCount++
		if lineCount == 8 {
			lineCount = 0
			wordNum++
		}
	}
	if lineCount != 0 {
		fmt.Printf("Error: Incomplete input at line %d\n", counter)
		os.Exit(0)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return true
}