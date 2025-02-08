package functions

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetConsoleWidth() int {
	command := exec.Command("stty", "size")
	command.Stdin = os.Stdin
	terminalOutput, err := command.Output()
	if err != nil {
		fmt.Println(Errors("size"))
		os.Exit(1)
	}

	terminalOutputStr := string(terminalOutput)
	terminalOutputStr = strings.TrimSpace(terminalOutputStr)
	dimensions := strings.Split(terminalOutputStr, " ")
	width, err := strconv.Atoi(dimensions[1])
	if err != nil {
		fmt.Println(Errors("size"))
		os.Exit(1)
	}

	return width
}
