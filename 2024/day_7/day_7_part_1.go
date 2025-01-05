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
		equation := &Equation{Result: helper.ExtractInt(result), Operands: Operands, Operators: []rune{'+', '*'}}

		if equation.DoesEvaluateToResult() {
			sum += equation.Result
		}
	}

	return sum
}

/*
Comprised of a result, operands and operatrors
*/
type Equation struct {
	Result    int
	Operands  []int
	Operators []rune
}

/*
It can be any combination of additions and multiplys!!

This makes it somewhat more interesting. What would the best way of doing it be???

surely it makes sense to come up with some set of permutations which represents the full spectrum of possibilities based on the input...

For each pair of operands, we'll need to try every operation avaliable to us (in this case just addition or multiplication)
*/
func (e *Equation) DoesEvaluateToResult() bool {

	operand_positions := len(e.Operands) - 1 // the number of places the operands can go...

	// construct the set containing all possible permutations of n elements, where n = operand_positions
	positionMap := make(map[int][]rune)

	for i := 0; i <= operand_positions; i++ {

	}

	return false
}
