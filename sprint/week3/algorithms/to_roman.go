/*
Write a Go function that takes an integer and converts it into a Roman numeral.
The function should return "Invalid input" if the input integer is less than 1 or more than 3999.
Otherwise, it should return the Roman numeral representation of the valid input integer.
Expected function signature
func ToRoman(num int) string {

}
Usage
ToRoman(128)
>> "CXXVIII"
*/

package sprint

func ToRoman(num int) string {
	if num < 1 || num > 3999 {
		return ""
	}

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	// string to store final result
	roman := ""

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			roman += symbols[i]
			num -= values[i]
		}
	}

	return roman
}
