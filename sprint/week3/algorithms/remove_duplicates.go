/*
Write a Go function that takes an array of integers as input and returns the same array with
duplicate elements removed, preserving the order of the remaining elements.
For example, if the input array is [3, 5, 2, 3, 8, 5, 2], the function should return [3, 5, 2, 8].
Expected function signature
func RemoveDuplicates(arr []int) []int {

}
Usage
RemoveDuplicates([]int{1, 2, 3, 2, 4, 8, 8, 1, 2, 0, 8})
>> [1, 2, 3, 4, 8, 0]
*/

package sprint

func RemoveDuplicates(arr []int) []int {
	seen := make(map[int]bool) // to store seen elements
	result := []int{}          // to store unique

	for _, value := range arr {
		if !seen[value] {
			result = append(result, value)
			seen[value] = true // flag to mark element
		}
	}

	return result
}
