/*
Create a Go function that takes a string as its parameter.
The function should return true if the string contains only lowercase characters, and false otherwise.
*/

package sprint

func IsLower(s string) bool {
	var counter int
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			counter++
		}
	}
	return counter == len(runes)
}
