/* Given an integer slice and boolean, return a slice that orders the integers
depending on the boolean (either ascending if true or descending if false).
Expected function signature
package sprint

func SortInts(ints []int, asc bool) []int {
    // solution code here
}
Usage
SortInts([]int{1,5,2,8}, false)
>> [8,5,2,1] */

package sprint

func SortInts(ints []int, asc bool) []int {
	n := len(ints)
	if asc == true {
		// Bubble sorting
		for i := 0; i < n-1; i++ {
			for j := 0; j < n-i-1; j++ {
				if ints[j] > ints[j+1] {
					ints[j], ints[j+1] = ints[j+1], ints[j]
				}
			}
		}
	} else if asc == false {
		// bubble sort but in reverse
		for i := 0; i < n-1; i++ {
			for j := 0; j < n-i-1; j++ {
				if ints[j] < ints[j+1] {
					ints[j], ints[j+1] = ints[j+1], ints[j]
				}
			}
		}
	}
	return ints
}
