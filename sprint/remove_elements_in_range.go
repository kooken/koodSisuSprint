package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	if from > to {
		from, to = to, from // swapping values
	}

	if from < 0 {
		from = 0 // setting a 'from' to '0' if it's negative
	}
	if to > len(arr) {
		to = len(arr) // setting a 'to' to the array length value if it's more than the arrays length
	}

	return append(arr[:from], arr[to:]...) // making a new array
}
