/*
Given a slice of integer values, return the same slice but with all the values
replaced with either the reverse factorial if possible or -1 if not.
Expected function signature:
package sprint

func InvFactArr(intSlice []int) []int {
    // solution code here
}
Usage
InvFactArr([]int{120, 5})
>> [5 -1]
Hint
A map in Go could be very useful...
*/

package sprint

func checkFactorial(x int) int {
	// number can't be negative
	if x <= 0 {
		return -1
	} else if x == 1 {
		return 1
	}
	n := 1 // initial value for n!
	currentFact := 1

	for currentFact < x {
		n++
		currentFact *= n
		if currentFact == x {
			return n
		}
	}
	return -1
}

func InvFactArr(intSlice []int) []int {
	// making a new empty slice
	result := make([]int, len(intSlice))
	for i, x := range intSlice {
		// filling slice with factorials or -1
		result[i] = checkFactorial(x)
	}
	return result
}
