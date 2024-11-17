/*
Write a Go function that takes a string as input and compresses it by replacing consecutive
repeating characters with a count of repetitions followed by the character.
However, if there is only one character in a sequence, it should not be prepended with a number.
For example, "aaabbaac" should be compressed to "3a2b2ac", but "abcde" should remain unchanged as "abcde".
Allowed packages
For this task you're allowed to use fmt package.
Expected function signature
func StrCompress(input string) string {

}
Usage
StrCompress("aaabbaac")
>> "3a2b2ac"
*/

package sprint

import "fmt"

func StrCompress(input string) string {
	if len(input) == 0 {
		return ""
	}

	compressed := ""
	count := 1

	for i := 1; i < len(input); i++ {

		if input[i] == input[i-1] {
			count++
		} else {
			// print counter
			if count > 1 {
				compressed += fmt.Sprintf("%d", count)
			}
			compressed += string(input[i-1])
			count = 1 // reset counter
		}
	}
	if count > 1 {
		compressed += fmt.Sprintf("%d", count)
	}
	compressed += string(input[len(input)-1])

	return compressed
}
