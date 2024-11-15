/* Implement a !!Fuck interpreter, but the ! is replaced with zeros and #
with ones. Return the output as a string.
Expected function signature
package sprint

func Fuck(input string) string {
    // solution code here
}
Usage
Fuck("0000000100000000000100000000100000000000100000010000001000000000001000000010000000000010000000010000000010000000010000010000000000001000000001000000000001000000100000010000001000000000000100000000000010000000000001000001000001000001000000001000000000000100000100000000100000000001000000001000000001000000001000000000010000010000000000100000000001000001000000000010000001000000100000010000001000000001000000000010000001000000010000000000100000100000100000100000100000100000000001000001000000000010000001000000100000000001000000100000000100000000001")
>> "hello world" */

package sprint

func Fuck(input string) string {
	const (
		inc = "00000001" // Increment command
		out = "00000010" // Output command
	)

	var memory [30000]byte // Brainfuck memory tape
	pointer := 0           // Pointer to the memory cell
	result := ""           // Output string

	for i := 0; i < len(input); {
		if len(input[i:]) >= len(inc) && input[i:i+len(inc)] == inc {
			// Increment the current memory cell
			memory[pointer]++
			i += len(inc)
		} else if len(input[i:]) >= len(out) && input[i:i+len(out)] == out {
			// Output the current memory cell as a character
			result += string(memory[pointer])
			i += len(out)
		} else {
			// Skip any unrecognized commands
			i++
		}
	}

	return result
}
