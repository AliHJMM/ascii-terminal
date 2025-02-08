package functions

func AreAllStringsEmpty(strings []string) bool {
	for _, str := range strings {
		if str != "" {
			return false
		}
	}
	return true
}
