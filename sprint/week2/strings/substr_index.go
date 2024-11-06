/*
Create a Go function that behaves like the Index function.
It takes two strings as input: s and toFind.
The function should find the index of the first occurrence of toFind in s and return that index as an integer.
If toFind is not present in s, the function should return -1.
*/

package sprint

func SubstrIndex(s string, toFind string) int {
	if len(toFind) > len(s) {
		return -1
	}
	if s == "" || toFind == "" { // check for empty strings
		return 0
	}
	for i := 0; i <= len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind { // if slice of 's' id equal to 'toFind'
			return i // returning index value
		}
	}

	return -1
}
