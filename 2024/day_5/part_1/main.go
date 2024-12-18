package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
The output should just be the sum of the middle values
of the rows that are determined to be valid!
*/
func main() {
	// update_instructions := ReadChallengeInput("../page_ordering_rules.txt")
	// ordering_rules := ReadChallengeInput("../update_orders.txt")
	ordering_rules := []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
	}

	update_instructions := []string{
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	}

	legal_followers := make(map[int][]int, len(ordering_rules))

	// populate the map
	for _, rule := range ordering_rules {
		before := extractInt(rule[0:2])
		after := extractInt(rule[3:5])

		legal_followers[before] = append(legal_followers[before], after)
	}

	fmt.Printf("final map: %v\n", legal_followers)

	sum := 0

	for _, instruction := range update_instructions {
		update_instructions_to_process := strings.Split(instruction, ",")
		updates_length := len(update_instructions_to_process) - 1
		valid_instructions := true

		// Case where all instructions are present.
		if len(update_instructions) == len(legal_followers)-1 {

			// Seeing as the keys represent a given print page, and all corresponding values that come after it
			// if the update instructions line is valid, the final element's key->value pair won't exist in the map
			// as nothing proceeds it!
			if _, exist := legal_followers[extractInt(update_instructions_to_process[updates_length])]; exist {
				valid_instructions = false
			}
		}

		if valid_instructions { // Take the middle element as instructed.
			sum += extractInt(update_instructions_to_process[len(update_instructions_to_process)/2])
		}
	}

	fmt.Printf("sum of lines: %v", sum)
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
