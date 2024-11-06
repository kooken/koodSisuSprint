package sprint

func SubstrIndex(s string, toFind string) int {
	if len(toFind) > len(s) {
		return -1
	}
	if s == "" || toFind == "" { // check for empty strings
		return 0
	}
	for i := 0; i <= len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind { // if slice of 's' id equal to 'toFind'
			return i // returning index value
		}
	}

	return -1
}
