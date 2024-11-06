/*
Create a Go function that takes two integers, min and max, as input.
The function should return a slice of integers containing all values between min (inclusive) and max (exclusive)
using the make function to create the slice.
If min is greater than or equal to max, the function should return a nil slice.
*/

package sprint

func GenerateRange(min, max int) []int {
	if min >= max {
		return nil // returning empty array (== nil)
	}

	length := max - min
	arr := make([]int, length) // making a new array with a fixed length

	for i := 0; i < length; i++ {
		arr[i] = min + i // adding value to the new array
	}
	return arr
}
