package day3

import (
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
	"regexp"
	"sort"
)

func SolveDay3Part2() {
	input := helper.ReadChallengeInputContigious("../day_3_input.txt")
	var tokensToProcess []*Token
	multiply_re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`) // Use extra parenthesis to created nested match groups
	multiply_matches_index := multiply_re.FindAllStringSubmatchIndex(input, -1)

	for _, match := range multiply_matches_index {
		calculatedValue := helper.ExtractInt(input[match[2]:match[3]]) * helper.ExtractInt(input[match[4]:match[5]])
		tokensToProcess = append(tokensToProcess, &Token{TokenType: "mult", Position: match[0], Value: calculatedValue})
	}

	do_re := regexp.MustCompile(`do\(\)`)
	dos := do_re.FindAllStringIndex(input, -1)
	for _, do := range dos {
		tokensToProcess = append(tokensToProcess, &Token{TokenType: "do", Position: do[0]})
	}

	dont_re := regexp.MustCompile(`don't\(\)`)
	donts := dont_re.FindAllStringIndex(input, -1)
	for _, dont := range donts {
		tokensToProcess = append(tokensToProcess, &Token{TokenType: "dont", Position: dont[0]})
	}

	sort.Sort(SortToken(tokensToProcess)) // get the tokens in the correct order.

	processTokens(tokensToProcess)
}
