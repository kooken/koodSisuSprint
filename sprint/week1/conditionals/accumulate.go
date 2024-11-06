/*
Create a Go function that takes a non-negative integer n.
If n is positive, return the sum of all the digits from 0 to n, including n itself.
If n is negative, return 0.
For example, in case of 4, the sum would be 0 + 1 + 2 + 3 + 4 = 10.
*/

package sprint

func Accumulate(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}
