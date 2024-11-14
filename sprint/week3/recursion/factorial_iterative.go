/*
You're tasked with creating an iterative function that calculates the factorial of an integer passed as a parameter.
Make sure to handle errors, returning 0 for non-possible values or overflows.
*/
package sprint

func FactorialIterative(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}
