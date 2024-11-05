package sprint

func IsLower(s string) bool {
	var counter int
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			counter++
		}
	}
	return counter == len(runes)
}
