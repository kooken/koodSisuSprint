/*
Write a function that sorts a slice of integers in ascending order.
*/

package sprint

func SortIntegerTable(table []int) []int {
	n := len(table)

	// Bubble sorting
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}

	return table
}
