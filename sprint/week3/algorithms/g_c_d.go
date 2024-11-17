/*
Write a Go function that takes two integers as input and returns their greatest common divisor (GCD).
The GCD is the largest positive integer that divides both of the input integers without leaving a remainder.
Expected function signature
func GCD(a, b int) int {

}
Usage
GCD(12, 18)
>> 6
*/

package sprint

func GCD(a, b int) int {
	for a != 0 && b != 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}
	return a + b
}
