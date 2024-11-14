/*
You're tasked with creating an recursive function that calculates the power of an integer n to the given power.
Handle negative powers by returning 0.
NB! You should achieve this without using a for loop.
*/

package sprint

func ToThePowerRecursive(n int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	return ToThePowerRecursive(n, power-1) * n
}
