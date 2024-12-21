package day3

import (
	"fmt"
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
	"regexp"
)

func SolveDay3Part1() {
	input := helper.ReadChallengeInputContigious("day_3_input.txt")
	multiply_re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`) // Use extra parenthesis to created nested match groups
	multiply_matches := multiply_re.FindAllStringSubmatch(input, -1)

	sum := 0

	for _, match := range multiply_matches {
		sum += helper.ExtractInt(string(match[1])) * helper.ExtractInt(string(match[2]))
	}
	fmt.Println(sum)
}
