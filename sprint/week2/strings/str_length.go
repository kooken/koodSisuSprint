/*
Create a Go function that accepts a string and returns two integers.
The first integer represents the number of runes in the string, and the second integer represents the byte length of the string.
The function should handle UTF-8 encoding for all characters.
*/

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
