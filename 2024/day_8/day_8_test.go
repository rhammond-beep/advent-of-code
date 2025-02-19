package day_8

import "testing"

func TestExampleCaseWorks(t *testing.T) {
	cityMap := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}

	expected := 14
	result := SolveDay8Part1(cityMap)

	if expected != result {
		t.Fatalf("Result was not the same as expected")
	}

}

func TestExampleCase2Works(t *testing.T) {
	cityMap := []string{
		"..........",
		"..........",
		"..........",
		"....a.....",
		"..........",
		".....a....",
		"..........",
		"..........",
		"..........",
		"..........",
	}

	expected := 2
	result := SolveDay8Part1(cityMap)

	if expected != result {
		t.Fatalf("Result was not the same as expected")
	}

}

func TestExampleCase3Works(t *testing.T) {
	cityMap := []string{
		"..........",
		"..........",
		"..........",
		"....a.....",
		"........a.",
		".....a....",
		"..........",
		"..........",
		"..........",
		"..........",
	}

	expected := 4
	result := SolveDay8Part1(cityMap)

	if expected != result {
		t.Fatalf("Result was not the same as expected")
	}
}

func TestExampleCase4Works(t *testing.T) {
	cityMap := []string{
		"..........",
		"...,......",
		"..........",
		"....a.....",
		"........a.",
		".....a....",
		"..,.......",
		"......A...",
		"..........",
		"..........",
	}

	expected := 4
	result := SolveDay8Part1(cityMap)

	if expected != result {
		t.Fatalf("Result was not the same as expected")
	}
}
