package day5

import (
	"fmt"
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
	"sort"
	"strings"
)

func SolveDay5Part2() {
	update_instructions := helper.ReadChallengeInput("../update_orders.txt")
	legal_followers := buildLegalFollowers()

	sum := 0

	for _, instructions := range update_instructions {
		update := createUpdate(strings.Split(instructions, ","), legal_followers)
		if !update.isValidUpdate() {
			sort.Slice(update.PageNumbers, func(i, j int) bool {
				num1 := update.PageNumbers[i]
				num2 := update.PageNumbers[j]
				return checkIfPageShouldFollow(num2, update.PageOrderingRules[num1])
			})
			sum += update.getMiddleElement()

		}
	}

	fmt.Printf("sum of lines: %v", sum)
}
