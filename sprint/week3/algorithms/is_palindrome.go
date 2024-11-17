/*
Write a Go function that takes a string as input and returns a boolean value indicating whether
the input string is a palindrome or not. A palindrome is a string that reads the same forwards and backward,
considering character case and white spaces.
Expected function signature
func IsPalindrome(s string) bool {

}
Usage
IsPalindrome("level")
>> true
*/

package sprint

func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
