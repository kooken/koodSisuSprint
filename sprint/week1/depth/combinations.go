/*
Create a Go function that generates a series of unique triplets of digits.
Each triplet should consist of different digits, and they should be in ascending order, from the smallest to the largest.
The triplets should be separated by a comma and a space.
For example, the function should return triplets like 012, 013, 014, 015....
Avoid combinations with all the same digits, like 000 or 999, and exclude triplets in descending order, like 987.
For this task you're allowed to use fmt package.
*/

package sprint

import "fmt"

func Combinations() string {
	var result string
	for a := 0; a <= 7; a++ {
		for b := a + 1; b <= 8; b++ {
			for c := b + 1; c <= 9; c++ {
				result += fmt.Sprintf("%d%d%d", a, b, c)
				if a == b && b == c {
					result += ""
				}
				if a == b+1 && b == c+1 {
					result += ""
				}
				if a == 7 && b == 8 && c == 9 {
					result += ""
				} else {
					result += ", "
				}
			}
		}
	}
	return result
}
