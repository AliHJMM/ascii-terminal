package functions

func Errors(errorType string) string {
	var errorMessages = map[string]string{
		"Usage": "\nUsage: go run main.go, followed by one of the options listed below as separate arguments:\n\n*Arguments in the same order*\n1- [STRING]\n2- [STRING] [BANNER]\n3- --color=<color> [STRING]\n4- --color=<color> [STRING] [BANNER]\n5- --color=<color> <letters to be colored> [STRING]\n6- --color=<color> <letters to be colored> [STRING] [BANNER]\n\n*Arguments in any order*\n7- --output=<fileName.txt>\n8- --align=<type>  ( left, right, center, justify )\n\n**********\nTo enable reverse functionality, activate it separately:\n9- --reverse=<fileName>",
		"availableColor": "\nThe available colors are:\nred, green, yellow, blue, purple, cyan, white, gray, darkred, orange, pink,\ngold, teal, lavender, brown, lightblue, magenta, olive, salmon, skyblue, darkpurple, lime\n",
		"bannerError": "The selected Banner input is not available. Choose from the following options: shadow, standard, or thinkertoy.",
		"alignmentError": "Invalid alignment input. Please use one of the following options: center, left, right, justify.",
		"fileContentError": "Unable to read the Banner file (standard.txt). Please ensure the file name is correct.",
		"reverseUsage": "Usage: go run main.go [OPTION]\n\nExample: go run main.go --reverse=<fileName>",
		"textFileError": "Unable to read the input file. Please verify the file name.",
		"terminalFitError": "The text does not fit within the terminal size.",
		"consoleSizeError": "Error while measuring console size:",
		"simpleUsage": "\nUsage: go run main.go [STRING]",
		"simpleUsageFile": "\nUsage: go run main.go, followed by one of these options:\n\n*Arguments in the same order*\n1- [STRING]\n2- [STRING] [BANNER]",
		"usageWithColor": "\nUsage: go run main.go, followed by one of the options listed below:\n\n*Arguments in the same order*\n1- [STRING]\n2- [STRING] [BANNER]\n3- --color=<color> [STRING]\n4- --color=<color> [STRING] [BANNER]\n5- --color=<color> <letters to be colored> [STRING]\n6- --color=<color> <letters to be colored> [STRING] [BANNER]",
		"usageWithAlign": "\nUsage: go run main.go, followed by one of the options listed below:\n\n*Arguments in the same order*\n1- [STRING]\n2- [STRING] [BANNER]\n3- --color=<color> [STRING]\n4- --color=<color> [STRING] [BANNER]\n5- --color=<color> <letters to be colored> [STRING]\n6- --color=<color> <letters to be colored> [STRING] [BANNER]\n\n*Arguments in any order*\n7- --align=<type> ( left, right, center, justify )",
		"usageWithOutput": "\nUsage: go run main.go, followed by one of the options listed below:\n\n*Arguments in the same order*\n1- [STRING]\n2- [STRING] [BANNER]\n3- --color=<color> [STRING]\n4- --color=<color> [STRING] [BANNER]\n5- --color=<color> <letters to be colored> [STRING]\n6- --color=<color> <letters to be colored> [STRING] [BANNER]\n\n*Arguments in any order*\n7- --output=<fileName.txt>\n8- --align=<type> ( left, right, center, justify )",
	}

	if errorMessage, ok := errorMessages[errorType]; ok {
		return errorMessage
	}
	return ""
}
