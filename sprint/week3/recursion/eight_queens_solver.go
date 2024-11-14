/*
You're tasked with creating a function that returns the solutions to the eight queens puzzle.
Use recursivity to solve this problem.
The function returns the solutions as a string, with each solution on a single line.
The index of the placement of a queen starts at 1 and reads from left to right, with each digit representing the position for each column.
The solutions should be in ascending order.
For this task you're allowed to use strconv & strings packages.
*/

package sprint

import (
	"strconv"
	"strings"
)

func EightQueensSolver() string {
	var result []string
	cols := make([]bool, 8)   // Tracks if a column is under attack
	diag1 := make([]bool, 15) // Tracks the first diagonal (row-col)
	diag2 := make([]bool, 15) // Tracks the second diagonal (row+col)
	board := make([]int, 8)   // Tracks the queen positions in each row

	// Helper function to solve the puzzle recursively
	var solve func(row int)
	solve = func(row int) {
		if row == 8 {
			// Convert the solution to a string and add it to the result
			var sb strings.Builder
			for i := 0; i < 8; i++ {
				sb.WriteString(strconv.Itoa(board[i] + 1)) // Convert 0-index to 1-index
			}
			result = append(result, sb.String())
			return
		}

		for col := 0; col < 8; col++ {
			// Check if the column or diagonals are attacked
			if cols[col] || diag1[row-col+7] || diag2[row+col] {
				continue
			}

			// Place the queen
			board[row] = col
			cols[col] = true
			diag1[row-col+7] = true
			diag2[row+col] = true

			// Recurse to place the next queen
			solve(row + 1)

			// Backtrack
			cols[col] = false
			diag1[row-col+7] = false
			diag2[row+col] = false
		}
	}

	// Start solving from the first row
	solve(0)

	// Return all solutions as a single string
	return strings.Join(result, "\n")
}
