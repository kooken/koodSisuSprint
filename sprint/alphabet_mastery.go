/*
Create a function that takes a positive integer as input and returns the corresponding number of letters from the Latin alphabet.
The input integer won't be larger than the alphabet's length.
*/

package sprint

func AlphabetMastery(n int) string {
	var result string
	counter := 0
	for i := 'a'; i <= 'z' && counter < n; i++ {
		result += string(i)
		counter++
	}
	return result
}
