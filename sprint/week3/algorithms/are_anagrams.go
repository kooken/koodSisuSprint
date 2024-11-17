/*
Write a Go function that takes two strings as input and returns a boolean value indicating whether
the input strings are anagrams or not. Anagrams are words or phrases formed by rearranging the letters of another,
and in this task, you should ignore differences in character case and whitespace.
For this task you're allowed to use sort & strings packages.
Expected function signature
func AreAnagrams(str1, str2 string) bool {

}
Usage
AreAnagrams("Listen", "S i l e n t")
>> true
*/

package sprint

import (
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	str1 = strings.ReplaceAll(strings.ToLower(str1), " ", "")
	str2 = strings.ReplaceAll(strings.ToLower(str2), " ", "")

	// checking length
	if len(str1) != len(str2) {
		return false
	}

	s1 := []rune(str1)
	s2 := []rune(str2)
	// anonymous functions
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })

	return string(s1) == string(s2)
}
