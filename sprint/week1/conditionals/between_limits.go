package sprint

func BetweenLimits(from, to rune) string {
	if from > to {
		from, to = to, from // swapping the values
	}

	var result string
	for i := from + 1; i < to; i++ {
		result += string(i)
	}
	return result
}
