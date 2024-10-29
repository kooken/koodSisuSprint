/*
Write a function that accepts a float64 value, rounds it to the nearest integer, and then returns the result as an int.
For this task you're allowed to use math package.
*/

package sprint

import "math"

func Casting(n float64) int {
	var result float64 = math.Round(n)
	return int(result)
}
