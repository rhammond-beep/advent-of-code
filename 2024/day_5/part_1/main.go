package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	update_instructions := ReadChallengeInput("../page_ordering_rules.txt")
	// ordering_rules := ReadChallengeInput("../update_orders.txt")
	priority_map := make(map[int][]int, len(update_instructions))

	// populate the map
	for _, instruction := range update_instructions {
		before := extractInt(instruction[0:2])
		after := extractInt(instruction[3:5])

		priority_map[before] = append(priority_map[before], after)
		fmt.Printf("key %v, map: %v\n", before, priority_map[before])
	}
}

func ReadChallengeInput(filepath string) (fileContents []string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileContents = append(fileContents, scanner.Text())
	}
	return
}

func extractInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		os.Exit(-1)
	}
	return int(i)
}
