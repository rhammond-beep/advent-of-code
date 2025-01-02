package day6

import "testing"

func TestGuardSouthExit(t *testing.T) {
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

	actualUniquePositions := runSimulation(guards_traversal)

	if expectedUniquePositions != actualUniquePositions {
		t.Fatalf("incorrect number of positions, actual: %v\t wanted: %v", actualUniquePositions, expectedUniquePositions)
	}
}

func TestGuardNorthExit(t *testing.T) {
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

	expectedUniquePositions := 7

	actualUniquePositions := runSimulation(guards_traversal)

	if expectedUniquePositions != actualUniquePositions {
		t.Fatalf("incorrect number of positions, actual: %v\t wanted: %v", actualUniquePositions, expectedUniquePositions)
	}
}

func TestGuardEastExit(t *testing.T) {
	guards_traversal := []string{
		"....#.....",
		"..........",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	expectedUniquePositions := 11

	actualUniquePositions := runSimulation(guards_traversal)

	if expectedUniquePositions != actualUniquePositions {
		t.Fatalf("incorrect number of positions, actual: %v\t wanted: %v", actualUniquePositions, expectedUniquePositions)
	}
}

func TestGuardWestExit(t *testing.T) {
	guards_traversal := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		"....^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	expectedUniquePositions := 22

	actualUniquePositions := runSimulation(guards_traversal)

	if expectedUniquePositions != actualUniquePositions {
		t.Fatalf("incorrect number of positions, actual: %v\t wanted: %v", actualUniquePositions, expectedUniquePositions)
	}
}
