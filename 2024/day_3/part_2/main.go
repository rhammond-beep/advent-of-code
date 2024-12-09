package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/*
Day 3 - Mull It Over

Goal of the program seems to be: Just multiply some numbers:

It does so with instructions like mul(X, Y)  Where X and Y are each 3 digit numbers:

I:E - mul(44, 46) Multiplies 44 by 56 to get a result of 2024

Because the program's memory has been corrupted, there are also many invalid characters that should be ignored...

Sequences like mul(4*, mul(6,9!, ?(12,34) or mul (2, 4) do nothing.

So the stages of this one seem pretty simple:

 1. We have a strict pattern to adehere to, which can be defined as a regular expression.

 2. This will then form the basis of the input to the following phase, parsing the sanatised resultant expressions
    to extract out integers from the string to perform the operation on

 3. perform the operations.
*/
func main() {
	input := ReadChallengeInput("../day_3_input.txt")
	var tokensToProcess []*Token

	multiply_re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`) // Use extra parenthesis to created nested match groups
	multiply_matches_index := multiply_re.FindAllStringIndex(input, -1)

	for _, match := range multiply_matches_index {
		tokensToProcess = append(tokensToProcess, &Token{TokenType: "mult", Position: match[0]})
	}

	do_re := regexp.MustCompile(`do()`)
	dos := do_re.FindAllStringIndex(input, -1)
	for _, do := range dos {
		tokensToProcess = append(tokensToProcess, &Token{TokenType: "do", Position: do[0]})
	}

	dont_re := regexp.MustCompile(`don't()`)
	donts := dont_re.FindAllStringIndex(input, -1)
	for _, dont := range donts {
		tokensToProcess = append(tokensToProcess, &Token{TokenType: "dont", Position: dont[0]})
	}

	sort.Sort(SortToken(tokensToProcess))

	for _, token := range tokensToProcess {
		fmt.Println(token.String())
	}

}

type Token struct {
	TokenType string
	Value     string
	Position  int
}

type SortToken []*Token

func (t SortToken) Len() int {
	return len(t)
}

func (t SortToken) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t SortToken) Less(i, j int) bool {
	return t[i].Position < t[j].Position
}

func (t Token) String() string {
	return fmt.Sprintf("Type: %v\tValue: %v\tPosition: %v\t", t.TokenType, t.Value, t.Position)
}

type Scanner interface {
	ScanTokens() []Token
}

func extractInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		os.Exit(-1)
	}
	return int(i)
}

func ReadChallengeInput(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sb strings.Builder
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}
	return sb.String()
}
