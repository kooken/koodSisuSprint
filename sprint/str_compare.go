package sprint

func StrCompare(a, b string) int {
	minLen := len(a) // defining a shorter string
	if len(b) < minLen {
		minLen = len(b) // if other string is shorter - redifining a variable
	}

	for i := 0; i < minLen; i++ { // going symbol by symbol
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}

	// if strings are equal till the end of the shortest - check their length
	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}

	// strings are completely equal
	return 0
}
