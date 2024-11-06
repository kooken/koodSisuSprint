/*
Create a Go function that takes a string as its parameter.
The function should split the string into words and store them in a string slice.
Words are separated by spaces, tabs, and newlines.
*/

package sprint

func SplitWhitespaces(s string) []string {
	var result []string
	var word []rune  // temporary variable to store word before separator (space, \t, \n)
	if len(s) == 0 { // for empty strings
		result = append(result, "")
		return result
	}
	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' { // if character is separator
			if len(word) > 0 { // only if see a separator and a word is already created out of symbols
				result = append(result, string(word)) // adding a word to array
				word = []rune{}                       // clearing temp variable
			}
		} else { // creating a word out of symbols
			word = append(word, char)
		}
	}
	// adding last stored symbols in word to result string
	if len(word) > 0 {
		result = append(result, string(word))
	}
	return result
}
