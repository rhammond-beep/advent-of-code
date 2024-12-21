package day4

import (
	"fmt"
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
)

func SolveDay4Part1Working() {
	input := helper.ReadChallengeInput("day_4_input.txt")

	ans := 0

	for i := range input {
		for j := range input {
			letter := string(input[i][j])
			if letter != "X" {
				continue
			}

			// Backwards Horizontal Case
			if j > 2 && input[i][j-1] == 'M' && input[i][j-2] == 'A' && input[i][j-3] == 'S' {
				ans += 1
			}

			// Forwards Horizontal Case
			if j < len(input)-3 && input[i][j+1] == 'M' && input[i][j+2] == 'A' && input[i][j+3] == 'S' {
				ans += 1
			}

			// Forward vertical Case
			if i > 2 && input[i-1][j] == 'M' && input[i-2][j] == 'A' && input[i-3][j] == 'S' {
				ans += 1
			}

			// Backward vertical Case
			if i < len(input)-3 && input[i+1][j] == 'M' && input[i+2][j] == 'A' && input[i+3][j] == 'S' {
				ans += 1
			}

			// top left-diagonal
			if i > 2 && j > 2 && input[i-1][j-1] == 'M' && input[i-2][j-2] == 'A' && input[i-3][j-3] == 'S' {
				ans += 1
			}

			// top right-diagonal
			if i > 2 && j < len(input)-3 && input[i-1][j+1] == 'M' && input[i-2][j+2] == 'A' && input[i-3][j+3] == 'S' {
				ans += 1
			}

			// bottom right-diagonal
			if i < len(input)-3 && j < len(input)-3 && input[i+1][j+1] == 'M' && input[i+2][j+2] == 'A' && input[i+3][j+3] == 'S' {
				ans += 1
			}

			// bottom left-diagonal
			if i < len(input)-3 && j > 2 && input[i+1][j-1] == 'M' && input[i+2][j-2] == 'A' && input[i+3][j-3] == 'S' {
				ans += 1
			}

		}
	}

	fmt.Println(ans)
}
