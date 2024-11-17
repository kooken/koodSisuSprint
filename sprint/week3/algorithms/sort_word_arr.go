/*
Write a Go function that takes an array of integers and returns the longest contiguous subarray in the array
in which the elements are increasing. The function should return the contiguous subarray.
Expected function signature
func LongestClimb(arr []int) []int {

}
Usage
LongestClimb([]int{8, 4, 2, 1, 2, 4, 8, 2, 4, 8})
>> [1 2 4 8]
*/

package sprint

func SortWordArr(a []string) []string {
	return advancedSortWordArr(a, strCompare2)
}

func advancedSortWordArr(a []string, f func(a, b string) int) []string {
	n := len(a)
	// bubble sorting
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if f(a[j], a[j+1]) > 0 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func strCompare2(a, b string) int {
	minLen := len(a) // defining a shorter string
	if len(b) < minLen {
		minLen = len(b) // if other string is shorter - redifining a variable
	}

	for i := 0; i < minLen; i++ { // going symbol by symbol
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}

	// if strings are equal till the end of the shortest - check their length
	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}

	// strings are completely equal
	return 0
}
