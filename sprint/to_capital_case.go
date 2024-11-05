package sprint

func ToCapitalCase(s string) string {
	if len(s) == 0 {
		return ""
	}
	runes := []rune(s)
	if (runes[0] >= 'a') && (runes[0] <= 'z') { // checking the very first element [0 index] of a rune array
		runes[0] -= 32 // to uppercase
	}
	for i := 1; i < len(runes); i++ { // starting from the 1st index
		if runes[i] >= 'a' && runes[i] <= 'z' {
			if !((runes[i-1] >= 'A' && runes[i-1] <= 'Z') || (runes[i-1] >= 'a' && runes[i-1] <= 'z') || (runes[i-1] >= '0' && runes[i-1] <= '9')) {
				runes[i] -= 32 // to uppercase
			}
		} else if runes[i] >= 'A' && runes[i] <= 'Z' {
			// Lowercase all letters that aren't the first letter of a word
			if (runes[i-1] >= 'A' && runes[i-1] <= 'Z') || (runes[i-1] >= 'a' && runes[i-1] <= 'z') || (runes[i-1] >= '0' && runes[i-1] <= '9') {
				runes[i] += 32 // to lowercase
			}
		}
	}
	return string(runes)
}
