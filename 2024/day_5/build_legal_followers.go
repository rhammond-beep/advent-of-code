package day5

import (
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
)

func buildLegalFollowers() map[int][]int {
	ordering_rules := helper.ReadChallengeInput("page_ordering_rules.txt")
	legal_followers := make(map[int][]int, len(ordering_rules))

	// populate the map
	for _, rule := range ordering_rules {
		before := helper.ExtractInt(rule[0:2]) // hardcoded assumption here that each page is represented by exactly 2 characters of the original string (and so therefore is two digits in size)
		after := helper.ExtractInt(rule[3:5])

		legal_followers[before] = append(legal_followers[before], after)
	}

	return legal_followers
}
