package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

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
	locationsTraversed := 0

	for {
		foundBarrier, err := obstacleMap.findBarrierOnLine()
		if err != nil { // We're done
			fmt.Println(locationsTraversed)
			break
		}

		locationsTraversed += obstacleMap.calculateDistance(foundBarrier)
	}

}

type Point struct {
	X int
	Y int
}

/*
Helper method for calculating the number of squares between a guard's location
and a given barrier
*/
func (m *Map) calculateDistance(barrierLocation Point) int {
	return int(math.Sqrt(math.Pow((float64(m.GuardLocation.X-barrierLocation.X)), 2) + (math.Pow(float64(m.GuardLocation.Y-barrierLocation.Y), 2))))
}

type Map struct {
	ObstacleLocations map[Point]bool
	GuardLocation     Point
	Direction         string
}

/*
Create a lookup reference for us to refer back to
as the Guard traverses the lab
*/
func buildMap(input []string) Map {
	obstacleMap := make(map[Point]bool)
	var guardLocation Point

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			point := Point{X: i, Y: j}
			if input[i][j] == '#' {
				obstacleMap[point] = true
			} else if input[i][j] == '^' {
				guardLocation = point
			} else {
				obstacleMap[point] = false
			}
		}
	}

	return Map{ObstacleLocations: obstacleMap, GuardLocation: guardLocation, Direction: "north"}
}

/*
draw out a Ray based on where the guard is facing and return a point to that barrier
*/
func (m *Map) findBarrierOnLine() (closestPoint Point, err error) {

	smallest := 1 << 8 // initalise to biggest possible value for 8 bit int

	// find the closest valid value to the guard's current location
	switch m.Direction {

	case "north":
		for key, barrier := range m.ObstacleLocations {
			if key.X < smallest && key.Y == m.GuardLocation.Y && barrier {
				closestPoint = key
			}
		}
		m.Direction = "east"
	case "east":
		m.Direction = "south"
	case "south":
		m.Direction = "west"
	case "west":

		m.Direction = "north"
	default:
		os.Kill.Signal()
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
