/*
Create a Go function that takes a string as its parameter and returns a new string with each letter capitalized.
*/

package sprint

func ToUpperCase(s string) string {
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			runes[i] -= 32
		}
	}
	return string(runes)
}
