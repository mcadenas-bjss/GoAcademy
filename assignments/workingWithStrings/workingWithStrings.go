package workingWithStrings

// ConcatenateStrings concatenates all string arguments passed in.
func ConcatenateStrings(args ...string) string {
	var result string
	for _, str := range args {
		result += str
	}
	return result
}
