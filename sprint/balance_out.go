package sprint

func BalanceOut(arr []bool) []bool {
	counterTrue := 0
	counterFalse := 0

	// Defining a number of 'true' and 'false' values in the array
	for _, element := range arr {
		if element == true {
			counterTrue++
		} else {
			counterFalse++
		}
	}
	if counterTrue > counterFalse {
		arrayToAppend := make([]bool, counterTrue-counterFalse) // 'make' sets bool values to false by default
		arr = append(arr, arrayToAppend...)                     // adding additional false values to the given array
	} else if counterFalse > counterTrue {
		arrayToAppend := make([]bool, counterFalse-counterTrue)
		for i := range arrayToAppend { // going through each index of the array
			arrayToAppend[i] = true // changing each element of a new array from false (default value) to true
		}
		arr = append(arr, arrayToAppend...) // adding additional true values to the given array
	}
	return arr
}
