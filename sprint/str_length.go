package sprint

func StrLength(s string) []int {
	var result []int
	var counterRunes int
	var counterBytes int
	for range []rune(s) {
		counterRunes++
	}
	for range []byte(s) {
		counterBytes++
	}
	result = append(result, counterRunes)
	result = append(result, counterBytes)
	return result
}
