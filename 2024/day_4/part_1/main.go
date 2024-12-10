package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Day 4 - Ceres Search
*/
func main() {
	input := ReadChallengeInput("../day_4_input.txt")
	fmt.Println(input)
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
