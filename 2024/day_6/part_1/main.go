package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

/*
How many DISTINCT positions!! Not the pure number!!
*/
func main() {
	// guards_traversal := ReadChallengeInput("../day_6_input.txt")
	guards_traversal := []string{ // This config will result in there being 41 distinct positions visited
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

	obstacleMap := buildMap(guards_traversal)

	for {
		foundBarrier, exited := obstacleMap.WalkUntilBarrierFound()
		if exited {
			fmt.Println(obstacleMap.CountUniquePositionsVisited())
			break
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

type Point struct {
	X int
	Y int
}

/*
Helper method for calculating the number of squares between a guard's location
and a given barrier

Not even needed lol
*/
func (m *Map) CalculateDistance(barrierLocation Point) int {
	a := math.Pow(float64(m.GuardLocation.X-barrierLocation.X), 2)
	b := math.Pow(float64(m.GuardLocation.Y-barrierLocation.Y), 2)

	return int(math.Sqrt(a+b)) - 1
}

type Map struct {
	ObstacleLocations map[Point]bool
	PositionsVisited  map[Point]bool
	GuardLocation     *Point
	Direction         string
	YUpperBound       int
	XUpperBound       int
}

/*
Create a lookup reference for us to refer back to
as the Guard traverses the lab
*/
func buildMap(input []string) Map {
	obstacleMap := make(map[Point]bool)
	positionsVisited := make(map[Point]bool)
	var guardLocation Point

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			point := Point{X: i, Y: j}
			if input[i][j] == '#' {
				obstacleMap[point] = true
			} else if input[i][j] == '^' {
				guardLocation = point
			} else {
				positionsVisited[point] = false
				obstacleMap[point] = false
			}
		}
	}

	return Map{ObstacleLocations: obstacleMap, GuardLocation: &guardLocation, Direction: "north", YUpperBound: len(input), XUpperBound: len(input), PositionsVisited: positionsVisited}
}

/*
draw out a Ray based on where the guard is facing and return a point representing a barrier
*/
func (m *Map) WalkUntilBarrierFound() (Point, bool) {
	var closestPoint Point

	switch m.Direction {

	case "north":
		for i := m.GuardLocation.X - 1; i >= 0; i-- {
			point2Test := Point{X: i, Y: m.GuardLocation.Y}

			if !m.ObstacleLocations[point2Test] {
				m.PositionsVisited[point2Test] = true
			} else {
				closestPoint = point2Test
				return closestPoint, false
			}
		}
	case "east":
		for i := m.GuardLocation.Y + 1; i < m.YUpperBound; i++ {
			point2Test := Point{X: m.GuardLocation.X, Y: i}

			if !m.ObstacleLocations[point2Test] {
				m.PositionsVisited[point2Test] = true
			} else {
				closestPoint = point2Test
				return closestPoint, false
			}
		}
	case "south":
		for i := m.GuardLocation.X + 1; i < m.XUpperBound; i++ {
			point2Test := Point{X: i, Y: m.GuardLocation.Y}

			if !m.ObstacleLocations[point2Test] {
				m.PositionsVisited[point2Test] = true
			} else {
				closestPoint = point2Test
				return closestPoint, false
			}
		}
	case "west":
		for i := m.GuardLocation.Y - 1; i >= 0; i-- {
			point2Test := Point{X: m.GuardLocation.X, Y: i}

			if !m.ObstacleLocations[point2Test] {
				m.PositionsVisited[point2Test] = true
			} else {
				closestPoint = point2Test
				return closestPoint, false
			}
		}
	default:
		panic("we really shouldn't be here")
	}

	return closestPoint, true
}

/*
I can't think of a clever way of
*/
func (m *Map) CountUniquePositionsVisited() (positionsVisited int) {
	for _, value := range m.PositionsVisited {
		if value {
			positionsVisited += 1
		}
	}
	return
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
