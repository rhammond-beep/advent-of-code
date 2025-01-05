package day6

import (
	"fmt"
)

func SolveDay6Part2() {
	input := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	lab := BuildMap(input)
	lab.ObstacleLocations[*lab.GuardLocation] = true // ignore the guard's starting position

	obstacleSum := 0

	for point, obstaclePresent := range lab.ObstacleLocations { // For every viable location

		if obstaclePresent { // Then a node already exists, so move on.
			continue
		}

		lab.ObstacleLocations[point] = true

		graph := CreateEmptyGraphFromMap(&lab)
		cycle, path := graph.DetectCycle(&lab)
		fmt.Println("-----------------------------")
		fmt.Printf("Value of Path %v\n, makes cycle?: %v\n", path, cycle)
		fmt.Println("-----------------------------")

		if cycle {
			obstacleSum += 1
		}

		lab.ObstacleLocations[point] = false // reset the location after testing
	}

	fmt.Println(obstacleSum)

	// _, nodes, error := graph.WalkGraphFromNode(day6.Point{X: 6, Y: 4}, day6.Point{X: 3, Y: 2}, "north")
}
