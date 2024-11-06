/*
Create a Go function that takes a non-empty string as input and returns the last character of the string as a rune.
*/

package sprint

func GetLastRune(s string) rune {
	runes := []rune(s)
	return runes[len(runes)-1]
}
