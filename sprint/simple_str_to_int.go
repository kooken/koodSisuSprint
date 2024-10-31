/*
Create a Go function that converts a valid string representation of a number into an integer.
If the string is not a valid number, the function should return 0.
The input will only contain one or more digits, and you don't need to handle signs like + or -.
*/

package sprint

func SimpleStrToInt(s string) int {
	result := 0

	for _, char := range s {
		digit := int(char - '0')
		result = result*10 + digit
	}

	return result
}
