/*
You're tasked with creating a function that takes an integer as a parameter and returns its square root if the square root is a whole number.
Otherwise, it should return 0.
*/

package sprint

func Sqrt(n int) int {

	i := 1
	if n == 0 || n == 1 {
		return n
	}
	if n < 0 {
		return 0
	}
	for i*i <= n {
		if i*i == n {
			return i
		}
		i++
	}
	return 0
}
