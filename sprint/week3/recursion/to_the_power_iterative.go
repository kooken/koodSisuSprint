/*
You're tasked with creating an iterative function that calculates the power of an integer n to the given power.
Handle negative powers by returning 0.
*/

package sprint

func ToThePowerIterative(n int, power int) int {

	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}

	result := 1
	for power > 0 {
		result *= n
		power--
	}
	return result
}
