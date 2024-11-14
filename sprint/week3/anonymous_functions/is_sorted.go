/*
You're tasked with creating a function called IsSorted that returns true if the slice of string is sorted according to a provided comparison function,
and false otherwise.
The comparison function is passed as an argument func(a, b string) int, and it should return a positive int if the first argument
is greater than the second argument, 0 if they are equal, and a negative int otherwise.
Note that the slice can be sorted in both ascending and descending order.
For this task you will need to have in the same file StrCompare function you have previously created.
*/

package sprint

func IsSorted(f func(a, b string) int, arr []string) bool {
	if len(arr) < 2 {
		return true // A list with less than 2 elements is trivially sorted
	}

	// flags
	isAscending := true
	isDescending := true

	for i := 0; i < len(arr)-1; i++ {
		comparison := f(arr[i], arr[i+1])
		if comparison < 0 {
			isDescending = false
		} else if comparison > 0 {
			isAscending = false
		}
	}

	// The array is sorted if it's either entirely ascending or descending
	return isAscending || isDescending
}

func strCompare(a, b string) int {
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
