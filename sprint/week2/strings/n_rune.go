/*
Create a Go function that takes a non-empty string as its first argument and an index as its second argument.
The function should return the rune at the specified index in the string.
*/

package sprint

func NRune(s string, i int) rune {
	runes := []rune(s)
	return runes[i]
}
