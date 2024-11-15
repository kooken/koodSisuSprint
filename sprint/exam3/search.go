/*
Instructions
Given a slice of strings and another string, return the last index where the
string is found. If the string is not found in the slice, return -1.
Expected function signature
package sprint

func Search(strings []string, toSearch string) int {
    // solution code here
}
Usage
Search([]string{"Hello", "How", "Are", "You", "Hello"}, "Hello")
>> 4
*/

package sprint

func Search(strings []string, toSearch string) int {
	// starting from the end and go one by one
	for i := len(strings) - 1; i >= 0; i-- {
		if strings[i] == toSearch {
			return i // return first match
		}
	}
	return -1
}
