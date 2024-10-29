/*
Create a function that takes two strings and a delimiter.
The function should combine the two strings, placing the delimiter between them, and return the combined result as a single string.
*/

package sprint

func StrConcat(s1, s2, delim string) string {
	result := s1 + delim + s2
	return result
}
