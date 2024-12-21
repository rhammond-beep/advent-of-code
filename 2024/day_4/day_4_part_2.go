package day4

import (
	"fmt"
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
)

func SolveDay4Part2() {
	input := helper.ReadChallengeInput("day_4_input.txt")

	ans := 0

	for i := range len(input) - 1 {
		for j := range len(input) - 1 {
			letter := string(input[i][j])
			if letter != "A" {
				continue
			}
			// Regular X
			if j > 0 && i > 0 && input[i-1][j-1] == 'M' && input[i-1][j+1] == 'S' && input[i+1][j-1] == 'M' && input[i+1][j+1] == 'S' {
				ans += 1
			}

			// Top Left diagonal swapped
			if j > 0 && i > 0 && input[i-1][j-1] == 'S' && input[i-1][j+1] == 'M' && input[i+1][j-1] == 'S' && input[i+1][j+1] == 'M' {
				ans += 1
			}

			// Top Right diagonal swapped
			if j > 0 && i > 0 && input[i-1][j-1] == 'M' && input[i-1][j+1] == 'M' && input[i+1][j-1] == 'S' && input[i+1][j+1] == 'S' {
				ans += 1
			}

			// Both diagonal swapped
			if j > 0 && i > 0 && input[i-1][j-1] == 'S' && input[i-1][j+1] == 'S' && input[i+1][j-1] == 'M' && input[i+1][j+1] == 'M' {
				ans += 1
			}

		}
	}

	fmt.Println(ans)
}
