package main

import (
	"fmt"
	"koodSisuSprint/sprint"
)

func main() {
	fmt.Println(sprint.Abacus(8, 3))
	fmt.Println(sprint.Abacus(9, 2))
	fmt.Println(sprint.Mean(1.15, -2.0, 8.35))
	fmt.Println(sprint.Casting(1.15))
	fmt.Println(sprint.ShiftBy('a', 4))
	fmt.Printf("%c\n", sprint.ShiftBy('a', 4))
	fmt.Println(sprint.StrConcat("r", "ProgrammerHumour", "/"))
	fmt.Println(sprint.ReverseAlphabetValue('a'))
	fmt.Println(sprint.IsNegative(1))
	fmt.Print((sprint.IntVsFloat(5, 8.8)))
	fmt.Println(sprint.IntVsFloat(-8, -8.8))
	fmt.Println(sprint.Season("feb"))
	fmt.Println(sprint.Season("September"))
	fmt.Println(sprint.Accumulate(10))
	fmt.Println(sprint.BetweenLimits('j', 'f'))
	fmt.Println(sprint.Doop(5, "%", 3))
	fmt.Println(sprint.Doop(8, "/", 0))
}
