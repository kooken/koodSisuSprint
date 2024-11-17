/*
Write a Go function that takes a valid Roman numeral as input and converts it into an integer.
The function should return the integer representation of the valid Roman numeral input.
Expected function signature
func FromRoman(s string) int {

}
Usage
FromRoman("CXXVIII")
>> 128
*/
package sprint

func FromRoman(s string) int {
	// map that stores all possible values
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	n := len(s) // length of the initial string

	for i := 0; i < n; i++ {
		value := romanValues[s[i]] // finding value from map

		if i+1 < n && romanValues[s[i+1]] > value {
			result -= value
		} else {
			result += value
		}
	}

	return result
}
