package main

import (
	"fmt"
	sprint "koodSisuSprint/sprint/exam3"
)

func main() {
	// fmt.Println(sprint.Abacus(8, 3))
	// fmt.Println(sprint.Abacus(9, 2))
	// fmt.Println(sprint.Mean(1.15, -2.0, 8.35))
	// fmt.Println(sprint.Casting(1.15))
	// fmt.Println(sprint.ShiftBy('a', 4))
	// fmt.Printf("%c\n", sprint.ShiftBy('a', 4))
	// fmt.Println(sprint.StrConcat("r", "ProgrammerHumour", "/"))
	// fmt.Println(sprint.ReverseAlphabetValue('a'))
	// fmt.Println(sprint.IsNegative(1))
	// fmt.Print((sprint.IntVsFloat(5, 8.8)))
	// fmt.Println(sprint.IntVsFloat(-8, -8.8))
	// fmt.Println(sprint.Season("feb"))
	// fmt.Println(sprint.Season("September"))
	// fmt.Println(sprint.Accumulate(10))
	// fmt.Println(sprint.BetweenLimits('j', 'f'))
	// fmt.Println(sprint.Doop(5, "%", 3))
	// fmt.Println(sprint.Doop(8, "/", 0))
	// fmt.Println(sprint.IsLeapYear(2020))
	// fmt.Println(sprint.IsLeapYear(1900))
	// fmt.Println(sprint.CountDivisible(5, 17, 2, 3))
	// fmt.Println(sprint.FindDividend(5, 17, 4))
	// fmt.Println(sprint.FindDividend(10, 18, 9))
	// fmt.Println(sprint.AlphabetMastery(6))
	//fmt.Println(sprint.ReverseAlphabet(5))
	//fmt.Printf(sprint.Pairs())
	//fmt.Printf(sprint.Combinations())
	//fmt.Println(sprint.Countdown(5))
	//fmt.Println(sprint.AlphaNumber(1280))
	//fmt.Println(sprint.SimpleStrToInt("10203"))
	// fmt.Println(sprint.SimpleStrToInt("0000000010203"))
	//fmt.Println(sprint.StrToInt("--10203"))
	//fmt.Println(sprint.Swaps([]float64{1.4, 1.2, 2.5, 2.7, 3.1, 3.7, 2.5}))
	//fmt.Println(round(1.4))
	//fmt.Println(sprint.Compress("banana"))
	// fmt.Println(sprint.Decompress([]int{
	// 	72,
	// 	101,
	// 	108,
	// 	108,
	// 	111,
	// 	32,
	// 	70,
	// 	114,
	// 	105,
	// 	101,
	// 	110,
	// 	100,
	// 	33,
	// }))
	// input := "0000000100000000000100000000100000000000100000010000001000000000001000000010000000000010000000010000000010000000010000010000000000001000000001000000000001000000100000010000001000000000000100000000000010000000000001000001000001000001000000001000000000000100000100000000100000000001000000001000000001000000001000000000010000010000000000100000000001000001000000000010000001000000100000010000001000000001000000000010000001000000010000000000100000100000100000100000100000100000000001000001000000000010000001000000100000000001000000100000000100000000001"
	// result := sprint.Fuck(input)
	// fmt.Println(result) // Output: "hello world"
	fmt.Println(sprint.HashMap("thekey", []string{"5863446:value1", "193493707:value2",
		"6954055932175:value3", "210714995797:value4"}))
}

// func round(digit float64) int {
// 	decimal := int(digit)
// 	if digit-float64(decimal) >= 0.5 {
// 		return decimal + 1 // rounding up
// 	}
// 	return decimal // rounding down
// }
