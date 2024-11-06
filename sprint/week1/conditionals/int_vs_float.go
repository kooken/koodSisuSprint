/*
Write a Go function that takes an integer and a float as input and returns a string.
The function should compare the integer to the float and return one of the following:
Integer if the integer is larger, Float if the float is larger, or Same if they are equal.
*/

package sprint

func IntVsFloat(i int, f float32) string {
	if float32(i) > f {
		return "Integer"
	}
	if f > float32(i) {
		return "Float"
	} else {
		return "Same"
	}
}
