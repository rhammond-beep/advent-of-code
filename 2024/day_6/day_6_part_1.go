package day6

import (
	"fmt"
	helper "github.com/rhammond-beep/advent-of-code-go-helper"
)

func SolveDay6Part1() {
	guards_traversal := helper.ReadChallengeInput("day_6_input.txt")
	fmt.Println(runSimulation(guards_traversal))
}

func runSimulation(input []string) int {
	obstacleMap := BuildMap(input)

	for {
		foundBarrier, exited := obstacleMap.WalkUntilBarrierFound()
		if exited {
			return obstacleMap.CountUniquePositionsVisited()
		}

		obstacleMap.SetGuardDirection(foundBarrier)
	}

}
