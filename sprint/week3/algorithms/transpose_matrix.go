/*
Write a Go function that takes a matrix represented as a 2D array and returns the transposed matrix.
Transposing a matrix involves swapping its rows and columns.
The function should take the original matrix and return the resulting transposed 2D array.
Expected function signature
func TransposeMatrix(matrix [][]int) [][]int {

}
Usage
TransposeMatrix([][]int{{1, 2, 3}, {4, 5, 6}})
>> [[1 4], [2, 5], [3, 6]]
*/
package sprint

func TransposeMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	transposed := make([][]int, cols)
	for i := 0; i < cols; i++ {
		transposed[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}
