/*
You're tasked with creating a function that takes an integer as a parameter.
If the given integer is a prime, the function should return that integer; otherwise, it should find and return the next prime number.
You should optimize the function to avoid time-outs during testing.
Remember that prime numbers can be only natural numbers greater than 1.
*/

package sprint

func isPrime(n int) bool {
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

func NextPrime(n int) int {
	if isPrime(n) == true {
		return n
	}
	n++
	for isPrime(n) == false {
		n++
	}
	return n
}
