package day3

import "fmt"

type Token struct {
	TokenType string
	Value     int
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

func processTokens(tokens []*Token) {
	sum := 0
	shouldPerformMultiply := true
	for _, token := range tokens {
		switch token.TokenType {

		case "do":
			shouldPerformMultiply = true
		case "dont":
			shouldPerformMultiply = false
		case "mult":
			if shouldPerformMultiply {
				sum += token.Value
			}
		}
	}
	fmt.Println(sum)
}
