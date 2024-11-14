/*
You're tasked with creating a function that takes an integer as a parameter and returns true if this number is prime and false if it's not.
You should optimize the function to avoid time-outs during testing.
Remember that prime numbers can be only natural numbers greater than 1.
*/

package sprint

func IsPrime(n int) bool {
	i := 2
	if n > 1 {
		for i < n {
			if n%i == 0 {
				return false
			}
			i++
		}
		return true
	}
	return false
}
