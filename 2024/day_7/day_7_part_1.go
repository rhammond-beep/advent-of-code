package day7

import (
	"fmt"
	"strings"

	helper "github.com/rhammond-beep/advent-of-code-go-helper"
)

/*
Return the total number of the equations which could be made valid using only the + and * oeprators
*/
func SolveDay7Part1() int {

	puzzleInput := []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
	}

	for _, input := range puzzleInput {
		result, operands, _ := strings.Cut(input, ":")
		operands = strings.TrimSpace(operands)

		Operands := make([]int, 0)

		for _, operand := range strings.Split(operands, " ") {
			Operands = append(Operands, helper.ExtractInt(operand))
		}
		equation := &Equation{Result: helper.ExtractInt(result), Operands: Operands}

		fmt.Println(equation)
	}

	sum := 0

	return sum
}

type Equation struct {
	Result   int
	Operands []int
}

func (e *Equation) CanEvaluate() bool {
	return true
}
