package sprint

func GetLastRune(s string) rune {
	runes := []rune(s)
	return runes[len(runes)-1]
}
