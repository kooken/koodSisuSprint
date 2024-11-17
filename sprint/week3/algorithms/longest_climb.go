/*
Write a Go function that takes an array of integers and returns the longest contiguous subarray in
the array in which the elements are increasing. The function should return the contiguous subarray.
Expected function signature
func LongestClimb(arr []int) []int {

}
Usage
LongestClimb([]int{8, 4, 2, 1, 2, 4, 8, 2, 4, 8})
>> [1 2 4 8] */

package sprint

func LongestClimb(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	// start and length of the longest increasing subarray
	longestStart, longestLength := 0, 1
	// current increasing subarray
	currentStart, currentLength := 0, 1

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			// if current element is greater than previous:
			currentLength++
		} else {
			// reset current subarray when sequence is no longer increasing
			if currentLength > longestLength {
				longestStart = currentStart
				longestLength = currentLength
			}
			currentStart = i  // resetting to i
			currentLength = 1 // resetting to 1
		}
	}

	if currentLength > longestLength {
		longestStart = currentStart
		longestLength = currentLength
	}

	return arr[longestStart : longestStart+longestLength]

}
