package sprint

func FilterBySum(arr [][]int, limit int) [][]int {
	result := [][]int{}       // empty 2d array to store final array
	for _, row := range arr { // accessing rows in array [+][]
		sum := 0                  // sets 'sum' to 0
		for _, val := range row { // accessing columns in array [][+]
			sum += val // adding value to the sum
		}
		if sum >= limit { // if sum more than limit
			result = append(result, row) // adding 'row' value to the new array
		}
	}
	return result
}
