/*
You're tasked with creating a function ArrMap that applies a function of type func(int) bool to each element of an int slice and
returns a slice of the resulting return values.
For this task you will need to have in the same file some functions you have previously created:
IsNegative
IsPrime
*/

package sprint

func ArrMap(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, v := range a {
		result[i] = f(v)
	}
	return result
}

func isNegative(n int) bool {
	return n < 0
}

func isPrime1(n int) bool {
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
