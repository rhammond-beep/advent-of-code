package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := ReadChallengeInput("../day_4_input.txt")
	// input := []string{
	// 	".M.S......",
	// 	"..A..MSMS.",
	// 	".M.S.MAA..",
	// 	"..A.ASMSM.",
	// 	".M.S.M....",
	// 	"..........",
	// 	"S.S.S.S.S.",
	// 	".A.A.A.A..",
	// 	"M.M.M.M.M.",
	// 	"..........",
	// }

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

func ReadChallengeInput(filepath string) (searchSpace []string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		searchSpace = append(searchSpace, scanner.Text())
	}
	return
}
