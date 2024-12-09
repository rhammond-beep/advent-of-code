package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	multiply_re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`) // Use extra parenthesis to created nested match groups
	multiply_matches := multiply_re.FindAllStringSubmatch(input, -1)

	sum := 0

	for _, match := range multiply_matches {
		sum += extractInt(string(match[1])) * extractInt(string(match[2]))
	}
	fmt.Println(sum)
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

type CorruptedScanner struct {
	input    string
	position int
}

/*
	this approach will take too long, it's not really what i'm going here for at all,

# Should just use a quick and dirty regex approach

Scan through the input string and perform lexical analysis stage
We only want to parse over useful characters, discarding everything else
at this stage, we just want to capture the accuracte representation of valid tokens, nothing else ( we just chuck out
everything that's not
*/
func (c *CorruptedScanner) ScanTokens() []Token {
	var validTokens []Token
	for _, character := range strings.Split(c.input, "") {
		switch character {
		}
	}
	return validTokens
}

type Token struct {
	tokenType string
	lexeme    string
	value     string
	line      int
}

type Scanner interface {
	ScanTokens() []Token
}

func GetOutputFromCorruptedScanner(tokens []Token) int {
	return 0
}
