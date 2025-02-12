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

	puzzleInput := helper.ReadChallengeInput("./day_7/day7_input.txt")
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

/*
Recursively backtrack through the list of possibilities, returning true and winding up the
stack when a valid permutation is found, otherwise, after search space has been exhausted, return false
*/
func canMakeValidEquation2(target int, operands []int) bool {
	if len(operands) == 1 {
		return operands[0] == target
	}

	var tmp int

	tmp = operands[1]
	add := operands[0] + operands[1]
	operands[1] = add
	if canMakeValidEquation2(target, operands[1:]) {
		return true
	}
	operands[1] = tmp

	multiply := operands[0] * operands[1]
	tmp = operands[1]
	operands[1] = multiply
	if canMakeValidEquation2(target, operands[1:]) {
		return true
	}
	operands[1] = tmp

	var buffer bytes.Buffer // Cat strings
	buffer.WriteString(strconv.Itoa(operands[0]))
	buffer.WriteString(strconv.Itoa(operands[1]))
	cat_val := helper.ExtractInt(buffer.String())

	tmp = operands[1]
	operands[1] = cat_val
	if canMakeValidEquation2(target, operands[1:]) {
		return true
	}
	operands[1] = tmp

	return false
}
