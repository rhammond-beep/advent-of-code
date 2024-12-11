package main

import "testing"

func TestVerticalSearch(t *testing.T) {
	searchSpace := []string{
		"X...X....X",
		"M...X....M",
		"A...M....A",
		"S...A....S",
		"....S.....",
		"....S...X.",
		"X...S...M.",
		"M...S...A.",
		"A...S...S.",
		"S.........",
	}
	number_of_occurrences_wanted := 5

	wordSearch := &WordSearch{SearchSpace: searchSpace, SearchTerm: "XMAS"}
	vs := &VerticalSearch{}
	occurrences := vs.FindTermOccurrences(wordSearch)

	if number_of_occurrences_wanted != occurrences {
		t.Fatalf("Incorrect number of occurrences found in vertical search\nwanted: %v\nactual: %v", number_of_occurrences_wanted, occurrences)
	}
}

func TestHorizontalSearch(t *testing.T) {
	searchSpace := []string{
		"XMAS..XMAS",
		"..........",
		"..........",
		"..........",
		"XXMASSSSS.",
		"..........",
		"..........",
		"..........",
		".....XMAS.",
		"XMAS......",
	}
	number_of_occurrences_wanted := 5

	wordSearch := &WordSearch{SearchSpace: searchSpace, SearchTerm: "XMAS"}
	hs := &HorizontalSearch{}
	occurrences := hs.FindTermOccurrences(wordSearch)

	if number_of_occurrences_wanted != occurrences {
		t.Fatalf("Incorrect number of occurrences found in horizontal search\nwanted: %v\nactual: %v", number_of_occurrences_wanted, occurrences)
	}
}

func TestDiagonalSearch(t *testing.T) {
	searchSpace := []string{
		"X.........",
		".M........",
		"..A.......",
		"...S......",
		"XXMASSSSS.",
		"..X.......",
		"...M......",
		"....A.....",
		".....SMAS.",
		"XMAS..S...",
	}
	number_of_occurrences_wanted := 2

	wordSearch := &WordSearch{SearchSpace: searchSpace, SearchTerm: "XMAS"}
	ds := &DiagonalSearch{}
	occurrences := ds.FindTermOccurrences(wordSearch)

	if number_of_occurrences_wanted != occurrences {
		t.Fatalf("Incorrect number of occurrences found in diagonal search\nwanted: %v\nactual: %v", number_of_occurrences_wanted, occurrences)
	}
}
