/*
Write a Go function that takes a string containing various characters, including parentheses ()[]{},
and checks if the parentheses are balanced. The function should return a boolean value indicating whether
the parentheses are balanced or not.
Expected function signature
func BalancedParentheses(input string) bool {

}
Usage
BalancedParentheses("Everything { is [ fine ()]}()")
>> true */

package sprint

func BalancedParentheses(input string) bool {
	stack := []rune{}

	// Matching pairs of parentheses
	matches := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range input {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != matches[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
