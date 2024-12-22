package day5

import (
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
)

type Update struct {
	PageNumbers       []int
	Size              int
	PageOrderingRules map[int][]int
}

func createUpdate(update []string, ordering_rules map[int][]int) Update {
	numbers := make([]int, 0, len(update))
	for _, s := range update {
		numbers = append(numbers, helper.ExtractInt(s))
	}

	return Update{PageNumbers: numbers, Size: len(numbers) - 1, PageOrderingRules: ordering_rules}
}

func (update *Update) getMiddleElement() int {
	return update.PageNumbers[update.Size/2]
}

/*
 */
func (update *Update) isValidUpdate() bool {
	valid_instructions := true

	for i := 0; i < update.Size && valid_instructions; i++ {
		subject_update := update.PageNumbers[i] // this is always going to be the ith element in the list, then take the rest of the list and check against the allowed followers
		allowed_followers := update.PageOrderingRules[subject_update]

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
	return valid_instructions
}

func checkIfPageShouldFollow(page int, allowedFollowers []int) bool {
	allowed := false
	for _, allowedFollower := range allowedFollowers {
		if page == allowedFollower {
			allowed = true
			break
		}
	}
	return allowed
}
