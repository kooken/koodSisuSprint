package sprint

func StrToInt(s string) int {
	result := 0
	sign := 1
	if s[0] == '-' && s[1] != '-' { // checking '-' sign
		sign *= -1
	} else if s[0] == '+' && s[1] != '+' { // checking '+' sign
		sign = 1
	} else {
		for i := 0; i < len(s); i++ {
			if s[i] == '-' || s[i] == '+' || s[i] == ' ' {
				return 0
			}
		}
	}
	for _, char := range s {
		digit := int(char - '0')
		if digit >= 0 && digit <= 9 {
			result = result*10 + digit
		}
	}
	return sign * result
}
