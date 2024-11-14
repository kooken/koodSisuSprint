/*
You're tasked with creating a function ArrAny that returns true for a string slice if at least one element returns true when passed through the f function.
For this task you will need to have in the same file IsLower & IsNumeric functions you have previously created.
You will also need to implement two new functions:
IsUpper - Returns true if the string given contains only uppercase characters, and false otherwise.
IsAlphanumeric - Returns true if the string contains only alphanumeric characters, and false otherwise.
*/

package sprint

func ArrAny(f func(string) bool, a []string) bool {
	for _, s := range a {
		if f(s) {
			return true
		}
	}
	return false
}

func isUpper(s string) bool {
	var counter int
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			counter++
		}
	}
	return counter == len(runes)

}

func isAlphanumeric(s string) bool {
	var counter int
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		if runes[i] >= '0' && runes[i] <= '9' || runes[i] >= 'a' && runes[i] <= 'z' || runes[i] >= 'A' && runes[i] <= 'Z' {
			counter++
		}
	}
	return counter == len(runes)
}

func isLower(s string) bool {
	var counter int
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			counter++
		}
	}
	return counter == len(runes)
}

func isNumeric(s string) bool {
	var counter int
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		if runes[i] >= '0' && runes[i] <= '9' {
			counter++
		}
	}
	return counter == len(runes)
}
