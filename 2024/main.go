package main

import (
	"fmt"
	"rupert-hammond-aoc/day_8"

	helper "github.com/rhammond-beep/advent-of-code-go-helper"
)

func main() {
	cityMap := helper.ReadChallengeInput("./day_8/day_8_input.txt")
	fmt.Println(day_8.SolveDay8Part1(cityMap))
}
