/*
You're tasked with creating a function ArrCountIf that returns the number of elements in a string slice that return true when passed through the function f.
For this task you will need to have in the same file some functions you have previously created:
IsLower
IsUpper
IsNumeric
IsAlphanumeric
*/

package sprint

func ArrCountIf(f func(string) bool, a []string) int {
	counter := 0
	for _, s := range a {
		if f(s) {
			counter++
		}
	}
	return counter
}

func IsUpper(s string) bool {
	var counter int
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			counter++
		}
	}
	return counter == len(runes)

}

func IsAlphanumeric(s string) bool {
	var counter int
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		if runes[i] >= '0' && runes[i] <= '9' || runes[i] >= 'a' && runes[i] <= 'z' || runes[i] >= 'A' && runes[i] <= 'Z' {
			counter++
		}
	}
	return counter == len(runes)
}

func isLower1(s string) bool {
	var counter int
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			counter++
		}
	}
	return counter == len(runes)
}

func isNumeric1(s string) bool {
	var counter int
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		if runes[i] >= '0' && runes[i] <= '9' {
			counter++
		}
	}
	return counter == len(runes)
}
