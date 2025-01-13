package day7

import (
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
	"strings"
)

/*
Return the total number of the equations which could be made valid using only the + and * oeprators
*/
func SolveDay7Part1() int {

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

		// for a given input of len(Operands) - 1 call it n.
		// construct a list of binary strings, where each element represents a valid permutation E.G:
		// len(operands) = 4, therefore n = 3, which maps to 2^3 possible values, as for each position we have
		// two choices (1 or 0) giving 2! * 2! * 2! = 2^n
		// hence when mapping this out to the set of possible values we get:
		// [000, 001, 010, 011, 100, 101, 110, 111]
		// if we let 0 = + and 1 = *, this maps to:
		// [+++, ++*, +*+, +**, *++, *+*, **+, ***]
		// which can be used to construct the set of all valid equations
		// This binary representation is nice, and easy enough to extend out, as for each supported operator, we increment
		// the number of distinct choices per operand, for example, lets say for arguments sake, that we now support - and %
		// for each pair of operand, we now have 4 choices, meaning in our previous example, we now have (4! * 4! * 4!). A very big number
		// indeed - 13824 possibilities. to be precise. This seems quite bad actually.
		// This smells like a backtracking problem!

		if canMakeValidEquation(result_int, int_operands) {
			calculation += result_int
		}

	}

	return calculation
}

/*
This seems to fit the backtracking problem definition
*/
func canMakeValidEquation(target int, operands []int) bool {
	if len(operands) == 0 { // if after all our operations are performed the target is 0, then it's valid
		return target == 0
	}

	for i := len(operands) - 1; i >= 0; i-- {
		left_branch := canMakeValidEquation(target-operands[i], operands[:i])

		valid_division := target%(operands[i]) == 0
		right_branch := valid_division && canMakeValidEquation(target/operands[i], operands[:i])

		return left_branch || (right_branch && valid_division)
	}

	return false
}

type Permutations struct {
	Input  []rune
	Answer [][]rune
}

func (p *Permutations) ConstructPermutations(index int) {
	if index == len(p.Input) { // we reached some unique permutation, so add it
		newPerm := make([]rune, len(p.Input)) // need to make sure the slice is the same size or no elements will be copied!
		copy(newPerm, p.Input)
		p.Answer = append(p.Answer, newPerm)
		return
	}

	for i := index; i < len(p.Input); i++ {
		swap(p.Input, index, i)
		p.ConstructPermutations(index + 1)
		swap(p.Input, index, i) // undo
	}
}

/*
Simply swap two elements at a given index
*/
func swap(s []rune, i, j int) {
	s[i], s[j] = s[j], s[i]
}

/*
For a given input string, create a list of all the possible permutations
As always, with recursion, we need to know where to stop.
*/
// func ConstructPermutations(index int, s []rune, answer [][]rune) {
// 	if index == len(s) { // we reached some unique permutation, so add it
// 		answer = append(answer, s)
// 		return
// 	}
//
// 	for i := index; i < len(s); i++ {
// 		swap(s, index, i)
// 		ConstructPermutations(index+1, s, answer)
// 		swap(s, index, i) // undo
// 	}
// }

/*
Comprised of a result, operands and operatrors
*/
type Equation struct {
	Result    int
	Operands  []int
	Operators []rune
}

type Expression struct {
	Left     int
	Right    int
	Operator rune
}

/*
Make all of the possible expressions for a given input string
returns a slice containing all of the potential expressions in order I:E -

input: operands = [1 2 3]

output:

	{
		1 -> &{Left: 1, Right: 2, Operator: '+'}, &{Left: 1, Right: 2, Operator: '*'}
		2 -> &{Left: 2, Right: 3, Operator: '+'}, &{Left: 2, Right: 3, Operator: '*'}
	}

But then this doesn't get around the issue of making sure we evalue things in order.

We need to test every possible combination
*/
func constructExpressions(operands []int) map[int][]Expression {
	expressionsOrderMap := make(map[int][]Expression)

	for i := 1; i < len(operands)-1; i++ {
		expressions := make([]Expression, 0, 2)
		expressions = append(expressions, Expression{Left: operands[i-1], Right: operands[i], Operator: '+'})
		expressions = append(expressions, Expression{Left: operands[i-1], Right: operands[i], Operator: '*'})
		expressionsOrderMap[i-1] = expressions
	}

	return expressionsOrderMap
}

func (e *Expression) Evaluate() int {

	if e.Operator == '+' {
		return e.Left + e.Right
	}

	if e.Operator == '*' {
		return e.Left * e.Right
	}

	panic("we really shouldn't be here")
}

/*
It can be any combination of additions and multiplys!!

This makes it somewhat more interesting. What would the best way of doing it be???

surely it makes sense to come up with some set of permutations which represents the full spectrum of possibilities based on the input...

For each pair of operands, we'll need to try every operation avaliable to us (in this case just addition or multiplication)
*/
func (e *Equation) DoesEvaluateToResult() bool {

	// operand_positions := len(e.Operands) - 1 // the number of places the operands can go...

	// construct the set containing all possible permutations of n elements, where n = operand_positions
	// positionMap := make(map[int][]rune)

	// cycle through possible combinations of the Operators (each operator needs a go in a given position)
	// if you get a match then return true
	// otherwise return false

	return false
}
