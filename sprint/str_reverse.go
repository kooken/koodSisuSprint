package sprint

func StrReverse(s string) string {
	runes := []rune(s)
	var result string
	for i := len(runes) - 1; i >= 0; i-- {
		result += string(runes[i])
	}
	return result
}
