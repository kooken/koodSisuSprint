/*
Create a Go function that takes two parameters:
s: a numeric string in a specific base.
base: a string containing unique characters representing the available digits in that base.

The function should return the integer value of s in the given base. If the base is not valid, it returns 0.

Here are the validity rules for the base:
- The base must contain at least 2 unique characters.
- The base should not contain the characters + or -.
- The function only works with valid string numbers that consist of elements present in the base. It does not handle negative numbers.

For this task you're allowed to use math & strings packages.
*/

package sprint

import (
	"math"
	"strings"
)

func isValid(base string) bool {
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

func ConvertAnyToDec(s string, base string) int {
	if !isValid(base) {
		return 0
	}

	baseLen := len(base)
	result := 0

	// going symbol by symbol in s
	for i, char := range s {
		val := strings.IndexRune(base, char) // looking for 1st seen rune in base and return its index
		if val == -1 {                       // by default IndexRune returns -1 if not found
			return 0 // not found -> return 0
		}
		// Calculating decimal value
		exponent := len(s) - i - 1
		result += val * int(math.Pow(float64(baseLen), float64(exponent))) // .Pow - exponentiation
	}

	return result
}
