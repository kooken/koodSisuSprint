/*
Create a Go function that takes a non-empty string as input and returns the first character of the string as a rune.
*/

package sprint

func GetFirstRune(s string) rune {
	return []rune(s)[0] // converts 1st element of a s [index 0] to a rune slice
}
