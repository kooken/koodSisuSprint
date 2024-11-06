/*
Create a Go function that takes a single lowercase letter as a rune and an int 'step'.
The function should shift the letter in the alphabet by the specified 'step' value, and return the resulting letter.
For example, if you shift 'a' by 4 steps, it should become 'e'.
Remember to handle looping back to the start of the alphabet if needed.
Only lower case characters are tested.
*/

package sprint

func ShiftBy(r rune, step int) rune {
	for i := 0; i < step; i++ {
		if r == 'z' {
			r = 'a'
		} else {
			r++
		}
	}
	return r
}
