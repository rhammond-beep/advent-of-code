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
	update_instructions := ReadChallengeInput("../update_orders.txt")
	ordering_rules := ReadChallengeInput("../page_ordering_rules.txt")

	legal_followers := make(map[int][]int, len(ordering_rules))

	// populate the map
	for _, rule := range ordering_rules {
		before := extractInt(rule[0:2])
		after := extractInt(rule[3:5])

		legal_followers[before] = append(legal_followers[before], after)
	}

	sum := 0

	for _, instructions := range update_instructions {
		update_instructions_to_process := strings.Split(instructions, ",")
		n := len(update_instructions_to_process) - 1
		valid_instructions := true

		for i := 0; i < n && valid_instructions; i++ {
			subject_update := extractInt(update_instructions_to_process[i]) // this is always going to be the ith element in the list, then take the rest of the list and check against the allowed followers
			allowed_followers := legal_followers[subject_update]

			followers_to_test := update_instructions_to_process[i+1 : n+1]

			// Check that each follower is actually allowed to be there
			for _, follower := range followers_to_test {

				present := false
				for _, allowed_follower := range allowed_followers {
					if extractInt(follower) == allowed_follower {
						present = true
						break
					}
				}
				if !present {
					valid_instructions = false
					break
				}

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
