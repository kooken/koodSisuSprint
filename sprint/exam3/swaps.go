/* Given a float64 slice, return a slice but with every float64 value swapped with
next value if both values rounded up/down into integers are equal. If there's an
uneven amount, then the last odd one should remain as is.
For the purposes of this task, a float value is rounded up if its decimal points
are greater than or equal to .5. If not, it's rounded down.
Expected function signature
package sprint

func Swaps(slice []float64) []float64 {
   // solution code here
}
Usage
Swaps([]float64
   {1.4, 1.2, 2.5, 2.7, 3.1, 3.7, 2.5})
>> [1.2 1.4 2.7 2.5 3.1 3.7 2.5] */

package sprint

// func round(digit float64) int {
// 	decimal := int(digit)
// 	if digit-float64(decimal) >= 0.5 {
// 		return decimal + 1 // rounding up
// 	}
// 	return decimal // rounding down
// }

// func Swaps(slice []float64) []float64 {
// 	n := len(slice)
// 	copySlice := slice
// 	result := make([]float64, n)

// 	for i := 0; i < n-1; i++ {
// 		// rounding values
// 		currentDigit := round(copySlice[i])
// 		nextDigit := round(copySlice[i+1])

// 		if currentDigit == nextDigit {
// 			// swapping
// 			result[i], result[i+1] = result[i+1], result[i]
// 			i++ // to the next element
// 		}
// 	}
// 	return result
// }

func Swaps(slice []float64) []float64 {
	n := len(slice)
	result := make([]float64, n)
	copy(result, slice) // Create a copy to avoid modifying the original slice

	round := func(digit float64) int {
		decimal := int(digit)
		if digit-float64(decimal) >= 0.5 {
			return decimal + 1 // Round up
		}
		return decimal // Round down
	}

	for i := 0; i < n-1; i++ {
		// Round the current and next values
		currentRounded := round(slice[i])
		nextRounded := round(slice[i+1])

		// Swap if the rounded values are equal
		if currentRounded == nextRounded {
			result[i], result[i+1] = result[i+1], result[i]
			i++ // Skip the next element since we just swapped
		}
	}

	return result
}
