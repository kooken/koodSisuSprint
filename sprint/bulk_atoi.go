package sprint

func strToInt(s string) int {
	if len(s) == 0 {
		return 0 // checking if the string is not empty
	}
	result := 0
	sign := 1
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1 // if the string starts with the sign, skip the sign and start with [1] index
	} else if s[0] == '+' {
		start = 1
	}

	for i := start; i < len(s); i++ { // going through each symbol in the string
		char := s[i] // start from the start index
		if char < '0' || char > '9' {
			return 0 // if it includes NOT numbers - skip and return 0
		}
		digit := int(char - '0') // converting into number
		result = result*10 + digit
	}

	return sign * result
}

func BulkAtoi(arr []string) []int {
	result := []int{} // empty array to store result

	for _, char := range arr { // going through each of the element of the array
		result = append(result, strToInt(char)) // append everything to the new resulting array
	}
	return result
}
