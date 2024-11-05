package sprint

func IsNumeric(s string) bool {
	var counter int
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		if runes[i] >= '0' && runes[i] <= '9' {
			counter++
		}
	}
	return counter == len(runes)
}
