package day7

import (
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

	sum := 0

	for _, input := range puzzleInput {
		result, operands, _ := strings.Cut(input, ":")
		operands = strings.TrimSpace(operands)

		Operands := make([]int, 0)

		for _, operand := range strings.Split(operands, " ") {
			Operands = append(Operands, helper.ExtractInt(operand))
		}
		equation := &Equation{Result: helper.ExtractInt(result), Operands: Operands}

		if equation.DoesEvaluateToResult() {
			sum += equation.Result
		}
	}

	return sum
}

type Equation struct {
	Result   int
	Operands []int
}

/*
It can be any combination of additions and multiplys!!

This makes it somewhat more interesting. What would the best way of doing it be???

surely it makes sense to come up with some set of permutations which represents the full spectrum of possibilities based on the input...
*/
func (e *Equation) DoesEvaluateToResult() bool {

	sum_result := 0
	// try adding
	for _, operand := range e.Operands {
		sum_result += operand
	}

	if sum_result != e.Result {
		sum_result = e.Operands[0] * e.Operands[1]
		if len(e.Operands) > 2 {
			for i := 2; i < len(e.Operands)-1; i++ {
				sum_result *= e.Operands[i]

			}
		}
	}

	if sum_result == e.Result {
		return true
	}

	return false
}
