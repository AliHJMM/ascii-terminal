package main

import (
	"ASCII-TERMINAL/functions"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 8 {
		fmt.Println(functions.Errors("Usage"))
		fmt.Println(functions.Errors("availableColor"))
		return
	}
	inputargs := os.Args[1:]
	filename := "standard.txt"
	coloredLetters := ""
	color := ""
	output := ""
	align := ""
	restoprint := ""

	if functions.AreAllStringsEmpty(inputargs) {
		os.Exit(0)
	}
	for i := 0; i < len(inputargs); i++ {
		if strings.HasPrefix(inputargs[i], "--reverse") {
			functions.ReverseArt()
			os.Exit(0)
		}
		if strings.HasPrefix(inputargs[i], "--align") {
			align = strings.TrimPrefix(inputargs[i], "--align=")
			inputargs[i] = ""
		}
		if strings.HasPrefix(inputargs[i], "--output") {
			output = strings.TrimPrefix(inputargs[i], "--output=")
			inputargs[i] = ""
		}
	}
	inputargs = functions.FilterEmptyArgs(inputargs)
	lastindex := len(inputargs)
	if len(inputargs) >= 1 {
		str := inputargs[0]

		if strings.HasPrefix(inputargs[0], "--color") {
			color = strings.TrimPrefix(inputargs[0], "--color=")
			if functions.GetColorCode(color) == "" {
				fmt.Println(functions.Errors("availableColor"))
				os.Exit(0)
			}
			if lastindex == 4 {
				coloredLetters = inputargs[1]
				str = inputargs[2]
				filename = functions.GetBannerFile(inputargs[3])
			} else if lastindex == 3 {
				if functions.GetBannerFile(inputargs[2]) != "" {
					coloredLetters = "   "
					str = inputargs[1]
					filename = functions.GetBannerFile(inputargs[2])
				} else {
					coloredLetters = inputargs[1]
					str = inputargs[2]
				}
			} else if lastindex == 2 {
				coloredLetters = "   "
				str = inputargs[1]
			} else {
				fmt.Println(functions.Errors("Usage"))
				fmt.Println(functions.Errors("availableColor"))
				os.Exit(0)
			}
		} else if lastindex == 2 {
			str = inputargs[0]
			filename = functions.GetBannerFile(inputargs[1])
		} else if len(inputargs) != 1 {
			fmt.Println(functions.Errors("Usage"))
			fmt.Println(functions.Errors("availableColor"))
			os.Exit(0)
		}

		str = functions.FormatAlphabetString(str)
		res := ""
		restemp := ""
		args := strings.Split(str, "\\n")
		if functions.AreAllStringsEmpty(args) {
			args = args[1:]
		}
		if align == "justify" {
			restoprint = functions.JustifyText(args, filename, color, coloredLetters, align, output)
		} else {
			for _, word := range args {
				if word == "" {
					fmt.Println()
					restoprint = restoprint + "\n"
					continue
				}

				for i := 0; i < 8; i++ {
					res, restemp = functions.FormatAndColorWord(i, word, filename, color, coloredLetters, align, res, restemp)
					if align != "" {
						space := functions.AlignText(restemp, align, args)
						res = space + res
						fmt.Println(res)
					} else {
						fmt.Println(res)
					}
					restoprint = restoprint + "\n" + restemp
					res = ""
					restemp = ""
				}
			}
		}
	} else {
		fmt.Println(functions.Errors("Usage"))
		fmt.Println(functions.Errors("availableColor"))
		return
	}
	if output != "" {
		err := functions.SaveContentToFile(restoprint[1:], output)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
		}
	}
	os.Exit(0)
}
