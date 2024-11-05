package sprint

func StrConcatWith(strs []string, sep string) string {
	var result string
	for i, char := range strs {
		result += char
		if i < len(strs)-1 { // if it's not the last element
			result += sep // add separator
		}
	}
	return result
}
