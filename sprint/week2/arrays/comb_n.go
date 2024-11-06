/*
Dive into the world of number combinations! Your goal is to create a function that reveals all possible combinations of ascending
digits of length n (see example).
For this task you're allowed to use fmt package.
*/

package sprint

func CombN(n int) []string {
	var result []string // string to store final result
	var combination []rune
	var generate func(start int)

	// Recursive function to generate combinations
	generate = func(start int) { // start == 0 in the beginning (line 21)
		if len(combination) == n { // if length of comb == n -> putting this into result
			result = append(result, string(combination))
			return // return the result only when the length of a combination == n
		}
		for i := start; i <= 9; i++ { // starts with '0'
			combination = append(combination, rune(i+'0')) // Convert current digit to rune
			generate(i + 1)                                // Move to the next digit using recursion
			combination = combination[:len(combination)-1] // Backtrack (deleting current digit and go to the next one)
		}
	}

	generate(0) // start generating combinations with 0
	return result
}
