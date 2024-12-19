package main

import (
	"bufio"
	"os"
)

func main() {
	// guards_traversal := ReadChallengeInput("../day_6_input.txt")
	guards_traversal := []string{
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

	buildObstacleMap(guards_traversal)
}

type Point struct {
	X int
	Y int
}

func buildObstacleMap(input []string) map[Point]bool {
	obstacleMap := make(map[Point]bool)

	return obstacleMap
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
