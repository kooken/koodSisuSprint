package sprint

import (
	"math"
	"strings"
)

func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {
	// Converting number from 'baseFrom' to decimal
	decimalValue := 0
	baseFromLen := len(baseFrom)
	for i, char := range nbr {
		val := strings.IndexRune(baseFrom, char) // looking for 1st seen rune in baseFrom and return its index
		if val == -1 {                           // by default IndexRune returns -1 if not found
			return "" // not found -> return empty string
		}

		// Calculating decimal value
		exponent := len(nbr) - i - 1
		decimalValue += val * int(math.Pow(float64(baseFromLen), float64(exponent)))
	}

	// Converting decimal to baseTo
	if decimalValue == 0 {
		return string(baseTo[0]) // If decimal is 0 -> returning first value from baseTo
	}

	result := []rune{}
	baseToRunes := []rune(baseTo)
	baseToLen := len(baseTo)

	for decimalValue > 0 {
		remainder := decimalValue % baseToLen
		result = append([]rune{baseToRunes[remainder]}, result...) // Adding number to the beggining (1st position)
		decimalValue /= baseToLen
	}

	return string(result)
}
