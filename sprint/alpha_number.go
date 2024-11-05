package sprint

func AlphaNumber(n int) string {
	var result string
	var sign string
	alphabet := "abcdefghij" // alphabet from 0 to 9
	number := n              // a copy of an initial integer (n), so it won't be changed during next steps
	if n < 0 {
		n *= -1
	}

	if n == 0 {
		return string(alphabet[0]) // returning 'a' if integer == 0
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