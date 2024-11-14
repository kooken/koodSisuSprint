/*
You're tasked with creating a recursive function that returns the value at a given position index in the Fibonacci sequence.
The first value is at index 0.
The Fibonacci sequence starts with 0 and 1, and each subsequent value is the sum of the two preceding values: 0, 1, 1, 2, 3, etc.
Negative indices should be handled by returning -1.
*/

package sprint

func NthFibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	if index == 2 || index == 1 {
		return 1
	}
	return NthFibonacci(index-1) + NthFibonacci(index-2)
}
