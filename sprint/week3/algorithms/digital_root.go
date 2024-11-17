/*
Write a Go function that solves the digital root problem.
The function should take an integer as input and return an integer.
The digital root is the recursive sum of the digits of a number until a single-digit result is achieved.
For example, the digital root of 9875 is 2 because 9 + 8 + 7 + 5 = 29, and 2 + 9 = 11, and finally 1 + 1 = 2.
Expected function signature
func DigitalRoot(n int) int {

}
Usage
DigitalRoot(123456)
>> 3 */

package sprint

func DigitalRoot(n int) int {
	if n == 0 {
		return 0
	}

	if n%9 == 0 {
		return 9
	}

	return n % 9
}
