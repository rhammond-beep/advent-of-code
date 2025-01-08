package main

import (
	"fmt"
	day7 "rupert-hammond-aoc/day_7"
)

func main() {
	output := [][]rune{}
	day7.ConstructPermutations(0, []rune{'A', 'B'}, &output)

	for _, perm := range output {
		fmt.Printf("Computed Permutation %v", perm)
	}
}
