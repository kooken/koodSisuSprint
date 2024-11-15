/* Instructions
Implement a decoder for the Lempel–Ziv–Welch compression algorithm.
Expected function signature
package sprint

func Decompress(compressed []int) string {
    // solution code here
}
Usage
Decompress([]int{
    72,
    101,
    108,
    108,
    111,
    32,
    70,
    114,
    105,
    101,
    110,
    100,
    33,
})
>> "Hello Friend!" */

package sprint

func Decompress(compressed []int) string {
	// making a dict with ascii values
	dictionary := make(map[int]string)
	for i := 0; i < 256; i++ {
		dictionary[i] = string(i)
	}

	output := ""
	currentDictIndex := 256

	// Handle the first symbol (it corresponds directly to a character in the dictionary)
	if len(compressed) > 0 {
		prev := compressed[0]      // start
		output += dictionary[prev] // adding to the string

		for i := 1; i < len(compressed); i++ {
			curr := compressed[i]

			var entry string
			if _, exists := dictionary[curr]; exists {
				entry = dictionary[curr]
			} else if curr == currentDictIndex {
				entry = dictionary[prev] + string(dictionary[prev][0])
			}

			// adding to the string
			output += entry

			// updating dict
			dictionary[currentDictIndex] = dictionary[prev] + string(entry[0])
			currentDictIndex++

			// updating element
			prev = curr
		}
	}

	return output
}
