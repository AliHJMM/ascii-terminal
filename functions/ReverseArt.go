package functions

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// 1) Create a global or file-level string to accumulate reversed characters
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
		// Reset or empty any previous accumulation.
		ReversedResult = ""

		fmt.Println("Art file:", artFile)
		artContent, err := os.ReadFile(artFile)
		if err != nil {
			fmt.Println(Errors("textFileError"))
			os.Exit(0)
		}
		artData := string(artContent)
		artDataTest := strings.TrimSpace(artData)
		asciiArt := strings.Split(artData, "\n")

		// Detect which banner file to use
		BannerFileName := "standard.txt"
		for _, symbol := range artDataTest {
			if strings.ContainsRune("o", symbol) {
				BannerFileName = "thinkertoy.txt"
				break
			} else if strings.ContainsRune("V)(/\\<>'", symbol) {
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
					// Move forward 8 lines (one ASCII "character" block)
					i = i + 8
				} else {
					// Could treat a blank ASCII line as a word boundary or newline
					// For instance, if you want a space or a new line:
					// ReversedResult += " "
					i++
				}
			}
		} else {
			reverseArt(BannerLines, asciiArt, 0, 0, 1)
		}

		// 4) Print the entire reversed result at once
		fmt.Println("Reverse process complete.")
		fmt.Println("Reconstructed text (all at once):")
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
			// still scanning the 8 lines of a single character
			if BannerLines[BannerLine+ArtLine] == asciiArt[ArtLine][SymbolFound:SymbolFound+bannerWidth] {
				// match found for one line of the character, keep going
				reverseArt(BannerLines, asciiArt, SymbolFound, ArtLine+1, BannerLine)
			} else {
				// jump to next 9 lines in banner
				idx = BannerLine + 9
				if len(BannerLines)-1-idx < 9 && len(BannerLines)-1-idx != 0 {
					fmt.Println("Reverse process complete.")
					os.Exit(0)
				}
				reverseArt(BannerLines, asciiArt, SymbolFound, 0, idx)
			}
		} else {
			// 2) Instead of printing "Character: X", accumulate it in ReversedResult
			r := ((BannerLine - 1) / 9) + 32
			ReversedResult += string(r)

			// Move SymbolFound to the next chunk
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
