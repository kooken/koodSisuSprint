/*
Create a function that takes a one-digit integer as input and returns a countdown string.
The countdown should start from the given number, skip every second number, and end with 0!.
For example, if the input is 7, the function should return "7, 5, 3, 1, 0!".
*/

package sprint

func Countdown(n int) string {
	var result string
	for i := rune(n); i >= rune(1); i -= rune(2) {
		result += string('0'+i) + ", "
	}
	return result + "0!"
}
