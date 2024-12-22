package day5

import (
	"fmt"
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
	"strings"
)

func SolveDay5Part1() {
	update_instructions := helper.ReadChallengeInput("update_orders.txt")

	legal_followers := buildLegalFollowers()

	sum := 0

	for _, instructions := range update_instructions {
		update := createUpdate(strings.Split(instructions, ","), legal_followers)

		if update.isValidUpdate() {
			sum += update.getMiddleElement()
		}
	}

	fmt.Printf("sum of lines: %v", sum)
}
