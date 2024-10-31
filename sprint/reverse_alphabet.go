/*
Create a function that takes a step value n as input.
Starting from z in the Latin alphabet, it should return the lowercase alphabet in reverse order as a string.
For each letter, you skip n-1 letters. If n is 0 or negative, use a default step of 1.
*/

package sprint

func ReverseAlphabet(step int) string {
	if step <= 0 {
		step = 1
	}

	var result string
	for i := 'z'; i >= 'a'; i -= rune(step) {
		result += string(i)
	}
	return result
}
