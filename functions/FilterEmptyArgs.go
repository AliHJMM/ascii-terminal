package functions

func FilterEmptyArgs(args []string) []string {
	filteredArgs := make([]string, 0, len(args))
	for _, argument := range args {
		if argument != "" {
			filteredArgs = append(filteredArgs, argument)
		}
	}

	return filteredArgs
}
