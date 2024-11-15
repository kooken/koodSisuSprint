/* Instructions
Given a string, apply the Burrows-Wheeler algorithm to it. For testing purposes,
the character ^ will symbolize the start of the string, and $ the end.
Expected function signature
package sprint

func Compress(text string) string {
    // solution code here
}
Usage
Compress("banana")
>>"a$nnb^aa" */

package sprint

// all cyclic rotations of a string
func generateRotations(input string) []string {
	n := len(input)
	rotations := make([]string, n) // new empty string
	for i := 0; i < n; i++ {
		rotations[i] = input[i:] + input[:i] // cyclic rotation
	}
	return rotations
}

// bubble sorting string
func sortStrings(arr []string) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// swapping
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func Compress(text string) string {
	text = "^" + text + "$" // start and end + text itself

	rotations := generateRotations(text)

	sortStrings(rotations)

	var bwtResult string
	for _, rotation := range rotations {
		bwtResult += string(rotation[len(rotation)-1]) // adding last element to the resulting string
	}

	return bwtResult
}
