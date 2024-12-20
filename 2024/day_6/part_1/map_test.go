package main

import "testing"

func TestMapSouthExit(t *testing.T) {
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

	expectedUniquePositions := 41

	actualUniquePositions := RunSimulation(guards_traversal)

	if expectedUniquePositions != actualUniquePositions {
		t.Fatalf("incorrect number of positions, actual: %v\t wanted: %v", actualUniquePositions, expectedUniquePositions)
	}
}

func TestMapTestNorthExit(t *testing.T) {
	guards_traversal := []string{
		"..........",
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

	expectedUniquePositions := 6

	actualUniquePositions := RunSimulation(guards_traversal)

	if expectedUniquePositions != actualUniquePositions {
		t.Fatalf("incorrect number of positions, actual: %v\t wanted: %v", actualUniquePositions, expectedUniquePositions)
	}
}
