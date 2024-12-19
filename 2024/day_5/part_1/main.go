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

	legal_followers := buildLegalFollowers()

	sum := 0

	for _, instructions := range update_instructions {
		update := createUpdate(strings.Split(instructions, ","))

		valid_instructions := true

		for i := 0; i < update.Size && valid_instructions; i++ {
			subject_update := update.PageNumbers[i] // this is always going to be the ith element in the list, then take the rest of the list and check against the allowed followers
			allowed_followers := legal_followers[subject_update]

			followers_to_test := update.PageNumbers[i+1 : update.Size+1]

			// Check that each follower is actually allowed to be there
			for _, follower := range followers_to_test {

				present := false
				for _, allowed_follower := range allowed_followers {
					if follower == allowed_follower {
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

		if valid_instructions {
			sum += update.getMiddleElement()
		}
	}

	fmt.Printf("sum of lines: %v", sum)
}

func buildLegalFollowers() map[int][]int {
	ordering_rules := ReadChallengeInput("../page_ordering_rules.txt")
	legal_followers := make(map[int][]int, len(ordering_rules))

	// populate the map
	for _, rule := range ordering_rules {
		before := extractInt(rule[0:2])
		after := extractInt(rule[3:5])

		legal_followers[before] = append(legal_followers[before], after)
	}

	return legal_followers
}

/*
 * Make the update struct
 */
func createUpdate(update []string) Update {
	numbers := make([]int, 0, len(update))
	for _, s := range update {
		numbers = append(numbers, extractInt(s))
	}

	return Update{PageNumbers: numbers, Size: len(numbers) - 1}
}

func (update *Update) getMiddleElement() int {
	return update.PageNumbers[update.Size/2]
}

/*
 */
func (update *Update) isValidUpdate() bool {
	return true
}

type Update struct {
	PageNumbers []int
	Size        int
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
