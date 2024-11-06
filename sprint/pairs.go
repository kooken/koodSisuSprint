package sprint

import "fmt"

func Pairs() string {
	var result string
	for a := 0; a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {
			result += fmt.Sprintf("%02d %02d", a, b)
			if a != 98 || b != 99 {
				result += ", "
			}
		}
	}
	return result
}
