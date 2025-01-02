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

		// set the guard's location for the next iteration
		switch obstacleMap.Direction {
		case "north":
			obstacleMap.GuardLocation = &Point{X: foundBarrier.X + 1, Y: foundBarrier.Y}
			obstacleMap.Direction = "east"
		case "east":
			obstacleMap.GuardLocation = &Point{X: foundBarrier.X, Y: foundBarrier.Y - 1}
			obstacleMap.Direction = "south"
		case "south":
			obstacleMap.GuardLocation = &Point{X: foundBarrier.X - 1, Y: foundBarrier.Y}
			obstacleMap.Direction = "west"
		case "west":
			obstacleMap.GuardLocation = &Point{X: foundBarrier.X, Y: foundBarrier.Y + 1}
			obstacleMap.Direction = "north"

		}
	}

}
