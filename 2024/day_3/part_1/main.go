package main

import (
	"bufio"
	"fmt"
	"os"
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
 1. As always process the input, in this case, just go line by line, appending the strings into one structure
 2. Create a Scanner class, which Parses out valid tokens, essentially we're looking anything which meets a valid regular expression.
    I don't think I'll do it like that though as I'd like to implement some kind of recursive decent parser..., although that could well just be overkill.
 3. Apply the multiplications and sum the result.
*/
func main() {
	var scanner Scanner
	scanner = &CorruptedScanner{input: ReadChallengeInput("../day_3_input.txt")}
	tokens := scanner.ScanTokens()
	fmt.Println(tokens)
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

type CorruptedScanner struct {
	input   string
	pointer int
}

func (c *CorruptedScanner) ScanTokens() []Token {

}

type Token struct {
	name string
}

type Scanner interface {
	ScanTokens() []Token
}

func GetOutputFromCorruptedScanner(tokens []Token) int {
	return 0
}
