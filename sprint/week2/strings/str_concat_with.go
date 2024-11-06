/*
Create a Go function that takes a slice of strings and a separator string as its parameters.
The function should return a new string by concatenating all the strings in the slice,
with each string separated by the provided separator.
*/

package sprint

func StrConcatWith(strs []string, sep string) string {
	var result string
	for i, char := range strs {
		result += char
		if i < len(strs)-1 { // if it's not the last element
			result += sep // add separator
		}
	}
	return result
}
