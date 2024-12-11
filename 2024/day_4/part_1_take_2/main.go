package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := ReadChallengeInput("../day_4_input.txt")
	targetSet := []string{"M", "A", "S"}

	ans := 0

	for i := range input {
		for j := range input {
			letter := string(input[i][j])
			if letter != "X" {
				continue
			}

			if j > 2 && input[i][j-1] == 'M' && input[i][j-2] == 'A' && input[i][j-3] == 'S' {
				ans += 1
			}

			if i > 2 && input[i-1][j] == 'M' && input[i-2][j] == 'A' && input[i-3][j] == 'S' {
				ans += 1
			}

		}
	}

	fmt.Println(ans)
}

func extractInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		os.Exit(-1)
	}
	return int(i)
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
