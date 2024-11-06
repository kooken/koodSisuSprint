/*
Create a Go function that takes an integer as input and returns a boolean value (true or false)
to indicate whether the input integer is negative or not.
*/

package sprint

func IsNegative(n int) bool {
	if n < 0 {
		return true
	} else {
		return false
	}
}
