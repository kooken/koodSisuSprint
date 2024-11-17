/*
Write a Go function that takes two arrays of integers and returns an array containing the elements
that are common in both input arrays, sorted in ascending order. If an element occurs multiple times in
both arrays, it should be included in the result array as many times as it occurs.
Expected function signature
func Overlap(arr1, arr2 []int) []int {

}
Usage
arr1 := []int{1, 2, 2, 3, 4, 5}
arr2 := []int{2, 3, 4, 4, 5, 6}
Overlap(arr1, arr2)
>> [2 3 4 5]
*/

package sprint

func Overlap(arr1, arr2 []int) []int {
	var result []int

	for _, num1 := range arr1 {
		for i, num2 := range arr2 {
			if num1 == num2 {
				result = append(result, num1)
				arr2 = append(arr2[:i], arr2[i+1:]...)
				break
			}
		}
	}
	if len(result) == 0 {
		return []int{}
	}

	// bubble sorting
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result)-1-i; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}

	return result
}
