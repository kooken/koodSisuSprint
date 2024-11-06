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
