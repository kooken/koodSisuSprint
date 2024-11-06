/*
Write a Go function that takes a 2D array of integers and an integer value.
The function should filter out all subarrays from the array (2D) in which the sum of elements is lower than the given value.
The resulting modified 2D array should be returned.
*/

package sprint

func FilterBySum(arr [][]int, limit int) [][]int {
	result := [][]int{}       // empty 2d array to store final array
	for _, row := range arr { // accessing rows in array [+][]
		sum := 0                  // sets 'sum' to 0
		for _, val := range row { // accessing columns in array [][+]
			sum += val // adding value to the sum
		}
		if sum >= limit { // if sum more than limit
			result = append(result, row) // adding 'row' value to the new array
		}
	}
	return result
}
