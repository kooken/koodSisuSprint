/*
Create a Go function that takes an integer n and a string base as parameters.
The function should return the integer n represented in the specified base as a string.

Here are the validity rules for the base:
- The base must contain at least 2 unique characters.
- The base should not contain the characters + or -.

The function should handle negative numbers as well (see examples in the usage section).
If the base is not valid, the function should return "NV" (Not Valid).
*/

package sprint

func isValidBase(base string) bool {
	// checking the base length
	if len(base) < 2 {
		return false
	}

	for i := 0; i < len(base); i++ {
		// checking for '-' and '+' symbols
		if base[i] == '+' || base[i] == '-' {
			return false
		}

		// checking for doubles (2 same symbols in a row)
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}

	return true
}

func NbrBase(n int, base string) string {
	// Checking for validity
	if !isValidBase(base) { // if false -> return "NV"
		return "NV"
	}

	baseRunes := []rune(base) // converting base to runes
	baseLen := len(baseRunes) // checking length
	var result []rune         // var to store result

	// for negative numbers
	isNegative := n < 0 // boolean
	if isNegative {     // if true
		n = -n // converting a number
	}

	// processing number into a base
	for n > 0 {
		remainder := n % baseLen                                 // finding where's the symbol in the base system
		result = append([]rune{baseRunes[remainder]}, result...) // adding the remainder to the beginning of the string
		n /= baseLen                                             // preparing n for the next iteration
	}

	// if isNegative == true -> adding '-' sign
	if isNegative {
		result = append([]rune{'-'}, result...)
	}

	return string(result) // converting rune result to string
}
