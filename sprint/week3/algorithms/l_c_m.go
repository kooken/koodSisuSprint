/*
Write a Go function that takes two integers as input and returns their least common multiple (LCM).
The LCM is the smallest positive integer that is divisible by both of the input integers.
Expected function signature
func LCM(a, b int) int {

}
Usage
LCM(12, 18)
>> 36
*/
package sprint

func LCM(a, b int) int {
	m := a * b
	for a != 0 && b != 0 {
		if a > b {
			a %= b
		} else {
			b %= a
		}
	}
	return m / (a + b)
}
