/*
Build a function that accepts two runes as input and returns a string containing all the runes that come
between these two runes in the alphabet, in the correct order.
For instance, if the input runes are 'f' and 'j', the function should return "ghi" regardless of the order in which the runes are given.
*/

package sprint

func BetweenLimits(from, to rune) string {
	if from > to {
		from, to = to, from
	}

	var result string
	for i := from + 1; i < to; i++ {
		result += string(i)
	}
	return result
}
