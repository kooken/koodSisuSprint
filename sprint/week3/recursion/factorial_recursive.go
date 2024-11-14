/*
You're tasked with creating a recursive function that calculates the factorial of an integer passed as a parameter.
Make sure to handle errors, returning 0 for non-possible values or overflows.
NB! You should achieve this without using a for loop.
*/

package sprint

func FactorialRecursive(n int) int {
	if n == 0 {
		return 1
	} else if n == -1 {
		return 0
	}
	return n * FactorialRecursive(n-1)
}
