/*
Create a Go function that takes an integer as input and returns a string representing that integer.
If the number is negative, preserve the minus sign. Replace the digits with lowercase letters, where 0 becomes a, 1 becomes b,
and so on up to 9, which becomes j.
*/

package sprint

func AlphaNumber(n int) string {
	var result string
	var sign string
	alphabet := "abcdefghij"
	number := n
	if n < 0 {
		n *= -1
	}

	if n == 0 {
		return string(alphabet[0])
	}

	for n > 0 {
		digit := n % 10
		result = string(alphabet[digit]) + result
		n /= 10
	}

	if number < 0 {
		sign = "-"
	} else {
		sign = ""
	}
	return sign + result
}
