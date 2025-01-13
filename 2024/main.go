package main

import (
	"fmt"
	day7 "rupert-hammond-aoc/day_7"
)

func main() {
	//perm := &day7.Permutations{Input: []rune{'A', 'B', 'C'}, Answer: make([][]rune, 0)}
	//perm.ConstructPermutations(0)

	//for _, inst := range perm.Answer {
	//	fmt.Printf("Computed Permutation %v\n", string(inst))
	//}

	fmt.Println(day7.SolveDay7Part2())
}
