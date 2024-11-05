package sprint

func GenerateRange(min, max int) []int {
	if min >= max {
		return nil // returning empty array (== nil)
	}

	length := max - min
	arr := make([]int, length) // making a new array with a fixed length

	for i := 0; i < length; i++ {
		arr[i] = min + i // adding value to the new array
	}
	return arr
}
