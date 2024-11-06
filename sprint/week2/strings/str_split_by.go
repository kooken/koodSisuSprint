/*
Create a Go function that takes a string s and a separator sep as parameters.
The function should split the string s into substrings using the separator sep and return a slice of strings
containing the resulting substrings.
*/

package sprint

func StrSplitBy(s, sep string) []string {
	var result []string
	var temp []rune           // to temporary store current word (before separator)
	sepRunes := []rune(sep)   // converting sep to runes
	sepLen := len(sepRunes)   // sep length
	i := 0                    // start index of 's'
	if s == "" || sep == "" { // for empty 's'
		return []string{}
	}
	for i < len(s) { // working until the end of the string
		// checking if symbols in string match with sep
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep { // identifying if a word is separator
			result = append(result, string(temp)) // adding word to a resulting string
			temp = []rune{}                       // clearing temp word to store next word after separator
			i += sepLen                           // skipping sep and going to the next word
		} else { // if not a sep -> adding symbol to word
			temp = append(temp, rune(s[i]))
			i++
		}
	}

	// adding last stored symbols in temp to result string
	if len(temp) > 0 {
		result = append(result, string(temp))
	}

	return result
}
