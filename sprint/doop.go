/*
You're tasked with creating a function called Doop. This function takes three parameters:
A value (integer).
An operator, which can be one of the following: +, -, /, *, %.
Another value (integer).
In case of an invalid operator, invalid values, or incorrect number of arguments, the function returns 0.
The program must also handle modulo and division by 0.
*/

package sprint

func Doop(a int, op string, b int) int {
	if op == "+" {
		return a + b
	}
	if op == "-" {
		return a - b
	}
	if op == "/" && b != 0 {
		return a / b
	}
	if op == "*" {
		return a * b
	}
	if op == "%" {
		return a % b
	} else {
		return 0
	}
}
