package main

import (
	"bufio"
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

	//

	obstacleMap := buildMap(guards_traversal)

	// While the guard has not tried to walk of the edge of the matrix (i:e X, Y > 0)
	// Move the guard "forward", making a count of the squares visited
	// If there's some obstacle in the way turn 90 Degrees to the right.
	//
}

type Point struct {
	X int
	Y int
}

type Map struct {
	ObstacleLocations map[Point]bool
}

/*
Create a lookup reference for us to refer back to
as the Guard traverses the lab

I could do this as a linked list. Ah no but I'd have no way of determining which direction to go in?
Feels like there should be a better way than making a list of points and obstacles.

If i know where the guard is, and all the points are. I should be able to draw a ray, based on his current posisiton
if there's any obstacles which lie on it's path.

if there are. Draw another ray, from a 90 degree angle to the right.

Keep on doing this until the ray intersects with a matrix limit (I:E X or Y is 0)
*/
func buildMap(input []string) Map {
	obstacleMap := make(map[Point]bool)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			point := Point{X: i, Y: j}
			if input[i][j] == '#' {
				obstacleMap[point] = true
			} else {

				obstacleMap[point] = false
			}
		}
	}

	return Map{ObstacleLocations: obstacleMap}
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
