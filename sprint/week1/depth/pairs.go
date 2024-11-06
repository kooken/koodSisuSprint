/*
Create a Go function that finds all possible combinations of two two-digit numbers.
The pairs should be in ascending order, and each number should be padded with a leading zero if it's less than 10.
The pairs should be separated by a comma and a space. Each number within a pair should be separated by a space.
The function should return a string containing these pairs.
For this task you're allowed to use fmt package.
*/

package sprint

import "fmt"

func Pairs() string {
	var result string
	for a := 0; a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {
			result += fmt.Sprintf("%02d %02d", a, b)
			if a != 98 || b != 99 {
				result += ", "
			}
		}
	}
	return result
}
