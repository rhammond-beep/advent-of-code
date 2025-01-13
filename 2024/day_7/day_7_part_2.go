package day7

import (
	"bytes"
	"strconv"
	"strings"

	helper "github.com/rhammond-beep/advent-of-code-go-helper"
)

/*
Return the total number of the equations which could be made valid using only the + and * oeprators
*/
func SolveDay7Part2() int {

	// puzzleInput := helper.ReadChallengeInput("./day_7/day7_input.txt")
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
		"156: 15 6",
		"7290: 6 8 6 15",
		"192: 17 8 14",
	}
	calculation := 0

	for _, input := range puzzleInput {
		result, operands, _ := strings.Cut(input, ":")
		operands = strings.TrimSpace(operands)

		int_operands := make([]int, 0)
		result_int := helper.ExtractInt(result)

		for _, operand := range strings.Split(operands, " ") {
			int_operands = append(int_operands, helper.ExtractInt(operand))
		}

		if canMakeValidEquation2(result_int, int_operands) {
			calculation += result_int
		}

	}

	return calculation
}

func canMakeValidEquation2(target int, operands []int) bool {
	if len(operands) == 0 { // if after all our operations are performed the target is 0, then it's valid
		return target == 0
	}

	for i := len(operands) - 1; i >= 0; i-- {
		if i > 0 { // handle the append operation first, we want to try to explore concatenated combinations first before moving on
			var buffer bytes.Buffer
			buffer.WriteString(strconv.Itoa(operands[i-1]))
			buffer.WriteString(strconv.Itoa(operands[i]))

			conceternate_branch := helper.ExtractInt(buffer.String())

			cat_operands := make([]int, len(operands))
			copy(cat_operands, operands)

			cat_operands = append(cat_operands[:i], cat_operands[i+1:]...) // remove the element and shift all to the right
			cat_operands[i-1] = conceternate_branch

			if canMakeValidEquation2(target, cat_operands) { // this should generate a bunch more combinations for our core operations but it doesn't work properly or as you would expect it to.
				return true
			}
		}

		plus_branch := canMakeValidEquation2(target-operands[i], operands[:i])
		valid_division := target%(operands[i]) == 0
		multiply_branch := valid_division && canMakeValidEquation2(target/operands[i], operands[:i])

		return plus_branch || (multiply_branch && valid_division)
	}

	return false
}
