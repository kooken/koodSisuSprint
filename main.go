package main

import (
	"fmt"
	"koodSisuSprint/sprint" // This should match your go.mod module name
)

func main() {
	fmt.Println(sprint.Abacus(8, 3))
	fmt.Println(sprint.Abacus(9, 2))
	fmt.Println(sprint.Mean(1.15, -2.0, 8.35))
	fmt.Println(sprint.Casting(1.15))
}
