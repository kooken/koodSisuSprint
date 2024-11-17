/*
Write a Go function that takes an integer as input and returns a 2D integer array representing
Pascal's triangle up to the specified number of levels. Pascal's triangle is a triangular array of binomial
coefficients where each number is the sum of the two directly above it.
The function should generate the specified number of levels of Pascal's triangle.
Expected function signature
func PascalsTriangle(n int) [][]int {

}
Usage
PascalsTriangle(2)
>> [[1] [1 1]]
*/

package sprint

func PascalsTriangle(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}

	triangle := make([][]int, n)

	for i := 0; i < n; i++ {
		row := make([]int, i+1)

		row[0], row[i] = 1, 1

		for j := 1; j < i; j++ {
			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}

		triangle[i] = row
	}

	return triangle
}
