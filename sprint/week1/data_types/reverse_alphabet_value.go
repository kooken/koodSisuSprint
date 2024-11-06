/*
Write a Go function that takes a lowercase letter rune ('a'-'z') as input and returns its reverse letter in the Latin alphabet
as a rune integer value. For example, 'a' should be transformed to 'z', 'y' should become 'b', and so on.
The function should maintain the case of the input letter.
*/

package sprint

func ReverseAlphabetValue(ch rune) rune {
	return 'a' + ('z' - ch)
}
